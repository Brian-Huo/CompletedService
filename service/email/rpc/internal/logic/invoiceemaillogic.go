package logic

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"cleaningservice/common/variables"
	"cleaningservice/service/email/rpc/internal/svc"
	"cleaningservice/service/email/rpc/types/email"

	generator "github.com/angelodlfrtr/go-invoice-generator"
	"github.com/zeromicro/go-zero/core/logx"
	"gopkg.in/gomail.v2"
)

type InvoiceEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const invoiceEmailMsg = ""

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
	})

	doc.SetHeader(&generator.HeaderFooter{
		Text:       "<center>CleaningService</center>",
		Pagination: true,
	})

	doc.SetFooter(&generator.HeaderFooter{
		Text:       "<center>Important notice: <br> Powered By QME Technology Ptd. Lty.</center>",
		Pagination: true,
	})

	doc.SetRef(fmt.Sprintf("Invoice_%d", in.OrderInfo.OrderId))
	doc.SetVersion("version 0.1")

	doc.SetDescription("Invoice of Items")
	doc.SetNotes("Suppose to be payment details here! ")

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

	doc.SetCustomer(&generator.Contact{
		Name: in.CustomerInfo.CustomerName,
		Address: &generator.Address{
			Address:    in.AddressInfo.Formatted,
			PostalCode: in.AddressInfo.Postcode,
			City:       in.AddressInfo.City,
			Country:    in.AddressInfo.Country,
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

	doc.SetDefaultTax(&generator.Tax{
		Percent: "10",
	})

	// doc.SetDiscount(&generator.Discount{
	// Percent: "90",
	// })
	// doc.SetDiscount(&generator.Discount{
	// 	Amount: "1340",
	// })

	pdf, err := doc.Build()
	if err != nil {
		log.Fatal(err)
		return &email.InvoiceEmailResponse{
			Code: 500,
			Msg:  "Create invoice PDF failed",
		}, err
	}

	invoicePath := l.svcCtx.Config.InvoiceLocation + "/Invoice_" + fmt.Sprintf("%d", in.OrderInfo.OrderId) + ".pdf"
	err = pdf.OutputFileAndClose(invoicePath)

	if err != nil {
		log.Fatal(err)
		return &email.InvoiceEmailResponse{
			Code: 500,
			Msg:  "Save invoice PDF failed",
		}, err
	}

	// Send Invoice email
	// Set attributes
	subject := fmt.Sprintf("Invoice [%d] for [%s] due [%s]", in.GetOrderInfo().OrderId, in.GetCategoryInfo().CategoryName, doc.PaymentTerm)
	emailHi := fmt.Sprintf("<p>Hi %s,</p><br>", in.GetCustomerInfo().CustomerName)
	emailGreetings := fmt.Sprintf("<p>Thanks for choosing %s.</p><br>", variables.Business_name)
	emailMain := fmt.Sprintf("<p>Attached is your %s at address %s at %s. Please be awared of your reservation time.</p><br>", subject, in.GetAddressInfo().Formatted, in.GetOrderInfo().ReserveDate)
	emailPayment := fmt.Sprintf("<p>When making payment, you must include your Reference number <em>%s</em> stated in your invoice.</p><br>", doc.Ref)
	emailEnd := fmt.Sprintf("<p>If you have any questions and concerns, you can kindly reply to this email</p><br><p>Kind regards,</p><br>")
	emailSign := fmt.Sprintf("%s Support Team", variables.Business_name)

	// Send email
	m := gomail.NewMessage()
	m.SetHeader("From", variables.QME_email)
	m.SetHeader("To", in.GetCustomerInfo().CustomerEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", emailHi+emailGreetings+emailMain+emailPayment+emailEnd+emailSign)

	// Send the email
	d := gomail.NewDialer("smtp.gmail.com", 587, variables.QME_email, variables.QME_password)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal(err)
		return &email.InvoiceEmailResponse{
			Code: 500,
			Msg:  "Send invoice email failed",
		}, err
	}

	return &email.InvoiceEmailResponse{
		Code: 200,
		Msg:  "success",
	}, nil
}
