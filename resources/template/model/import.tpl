import (
	"context"
	"fmt"
	{{if .time}}"time"{{end}}
	"database/sql"

	"go-zero-base/common/pkg/gormx"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"gorm.io/gorm"
)
