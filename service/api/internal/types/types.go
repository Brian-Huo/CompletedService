// Code generated by goctl. DO NOT EDIT.
package types

type InitSystemRequest struct {
}

type InitSystemResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type VerifyCodeRequest struct {
	Contact_details string `json:"contact_details"`
}

type VerifyCodeResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LoginContractorRequest struct {
	Contact_details string `json:"contact_details"`
	VerifyCode      string `json:"verify_code,optional"`
	LinkCode        string `json:"link_code"`
}

type LoginContractorResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	AccessToken string `json:"access_token,optional"`
}

type LoginCompanyRequest struct {
	Contact_details string `json:"contact_details"`
	VerifyCode      string `json:"verify_code"`
}

type LoginCompanyResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	AccessToken string `json:"access_token,optional"`
}

type LoginCustomerRequest struct {
	Contact_details string `json:"contact_details"`
	VerifyCode      string `json:"verify_code"`
}

type LoginCustomerResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	AccessToken string `json:"access_token,optional"`
}

type CreateAddressRequest struct {
	Street     string  `json:"street"`
	Suburb     string  `json:"suburb"`
	Postcode   string  `json:"postcode"`
	City       string  `json:"city"`
	State_code string  `json:"state_code"`
	Country    string  `json:"country"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Formatted  string  `json:"formatted"`
}

type CreateAddressResponse struct {
	Address_id int64 `json:"address_id"`
}

type UpdateAddressRequest struct {
	Address_id int64   `json:"address_id"`
	Street     string  `json:"street"`
	Suburb     string  `json:"suburb"`
	Postcode   string  `json:"postcode"`
	City       string  `json:"city"`
	State_code string  `json:"state_code"`
	Country    string  `json:"country"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Formatted  string  `json:"formatted"`
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
	Address_id int64   `json:"address_id"`
	Street     string  `json:"street"`
	Suburb     string  `json:"suburb"`
	Postcode   string  `json:"postcode"`
	City       string  `json:"city"`
	State_code string  `json:"state_code"`
	Country    string  `json:"country"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Formatted  string  `json:"formatted"`
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

type CreateContractorRequest struct {
	Contractor_photo string `json:"contractor_photo"`
	Contractor_name  string `json:"contractor_name"`
	Contact_details  string `json:"contact_details"`
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
	Link_code        string               `json:"link_code,optional"`
	Work_status      int                  `json:"work_status"`
}

type UpdateContractorResponse struct {
}

type RemoveContractorRequest struct {
	Contractor_id int64 `json:"contractor_id,optional"`
}

type RemoveContractorResponse struct {
}

type DetailContractorRequest struct {
	Contractor_id int64 `json:"contractor_id,optional"`
}

type DetailContractorResponse struct {
	Contractor_id    int64                 `json:"contractor_id"`
	Contractor_photo string                `json:"contractor_photo"`
	Contractor_name  string                `json:"contractor_name"`
	Contractor_type  string                `json:"contractor_type"`
	Contact_details  string                `json:"contact_details"`
	Address_info     DetailAddressResponse `json:"address_info"`
	Finance_id       int64                 `json:"finance_id"`
	Link_code        string                `json:"link_code"`
	Work_status      int                   `json:"work_status"`
	Order_id         int64                 `json:"order_id"`
	Category_list    []int64               `json:"category_list"`
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
	Customer_info     CreateCustomerRequest  `json:"customer_info"`
	Address_info      CreateAddressRequest   `json:"address_info"`
	Category_id       int64                  `json:"category_id"`
	Service_list      []SelectServiceRequest `json:"service_list"`
	Deposite_info     CreatePaymentRequest   `json:"deposite_info"`
	Order_description string                 `json:"order_description,optional"`
	Reserve_date      string                 `json:"reserve_date"`
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
	Order_id        int64  `json:"order_id"`
	Contact_details string `json:"contact_details"`
}

type DetailOrderResponse struct {
	Order_id              int64                    `json:"order_id"`
	Customer_info         DetailCustomerResponse   `json:"customer_info"`
	Address_info          DetailAddressResponse    `json:"address_info"`
	Contractor_info       DetailContractorResponse `json:"contractor_info"`
	Finance_id            int64                    `json:"finance_id"`
	Category              DetailCategoryResponse   `json:"category"`
	Service_list          string                   `json:"service_list"`
	Deposite_payment      int64                    `json:"deposite_payment"`
	Deposite_amount       float64                  `json:"deposite_amount"`
	Deposite_date         string                   `json:"deposite_date"`
	Final_payment         int64                    `json:"final_payment"`
	Final_amount          float64                  `json:"final_amount"`
	Final_payment_date    string                   `json:"final_payment_date"`
	Current_deposite_rate int                      `json:"current_deposite_rate"`
	Gst_amount            float64                  `json:"gst_amount"`
	Total_fee             float64                  `json:"total_fee"`
	Order_description     string                   `json:"order_description"`
	Post_date             string                   `json:"post_date"`
	Reserve_date          string                   `json:"reserve_date"`
	Finish_date           string                   `json:"finish_date"`
	Status                int                      `json:"status"`
	Urgent_flag           int                      `json:"urgent_flag"`
}

type GetOrderDetailsRequest struct {
	Order_id int64 `json:"order_id"`
}

type GetOrderDetailsResponse struct {
	Order_id              int64                    `json:"order_id"`
	Customer_info         DetailCustomerResponse   `json:"customer_info"`
	Address_info          DetailAddressResponse    `json:"address_info"`
	Contractor_info       DetailContractorResponse `json:"contractor_info"`
	Finance_id            int64                    `json:"finance_id"`
	Category              DetailCategoryResponse   `json:"category"`
	Service_list          string                   `json:"service_list"`
	Deposite_payment      int64                    `json:"deposite_payment"`
	Deposite_amount       float64                  `json:"deposite_amount"`
	Deposite_date         string                   `json:"deposite_date"`
	Final_payment         int64                    `json:"final_payment"`
	Final_amount          float64                  `json:"final_amount"`
	Final_payment_date    string                   `json:"final_payment_date"`
	Current_deposite_rate int                      `json:"current_deposite_rate"`
	Gst_amount            float64                  `json:"gst_amount"`
	Total_fee             float64                  `json:"total_fee"`
	Order_description     string                   `json:"order_description"`
	Post_date             string                   `json:"post_date"`
	Reserve_date          string                   `json:"reserve_date"`
	Finish_date           string                   `json:"finish_date"`
	Status                int                      `json:"status"`
	Urgent_flag           int                      `json:"urgent_flag"`
}

type ListOrderRequest struct {
}

type ListOrderResponse struct {
	Items []DetailOrderResponse `json:"items"`
}

type RecommendOrderRequest struct {
}

type RecommendOrderResponse struct {
	Items []DetailOrderResponse `json:"items"`
}

type SelectServiceRequest struct {
	Service_id       int64 `json:"service_id"`
	Service_quantity int   `json:"service_quantity"`
}

type SelectServiceResponse struct {
}

type DetailServiceRequest struct {
	Service_id int64 `json:"service_id"`
}

type DetailServiceResponse struct {
	Service_id          int64                  `json:"service_id"`
	Service_type        DetailCategoryResponse `json:"service_type"`
	Service_scope       string                 `json:"service_scope"`
	Service_name        string                 `json:"service_name"`
	Service_photo       string                 `json:"service_photo"`
	Service_description string                 `json:"service_description"`
	Service_price       float64                `json:"service_price"`
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

type AcceptOperationRequest struct {
	Order_id int64 `json:"order_id"`
}

type AcceptOperationResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DeclineOperationRequest struct {
	Order_id int64 `json:"order_id"`
}

type DeclineOperationResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type TransferOperationRequest struct {
	Order_id int64 `json:"order_id"`
}

type TransferOperationResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
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

type UploadContractorPhotoRequest struct {
	Contractor_id    int64  `json:"contractor_id,optional"`
	Contractor_photo string `json:"contractor_photo"`
}

type UploadContractorPhotoResponse struct {
	Contractor_photo string `json:"contractor_photo"`
}

type GetContractorScheduleRequest struct {
}

type GetContractorScheduleResponse struct {
	Items []DetailOrderResponse `json:"items"`
}

type DetailCategoryRequest struct {
	Category_id int64 `json:"category_id"`
}

type DetailCategoryResponse struct {
	Category_id          int64  `json:"category_id"`
	Category_name        string `json:"category_name"`
	Category_description string `json:"category_description"`
}

type ListCategoryRequest struct {
}

type ListCategoryResponse struct {
	Items []DetailCategoryResponse `json:"items"`
}

type JoinSubscribeGroupRequest struct {
	Category_list []int64 `json:"category_list"`
	Location      string  `json:"location,optional"`
}

type JoinSubscribeGroupResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type LeaveSubscribeGroupRequest struct {
	Category_list []int64 `json:"category_list"`
	Location      string  `json:"location,optional"`
}

type LeaveSubscribeGroupResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}
