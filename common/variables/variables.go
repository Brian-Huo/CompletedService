package variables

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
)

// order status enum
const (
	Queuing int = iota
	Waiting
	Working
	Completed
	Cancelled
)
