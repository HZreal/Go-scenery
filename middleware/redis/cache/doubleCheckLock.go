package cache

/**
 * @Author elastic·H
 * @Date 2024-08-11
 * @File: doubleCheckLock.go
 * @Description:
 */

import (
	"sync"
)

// /////////////////////////////////////////////////////////////////////////
func DoubleLockCheck() {
	// 双重检查锁
}

// ///////////////////////////// 单例模式 ////////////////////////////////////////////

func getData() string {
	var data string
	var once sync.Once

	data = getCache()
	if data == "" {
		once.Do(func() {
			data = getFromDB()
		})
	}
	return data
}

func getCache() string {
	return "string"
}

func getFromDB() string {
	return "string"
}
