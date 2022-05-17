// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"fmt"
	"net/http"
	"strings"

	"cleaningservice/service/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/verify_code",
				Handler: VerifyCodeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login_employee",
				Handler: LoginEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login_company",
				Handler: LoginCompanyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/user/login_customer",
				Handler: LoginCustomerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/create",
				Handler: CreateOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/update",
				Handler: UpdateOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/cancel",
				Handler: CancelOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/pay",
				Handler: PayOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/detail",
				Handler: DetailOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/service/detail",
				Handler: DetailServiceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/service/list",
				Handler: ListServiceHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/address/create",
				Handler: CreateAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/address/update",
				Handler: UpdateAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/address/remove",
				Handler: RemoveAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/address/detail",
				Handler: DetailAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/address/list",
				Handler: ListAddressHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/company/create",
				Handler: CreateCompanyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/company/update",
				Handler: UpdateCompanyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/company/remove",
				Handler: RemoveCompanyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/company/detail",
				Handler: DetailCompanyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/company/list",
				Handler: ListCompanyHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/customer/create",
				Handler: CreateCustomerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/customer/update",
				Handler: UpdateCustomerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/customer/remove",
				Handler: RemoveCustomerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/customer/detail",
				Handler: DetailCustomerHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/employee/create",
				Handler: CreateEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/employee/update",
				Handler: UpdateEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/employee/remove",
				Handler: RemoveEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/employee/detail",
				Handler: DetailEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/employee/list",
				Handler: ListEmployeeHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/employee/photo",
				Handler: UploadEmployeePhotoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/finish",
				Handler: FinishOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/remove",
				Handler: RemoveOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/list",
				Handler: ListOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/payment/create",
				Handler: CreatePaymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/payment/update",
				Handler: UpdatePaymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/payment/remove",
				Handler: RemovePaymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/payment/detail",
				Handler: DetailPaymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/payment/list",
				Handler: ListPaymentHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/operation/create",
				Handler: CreateOperationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/operation/accept",
				Handler: AcceptOperationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/operation/decline",
				Handler: DeclineOperationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/operation/detail",
				Handler: DetailOperationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/operation/list",
				Handler: ListOperationHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/employee_service/create",
				Handler: CreateEmployeeServiceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/employee_service/remove",
				Handler: RemoveEmployeeServiceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/employee_service/list",
				Handler: ListEmployeeServiceHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	// Static files resgistration
	dirlevel := []string{":1", ":2", ":3", ":4"}
	// CSS files
	patern := "/css/"
	dirpath := "./static/css/"
	for i := 1; i < 2; i++ {
		path := patern + strings.Join(dirlevel[:i], "/")
		server.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: StaticHandler(patern,dirpath),
			},
		)
		fmt.Printf("registered css file %s  %s\n", path,dirpath)
	}

	// Images files
	patern = "/images/"
	dirpath = "./static/images/"
	for i := 1; i < 3; i++ {
		path := patern + strings.Join(dirlevel[:i], "/")
		server.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: StaticHandler(patern,dirpath),
			},
		)
		fmt.Printf("registered images file %s  %s\n", path,dirpath)
	}

	// JS files
	patern = "/js/"
	dirpath = "./static/js/"
	for i := 1; i < 2; i++ {
		path := patern + strings.Join(dirlevel[:i], "/")
		server.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: StaticHandler(patern,dirpath),
			},
		)
		fmt.Printf("registered JS file %s  %s\n", path,dirpath)
	}

	// SCSS files
	patern = "/scss/"
	dirpath = "./static/scss/"
	for i := 1; i < 2; i++ {
		path := patern + strings.Join(dirlevel[:i], "/")
		server.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: StaticHandler(patern,dirpath),
			},
		)
		fmt.Printf("registered SCSS file %s  %s\n", path,dirpath)
	}

	// Vendor files
	patern = "/vendor/"
	dirpath = "./static/vendor/"
	for i := 1; i < 2; i++ {
		path := patern + strings.Join(dirlevel[:i], "/")
		server.AddRoute(
			rest.Route{
				Method:  http.MethodGet,
				Path:    path,
				Handler: StaticHandler(patern,dirpath),
			},
		)
		fmt.Printf("registered Vendor file %s  %s\n", path,dirpath)
	}
}
