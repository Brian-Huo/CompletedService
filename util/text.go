package util

import (
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/payment"
	"strconv"
)

// Check address strings valid
func CheckAddressDetails(data *address.BAddress) bool {
	return true
}

// Check contact details valid
func CheckContactDetails(contact_details string) bool {
	if _, err := strconv.Atoi(contact_details); err != nil {
		return false
	}
	return true
}

// Check payment details valid
func CheckPaymentDetails(data *payment.BPayment) bool {
	return true
}
