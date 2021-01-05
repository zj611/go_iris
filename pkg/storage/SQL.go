/**
 * @Author: 张俊ZhangJun
 * @Date: 2020/12/16 下午4:37
 */
package storage
const (
	SQLGetByWaybillNo = `SELECT
			name as name,
			sex as sex,
			age as age,
			chinese as chinese,
			math as math,
			english as english,
			mobile_phone as mobile_phone,
            DATE_FORMAT(created_at,'%Y-%m-%d %H:%i:%s') as created_at,
            DATE_FORMAT(updated_at,'%Y-%m-%d %H:%i:%s') as updated_at
			FROM test_table `
)


//DATE_FORMAT(birthday,'%Y-%m-%d') as birthday,