package util

import (
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/payment"
)

// Check address strings valid
func CheckAddressDetails(data *address.BAddress) bool {
	return true
}

// Check contact details valid
func CheckContactDetails(contact_details string) bool {
	return true
}

// Check payment details valid
func CheckPaymentDetails(data *payment.BPayment) bool {
	return true
}
