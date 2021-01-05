/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 下午5:19
 */
package storage

type ResGetByWaybillNo struct {
	WaybillNo  string // 运单号
	FxsOrderNo string // 丰先生订单号
	BizType    int    // 业务类型，后续会扩展，1为电子回单业务
	OrderTm    string // 下单时间
	CreatedAt  string
	UpdatedAt  string
}

type TestTableRes struct {
	Name        string  `gorm:"column:name;type:varchar(20);primary_key" json:"name"`
	Sex         string  `gorm:"column:sex;type:varchar(10);NOT NULL" json:"sex"`
	Age         int     `gorm:"column:age;type:int(10);default:null" json:"age"`
	MobilePhone string  `gorm:"column:mobile_phone;type:varchar(20);default:null" json:"mobile_phone"`
	Chinese     float32 `gorm:"column:chinese;type:float(10,2);default:0" json:"chinese"`
	Math        float32 `gorm:"column:math;type:float(10,2);default:0" json:"math"`
	English     float32 `gorm:"column:english;type:float(10,2);default:0" json:"english"`
	BirthDay    string  `gorm:"column:birthday;type:date;default:null" json:"birthday"` // 下单时间
	CreatedAt   string  `gorm:"column:created_at;type:datetime" json:"created_at"`
	UpdatedAt   string  `gorm:"column:updated_at;type:datetime" json:"updated_at"`
}
