package constantx

// User roles for JWT
// role = {0: "company", 1: "contractor", 2: "customer"}
const (
	Company int = iota
	Contractor
	Customer
	Admin
)
