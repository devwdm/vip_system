package dao

import "vipSys/utils"

//AddJiFenRecord 添加积分变化记录
// @vipinfoID 会员ID
// @adminID 变更人ID
// @typeOP 会员ID 1增加 2减少
func AddJiFenRecord(vipInfoID int, adminID int, typeOP int, num float64) error {
	sqlstr := "INSERT INTO jiFenRecord(vipInfoid, adminid, typeop, num) VALUES(?,?,?,?) "
	_, err := utils.Db.Exec(sqlstr, vipInfoID, adminID, typeOP, num)
	if err != nil {
		return err
	}
	return nil
}
