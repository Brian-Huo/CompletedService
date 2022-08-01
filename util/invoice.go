package util

import (
	"fmt"
	"log"

	generator "github.com/angelodlfrtr/go-invoice-generator"
)

const invoiceFolder string = "../../../static/invoices/"

func SaveInvoice(doc *generator.Document, orderId int64) (string, error) {
	pdf, err := doc.Build()
	if err != nil {
		log.Fatal(err)
		return "Save invoice failed", err
	}

	invoicePath := invoiceFolder + "/Invoice_" + fmt.Sprintf("%d", orderId) + ".pdf"
	err = pdf.OutputFileAndClose(invoicePath)

	if err != nil {
		log.Fatal(err)
		return "Save invoice failed", err
	}

	return invoicePath, nil
}
