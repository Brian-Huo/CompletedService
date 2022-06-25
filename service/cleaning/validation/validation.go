package validation

import (
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/payment"
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

// Check customer phone valid
func CheckCustomerPhone(customer_phone string) bool {
	re := regexp.MustCompile(`^[+]*[(]{0,1}[0-9]{1,4}[)]{0,1}[-\s\./0-9]*$`)
	if len(customer_phone) <= 7 {
		return false
	}
	return re.MatchString(customer_phone)
}

// Check customer email valid
func CheckCustomerEmail(customer_mail string) bool {
	re := regexp.MustCompile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if len(customer_mail) <= 7 {
		return false
	}
	return re.MatchString(customer_mail)
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
