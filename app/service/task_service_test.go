package service

import (
	"pear-admin-golang/app/core/config"
	"pear-admin-golang/app/core/db"
	"testing"
)

func init() {
	config.InitConfig("../../config.toml")
	db.InitConn()
}

func TestRunTask(t *testing.T) {
	//task, _ := dao.NewTaskDaoImpl().FindOne(17)
	select {}
}
