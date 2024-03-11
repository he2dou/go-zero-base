脚手架生成api命令
goctl api go -api ocelink.api -dir ./

脚手架生成rpc命令
goctl rpc protoc .\ocelink.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

脚手架生成Docker文件命令
goctl docker --go .\ocelink.go --exe ocelink

根据模板生产
goctl docker --go ocelink.go --exe ocelink --home ......\deploy\goctl\template\

脚手架生成Model文件命令
goctl model mysql ddl -src="deploy/sql/*.sql" -dir="./app/ocelink/model" -c=true

连接数据库脚手架生成Model文件命令-可指定表名
goctl model mysql datasource -url="root:donson.123@tcp(192.168.0.170:3306)/draning" -table="mb_opr_taobao*" -dir="./app/ocelink/model"


# model
goctl model mysql ddl --src ./database/sql/customer.sql --dir ./app/user/internal/model/customer --style=goZero --home ./resources/template

# api
goctl api go -api ./app/user/api/user.api -dir ./app/user
