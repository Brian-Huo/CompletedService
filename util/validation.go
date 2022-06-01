package util

import (
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/payment"
	"errors"
	"regexp"
)

// Check address strings valid
func CheckAddressDetails(data *address.BAddress) (bool, error) {
	if len(data.Formatted) < 10 {
		return false, errors.New("invalid formatted address")
	}
	if data.Lat == 0.0 && data.Lng == 0.0 {
		return false, errors.New("invalid lat or lng")
	}
	return true, nil
}

// Check contact details valid
func CheckContactDetails(contact_details string) bool {
	re := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	if len(contact_details) <= 7 {
		return false
	}
	return re.MatchString(contact_details)
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
