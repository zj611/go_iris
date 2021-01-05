/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 下午5:19
 */
package storage

type ResGetByWaybillNo struct {
	WaybillNo           string           // 运单号
	FxsOrderNo          string           // 丰先生订单号
	BizType             int              // 业务类型，后续会扩展，1为电子回单业务
	OrderTm             string           // 下单时间
	CreatedAt            string
	UpdatedAt           string
}
