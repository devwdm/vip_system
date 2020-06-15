package model

//AdminInfo 结构体
type AdminInfo struct {
	ID        int
	Username  string
	Password  string
	Sex       int
	Mobile    string
	CreTime   string
	LastLogin string
	IsDelete  int
}

//AdminInfoJson Json结构体
type AdminInfoJson struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Sex       int    `json:"sex"`
	Mobile    string `json:"mobile"`
	CreTime   string `json:"cre_time"`
	LastLogin string `json:"last_login"`
}
