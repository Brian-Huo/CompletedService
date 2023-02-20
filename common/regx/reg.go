package regx

import "regexp"

const addressRegex = `^(.+?), ([a-zA-Z\s]+) (\b(?:VIC|QLD|NSW|WA|SA|TAS|ACT|NT)\b) ([0-9]{4})`

type addressMtr struct {
	Street    string `json:"street"`
	Suburb    string `json:"suburb"`
	StateCode string `json:"statecode"`
	Postcode  string `json:"postcode"`
	Formatted string `json:"formatted"`
}

func GetAddressAttributes(formatted string) *addressMtr {
	re := regexp.MustCompile(addressRegex)
	matched := re.FindStringSubmatch(formatted)

	newAddress := &addressMtr{
		Street:    matched[1],
		Suburb:    matched[2],
		StateCode: matched[3],
		Postcode:  matched[4],
		Formatted: matched[0],
	}

	return newAddress
}
