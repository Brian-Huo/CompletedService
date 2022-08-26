package variables

// system strings separator
const Separator string = "∑"

// system deposite rate
const Deposite_rate float64 = 10.0

// GST for cleaning services
const GST float64 = 10.0

// Factor of order surcharge
const Surcharge_factor int = 20

// Inwork distance
const Inwork_distance float64 = 1

// role = {0: "company", 1: "contractor", 2: "customer"}
const (
	Company int = iota
	Contractor
	Customer
	Admin
)

// Checking Timer in order queue
const Check_time_unit int64 = 3600
const Check_time_clock int = 24
