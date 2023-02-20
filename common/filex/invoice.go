package filex

import (
	"log"

	generator "github.com/angelodlfrtr/go-invoice-generator"
)

const invoiceFolder string = "../../../static/invoices/"

func SaveInvoice(doc *generator.Document, docRef string) (string, error) {
	pdf, err := doc.Build()
	if err != nil {
		log.Fatal(err)
		return "Save invoice failed", err
	}

	invoicePath := invoiceFolder + "/Invoice_" + docRef + ".pdf"
	err = pdf.OutputFileAndClose(invoicePath)

	if err != nil {
		log.Fatal(err)
		return "Save invoice failed", err
	}

	return invoicePath, nil
}
