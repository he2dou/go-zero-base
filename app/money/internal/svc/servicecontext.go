package svc

import (
	"fmt"
	"github.com/zeromicro/go-zero/rest"
	"go-zero-base/app/money/internal/config"
	"go-zero-base/app/money/internal/model/customer"
	"go-zero-base/app/money/internal/model/user"
	"go-zero-base/common/gormx/config/mysql"
	"go-zero-base/common/middleware"
	"log"
)

type ServiceContext struct {
	Config         config.Config
	UserModel      user.UserModel
	CustomerModel  customer.CustomerModel
	CorsMiddleware rest.Middleware
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

		// middleware
		CorsMiddleware: middleware.NewCorsMiddleware().Handle,
	}
}
