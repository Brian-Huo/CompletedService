package svc

import (
	"cleaningservice/service/api/internal/config"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/broadcast"
	"cleaningservice/service/model/category"
	"cleaningservice/service/model/company"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/customer"
	"cleaningservice/service/model/operation"
	"cleaningservice/service/model/order"
	"cleaningservice/service/model/orderdelay"
	"cleaningservice/service/model/payment"
	"cleaningservice/service/model/service"
	"cleaningservice/service/model/subscriberecord"
	"cleaningservice/service/model/subscription"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
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
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                c,
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
	}
}
