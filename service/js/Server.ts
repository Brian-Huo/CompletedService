import * as components from "./ServerComponents"
export * from "./ServerComponents"

/**
 * @description 
 * @param req
 */
export function verifyCode(req: components.VerifyCodeRequest) {
	return webapi.post<components.VerifyCodeResponse>("/api/user/verify_code", req)
}

/**
 * @description 
 * @param req
 */
export function loginService(req: components.LoginServiceRequest) {
	return webapi.post<components.LoginServiceResponse>("/api/user/login_service", req)
}

/**
 * @description 
 * @param req
 */
export function loginCustomer(req: components.LoginCustomerRequest) {
	return webapi.post<components.LoginCustomerResponse>("/api/user/login_customer", req)
}

/**
 * @description 
 * @param req
 */
export function createAddress(req: components.CreateAddressRequest) {
	return webapi.post<components.CreateAddressResponse>("/api/address/create", req)
}

/**
 * @description 
 * @param req
 */
export function updateAddress(req: components.UpdateAddressRequest) {
	return webapi.post<components.UpdateAddressResponse>("/api/address/update", req)
}

/**
 * @description 
 * @param req
 */
export function removeAddress(req: components.RemoveAddressRequest) {
	return webapi.post<components.RemoveAddressResponse>("/api/address/remove", req)
}

/**
 * @description 
 * @param req
 */
export function detailAddress(req: components.DetailAddressRequest) {
	return webapi.post<components.DetailAddressResponse>("/api/address/detail", req)
}

/**
 * @description 
 */
export function listAddress() {
	return webapi.post<components.ListAddressResponse>("/api/address/list")
}

/**
 * @description 
 * @param req
 */
export function createCompany(req: components.CreateCompanyRequest) {
	return webapi.post<components.CreateCompanyResponse>("/api/company/create", req)
}

/**
 * @description 
 * @param req
 */
export function updateCompany(req: components.UpdateCompanyRequest) {
	return webapi.post<components.UpdateCompanyResponse>("/api/company/update", req)
}

/**
 * @description 
 * @param req
 */
export function removeCompany(req: components.RemoveCompanyRequest) {
	return webapi.post<components.RemoveCompanyResponse>("/api/company/remove", req)
}

/**
 * @description 
 * @param req
 */
export function detailCompany(req: components.DetailCompanyRequest) {
	return webapi.post<components.DetailCompanyResponse>("/api/company/detail", req)
}

/**
 * @description 
 */
export function listCompany() {
	return webapi.post<components.ListCompanyResponse>("/api/company/list")
}

/**
 * @description 
 * @param req
 */
export function createCustomer(req: components.CreateCustomerRequest) {
	return webapi.post<components.CreateCustomerResponse>("/api/customer/create", req)
}

/**
 * @description 
 * @param req
 */
export function updateCustomer(req: components.UpdateCustomerRequest) {
	return webapi.post<components.UpdateCustomerResponse>("/api/customer/update", req)
}

/**
 * @description 
 * @param req
 */
export function removeCustomer(req: components.RemoveCustomerRequest) {
	return webapi.post<components.RemoveCustomerResponse>("/api/customer/remove", req)
}

/**
 * @description 
 * @param req
 */
export function detailCustomer(req: components.DetailCustomerRequest) {
	return webapi.post<components.DetailCustomerResponse>("/api/customer/detail", req)
}

/**
 * @description 
 * @param req
 */
export function createEmployee(req: components.CreateEmployeeRequest) {
	return webapi.post<components.CreateEmployeeResponse>("/api/employee/create", req)
}

/**
 * @description 
 * @param req
 */
export function updateEmployee(req: components.UpdateEmployeeRequest) {
	return webapi.post<components.UpdateEmployeeResponse>("/api/employee/update", req)
}

/**
 * @description 
 * @param req
 */
export function removeEmployee(req: components.RemoveEmployeeRequest) {
	return webapi.post<components.RemoveEmployeeResponse>("/api/employee/remove", req)
}

/**
 * @description 
 * @param req
 */
export function detailEmployee(req: components.DetailEmployeeRequest) {
	return webapi.post<components.DetailEmployeeResponse>("/api/employee/detail", req)
}

/**
 * @description 
 */
export function listEmployee() {
	return webapi.post<components.ListEmployeeResponse>("/api/employee/list")
}

/**
 * @description 
 * @param req
 */
export function createOrder(req: components.CreateOrderRequest) {
	return webapi.post<components.CreateOrderResponse>("/api/order/create", req)
}

/**
 * @description 
 * @param req
 */
export function updateOrder(req: components.UpdateOrderRequest) {
	return webapi.post<components.UpdateOrderResponse>("/api/order/update", req)
}

/**
 * @description 
 * @param req
 */
export function removeOrder(req: components.RemoveOrderRequest) {
	return webapi.post<components.RemoveOrderResponse>("/api/order/remove", req)
}

/**
 * @description 
 * @param req
 */
export function detailOrder(req: components.DetailOrderRequest) {
	return webapi.post<components.DetailOrderResponse>("/api/order/detail", req)
}

/**
 * @description 
 */
export function listOrder() {
	return webapi.post<components.ListOrderResponse>("/api/order/list")
}

/**
 * @description 
 * @param req
 */
export function detailService(req: components.DetailServiceRequest) {
	return webapi.post<components.DetailServiceResponse>("/api/service/detail", req)
}

/**
 * @description 
 */
export function listService() {
	return webapi.post<components.ListServiceResponse>("/api/service/list")
}

/**
 * @description 
 * @param req
 */
export function createDesign(req: components.CreateDesignRequest) {
	return webapi.post<components.CreateDesignResponse>("/api/design/create", req)
}

/**
 * @description 
 * @param req
 */
export function updateDesign(req: components.UpdateDesignRequest) {
	return webapi.post<components.UpdateDesignResponse>("/api/design/update", req)
}

/**
 * @description 
 * @param req
 */
export function removeDesign(req: components.RemoveDesignRequest) {
	return webapi.post<components.RemoveDesignResponse>("/api/design/remove", req)
}

/**
 * @description 
 * @param req
 */
export function detailDesign(req: components.DetailDesignRequest) {
	return webapi.post<components.DetailDesignResponse>("/api/design/detail", req)
}

/**
 * @description 
 * @param req
 */
export function listDesign(req: components.ListDesignRequest) {
	return webapi.post<components.ListDesignResponse>("/api/design/list", req)
}

/**
 * @description 
 * @param req
 */
export function createPayment(req: components.CreatePaymentRequest) {
	return webapi.post<components.CreatePaymentResponse>("/api/payment/create", req)
}

/**
 * @description 
 * @param req
 */
export function updatePayment(req: components.UpdatePaymentRequest) {
	return webapi.post<components.UpdatePaymentResponse>("/api/payment/update", req)
}

/**
 * @description 
 * @param req
 */
export function removePayment(req: components.RemovePaymentRequest) {
	return webapi.post<components.RemovePaymentResponse>("/api/payment/remove", req)
}

/**
 * @description 
 * @param req
 */
export function detailPayment(req: components.DetailPaymentRequest) {
	return webapi.post<components.DetailPaymentResponse>("/api/payment/detail", req)
}

/**
 * @description 
 */
export function listPayment() {
	return webapi.post<components.ListPaymentResponse>("/api/payment/list")
}

/**
 * @description 
 * @param req
 */
export function createOperation(req: components.CreateOperationRequest) {
	return webapi.post<components.CreateOperationResponse>("/api/operation/create", req)
}

/**
 * @description 
 * @param req
 */
export function detailOperation(req: components.DetailOperationRequest) {
	return webapi.post<components.DetailOperationResponse>("/api/operation/detail", req)
}

/**
 * @description 
 */
export function listOperation() {
	return webapi.post<components.ListOperationResponse>("/api/operation/list")
}

/**
 * @description 
 * @param req
 */
export function createEmployeeService(req: components.CreateEmployeeServiceRequest) {
	return webapi.post<components.CreateEmployeeServiceResponse>("/api/employee_service/create", req)
}

/**
 * @description 
 * @param req
 */
export function removeEmployeeService(req: components.RemoveEmployeeServiceRequest) {
	return webapi.post<components.RemoveEmployeeServiceResponse>("/api/employee_service/remove", req)
}

/**
 * @description 
 * @param req
 */
export function listEmployeeService(req: components.ListEmployeeServiceRequest) {
	return webapi.post<components.ListEmployeeServiceResponse>("/api/employee_service/list", req)
}
