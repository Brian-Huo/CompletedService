package svc

import (
	"cleaningservice/service/cleaning/api/internal/config"
	"cleaningservice/service/cleaning/model/address"
	"cleaningservice/service/cleaning/model/broadcast"
	"cleaningservice/service/cleaning/model/category"
	"cleaningservice/service/cleaning/model/company"
	"cleaningservice/service/cleaning/model/contractor"
	"cleaningservice/service/cleaning/model/customer"
	"cleaningservice/service/cleaning/model/operation"
	"cleaningservice/service/cleaning/model/order"
	"cleaningservice/service/cleaning/model/orderdelay"
	"cleaningservice/service/cleaning/model/payment"
	"cleaningservice/service/cleaning/model/service"
	"cleaningservice/service/cleaning/model/subscriberecord"
	"cleaningservice/service/cleaning/model/subscription"
	"cleaningservice/service/email/rpc/email"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	// models dao
	BAddressModel         address.BAddressModel
	BBroadcastModel       broadcast.BBroadcastModel
	BCategoryModel        category.BCategoryModel
	BCompanyModel         company.BCompanyModel
	BCustomerModel        customer.BCustomerModel
	BContractorModel      contractor.BContractorModel
	BOperationModel       operation.BOperationModel
	BOrderModel           order.BOrderModel
	ROrderDelayModel      orderdelay.ROrderDelayModel
	BPaymentModel         payment.BPaymentModel
	BServiceModel         service.BServiceModel
	RSubscribeRecordModel subscriberecord.RSubscribeRecordModel
	BSubscriptionModel    subscription.BSubscriptionModel

	// rpc api
	EmailRpc email.Email
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config: c,

		// models dao
		BAddressModel:         address.NewBAddressModel(conn, c.CacheRedis),
		BBroadcastModel:       broadcast.NewBBroadcastModel(c.RedisConf),
		BCategoryModel:        category.NewBCategoryModel(conn, c.CacheRedis),
		BCompanyModel:         company.NewBCompanyModel(conn, c.CacheRedis),
		BCustomerModel:        customer.NewBCustomerModel(conn, c.CacheRedis),
		BContractorModel:      contractor.NewBContractorModel(conn, c.CacheRedis),
		BOrderModel:           order.NewBOrderModel(conn, c.CacheRedis),
		ROrderDelayModel:      orderdelay.NewROrderDelayModel(c.RedisConf),
		BOperationModel:       operation.NewBOperationModel(conn, c.CacheRedis),
		BPaymentModel:         payment.NewBPaymentModel(conn, c.CacheRedis),
		BServiceModel:         service.NewBServiceModel(conn, c.CacheRedis),
		RSubscribeRecordModel: subscriberecord.NewRSubscribeRecordModel(conn),
		BSubscriptionModel:    subscription.NewBSubscriptionModel(c.RedisConf),

		// rpc api
		EmailRpc: email.NewEmail(zrpc.MustNewClient(c.EmailRpc)),
	}
}
