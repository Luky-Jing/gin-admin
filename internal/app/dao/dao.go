package dao

import (
	"strings"

	"github.com/google/wire"
	"gorm.io/gorm"

	"gin-admin/internal/app/config"
	"gin-admin/internal/app/dao/menu"
	"gin-admin/internal/app/dao/role"
	"gin-admin/internal/app/dao/user"
	"gin-admin/internal/app/dao/util"
) // end

// DaoSet dao injection
var DaoSet = wire.NewSet(
	util.TransSet,
	menu.MenuActionResourceSet,
	menu.MenuActionSet,
	menu.MenuSet,
	role.RoleMenuSet,
	role.RoleSet,
	user.UserRoleSet,
	user.UserSet,
) // end

// Define repo type alias
type (
	TransDao              = util.Trans
	MenuActionResourceDao = menu.MenuActionResourceDao
	MenuActionDao         = menu.MenuActionDao
	MenuDao               = menu.MenuDao
	RoleMenuDao           = role.RoleMenuDao
	RoleDao               = role.RoleDao
	UserRoleDao           = user.UserRoleDao
	UserDao               = user.UserDao
) // end

// Auto migration for given models
func AutoMigrate(db *gorm.DB) error {
	if dbType := config.C.Gorm.DBType; strings.ToLower(dbType) == "mysql" {
		db = db.Set("gorm:table_options", "ENGINE=InnoDB")
	}

	return db.AutoMigrate(
		new(menu.MenuActionResource),
		new(menu.MenuAction),
		new(menu.Menu),
		new(role.RoleMenu),
		new(role.Role),
		new(user.UserRole),
		new(user.User),
	) // end
}
