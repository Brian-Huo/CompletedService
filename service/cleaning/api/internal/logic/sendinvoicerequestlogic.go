package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"cleaningservice/common/orderqueue"
	"cleaningservice/service/cleaning/api/internal/svc"
	"cleaningservice/service/cleaning/api/internal/types"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/category"
	"cleaningservice/service/cleaning/model/customer"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/service/cleaning/model/region"
	"cleaningservice/service/cleaning/model/service"
	"cleaningservice/service/email/rpc/types/email"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendInvoiceRequestLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendInvoiceRequestLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendInvoiceRequestLogic {
	return &SendInvoiceRequestLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendInvoiceRequestLogic) SendInvoiceRequest() (err error) {
	// Get all request for creating invoice
	if orderqueue.IsEmpty() {
		return nil
	}
	processqueue := orderqueue.PullAll()

	// Get information and send request
	for order_id := range processqueue {
		fmt.Printf("%d", order_id)
		// Order Info
		order_item, err := l.svcCtx.BOrderModel.FindOne(l.ctx, order_id)
		if err != nil {
			if err == order.ErrNotFound {
				logx.Info("Order not found on order", order_id)
				orderqueue.DeleteOne(order_id)
				continue
			}
			return err
		}

		order_email := email.OrderMsg{
			OrderId:              order_item.OrderId,
			DepositeAmount:       order_item.DepositeAmount,
			FinalAmount:          order_item.FinalAmount,
			DepositeRate:         int32(order_item.CurrentDepositeRate),
			GstAmount:            order_item.GstAmount,
			ItemAmount:           order_item.ItemAmount,
			TotalAmount:          order_item.TotalAmount,
			ReserveDate:          order_item.ReserveDate.Format("02/01/2006 15:04:05"),
			SurchargeItem:        order_item.SurchargeItem,
			SurchargeRate:        int32(order_item.SurchargeRate),
			SurchargeAmount:      order_item.SurchargeAmount,
			SurchargeDescription: order_item.SurchargeDescription.String,
		}

		// Customer Info
		customer_item, err := l.svcCtx.BCustomerModel.FindOne(l.ctx, order_item.CustomerId)
		if err != nil {
			if err == customer.ErrNotFound {
				logx.Alert("Customer not found")
				orderqueue.DeleteOne(order_id)
				continue
			}
			return err
		}

		customer_email := email.CustomerMsg{
			CustomerId:    customer_item.CustomerId,
			CustomerName:  customer_item.CustomerName,
			CustomerType:  customer_item.CustomerType,
			CountryCode:   customer_item.CountryCode,
			CustomerPhone: customer_item.CustomerPhone,
			CustomerEmail: customer_item.CustomerEmail,
		}

		// Address Info
		address_item, err := l.svcCtx.BAddressModel.FindOne(l.ctx, order_item.AddressId)
		if err != nil {
			if err == address.ErrNotFound {
				logx.Alert("Address not found")
				orderqueue.DeleteOne(order_id)
				continue
			}
			return err
		}
		// Get region details
		region_item, err := l.svcCtx.BRgionModel.FindOneByPostcode(l.ctx, address_item.Postcode)
		if err != nil {
			if err == region.ErrNotFound {
				logx.Alert("Address not found")
				orderqueue.DeleteOne(order_id)
				continue
			}
			return err
		}

		address_email := email.AddressMsg{
			AddressId: address_item.AddressId,
			Street:    address_item.Street,
			Suburb:    address_item.Suburb,
			Postcode:  address_item.Postcode,
			City:      address_item.City,
			StateCode: region_item.StateCode,
			Formatted: address_item.Formatted,
		}

		// Category Info
		category_item, err := l.svcCtx.BCategoryModel.FindOne(l.ctx, order_item.CategoryId)
		if err != nil {
			if err == category.ErrNotFound {
				logx.Alert("Category not found")
				orderqueue.DeleteOne(order_id)
				continue
			}
			return err
		}

		category_email := email.CategoryMsg{
			CategoryId:          category_item.CategoryId,
			CategoryAbbr:        category_item.CategoryAddr,
			CategoryName:        category_item.CategoryName,
			CategoryDescription: category_item.CategoryDescription,
		}

		// Services Info
		var service_email []*email.ServiceMsg
		// Get Basic Service Details
		var basic_items types.SelectedServiceList
		err = json.Unmarshal([]byte(order_item.BasicItems), &basic_items)
		if err != nil {
			return err
		}

		// Append
		for _, basic_item := range basic_items.Items {
			service_item, err := l.svcCtx.BServiceModel.FindOne(l.ctx, basic_item.Service_id)
			if err != nil {
				if err == service.ErrNotFound {
					logx.Alert("Service not found [Basic items]")
					orderqueue.DeleteOne(order_id)
					continue
				}
				return err
			}
			service_email = append(service_email, &email.ServiceMsg{
				ServiceId:          basic_item.Service_id,
				ServiceScope:       basic_item.Service_scope,
				ServiceName:        basic_item.Service_name,
				ServiceDescription: service_item.ServiceDescription,
				ServiceQuantity:    int32(basic_item.Service_quantity),
				ServicePrice:       basic_item.Service_price,
			})
		}

		// Get Additional Service Details
		var additional_items types.SelectedServiceList
		err = json.Unmarshal([]byte(order_item.AdditionalItems.String), &additional_items)
		if err != nil {
			return err
		}

		for _, addition_item := range additional_items.Items {
			// Append
			service_item, err := l.svcCtx.BServiceModel.FindOne(l.ctx, addition_item.Service_id)
			if err != nil {
				logx.Info("Get additionl item(s) failed", err)
				continue
			}

			service_email = append(service_email, &email.ServiceMsg{
				ServiceId:          addition_item.Service_id,
				ServiceScope:       addition_item.Service_scope,
				ServiceName:        addition_item.Service_name,
				ServiceDescription: service_item.ServiceDescription,
				ServiceQuantity:    int32(addition_item.Service_quantity),
				ServicePrice:       addition_item.Service_price,
			})
		}

		// Send order Invoice
		_, err = l.svcCtx.EmailRpc.InvoiceEmail(l.ctx, &email.InvoiceEmailRequest{
			AddressInfo:  &address_email,
			CategoryInfo: &category_email,
			CustomerInfo: &customer_email,
			ServiceInfo:  service_email,
			OrderInfo:    &order_email,
		})
		if err != nil {
			logx.Alert("Send invoice email failded on order " + strconv.Itoa(int(order_id)))
		} else {
			logx.Info("Send invoice email success on order ", strconv.Itoa(int(order_id)))
			orderqueue.DeleteOne(order_id)
		}
	}
	orderqueue.IterationFinish()

	return nil
}
