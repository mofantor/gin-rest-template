package system

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	BaseApi
}

var (
	userService = service.ServiceGroupApp.SystemServiceGroup.UserService
)
