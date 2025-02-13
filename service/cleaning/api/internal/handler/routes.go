// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"fmt"
	"net/http"
	"strings"

	"cleaningservice/service/cleaning/api/internal/svc"

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
				Path:    "/api/user/login_contractor",
				Handler: LoginContractorHandler(serverCtx),
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
				Path:    "/api/service/enquire",
				Handler: EnquireServiceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/service/list",
				Handler: ListServiceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/category/detail",
				Handler: DetailCategoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/category/list",
				Handler: ListCategoryHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/user/refresh_token",
				Handler: RefreshTokenHandler(serverCtx),
			},
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
				Path:    "/api/contractor/create",
				Handler: CreateContractorHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/contractor/update",
				Handler: UpdateContractorHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/contractor/remove",
				Handler: RemoveContractorHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/contractor/detail",
				Handler: DetailContractorHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/contractor/list",
				Handler: ListContractorHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/contractor/photo",
				Handler: UploadContractPhotoHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/contractor/schedule",
				Handler: GetContractorScheduleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/contractor/history",
				Handler: GetContractorHistoryHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/start",
				Handler: StartOrderHandler(serverCtx),
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
				Path:    "/api/order/surcharge",
				Handler: SurchargeOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/getDetails",
				Handler: GetOrderDetailsHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/addService",
				Handler: AddOrderServiceHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/list",
				Handler: ListOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/confirm",
				Handler: ConfirmOrderHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/order/recommend",
				Handler: RecommendOrderHandler(serverCtx),
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
				Path:    "/api/operation/transfer",
				Handler: TransferOperationHandler(serverCtx),
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
				Path:    "/api/subscribegroup/join",
				Handler: JoinSubscribeGroupHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/subscribegroup/leave",
				Handler: LeaveSubscribeGroupHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
	)

	// Static files resgistration
	var patern string
	var dirpath string
	dirlevel := []string{":1", ":2", ":3", ":4"}
	// // CSS files
	// patern = "/css/"
	// dirpath = "./static/css/"
	// for i := 1; i < 2; i++ {
	// 	path := patern + strings.Join(dirlevel[:i], "/")
	// 	server.AddRoute(
	// 		rest.Route{
	// 			Method:  http.MethodGet,
	// 			Path:    path,
	// 			Handler: StaticHandler(patern,dirpath),
	// 		},
	// 	)
	// 	fmt.Printf("registered css file %s  %s\n", path,dirpath)
	// }
	
	// Images files
	patern = "/images/"
	dirpath = "../../../static/images/"
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
	
	// // JS files
	// patern = "/js/"
	// dirpath = "./static/js/"
	// for i := 1; i < 2; i++ {
	// 	path := patern + strings.Join(dirlevel[:i], "/")
	// 	server.AddRoute(
	// 		rest.Route{
	// 			Method:  http.MethodGet,
	// 			Path:    path,
	// 			Handler: StaticHandler(patern,dirpath),
	// 		},
	// 	)
	// 	fmt.Printf("registered JS file %s  %s\n", path,dirpath)
	// }
	
	// // SCSS files
	// patern = "/scss/"
	// dirpath = "./static/scss/"
	// for i := 1; i < 2; i++ {
	// 	path := patern + strings.Join(dirlevel[:i], "/")
	// 	server.AddRoute(
	// 		rest.Route{
	// 			Method:  http.MethodGet,
	// 			Path:    path,
	// 			Handler: StaticHandler(patern,dirpath),
	// 		},
	// 	)
	// 	fmt.Printf("registered SCSS file %s  %s\n", path,dirpath)
	// }
	
	// // Vendor files
	// patern = "/vendor/"
	// dirpath = "./static/vendor/"
	// for i := 1; i < 2; i++ {
	// 	path := patern + strings.Join(dirlevel[:i], "/")
	// 	server.AddRoute(
	// 		rest.Route{
	// 			Method:  http.MethodGet,
	// 			Path:    path,
	// 			Handler: StaticHandler(patern,dirpath),
	// 		},
	// 	)
	// 	fmt.Printf("registered Vendor file %s  %s\n", path,dirpath)
	// }
}
