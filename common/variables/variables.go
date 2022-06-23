package variables

// system strings separator
const Separator string = "∑"

// system deposite rate
const Deposite_rate float64 = 10.0

// GST for cleaning services
const GST float64 = 10.0

// Inwork distance
const Inwork_distance float64 = 1

// role = {0: "company", 1: "contractor", 2: "customer"}
const (
	Company int = iota
	Contractor
	Customer
	Admin
)

// Checking Timer in
const Check_time int64 = 86400

// Reception Email
const Reception_email string = "akknowall@gmail.com"

// QME Offical Email
const QME_email string = "qmetechnology@gmail.com"

// QME Offical Password
const QME_password string = "kwguicxjpjrabrbo"
