package svc

import (
	"fmt"
	"go-zero-base/app/user/internal/config"
	"go-zero-base/app/user/internal/model/customer"
	"go-zero-base/app/user/internal/model/user"
	"go-zero-base/common/gormx/config/mysql"
	"log"
)

type ServiceContext struct {
	Config        config.Config
	UserModel     user.UserModel
	CustomerModel customer.CustomerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := mysql.Connect(c.MySQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mysql connected...")
	return &ServiceContext{
		Config:        c,
		UserModel:     user.NewUserModel(db),
		CustomerModel: customer.NewCustomerModel(db),
	}
}
