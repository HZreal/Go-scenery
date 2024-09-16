package main

import "fmt"

/**
 * @Author elastic·H
 * @Date 2024-09-16
 * @File: funcOption.go
 * @Description: 函数选项模式
 */

type DB struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

type DBOption func(*DB)

func NewDB(opts ...DBOption) *DB {
	db := &DB{
		Host: "localhost",
		Port: 3306,
	}
	for _, opt := range opts {
		opt(db)
	}
	return db
}

func WithUsername(username string) DBOption {
	return func(db *DB) {
		db.Username = username
	}
}

func WithPassword(password string) DBOption {
	return func(db *DB) {
		db.Password = password
	}
}

func WithDatabase(database string) DBOption {
	return func(db *DB) {
		db.Database = database
	}
}

func main() {
	db := NewDB(
		WithUsername("admin"),
		WithPassword("secret"),
		WithDatabase("mydb"),
	)
	fmt.Printf("DB: %+v\n", db)
}
