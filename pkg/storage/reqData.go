/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 下午3:18
 */
package storage

//type ReqInsertData struct {
//	WaybillNo    string
//	Fxs_order_no string
//	Biz_type     int
//
//}

type ReqUpdateData struct {
	Name           string     `gorm:"column:name;type:varchar(20);primary_key" json:"name"`
	Chinese         float32   `gorm:"column:chinese;type:int(100);default:0" json:"chinese"`
	Math            float32   `gorm:"column:math;type:float(10,2);default:0" json:"math"`
	English         float32   `gorm:"column:english;type:float(10,2);default:0"`
}
