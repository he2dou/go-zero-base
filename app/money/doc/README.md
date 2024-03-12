# 格式化api
goctl api format --dir main.api

# 生成api代码
goctl api go -api ./app/money/doc/main.api -dir ./app/money

# 生产swagger文档
goctl api plugin -p goctl-swagger="swagger -filename swagger.json" --api main.api --dir .
