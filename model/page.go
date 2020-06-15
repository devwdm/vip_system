package model

//Page 结构
type Page struct {
	VipInfos    []*VipInfo `json:"vip_infos"`     //每页查询出来的会员存放的切片
	PageNo      int64      `json:"page_no"`       //当前页
	PageSize    int64      `json:"page_size"`     //每页显示的条数
	TotalPageNo int64      `json:"total_page_no"` //总页数，通过计算得到
	TotalRecord int64      `json:"total_record"`  //总记录数，通过查询数据库得到
	IsLogin     bool
	Username    string
}

//PageVIP 会员页面结构
type PageVip struct {
	VipInfos    []*VipInfoRetJson `json:"vip_infos"`     //每页查询出来的会员存放的切片
	PageNo      int64             `json:"page_no"`       //当前页
	PageSize    int64             `json:"page_size"`     //每页显示的条数
	TotalPageNo int64             `json:"total_page_no"` //总页数，通过计算得到
	TotalRecord int64             `json:"total_record"`  //总记录数，通过查询数据库得到
}

//IsHasPrev 判断是否有上一页
func (p *PageVip) IsHasPrev() bool {
	return p.PageNo > 1
}

//IsHasNext 判断是否有下一页
func (p *PageVip) IsHasNext() bool {
	return p.PageNo < p.TotalPageNo
}

//GetPrevPageNo 获取上一页
func (p *PageVip) GetPrevPageNo() int64 {
	if p.IsHasPrev() {
		return p.PageNo - 1
	}
	return 1
}

//GetNextPageNo 获取下一页
func (p *PageVip) GetNextPageNo() int64 {
	if p.IsHasNext() {
		return p.PageNo + 1
	}
	return p.TotalPageNo

}

//PageAdminInfo 用户页面结构
type PageAdminInfo struct {
	AdminInfos  []*AdminInfoJson `json:"admin_infos"`   //每页查询出来的用户存放的切片
	PageNo      int64            `json:"page_no"`       //当前页
	PageSize    int64            `json:"page_size"`     //每页显示的条数
	TotalPageNo int64            `json:"total_page_no"` //总页数，通过计算得到
	TotalRecord int64            `json:"total_record"`  //总记录数，通过查询数据库得到
}
