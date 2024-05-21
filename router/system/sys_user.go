package system

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (s *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	userRouter := Router.Group("user").Use()
	userRouterWithoutRecord := Router.Group("user")
	baseApi := v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		userRouter.POST("admin_register", baseApi.Register)       // 管理员注册账号
		userRouter.POST("changePassword", baseApi.ChangePassword) // 用户修改密码

	}
	{
		userRouterWithoutRecord.POST("getUserList", baseApi.GetUserList) // 分页获取用户列表
	}
}
