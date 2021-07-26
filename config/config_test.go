package config

import (
	"testing"
)

//使用t.Run确定测试顺序，可以写多个
func TestGetPort(t *testing.T) {
	t.Run("Init", func(t *testing.T) {
		InitConfig("../config.ini")
	})
	t.Run("Port", func(t *testing.T) {
		port := GetPort()
		if port != "9997" {
			t.Errorf("Error is not Equal")
		}
	})
	t.Run("Mysql", func(t *testing.T) {
		m := GetDataBaseConfig()
		if m["ip"] != "127.0.0.1:3306" {
			t.Errorf("Error is not Equal")
		}
	})
}

/*
func TestMain(m *testing.M){
	fmt.Println("main begins")
	//不能去掉，不然不会执行其他的TestXxx
	m.Run()
}
*/
