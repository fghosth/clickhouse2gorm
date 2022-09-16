package test

import (
	"fmt"
	"github.com/fghosth/clickhouse2gorm/gen"
	"net/url"
	"testing"
)

func TestClickhouse2Struct(t *testing.T) {
	ip := "127.0.0.1"
	port := "9000"
	dbName := "project_helper"
	username := "default"
	password := url.QueryEscape("zaq1xsw2CDE#")
	dsn := fmt.Sprintf("tcp://%s:%s?database=%s&username=%s&password=%s&read_timeout=10&write_timeout=20", ip, port, dbName, username, password)
	// 生成指定单表
	// tblName := "issues"
	// err := gen.GenerateOne(gen.CHGenConf{
	// 	Dsn:       dsn,
	// 	WritePath: "./model",
	// 	Stdout:    false,
	// 	Overwrite: true,
	// }, dbName, tblName)
	err := gen.GenerateAll(gen.CHGenConf{
		Dsn:       dsn,
		WritePath: "./model",
		Stdout:    false,
		Overwrite: true,
	}, dbName)
	if err != nil {
		return
	}
}
