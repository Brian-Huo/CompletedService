// Code generated by goctl. DO NOT EDIT.
package types

type VerifyCodeRequest struct {
	Contact_details string `json:"contact_details"`
}

type VerifyCodeResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type LoginEmployeeRequest struct {
	Contact_details string `json:"contact_details"`
	VerifyCode      string `json:"verify_code,optional"`
	LinkCode        string `json:"link_code"`
}

type LoginEmployeeResponse struct {
	Code        int    `json:"code"`
	Message     string `json:"message"`
	AccessToken string `json:"access_token,optional"`
}

type LoginCompanyRequest struct {
	Contact_details string `json:"contact_details"`
	VerifyCode      string `json:"verify_code"`
}

type LoginCompanyResponse struct {
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
	Street     string `json:"street"`
	Suburb     string `json:"suburb"`
	Postcode   string `json:"postcode"`
	State_code string `json:"state_code"`
	Country    string `json:"country"`
}

type CreateAddressResponse struct {
	Address_id int64 `json:"address_id"`
}

type UpdateAddressRequest struct {
	Address_id int64  `json:"address_id"`
	Street     string `json:"street"`
	Suburb     string `json:"suburb"`
	Postcode   string `json:"postcode"`
	State_code string `json:"state_code"`
	Country    string `json:"country"`
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
	Address_id int64  `json:"address_id"`
	Street     string `json:"street"`
	Suburb     string `json:"suburb"`
	Postcode   string `json:"postcode"`
	State_code string `json:"state_code"`
	Country    string `json:"country"`
}

type ListAddressRequest struct {
}

type ListAddressResponse struct {
	Items []DetailAddressResponse `json:"items"`
}

type CreateCompanyRequest struct {
	Company_name    string               `json:"company_name"`
	Payment_info    CreatePaymentRequest `json:"payment_info"`
	Director_name   string               `json:"director_name"`
	Contact_details string               `json:"contact_details"`
	Address_info    CreateAddressRequest `json:"address_info"`
	Deposite_rate   int                  `json:"deposite_rate"`
}

type CreateCompanyResponse struct {
	Company_id int64 `json:"company_id"`
}

type UpdateCompanyRequest struct {
	Company_name    string `json:"company_name"`
	Director_name   string `json:"director_name"`
	Contact_details string `json:"contact_details"`
}

type UpdateCompanyResponse struct {
}

type RemoveCompanyRequest struct {
	Company_id int64 `json:"company_id"`
}

type RemoveCompanyResponse struct {
}

type DetailCompanyRequest struct {
	Company_id int64 `json:"company_id,optional"`
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
	Customer_id int64 `json:"customer_id"`
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
}

type CreateEmployeeResponse struct {
	Employee_id int64 `json:"employee_id"`
}

type UpdateEmployeeRequest struct {
	Employee_id     int64   `json:"employee_id"`
	Employee_photo  string  `json:"employee_photo"`
	Employee_name   string  `json:"employee_name"`
	Contact_details string  `json:"contact_details"`
	Link_code       string  `json:"link_code"`
	Work_status     int     `json:"work_status"`
	New_services    []int64 `json:"new_services"`
	Remove_services []int64 `json:"remove_services"`
}

type UpdateEmployeeResponse struct {
}

type RemoveEmployeeRequest struct {
	Employee_id int64 `json:"employee_id,optional"`
}

type RemoveEmployeeResponse struct {
}

type DetailEmployeeRequest struct {
	Employee_id int64 `json:"employee_id,optional"`
}

type DetailEmployeeResponse struct {
	Employee_id      int64                       `json:"employee_id"`
	Employee_photo   string                      `json:"employee_photo"`
	Employee_name    string                      `json:"employee_name"`
	Contact_details  string                      `json:"contact_details"`
	Company_id       int64                       `json:"company_id"`
	Link_code        string                      `json:"link_code"`
	Work_status      int                         `json:"work_status"`
	Order_id         int64                       `json:"order_id"`
	Employee_service ListEmployeeServiceResponse `json:"employee_service"`
}

type ListEmployeeRequest struct {
}

type ListEmployeeResponse struct {
	Items []DetailEmployeeResponse `json:"items"`
}

type CreateContractorRequest struct {
	Contractor_photo string               `json:"contractor_photo"`
	Contractor_name  string               `json:"contractor_name"`
	Contractor_type  string               `json:"contractor_type"`
	Contact_details  string               `json:"contact_details"`
	Address_info     CreateAddressRequest `json:"address_info,optional"`
}

type CreateContractorResponse struct {
	Contractor_id int64 `json:"contractor_id"`
}

type UpdateContractorRequest struct {
	Contractor_id    int64                `json:"contractor_id"`
	Contractor_photo string               `json:"contractor_photo"`
	Contractor_name  string               `json:"contractor_name"`
	Contractor_type  string               `json:"contractor_type"`
	Contact_details  string               `json:"contact_details"`
	Address_info     UpdateAddressRequest `json:"address_info,optional"`
	Link_code        string               `json:"link_code"`
	Work_status      int                  `json:"work_status"`
	New_services     []int64              `json:"new_services"`
	Remove_services  []int64              `json:"remove_services"`
}

type UpdateContractorResponse struct {
}

type RemoveContractorRequest struct {
	Contractor_id int64 `json:"Contractor_id,optional"`
}

type RemoveContractorResponse struct {
}

type DetailContractorRequest struct {
	Contractor_id int64 `json:"Contractor_id,optional"`
}

type DetailContractorResponse struct {
	Contractor_id      int64                         `json:"contractor_id"`
	Contractor_photo   string                        `json:"contractor_photo"`
	Contractor_name    string                        `json:"contractor_name"`
	Contractor_type    string                        `json:"contractor_type"`
	Contact_details    string                        `json:"contact_details"`
	Address_info       DetailAddressResponse         `json:"address_info"`
	Finance_id         int64                         `json:"finance_id"`
	Link_code          string                        `json:"link_code"`
	Work_status        int                           `json:"work_status"`
	Order_id           int64                         `json:"order_id"`
	Contractor_service ListContractorServiceResponse `json:"contractor_service"`
}

type ListContractorRequest struct {
}

type ListContractorResponse struct {
	Items []DetailContractorResponse `json:"items"`
}

type CreateContractorServiceRequest struct {
	Contractor_id int64 `json:"contractor_id,optional"`
	Service_id    int64 `json:"service_id"`
}

type CreateContractorServiceResponse struct {
}

type RemoveContractorServiceRequest struct {
	Contractor_id int64 `json:"contractor_id,optional"`
	Service_id    int64 `json:"service_id"`
}

type RemoveContractorServiceResponse struct {
}

type ListContractorServiceRequest struct {
	Contractor_id int64 `json:"contractor_id"`
}

type ListContractorServiceResponse struct {
	Items []DetailServiceResponse `json:"items"`
}

type CreateOrderRequest struct {
	Customer_info     CreateCustomerRequest `json:"customer_info"`
	Address_info      CreateAddressRequest  `json:"address_info"`
	Service_list      []int64               `json:"service_list"`
	Deposite_info     CreatePaymentRequest  `json:"deposite_info"`
	Order_description string                `json:"order_description,optional"`
	Reserve_date      string                `json:"reserve_date"`
}

type CreateOrderResponse struct {
	Order_id int64 `json:"order_id"`
}

type UpdateOrderRequest struct {
	Order_id          int64                 `json:"order_id"`
	Customer_info     UpdateCustomerRequest `json:"customer_info"`
	Address_info      UpdateAddressRequest  `json:"address_info"`
	Order_description string                `json:"order_description"`
	Reserve_date      string                `json:"reserve_date"`
}

type UpdateOrderResponse struct {
}

type CancelOrderRequest struct {
	Order_id int64 `json:"order_id"`
}

type CancelOrderResponse struct {
}

type FinishOrderRequest struct {
	Order_id int64 `json:"order_id"`
}

type FinishOrderResponse struct {
}

type PayOrderRequest struct {
	Order_id   int64                `json:"order_id"`
	Final_info CreatePaymentRequest `json:"final_info"`
}

type PayOrderResponse struct {
}

type RemoveOrderRequest struct {
	Order_id int64 `json:"order_id"`
}

type RemoveOrderResponse struct {
}

type DetailOrderRequest struct {
	Order_id      int64  `json:"order_id"`
	Customer_name string `json:"customer_name"`
}

type DetailOrderResponse struct {
	Order_id              int64                  `json:"order_id"`
	Customer_info         DetailCustomerResponse `json:"customer_info"`
	Address_info          DetailAddressResponse  `json:"address_info"`
	Employee_info         DetailEmployeeResponse `json:"employee_info"`
	Company_id            int64                  `json:"company_id"`
	Service_list          string                 `json:"service_list"`
	Deposite_payment      int64                  `json:"deposite_payment"`
	Deposite_amount       float64                `json:"deposite_amount"`
	Current_deposite_rate int                    `json:"current_deposite_rate"`
	Deposite_date         string                 `json:"deposite_date"`
	Final_payment         int64                  `json:"final_payment"`
	Final_amount          float64                `json:"final_amount"`
	Final_payment_date    string                 `json:"final_payment_date"`
	Gst_amount            float64                `json:"gst_amount"`
	Total_fee             float64                `json:"total_fee"`
	Order_description     string                 `json:"order_description"`
	Post_date             string                 `json:"post_date"`
	Reserve_date          string                 `json:"reserve_date"`
	Finish_date           string                 `json:"finish_date"`
	Status                int                    `json:"status"`
}

type GetOrderDetailsRequest struct {
	Order_id int64 `json:"order_id"`
}

type GetOrderDetailsResponse struct {
	Order_id              int64                  `json:"order_id"`
	Customer_info         DetailCustomerResponse `json:"customer_info"`
	Address_info          DetailAddressResponse  `json:"address_info"`
	Employee_info         DetailEmployeeResponse `json:"employee_info"`
	Company_id            int64                  `json:"company_id"`
	Service_list          string                 `json:"service_list"`
	Deposite_payment      int64                  `json:"deposite_payment"`
	Deposite_amount       float64                `json:"deposite_amount"`
	Current_deposite_rate int                    `json:"current_deposite_rate"`
	Deposite_date         string                 `json:"deposite_date"`
	Final_payment         int64                  `json:"final_payment"`
	Final_amount          float64                `json:"final_amount"`
	Final_payment_date    string                 `json:"final_payment_date"`
	Gst_amount            float64                `json:"gst_amount"`
	Total_fee             float64                `json:"total_fee"`
	Order_description     string                 `json:"order_description"`
	Post_date             string                 `json:"post_date"`
	Reserve_date          string                 `json:"reserve_date"`
	Finish_date           string                 `json:"finish_date"`
	Status                int                    `json:"status"`
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
	Service_id          int64   `json:"service_id"`
	Service_type        string  `json:"service_type"`
	Service_scope       string  `json:"service_scope"`
	Service_name        string  `json:"service_name"`
	Service_description string  `json:"service_description"`
	Service_price       float64 `json:"service_price"`
}

type ListServiceRequest struct {
}

type ListServiceResponse struct {
	Items []DetailServiceResponse `json:"items"`
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
	Order_id  int64 `json:"order_id"`
	Operation int64 `json:"operation"`
}

type CreateOperationResponse struct {
	Operation_id int64 `json:"operation_id"`
}

type AcceptOperationRequest struct {
	Order_id int64 `json:"order_id"`
}

type AcceptOperationResponse struct {
	Operation_id int64 `json:"operation_id"`
}

type DeclineOperationRequest struct {
	Order_id int64 `json:"order_id"`
}

type DeclineOperationResponse struct {
	Operation_id int64 `json:"operation_id"`
}

type DetailOperationRequest struct {
	Operation_id int64 `json:"operation_id"`
}

type DetailOperationResponse struct {
	Operation_id  int64  `json:"operation_id"`
	Contractor_id int64  `json:"contractor_id"`
	Order_id      int64  `json:"order_id"`
	Operation     int64  `json:"operation"`
	Issue_date    string `json:"issue_date"`
}

type ListOperationRequest struct {
	Contractor_id int64 `json:"contractor_id"`
}

type ListOperationResponse struct {
	Items []DetailOperationResponse `json:"items"`
}

type CreateEmployeeServiceRequest struct {
	Employee_id int64 `json:"employee_id,optional"`
	Service_id  int64 `json:"service_id"`
}

type CreateEmployeeServiceResponse struct {
}

type RemoveEmployeeServiceRequest struct {
	Employee_id int64 `json:"employee_id,optional"`
	Service_id  int64 `json:"service_id"`
}

type RemoveEmployeeServiceResponse struct {
}

type ListEmployeeServiceRequest struct {
	Employee_id int64 `json:"employee_id"`
}

type ListEmployeeServiceResponse struct {
	Items []DetailServiceResponse `json:"items"`
}

type UploadContractorPhotoRequest struct {
	Contractor_id    int64  `json:"contractor_id,optional"`
	Contractor_photo string `json:"contractor_photo"`
}

type UploadContractorPhotoResponse struct {
	Contractor_photo string `json:"contractor_photo"`
}
