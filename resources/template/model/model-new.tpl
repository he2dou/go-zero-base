func ({{.upperStartCamelObject}}) TableName() string {
    return {{.table}}
}

func new{{.upperStartCamelObject}}Model(conn *gorm.DB{{if .withCache}}, c cache.CacheConf{{end}}) *default{{.upperStartCamelObject}}Model {
	return &default{{.upperStartCamelObject}}Model{
		{{if .withCache}}CachedConn: gormx.NewConn(conn, c){{else}}conn:conn{{end}},
		table: {{.table}},
	}
}
