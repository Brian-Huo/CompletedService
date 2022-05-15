package variables

// system strings separator
const Separator string = "%/%"

// system deposite rate
const Deposite_rate float64 = 10.0

// GST for cleaning services
const GST float64 = 10.0

// role = {0: "company", 1: "employee", 2: "customer"}
const (
	Company int = iota
	Employee
	Customer
)

// operation enum
const (
	Decline int = iota
	Accept
)

// workstatus enum
const (
	Vacant int = iota
	InWork
	Resigned
)

// company status enum
const (
	Abolished int = 0
	Active    int = 1
)

// order status enum
const (
	Queuing int = iota
	Pending
	Working
	Unpaid
	Completed
	Cancelled
)
