package model

//JiFenRecord 结构体
type JiFenRecord struct {
	ID        int
	VipInfoID int
	AdminID   int
	CreTime   string
	TypeOP    int8
	Num       float64
	IsDelete  int8
}
