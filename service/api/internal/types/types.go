// Code generated by goctl. DO NOT EDIT.
package types

type VerifyCodeRequest struct {
	Contact_details string `json:"contact_details"`
}

type VerifyCodeResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LoginServiceRequest struct {
	Contact_details string `json:"contact_details"`
	VerifyCode      string `json:"verify_code"`
}

type LoginServiceResponse struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token,optional"`
}

type LoginCustomerRequest struct {
	Contact_details string `json:"contact_details"`
	VerifyCode      string `json:"verify_code"`
}

type LoginCustomerResponse struct {
	Code        string `json:"code"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token,optional"`
}

type CreateAddressRequest struct {
	Address_details string `json:"address_details"`
	Suburb          string `json:"suburb"`
	Postcode        string `json:"postcode"`
	State_code      string `json:"state_code"`
	Country         string `json:"country,optional"`
}

type CreateAddressResponse struct {
	Address_id int64 `json:"address_id"`
}

type UpdateAddressRequest struct {
	Address_id      int64  `json:"address_id"`
	Address_details string `json:"address_details"`
	Suburb          string `json:"suburb"`
	Postcode        string `json:"postcode"`
	State_code      string `json:"state_code"`
	Country         string `json:"country"`
}

type UpdateAddressResponse struct {
}

type RemoveAddressRequest struct {
	Address_id int64 `json:"address_id"`
}

type RemoveAddressResponse struct {
}

type DetailAddressRequest struct {
	Address_id int64 `json:"address_id"`
}

type DetailAddressResponse struct {
	Address_id      int64  `json:"address_id"`
	Address_details string `json:"address_details"`
	Suburb          string `json:"suburb"`
	Postcode        string `json:"postcode"`
	State_code      string `json:"state_code"`
	Country         string `json:"country"`
}

type ListAddressRequest struct {
}

type ListAddressResponse struct {
	Items []DetailAddressResponse `json:"items"`
}

type CreateCompanyRequest struct {
	Company_name       string `json:"company_name"`
	Payment_id         int64  `json:"payment_id"`
	Director_name      string `json:"director_name"`
	Contact_details    string `json:"contact_details"`
	Registered_address int64  `json:"registered_address"`
	Deposite_rate      int    `json:"deposite_rate"`
}

type CreateCompanyResponse struct {
	Company_id int64 `json:"company_id"`
}

type UpdateCompanyRequest struct {
	Company_id         int64  `json:"company_id"`
	Company_name       string `json:"company_name"`
	Payment_id         int64  `json:"payment_id,optional"`
	Director_name      string `json:"director_name,optional"`
	Contact_details    string `json:"contact_details"`
	Registered_address int64  `json:"registered_address,optional"`
	Deposite_rate      int    `json:"deposite_rate"`
}

type UpdateCompanyResponse struct {
}

type RemoveCompanyRequest struct {
	Company_id int64 `json:"company_id"`
}

type RemoveCompanyResponse struct {
}

type DetailCompanyRequest struct {
	Company_id int64 `json:"company_id"`
}

type DetailCompanyResponse struct {
	Company_id         int64  `json:"company_id"`
	Company_name       string `json:"company_name"`
	Payment_id         int64  `json:"payment_id"`
	Director_name      string `json:"director_name"`
	Contact_details    string `json:"contact_details"`
	Registered_address int64  `json:"registered_address"`
	Deposite_rate      int    `json:"deposite_rate"`
}

type ListCompanyRequest struct {
}

type ListCompanyResponse struct {
	Items []DetailCompanyResponse `json:"items"`
}

type CreateCustomerRequest struct {
	Customer_name   string `json:"customer_name"`
	Country_code    string `json:"country_code"`
	Contact_details string `json:"contact_details"`
}

type CreateCustomerResponse struct {
	Customer_id int64 `json:"customer_id"`
}

type UpdateCustomerRequest struct {
	Customer_id     int64  `json:"customer_id"`
	Customer_name   string `json:"customer_name"`
	Country_code    string `json:"country_code"`
	Contact_details string `json:"contact_details"`
}

type UpdateCustomerResponse struct {
}

type RemoveCustomerRequest struct {
}

type RemoveCustomerResponse struct {
}

type DetailCustomerRequest struct {
	Customer_id int64 `json:"customer_id"`
}

type DetailCustomerResponse struct {
	Customer_id     int64  `json:"customer_id"`
	Customer_name   string `json:"customer_name"`
	Country_code    string `json:"country_code"`
	Contact_details string `json:"contact_details"`
}

type CreateEmployeeRequest struct {
	Employee_photo  string `json:"employee_photo"`
	Employee_name   string `json:"employee_name"`
	Contact_details string `json:"contact_details"`
	Company_id      int64  `json:"company_id"`
}

type CreateEmployeeResponse struct {
	Employee_id int64 `json:"employee_id"`
}

type UpdateEmployeeRequest struct {
	Employee_id     int64  `json:"employee_id"`
	Employee_photo  string `json:"employee_photo"`
	Employee_name   string `json:"employee_name"`
	Contact_details string `json:"contact_details"`
	Company_id      int64  `json:"company_id"`
}

type UpdateEmployeeResponse struct {
}

type RemoveEmployeeRequest struct {
	Employee_id int64 `json:"employee_id"`
}

type RemoveEmployeeResponse struct {
}

type DetailEmployeeRequest struct {
	Employee_id int64 `json:"employee_id"`
}

type DetailEmployeeResponse struct {
	Employee_id     int64  `json:"employee_id"`
	Employee_photo  string `json:"employee_photo"`
	Employee_name   string `json:"employee_name"`
	Contact_details string `json:"contact_details"`
	Company_id      int64  `json:"company_id"`
	Link_code       string `json:"link_code"`
	Work_status     int    `json:"work_status"`
	Order_id        int64  `json:"order_id"`
}

type ListEmployeeRequest struct {
}

type ListEmployeeResponse struct {
	Items []DetailEmployeeResponse `json:"items"`
}

type CreateOrderRequest struct {
	Customer_id        int64  `json:"customer_id"`
	Company_id         int64  `json:"company_id"`
	Address_id         int64  `json:"address_id"`
	Design_id          int64  `json:"desgin_id"`
	Deposite_payment   int64  `json:"deposite_payment"`
	Deposite_date      string `json:"deposite_date"`
	Final_payment      int64  `json:"final_payment,optional"`
	Final_payment_date string `json:"final_payment_date,optional"`
	Order_description  string `json:"order_description,optional"`
	Post_date          string `json:"post_date"`
	Reserve_date       string `json:"reserve_date"`
	Finish_date        string `json:"finish_date,optional"`
}

type CreateOrderResponse struct {
	Order_id int64 `json:"order_id"`
}

type UpdateOrderRequest struct {
	Order_id           int64  `json:"order_id"`
	Customer_id        int64  `json:"customer_id"`
	Final_payment      int64  `json:"final_payment,optional"`
	Final_payment_date string `json:"final_payment_date,optional"`
	Order_description  string `json:"order_description"`
	Reserve_date       string `json:"reserve_date"`
	Finish_date        string `json:"finish_date,optional"`
}

type UpdateOrderResponse struct {
}

type RemoveOrderRequest struct {
	Order_id int64 `json:"order_id"`
}

type RemoveOrderResponse struct {
}

type DetailOrderRequest struct {
	Order_id int64 `json:"order_id"`
}

type DetailOrderResponse struct {
	Order_id              int64   `json:"order_id"`
	Customer_id           int64   `json:"customer_id"`
	Company_id            int64   `json:"company_id"`
	Address_id            int64   `json:"address_id"`
	Design_id             int64   `json:"desgin_id"`
	Deposite_payment      int64   `json:"deposite_payment"`
	Deposite_amount       float64 `json:"deposite_amount"`
	Current_deposite_rate int   `json:"current_deposite_rate"`
	Deposite_date         string  `json:"deposite_date"`
	Final_payment         int64   `json:"final_payment"`
	Final_amount          float64 `json:"final_amount"`
	Final_payment_date    string  `json:"final_payment_date"`
	Total_fee             float64 `json:"total_fee"`
	Order_description     string  `json:"order_description"`
	Post_date             string  `json:"post_date"`
	Reserve_date          string  `json:"reserve_date"`
	Finish_date           string  `json:"finish_date"`
	Status                int     `json:"status"`
}

type ListOrderRequest struct {
}

type ListOrderResponse struct {
	Items []DetailOrderResponse `json:"items"`
}

type DetailServiceRequest struct {
	Service_id int64 `json:"service_id"`
}

type DetailServiceResponse struct {
	Service_id          int64  `json:"service_id"`
	Service_type        string `json:"service_type"`
	Service_description string `json:"service_description"`
}

type ListServiceRequest struct {
}

type ListServiceResponse struct {
	Items		[]DetailServiceResponse `json:"items"`
}

type CreateDesignRequest struct {
	Company_id int64   `json:"company_id"`
	Service_id int64   `json:"service_id"`
	Price      float64 `json:"price"`
	Comments   string  `json:"comments"`
}

type CreateDesignResponse struct {
	Design_id int64 `json:"design_id"`
}

type UpdateDesignRequest struct {
	Design_id  int64   `json:"design_id"`
	Company_id int64   `json:"company_id"`
	Service_id int64   `json:"service_id"`
	Price      float64 `json:"price"`
	Comments   string  `json:"comments"`
}

type UpdateDesignResponse struct {
}

type RemoveDesignRequest struct {
	Design_id int64 `json:"design_id"`
}

type RemoveDesignResponse struct {
}

type DetailDesignRequest struct {
	Design_id int64 `json:"design_id"`
}

type DetailDesignResponse struct {
	Design_id  int64   `json:"design_id"`
	Company_id int64   `json:"company_id"`
	Service_id int64   `json:"service_id"`
	Price      float64 `json:"price"`
	Comments   string  `json:"comments"`
}

type ListDesignRequest struct {
	Service_id int64 `json:"service_id,optional"`
}

type ListDesignResponse struct {
	Items []DetailDesignResponse `json:"items"`
}

type CreatePaymentRequest struct {
	Card_number   string `json:"card_number"`
	Holder_name   string `json:"holder_name"`
	Expiry_time   string `json:"expiry_time"`
	Security_code string `json:"security_code"`
}

type CreatePaymentResponse struct {
	Payment_id int64 `json:"payment_id"`
}

type UpdatePaymentRequest struct {
	Payment_id    int64  `json:"payment_id"`
	Card_number   string `json:"card_number"`
	Holder_name   string `json:"holder_name"`
	Expiry_time   string `json:"expiry_time"`
	Security_code string `json:"security_code"`
}

type UpdatePaymentResponse struct {
}

type RemovePaymentRequest struct {
	Payment_id int64 `json:"payment_id"`
}

type RemovePaymentResponse struct {
}

type DetailPaymentRequest struct {
	Payment_id int64 `json:"payment_id"`
}

type DetailPaymentResponse struct {
	Payment_id    int64  `json:"payment_id"`
	Card_number   string `json:"card_number"`
	Holder_name   string `json:"holder_name"`
	Expiry_time   string `json:"expiry_time"`
	Security_code string `json:"security_code"`
}

type ListPaymentRequest struct {
}

type ListPaymentResponse struct {
	Items []DetailPaymentResponse `json:"items"`
}

type CreateOperationRequest struct {
	Employee_id int64  `json:"employee_id"`
	Order_id    int64  `json:"order_id"`
	Operation   int64  `json:"operation"`
	Issue_date  string `json:"issue_date"`
}

type CreateOperationResponse struct {
	Operation_id int64 `json:"operation_id"`
}

type DetailOperationRequest struct {
	Operation_id int64 `json:"operation_id"`
}

type DetailOperationResponse struct {
	Operation_id int64  `json:"operation_id"`
	Employee_id  int64  `json:"employee_id"`
	Order_id     int64  `json:"order_id"`
	Operation    int64  `json:"operation"`
	Issue_date   string `json:"issue_date"`
}

type ListOperationRequest struct {
	Employee_id  int64  `json:"employee_id"`
}

type ListOperationResponse struct {
	Items []DetailOperationResponse `json:"items"`
}

type CreateEmployeeServiceRequest struct {
	Service_id int64 `json:"service_id"`
}

type CreateEmployeeServiceResponse struct {
}

type RemoveEmployeeServiceRequest struct {
	Employee_id int64 `json:"employee_id"`
	Service_id int64  `json:"service_id"`
}

type RemoveEmployeeServiceResponse struct {
}

type ListEmployeeServiceRequest struct {
	Employee_id int64 `json:"employee_id"`
}

type ListEmployeeServiceResponse struct {
	Items []DetailServiceResponse `json:"items"`
}
