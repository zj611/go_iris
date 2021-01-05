/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 上午10:38
 */
package storage

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
)

type TestTable struct {
	Name        string    `gorm:"column:name;type:varchar(20);primary_key" json:"name"`
	Sex         string    `gorm:"column:sex;type:varchar(10);NOT NULL" json:"sex"`
	Age         int       `gorm:"column:age;type:int(10);default:null" json:"age"`
	MobilePhone string    `gorm:"column:mobile_phone;type:varchar(20);default:null" json:"mobile_phone"`
	Chinese     float32   `gorm:"column:chinese;type:float(10,2);default:0" json:"chinese"`
	Math        float32   `gorm:"column:math;type:float(10,2);default:0" json:"math"`
	English     float32   `gorm:"column:english;type:float(10,2);default:0" json:"english"`
	BirthDay    time.Time `gorm:"column:birthday;type:date;default:null" json:"birthday"` // 下单时间
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}

func (p *TestTable) Insert(db *gorm.DB) error {
	err := db.Create(p).Error
	return err
}

func (p *TestTable) Update(db *gorm.DB, req *TestTable) error {
	if req.Name != "" {
		//if err3 := db.Model(p).Where("name = ?",req.Name).
		//	Update("chinese", req.Chinese).
		//	Update("math", req.Math).
		//	Update("english",req.English).Error; err3 != nil{
		//	return err3

		//if err3 := db.Model(p).Where("name = ?",req.Name).
		//	Updates(req).Error; err3 != nil{
		//	return err3
		//}

		if err3 := db.Model(p).
			Updates(req).Error; err3 != nil {
			return err3
		}
	}
	return nil
}

func (p *TestTable) GetByName(db *gorm.DB, name string) (*TestTableRes, error) {

	var result TestTableRes
	sql := SQLGetByWaybillNo
	sql += " WHERE name = '" + name + "';"
	err := db.Raw(sql).Scan(&result).Error

	//err := db.Where("name=?",name).Find(&result).Error

	fmt.Println(result)

	return &result, err
}
