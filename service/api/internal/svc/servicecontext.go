package svc

import (
	"cleaningservice/service/api/internal/config"
	"cleaningservice/service/model/address"
	"cleaningservice/service/model/company"
	"cleaningservice/service/model/contractor"
	"cleaningservice/service/model/contractorservice"
	"cleaningservice/service/model/customer"
	"cleaningservice/service/model/employee"
	"cleaningservice/service/model/employeeservice"
	"cleaningservice/service/model/operation"
	"cleaningservice/service/model/order"
	"cleaningservice/service/model/payment"
	"cleaningservice/service/model/schedule"
	"cleaningservice/service/model/service"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config

	// models dao
	BAddressModel           address.BAddressModel
	BCompanyModel           company.BCompanyModel
	BCustomerModel          customer.BCustomerModel
	BContractorModel        contractor.BContractorModel
	RContractorServiceModel contractorservice.RContractorServiceModel
	BEmployeeModel          employee.BEmployeeModel
	REmployeeServiceModel   employeeservice.REmployeeServiceModel
	BOperationModel         operation.BOperationModel
	BOrderModel             order.BOrderModel
	BPaymentModel           payment.BPaymentModel
	BServiceModel           service.BServiceModel
	BScheduleModel          schedule.BScheduleModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	return &ServiceContext{
		Config:                  c,
		BAddressModel:           address.NewBAddressModel(conn, c.CacheRedis),
		BCompanyModel:           company.NewBCompanyModel(conn, c.CacheRedis),
		BCustomerModel:          customer.NewBCustomerModel(conn, c.CacheRedis),
		BContractorModel:        contractor.NewBContractorModel(conn, c.CacheRedis),
		RContractorServiceModel: contractorservice.NewRContractorServiceModel(conn),
		BEmployeeModel:          employee.NewBEmployeeModel(conn, c.CacheRedis),
		REmployeeServiceModel:   employeeservice.NewREmployeeServiceModel(conn),
		BOrderModel:             order.NewBOrderModel(conn, c.CacheRedis),
		BOperationModel:         operation.NewBOperationModel(conn, c.CacheRedis),
		BPaymentModel:           payment.NewBPaymentModel(conn, c.CacheRedis),
		BServiceModel:           service.NewBServiceModel(conn, c.CacheRedis),
		BScheduleModel:          schedule.NewBScheduleModel(c.RedisConf),
	}
}
