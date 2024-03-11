package config

import (
	"github.com/zeromicro/go-zero/rest"
	"go-zero-base/common/gormx/config/mysql"
)

type Config struct {
	rest.RestConf
	MySQL mysql.Mysql
	Auth  struct {
		AccessSecret string
		AccessExpire int64
	}
}
