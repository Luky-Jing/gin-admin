package menu

import (
	"context"
	"github.com/google/wire"
	"gorm.io/gorm"

	"gin-admin/internal/app/dao/util"
	"gin-admin/internal/app/schema"
	"gin-admin/pkg/errors"
)

var MenuActionResourceSet = wire.NewSet(wire.Struct(new(MenuActionResourceDao), "*"))

type MenuActionResourceDao struct {
	DB *gorm.DB
}

func (a *MenuActionResourceDao) getQueryOption(opts ...schema.MenuActionResourceQueryOptions) schema.MenuActionResourceQueryOptions {
	var opt schema.MenuActionResourceQueryOptions
	if len(opts) > 0 {
		opt = opts[0]
	}
	return opt
}

func (a *MenuActionResourceDao) Query(ctx context.Context, params schema.MenuActionResourceQueryParam, opts ...schema.MenuActionResourceQueryOptions) (*schema.MenuActionResourceQueryResult, error) {
	opt := a.getQueryOption(opts...)

	db := GetMenuActionResourceDB(ctx, a.DB)
	if v := params.MenuID; v > 0 {
		subQuery := GetMenuActionDB(ctx, a.DB).
			Where("menu_id=?", v).
			Select("id")
		db = db.Where("action_id IN (?)", subQuery)
	}

	if v := params.MenuIDs; len(v) > 0 {
		subQuery := GetMenuActionDB(ctx, a.DB).Where("menu_id IN (?)", v).Select("id")
		db = db.Where("action_id IN (?)", subQuery)
	}

	if len(opt.OrderFields) > 0 {
		db = db.Order(util.ParseOrder(opt.OrderFields))
	}

	var list MenuActionResources
	pr, err := util.WrapPageQuery(ctx, db, params.PaginationParam, &list)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	qr := &schema.MenuActionResourceQueryResult{
		PageResult: pr,
		Data:       list.ToSchemaMenuActionResources(),
	}

	return qr, nil
}

func (a *MenuActionResourceDao) Get(ctx context.Context, id uint64, opts ...schema.MenuActionResourceQueryOptions) (*schema.MenuActionResource, error) {
	db := GetMenuActionResourceDB(ctx, a.DB).Where("id=?", id)
	var item MenuActionResource
	ok, err := util.FindOne(ctx, db, &item)
	if err != nil {
		return nil, errors.WithStack(err)
	} else if !ok {
		return nil, nil
	}

	return item.ToSchemaMenuActionResource(), nil
}

func (a *MenuActionResourceDao) Create(ctx context.Context, item schema.MenuActionResource) error {
	eitem := SchemaMenuActionResource(item).ToMenuActionResource()
	result := GetMenuActionResourceDB(ctx, a.DB).Create(eitem)
	return errors.WithStack(result.Error)
}

func (a *MenuActionResourceDao) Update(ctx context.Context, id uint64, item schema.MenuActionResource) error {
	eitem := SchemaMenuActionResource(item).ToMenuActionResource()
	result := GetMenuActionResourceDB(ctx, a.DB).Where("id=?", id).Updates(eitem)
	return errors.WithStack(result.Error)
}

func (a *MenuActionResourceDao) Delete(ctx context.Context, id uint64) error {
	result := GetMenuActionResourceDB(ctx, a.DB).Where("id=?", id).Delete(MenuActionResource{})
	return errors.WithStack(result.Error)
}

func (a *MenuActionResourceDao) DeleteByActionID(ctx context.Context, actionID uint64) error {
	result := GetMenuActionResourceDB(ctx, a.DB).Where("action_id=?", actionID).Delete(MenuActionResource{})
	return errors.WithStack(result.Error)
}

func (a *MenuActionResourceDao) DeleteByMenuID(ctx context.Context, menuID uint64) error {
	subQuery := GetMenuActionDB(ctx, a.DB).Where("menu_id=?", menuID).Select("id")
	result := GetMenuActionResourceDB(ctx, a.DB).Where("action_id IN (?)", subQuery).Delete(MenuActionResource{})
	return errors.WithStack(result.Error)
}
