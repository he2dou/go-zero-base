# 直接使用sql生成model
goctl model mysql ddl --src user.sql --dir ../../internal/model/user

# 使用本地template生成model
goctl model mysql ddl --src ./database/sql/customer.sql --dir ./app/user/internal/model/customer --style=goZero --home ./resources/template

# 使用数据库表生成model
goctl model mysql datasource -url="root:donson.123@tcp(192.168.1.137:3306)/test" -table="customer"  -dir="services/model" -cache=true --style=goZero --home ./template
