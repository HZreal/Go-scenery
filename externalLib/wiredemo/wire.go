//go:build wireinject
// +build wireinject

package wiredemo

import (
	"github.com/google/wire"
)

// InitializeUserController 初始化用户控制器
func InitializeUserController() (*UserController, error) {
	wire.Build(NewUserController, NewUserService, NewUserRepository, NewDatabase)
	return nil, nil
}
