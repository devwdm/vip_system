package dao

import (
	"database/sql"
	"fmt"
	"vipSys/model"
	"vipSys/utils"
)

//AddVip 新增会员
func AddVip(vip *model.VipInfo) (rowCount int64, err error) {
	sqlstr := "INSERT INTO vipInfo(name, sex, mobile,belong) VALUES(?,?,?,?)"
	result, err := utils.Db.Exec(sqlstr, vip.Name, vip.Sex, vip.Mobile, vip.Belong)
	if err != nil {
		return 0, err
	}
	rowCount, errRes := result.RowsAffected()
	return rowCount, errRes
}

//DelVipById  通过ID删除会员
func DelVipById(id int) (rowCount int64, err error) {
	sqlstr := "UPDATE vipInfo SET isDelete=1 WHERE id=?"
	result, err := utils.Db.Exec(sqlstr, id)
	if err != nil {
		return 0, err
	}
	rowCount, errRes := result.RowsAffected()
	return rowCount, errRes
}

//UpdateVip 更新会员信息
func UpdateVip(vip *model.VipInfo) (rowCount int64, err error) {
	sqlstr := "UPDATE vipInfo set name=?, sex=?, mobile=?, status=? WHERE ID=?"
	result, err := utils.Db.Exec(sqlstr, vip.Name, vip.Sex, vip.Mobile, vip.Status, vip.ID)
	if err != nil {
		return 0, err
	}
	rowCount, errRes := result.RowsAffected()
	return rowCount, errRes
}

//AddJiFen 添加积分
// @ID 会员ID
func AddJiFen(ID int, jiFen float64) (rowCount int64, err error) {
	sqlstr := "UPDATE vipInfo set jiFen=jifen+?, JiFenCount=JiFenCount+? WHERE ID=?"
	result, err := utils.Db.Exec(sqlstr, jiFen, jiFen, ID)
	if err != nil {
		return 0, err
	}
	rowCount, errRes := result.RowsAffected()
	return rowCount, errRes
}

//ReduceJiFen 减少积分
// @ID 会员ID
func ReduceJiFen(ID int, jiFen float64) (rowCount int64, err error) {
	sqlstr := "UPDATE vipInfo set jiFen=jifen-?, WHERE ID=?"
	result, err := utils.Db.Exec(sqlstr, jiFen, ID)
	if err != nil {
		return 0, err
	}

	rowCount, errRes := result.RowsAffected()
	return rowCount, errRes
}

//CheckVipName 判断会员名是否存在
func CheckVipName(name string) bool {
	sqlstr := "SELECT ID FROM vipInfo WHERE Name=?"
	row := utils.Db.QueryRow(sqlstr, name)
	var ID int
	err := row.Scan(&ID)
	if err == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}

//CheckVipMobile 判断会员手机号是否存在
func CheckVipMobile(mobile string) bool {
	sqlstr := "SELECT ID FROM vipInfo WHERE mobile=?"
	row := utils.Db.QueryRow(sqlstr, mobile)
	var ID int
	err := row.Scan(&ID)
	if err == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}

//GetPageVipInfo 获取带分页的会员信息
//@iPageNo 页号
//@iPageSize 页面大小
func GetPageVipInfo(iPageNo int64, iPageSize int64) (*model.PageVip, error) {
	sqlStr := "SELECT count(*) FROM vipInfo WHERE isdelete=0"
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	var totalPageNo int64
	if totalRecord%iPageSize == 0 {
		totalPageNo = totalRecord / iPageSize
	} else {
		totalPageNo = totalRecord/iPageSize + 1
	}
	//	获取当前页中的会员
	sqlStr2 := "SELECT ID, Name, Sex, Mobile, JiFen, JiFenCount, Status, CreTime, Belong FROM vipInfo WHERE isdelete=0 LIMIT ?,?"
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*iPageSize, iPageSize)
	if err != nil {
		return nil, err
	}
	var vipInofs []*model.VipInfoRetJson
	for rows.Next() {
		vipInfo := &model.VipInfoRetJson{}
		rows.Scan(&vipInfo.ID, &vipInfo.Name, &vipInfo.Sex, &vipInfo.Mobile, &vipInfo.JiFen, &vipInfo.JiFenCount, &vipInfo.Status, &vipInfo.CreTime, &vipInfo.Belong)
		vipInofs = append(vipInofs, vipInfo)
	}
	//	创建page
	page := &model.PageVip{
		VipInfos:    vipInofs,
		PageNo:      iPageNo,
		PageSize:    iPageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

//GetVipInfoByWhere 根据条件来获取带分页的会员信息
func GetPageVipInfoByWhere(id int64, name string, mobile string, iPageNo int64, iPageSize int64) (*model.PageVip, error) {
	sqlWhere := "IsDelete=0 "
	if id != 0 {
		sqlWhere = fmt.Sprintf("%s AND id=%d", sqlWhere, id)
	}
	if name != "" {
		sqlWhere = fmt.Sprintf("%s AND name='%s'", sqlWhere, name)
	}
	if mobile != "" {
		sqlWhere = fmt.Sprintf("%s AND mobile='%s'", sqlWhere, mobile)
	}
	sqlStrC := "SELECT count(*) FROM vipInfo WHERE "
	sqlStrC += sqlWhere
	fmt.Println(sqlStrC)
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStrC)
	row.Scan(&totalRecord)
	//设置每页记录数
	//设置一个变量接收总页数
	var totalPageNo int64
	if totalRecord%iPageSize == 0 {
		totalPageNo = totalRecord / iPageSize
	} else {
		totalPageNo = totalRecord/iPageSize + 1
	}
	//	获取当前页中的会员
	sqlStrS := "SELECT ID, Name, Sex, Mobile, JiFen, JiFenCount, Status, CreTime, Belong FROM vipInfo WHERE "
	sqlStrS += sqlWhere
	sqlOrder := " AND isdelete=0 LIMIT ?,?"
	sqlStrS += sqlOrder
	rows, err := utils.Db.Query(sqlStrS, (iPageNo-1)*iPageSize, iPageSize)
	if err != nil {
		return nil, err
	}
	var vipInofs []*model.VipInfoRetJson
	for rows.Next() {
		vipInfo := &model.VipInfoRetJson{}
		rows.Scan(&vipInfo.ID, &vipInfo.Name, &vipInfo.Sex, &vipInfo.Mobile, &vipInfo.JiFen, &vipInfo.JiFenCount, &vipInfo.Status, &vipInfo.CreTime, &vipInfo.Belong)
		vipInofs = append(vipInofs, vipInfo)
	}
	//	创建page
	page := &model.PageVip{
		PageNo:      iPageNo,
		PageSize:    iPageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
		VipInfos:    vipInofs,
	}
	return page, nil

}

//GetOneVipInfo 获取单个会员
func GetOneVipInfoById(id int64) ([]*model.VipInfoRetJson, error) {
	sqlstr := "SELECT ID, Name, Sex, Mobile, JiFen, JiFenCount, Status, CreTime, Belong FROM vipInfo WHERE isdelete=0 AND id=?"
	rows, err := utils.Db.Query(sqlstr, id)
	if err != nil {
		return nil, err
	}
	var vipInfos []*model.VipInfoRetJson
	for rows.Next() {
		vipInfo := &model.VipInfoRetJson{}
		rows.Scan(&vipInfo.ID, &vipInfo.Name, &vipInfo.Sex, &vipInfo.Mobile, &vipInfo.JiFen, &vipInfo.JiFenCount, &vipInfo.Status, &vipInfo.CreTime, &vipInfo.Belong)
		vipInfos = append(vipInfos, vipInfo)
	}
	return vipInfos, nil
}
