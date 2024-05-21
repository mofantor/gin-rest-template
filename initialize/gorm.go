package initialize

import (
	"os"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Gorm() *gorm.DB {
	return GormSqlite()
}

func RegisterTables() {
	db := global.GVA_DB
	err := db.AutoMigrate(
		system.SysUser{},
		system.JwtBlacklist{},
	)
	if err != nil {
		global.GVA_LOG.Error("register table failed", zap.Error(err))
		os.Exit(0)
	}
	global.GVA_LOG.Info("register table success")
}
