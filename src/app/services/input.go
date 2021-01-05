/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 上午10:18
 */
package services

import (
	"example.com/m/pkg/storage"
	//"github.com/golang-sql/civil"
	"github.com/jinzhu/gorm"
)

type InputI interface {
	Insert(insertData *storage.TestTable)(res string)
	//UpdateDataFromFxs(waybillNo string, fxsOrderNo string, wg *sync.WaitGroup)
	Update(data *storage.TestTable)(res string)
}

type MysqlInputI struct {
	*gorm.DB
}

func NewInputI(db *gorm.DB) InputI {
	return &MysqlInputI{
		db,
	}
}

func (mu *MysqlInputI)Insert(insertData *storage.TestTable)(res string){
	if err := insertData.Insert(mu.DB); err != nil {
		return insertData.Name + " is failed"
	}
	return insertData.Name + " is good!"
}

func (mu *MysqlInputI)Update(req *storage.TestTable)(res string){
	 if err := new(storage.TestTable).Update(mu.DB,req);err != nil{
		 return "update failed"
	 }
	return "update success"
}