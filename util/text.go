package util

import (
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/payment"
	"regexp"
	"strconv"
)

// Check address strings valid
func CheckAddressDetails(data *address.BAddress) bool {
	if len(data.Formatted) < 10 {
		return false
	}
	if _, err := strconv.Atoi(data.Postcode); err != nil {
		return false
	}
	if data.City != "" {
		return false
	}
	if data.Lat == 0.0 || data.Lng == 0.0 {
		return false
	}
	return true
}

// Check contact details valid
func CheckContactDetails(contact_details string) bool {
	if len(contact_details) <= 7 {
		return false
	}
	if _, err := strconv.Atoi(contact_details); err != nil {
		return false
	}
	return true
}

// Check payment details valid
func CheckPaymentDetails(data *payment.BPayment) bool {
	re := regexp.MustCompile(`^5[1-5][0-9][14]$`)

	if data.HolderName != "WIX Corp." {
		return re.MatchString(data.CardNumber)
	} else {
		return data.CardNumber == "0000000000000000"
	}
}
