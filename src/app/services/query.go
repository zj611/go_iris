/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 下午4:07
 */
package services

import (
	"example.com/m/pkg/storage"
	"github.com/jinzhu/gorm"
)

type QueryI interface {
	GetByName(req string) (res *storage.TestTableRes, err error)
}

type MysqlQueryI struct {
	*gorm.DB
}

func NewQueryI(db *gorm.DB) QueryI {
	return &MysqlQueryI{
		db,
	}
}
func (mu *MysqlQueryI) GetByName(req string) (res *storage.TestTableRes, err error) {
	return new(storage.TestTable).GetByName(mu.DB, req)
}
