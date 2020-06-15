package model

//log 结构体
type Log struct {
	ID       int
	AdminID  int
	TypeOP   int8
	Comment  string
	CreTime  string
	IsDelete int8
}
