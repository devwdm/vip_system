package dao

import "vipSys/utils"

//AddLog 添加日志信息
// @adminId 用户ID
// @typeOp 操作类型 1增加 2修改 3删除
// @comment 操作内容
func AddLog(adminId int, typeOP int, comment string) error {
	sqlstr := "INSERT INTO log( adminid, typeop, comment) VALUES (?,?,?)"
	_, err := utils.Db.Exec(sqlstr, adminId, typeOP, comment)
	if err != nil {
		return err
	}
	return nil
}
