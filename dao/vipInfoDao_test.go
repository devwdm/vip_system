package dao

import (
	"testing"
)

//func TestMain(m *testing.M) {
//	fmt.Println("测试Info中的方法")
//	m.Run()
//}

func TestUpdateLastLogin(t *testing.T) {
	UpdateLastLogin(3)
}
