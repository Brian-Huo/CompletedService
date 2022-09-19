// Code generated by goctl. DO NOT EDIT.
package types

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

type RefreshTokenRequest struct {
}

type RefreshTokenResponse struct {
	Code        int    `json:"code"`
	Msg         string `json:"msg"`
	AccessToken string `json:"access_token"`
}

type CreateAddressRequest struct {
	Street     string  `json:"street"`
	Suburb     string  `json:"suburb"`
	Postcode   string  `json:"postcode"`
	Property   string  `json:"property"`
	City       string  `json:"city"`
	State_code string  `json:"state_code"`
	State_name string  `json:"state_name"`
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
	Property   string  `json:"property"`
	City       string  `json:"city"`
	State_code string  `json:"state_code"`
	State_name string  `json:"state_name"`
	Lat        float64 `json:"lat"`
	Lng        float64 `json:"lng"`
	Formatted  string  `json:"formatted"`
}

type UpdateAddressResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RemoveAddressRequest struct {
	Address_id int64 `json:"address_id"`
}

type RemoveAddressResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DetailAddressRequest struct {
	Address_id int64 `json:"address_id"`
}

type DetailAddressResponse struct {
	Address_id int64   `json:"address_id"`
	Street     string  `json:"street"`
	Suburb     string  `json:"suburb"`
	Postcode   string  `json:"postcode"`
	Property   string  `json:"property"`
	City       string  `json:"city"`
	State_code string  `json:"state_code"`
	State_name string  `json:"state_name"`
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
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RemoveCompanyRequest struct {
	Company_id int64 `json:"company_id"`
}

type RemoveCompanyResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
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
	Customer_name  string `json:"customer_name"`
	Country_code   string `json:"country_code"`
	Customer_phone string `json:"customer_phone"`
	Customer_email string `json:"customer_email"`
}

type CreateCustomerResponse struct {
	Customer_id int64 `json:"customer_id"`
}

type UpdateCustomerRequest struct {
	Customer_id    int64  `json:"customer_id"`
	Customer_name  string `json:"customer_name"`
	Country_code   string `json:"country_code"`
	Customer_phone string `json:"customer_phone"`
	Customer_email string `json:"customer_email"`
}

type UpdateCustomerResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RemoveCustomerRequest struct {
	Customer_id int64 `json:"customer_id"`
}

type RemoveCustomerResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type DetailCustomerRequest struct {
	Customer_id int64 `json:"customer_id"`
}

type DetailCustomerResponse struct {
	Customer_id    int64  `json:"customer_id"`
	Customer_name  string `json:"customer_name"`
	Country_code   string `json:"country_code"`
	Customer_phone string `json:"customer_phone"`
	Customer_email string `json:"customer_email"`
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
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RemoveContractorRequest struct {
	Contractor_id int64 `json:"contractor_id,optional"`
}

type RemoveContractorResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
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

type CreateOrderRequest struct {
	Customer_info     CreateCustomerRequest    `json:"customer_info"`
	Address_info      CreateAddressRequest     `json:"address_info"`
	Category_id       int64                    `json:"category_id"`
	Base_items        SelectedServiceStructure `json:"base_items"`
	Additional_items  SelectedServiceList      `json:"additional_items"`
	Deposite_info     CreatePaymentRequest     `json:"deposite_info,optional"`
	Order_description string                   `json:"order_description,optional"`
	Reserve_date      string                   `json:"reserve_date"`
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
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type CancelOrderRequest struct {
	Order_id int64 `json:"order_id"`
}

type CancelOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type StartOrderRequest struct {
	Order_id int64   `json:"order_id"`
	Lat      float64 `json:"lat"`
	Lng      float64 `json:"lng"`
}

type StartOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type FinishOrderRequest struct {
	Order_id int64 `json:"order_id"`
}

type FinishOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type PayOrderRequest struct {
	Order_id   int64                `json:"order_id"`
	Pay_info   CreatePaymentRequest `json:"pay_info"`
	Pay_amount float64              `json:"pay_amount"`
}

type PayOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RemoveOrderRequest struct {
	Order_id int64 `json:"order_id"`
}

type RemoveOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type SurchargeOrderRequest struct {
	Order_id         int64   `json:"order_id"`
	Surcharge_item   string  `json:"surcharge_item"`
	Surcharge_rate   int     `json:"surcharge_rate,range=[1:3]"`
	Surcharge_amount float64 `json:"surcharge_amount,optional"`
}

type SurchargeOrderResponse struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data DetailOrderResponse `json:"data"`
}

type AddOrderServiceRequest struct {
	Order_id         int64               `json:"order_id"`
	Additional_items SelectedServiceList `json:"additional_items"`
}

type AddOrderServiceResponse struct {
	Code int                 `json:"code"`
	Msg  string              `json:"msg"`
	Data DetailOrderResponse `json:"data"`
}

type GetOrderDetailsRequest struct {
	Order_id int64 `json:"order_id"`
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
	Basic_items           SelectedServiceStructure `json:"basic_items"`
	Additional_items      SelectedServiceList      `json:"additional_items"`
	Order_description     string                   `json:"order_description"`
	Order_comments        string                   `json:"order_comments"`
	Current_deposite_rate int                      `json:"current_deposite_rate"`
	Deposite_amount       float64                  `json:"deposite_amount"`
	Final_amount          float64                  `json:"final_amount"`
	Item_amount           float64                  `json:"item_amount"`
	Gst_amount            float64                  `json:"gst_amount"`
	Surcharge_item        string                   `json:"surcharge_item"`
	Surcharge_rate        int                      `json:"surcharge_rate"`
	Surcharge_amount      float64                  `json:"surcharge_amount"`
	Total_amount          float64                  `json:"total_amount"`
	Balance_amount        float64                  `json:"balance_amount"`
	Post_date             string                   `json:"post_date"`
	Reserve_date          string                   `json:"reserve_date"`
	Finish_date           string                   `json:"finish_date"`
	Payment_date          string                   `json:"payment_date"`
	Status                int                      `json:"status"`
	Urgent_flag           int                      `json:"urgent_flag"`
}

type ListOrderRequest struct {
}

type ListOrderResponse struct {
	Items []DetailOrderResponse `json:"items"`
}

type ConfirmOrderRequest struct {
	Order_list []int64 `json:"order_list"`
}

type ConfirmOrderResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RecommendOrderRequest struct {
}

type RecommendOrderResponse struct {
	Items []DetailOrderResponse `json:"items"`
}

type SelectedServiceStructure struct {
	Service_id       int64   `json:"service_id"`
	Service_scope    string  `json:"service_scope,optional"`
	Service_name     string  `json:"service_name,optional"`
	Service_price    float64 `json:"service_price,optional"`
	Service_quantity int     `json:"service_quantity"`
}

type SelectedServiceList struct {
	Items []SelectedServiceStructure `json:"items"`
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

type EnquireServiceRequest struct {
	Category_id int64  `json:"category_id"`
	Postcode    string `json:"postcode"`
	Property    string `json:"property"`
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
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

type RemovePaymentRequest struct {
	Payment_id int64 `json:"payment_id"`
}

type RemovePaymentResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
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

type GetContractorHistoryRequest struct {
}

type GetContractorHistoryResponse struct {
	Items []DetailOrderResponse `json:"items"`
}

type DetailCategoryRequest struct {
	Category_id int64 `json:"category_id"`
}

type DetailCategoryResponse struct {
	Category_id          int64  `json:"category_id"`
	Category_addr        string `json:"category_addr"`
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
