package wiredemo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Database 是数据库连接
type Database struct {
	*sql.DB
}

// NewDatabase 创建一个新的数据库连接
func NewDatabase() (*Database, error) {
	db, err := sql.Open("mysql", "huang:123456@tcp(127.0.0.1:3306)/demo_test")
	if err != nil {
		return nil, err
	}
	return &Database{db}, nil
}
