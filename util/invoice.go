package util

import (
	"fmt"
	"log"
	"time"

	generator "github.com/angelodlfrtr/go-invoice-generator"
)

const invoiceFolder string = "../../../static/invoices/"

func SaveInvoice(doc *generator.Document, orderId int64) (string, error) {
	pdf, err := doc.Build()
	if err != nil {
		log.Fatal(err)
		return "Save invoice failed", err
	}

	invoicePath := invoiceFolder + fmt.Sprintf("/Invoice_%d%04d.pdf", time.Now().Unix(), orderId)
	err = pdf.OutputFileAndClose(invoicePath)

	if err != nil {
		log.Fatal(err)
		return "Save invoice failed", err
	}

	return invoicePath, nil
}
