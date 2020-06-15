package model

//VipInfo 结构体
type VipInfo struct {
	ID         int
	Name       string
	Sex        int
	Mobile     string
	JiFen      float64
	JiFenCount float64
	Status     int
	CreTime    string
	Belong     int
	IsDelete   int
}

//VipInfoRetJson Json返回结构体
type VipInfoRetJson struct {
	ID         int
	Name       string
	Sex        int
	Mobile     string
	JiFen      float64
	JiFenCount float64
	Status     int
	CreTime    string
	Belong     int
}
