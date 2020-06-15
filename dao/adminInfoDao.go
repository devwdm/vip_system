package dao

import (
	"database/sql"
	"vipSys/model"
	"vipSys/utils"
)

//GetPageAdminInfo 获取带分页的用户信息
func GetPageAdminInfo(iPageNo int64, iPageSize int64) (*model.PageAdminInfo, error) {
	sqlStr := "SELECT count(*) FROM adminInfo WHERE isdelete=0"
	var totalRecord int64
	row := utils.Db.QueryRow(sqlStr)
	row.Scan(&totalRecord)
	var totalPageNo int64
	if totalRecord%iPageSize == 0 {
		totalPageNo = totalRecord / iPageSize
	} else {
		totalPageNo = totalRecord/iPageSize + 1
	}
	//	获取当前页中的用户信息
	sqlStr2 := "SELECT id, username, sex, mobile, cretime, lastlogin FROM adminInfo WHERE isDelete=0 LIMIT ?,?"
	rows, err := utils.Db.Query(sqlStr2, (iPageNo-1)*iPageSize, iPageSize)
	if err != nil {
		return nil, err
	}
	var adminInfos []*model.AdminInfoJson
	for rows.Next() {
		adminInfo := &model.AdminInfoJson{}
		rows.Scan(&adminInfo.ID, &adminInfo.Username, &adminInfo.Sex, &adminInfo.Mobile, &adminInfo.CreTime, &adminInfo.LastLogin)
		adminInfos = append(adminInfos, adminInfo)
	}
	//	创建page
	page := &model.PageAdminInfo{
		AdminInfos:  adminInfos,
		PageNo:      iPageNo,
		PageSize:    iPageSize,
		TotalPageNo: totalPageNo,
		TotalRecord: totalRecord,
	}
	return page, nil
}

//GetAllAdminInfo 获取所有用户信息
func GetAllAdminInfo() ([]*model.AdminInfo, error) {
	sqlstr := "SELECT id, username, sex, mobile, cretime, lastlogin FROM adminInfo WHERE isDelete=0"
	rows, err := utils.Db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	var admins []*model.AdminInfo
	for rows.Next() {
		admin := &model.AdminInfo{}
		rows.Scan(&admin.ID, &admin.Username, &admin.Sex, &admin.Mobile, &admin.CreTime, &admin.LastLogin)
		admins = append(admins, admin)
	}
	return admins, nil
}

//CheckAdminNameAndPassword 根据用户名(手机号)和密码从数据库中查询一条记录
func CheckAdminNameAndPassword(username string, password string) (*model.AdminInfoJson, error) {
	//写sql语句
	sqlStr := "SELECT ID, Username, Sex, Mobile, CreTime, Lastlogin FROM adminInfo WHERE  username = ? AND password = ? AND isDelete=0"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	adminInfo := &model.AdminInfoJson{}
	err := row.Scan(&adminInfo.ID, &adminInfo.Username, &adminInfo.Sex, &adminInfo.Mobile, &adminInfo.CreTime, &adminInfo.LastLogin)
	if err == sql.ErrNoRows {
		sqlStr := "SELECT ID, Username, Sex, Mobile, CreTime, Lastlogin FROM adminInfo WHERE mobile = ? AND password = ? AND isDelete=0"
		row := utils.Db.QueryRow(sqlStr, username, password)
		//adminInfo := &model.AdminInfo{}
		err = row.Scan(&adminInfo.ID, &adminInfo.Username, &adminInfo.Sex, &adminInfo.Mobile, &adminInfo.CreTime, &adminInfo.LastLogin)
	}
	return adminInfo, err
}

//CheckAdminUsername 判断用户是否存在
func CheckAdminUsername(username string) bool {
	sqlstr := "SELECT ID FROM adminInfo WHERE username=?"
	row := utils.Db.QueryRow(sqlstr, username)
	var ID int
	err := row.Scan(&ID)
	if err == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}

//CheckAdminMobile 判断用户手机号是否存在
func CheckAdminMobile(mobile string) bool {
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

//CheckUsernameAndMobile 验证用户名或手机号是否存在
func CheckUsernameAndMobile(username string) bool {
	//写sql语句
	sqlStrUsername := "SELECT id FROM adminInfo WHERE username = ? AND isDelete=0 "
	sqlStrMobile := "SELECT id FROM adminInfo WHERE mobile = ? AND isDelete=0 "
	//执行
	rowUsername := utils.Db.QueryRow(sqlStrUsername, username)
	rowMobile := utils.Db.QueryRow(sqlStrMobile, username)
	var idUsername int
	var idMobile int
	errUsername := rowUsername.Scan(&idUsername)
	errMobile := rowMobile.Scan(&idMobile)
	if (errUsername != sql.ErrNoRows) || (errMobile != sql.ErrNoRows) {
		return true
	} else {
		return false
	}
}

//AddAdmin 添加用户
func AddAdmin(admin *model.AdminInfo) (rowCount int64, err error) {
	//写sql语句
	sqlStr := "INSERT INTO adminInfo(username, password, sex, mobile) VALUE ( ?, ?, ?, ?)"
	//执行
	result, err := utils.Db.Exec(sqlStr, admin.Username, admin.Password, admin.Sex, admin.Mobile)
	if err != nil {
		return 0, err
	}
	rowCount, errRes := result.RowsAffected()
	return rowCount, errRes
}

//DelAdminById  通过ID删除用户r
func DelAdminById(id int) (rowCount int64, err error) {
	sqlStr := "UPDATE adminInfo SET isDelete=1 WHERE id=?"
	result, err := utils.Db.Exec(sqlStr, id)
	if err != nil {
		return 0, err
	}
	rowCount, errRes := result.RowsAffected()
	return rowCount, errRes
}

//UpdateAdminPass 修改密码
func UpdateAdminPassword(password string, ID int) (rowCount int64, err error) {
	//写sql语句
	sqlStr := "UPDATE adminInfo SET Password=? WHERE ID=?;"
	//执行
	result, err := utils.Db.Exec(sqlStr, password, ID)
	if err != nil {
		return 0, err
	}
	rowCount, errRes := result.RowsAffected()
	return rowCount, errRes
}

//CheckAdminPassword 验证当前密码是否正确
func CheckAdminPassword(password string, id int) bool {
	//写sql语句
	sqlstr := "SELECT ID FROM adminInfo WHERE ID=? and password=?;"
	row := utils.Db.QueryRow(sqlstr, id, password)
	var adminID int
	err := row.Scan(&adminID)
	if err == sql.ErrNoRows {
		return false
	} else {
		return true
	}
}

func UpdateLastLogin(id int) {
	sqlstr := "UPDATE adminInfo SET lastlogin=now() WHERE ID=?"
	utils.Db.QueryRow(sqlstr, id)
}
