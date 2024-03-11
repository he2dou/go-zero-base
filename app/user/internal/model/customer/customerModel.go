package customer

import (
	"context"
	"go-zero-base/app/user/internal/types"
	"gorm.io/gorm"
)

var _ CustomerModel = (*customCustomerModel)(nil)

type (
	// CustomerModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerModel.
	CustomerModel interface {
		customerModel
		customCustomerLogicModel
	}

	customCustomerModel struct {
		*defaultCustomerModel
	}

	customCustomerLogicModel interface {
		FindPage(ctx context.Context, tx *gorm.DB, req *types.ListCustomerReq, data interface{}) (int64, error)
	}
)

// NewCustomerModel returns a model for the database table.
func NewCustomerModel(conn *gorm.DB) CustomerModel {
	return &customCustomerModel{
		defaultCustomerModel: newCustomerModel(conn),
	}
}

func (m *defaultCustomerModel) customCacheKeys(data *Customer) []string {
	if data == nil {
		return []string{}
	}
	return []string{}
}

func (m *defaultCustomerModel) FindPage(ctx context.Context, tx *gorm.DB, req *types.ListCustomerReq, data interface{}) (int64, error) {
	db := m.conn
	if tx != nil {
		db = tx
	} else {
		db = db.WithContext(ctx).Model(&Customer{})
	}
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	} else if count == 0 {
		return count, nil
	}

	if req.UserName != "" {
		db = db.Where("user_name = ?", req.UserName)
	}

	if req.Page > 0 && req.PageSize > 0 {
		db = db.Offset((req.Page - 1) * req.PageSize).Limit(req.PageSize)
	} else if req.PageSize > 0 {
		db = db.Limit(req.PageSize)
	}

	err = db.Find(data).Error
	return count, err
}
