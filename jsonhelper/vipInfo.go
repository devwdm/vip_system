package jsonhelper

import "vipSys/model"

type VipInfoJson struct {
	Code        string                  `json:"code"`
	TotalRecord int64                   `json:"total_record"`  //总记录数，通过查询数据库得到
	TotalPageNo int64                   `json:"total_page_no"` //总页数，通过计算得到
	PageNo      int64                   `json:"page_no"`       //当前页
	PageSize    int64                   `json:"page_size"`     //每页显示的条数
	VipInfos    []*model.VipInfoRetJson `json:"vip_infos"`     //每页查询出来的会员存放的切片
}
