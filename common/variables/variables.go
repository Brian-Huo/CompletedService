package variables

// system strings separator
const Separator string = "∑"

// system deposite rate
const Deposite_rate float64 = 10.0

// GST for cleaning services
const GST float64 = 10.0

// role = {0: "company", 1: "contractor", 2: "customer"}
const (
	Company int = iota
	Contractor
	Customer
)
