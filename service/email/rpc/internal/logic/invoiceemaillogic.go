package logic

import (
	"context"
	"fmt"
	"io/ioutil"
	"math"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/email/rpc/internal/svc"
	"cleaningservice/service/email/rpc/types/email"
	"cleaningservice/util"

	generator "github.com/angelodlfrtr/go-invoice-generator"
	"github.com/zeromicro/go-zero/core/logx"
)

type InvoiceEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewInvoiceEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InvoiceEmailLogic {
	return &InvoiceEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *InvoiceEmailLogic) InvoiceEmail(in *email.InvoiceEmailRequest) (*email.InvoiceEmailResponse, error) {
	// Create Invoice PDF
	doc, _ := generator.New(generator.Invoice, &generator.Options{
		TextTypeInvoice: "QUETO",
		AutoPrint:       true,
		CurrencySymbol:  "$",
	})

	// Set constant variables for invoice
	doc.SetHeader(&generator.HeaderFooter{
		Text:       "<center>CleaningService</center>",
		Pagination: true,
	})

	doc.SetFooter(&generator.HeaderFooter{
		Text:       "<center>Important notice:</center><br><center>Powered By QME Technology Ptd. Lty.</center>",
		Pagination: true,
	})

	doc.SetVersion("version 0.2")
	doc.SetDescription("Invoice of Items")
	doc.SetNotes(fmt.Sprintf("<b>BSB: %s </b><br><b>Account No: %s</b><br><b> Account Name: %s</b>", variables.BSB, variables.Account_number, variables.Account_name))

	doc.SetDate(time.Now().Format("02/01/2006"))
	doc.SetPaymentTerm(time.Now().Add(time.Hour * 168).Format("02/01/2006"))

	logoBytes, _ := ioutil.ReadFile(variables.Business_logo)

	doc.SetCompany(&generator.Contact{
		Name: variables.Business_name,
		Logo: &logoBytes,
		Address: &generator.Address{
			Address:    variables.Unit,
			Address2:   variables.Street,
			PostalCode: variables.PostalCode,
			City:       variables.City,
			Country:    variables.Country,
		},
	})

	doc.SetDefaultTax(&generator.Tax{
		Percent: "10",
	})

	// doc.SetDiscount(&generator.Discount{
	// Percent: "90",
	// })
	// doc.SetDiscount(&generator.Discount{
	// 	Amount: "1340",
	// })

	// Set invoice variables
	doc.SetRef(fmt.Sprintf("Invoice_%d", in.OrderInfo.OrderId))

	doc.SetCustomer(&generator.Contact{
		Name: in.CustomerInfo.CustomerName,
		Address: &generator.Address{
			Address:    in.AddressInfo.Formatted,
			PostalCode: in.AddressInfo.Postcode,
			City:       in.AddressInfo.City,
			Country:    variables.Country,
		},
	})

	for _, service_item := range in.ServiceInfo {
		doc.AppendItem(&generator.Item{
			Name:        service_item.ServiceName,
			Description: service_item.ServiceDescription,
			UnitCost:    fmt.Sprintf("%f", service_item.ServicePrice),
			Quantity:    fmt.Sprintf("%d", service_item.ServiceQuantity),
			Tax: &generator.Tax{
				Percent: fmt.Sprintf("%.2f", variables.GST),
			},
			// Discount: &generator.Discount{
			// 	Percent: "0",
			// },
		})
	}

	// Set order surcharge
	doc.AppendItem(&generator.Item{
		Name:        in.OrderInfo.SurchargeItem,
		Description: in.OrderInfo.SurchargeDescription,
		UnitCost:    fmt.Sprintf("%f", in.OrderInfo.SurchargeAmount),
		Quantity:    "1",
		Tax: &generator.Tax{
			Percent: fmt.Sprintf("%.2f", variables.GST),
		},
	})

	// Discount all decimals
	doc.SetDiscount(&generator.Discount{
		Amount: fmt.Sprintf("%f", in.OrderInfo.TotalAmount-math.Trunc(in.OrderInfo.TotalAmount)),
	})

	// Set invoice pdf
	invoiceLocation, err := util.SaveInvoice(doc, in.OrderInfo.OrderId)
	if err != nil {
		return &email.InvoiceEmailResponse{
			Code: 500,
			Msg:  "Save invoice PDF failed",
		}, err
	}

	// Send Invoice email
	// Set attributes
	target := in.GetCustomerInfo().CustomerEmail
	subject := fmt.Sprintf("Invoice [%d] for [%s] due [%s]", in.GetOrderInfo().OrderId, in.GetCategoryInfo().CategoryName, doc.PaymentTerm)
	emailHi := fmt.Sprintf("<p>Hi %s,</p><br>", in.GetCustomerInfo().CustomerName)
	emailGreetings := fmt.Sprintf("<p>Thanks for choosing %s.</p><br>", variables.Business_name)
	emailMain := fmt.Sprintf("<p>Attached is your %s at address %s at %s. Please be awared of your reservation time.</p><br>", subject, in.GetAddressInfo().Formatted, in.GetOrderInfo().ReserveDate)
	emailPayment := fmt.Sprintf("<p>When making payment, you must include your Reference number <em>%s</em> stated in your invoice.</p><br>", doc.Ref)

	// Send email
	go util.SendToClient(target, subject, emailHi+emailGreetings+emailMain+emailPayment, invoiceLocation)

	return &email.InvoiceEmailResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
