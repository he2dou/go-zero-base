# 格式化api
goctl api format --dir user.api

# 生成api代码
goctl api go -api ./app/user/api/user.api -dir ./app/user

# 生产swagger文档
goctl api plugin -p goctl-swagger="swagger -filename swagger.json" --api main.api --dir .
