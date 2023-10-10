// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package yoaqi

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"yqnft/NFTManage/test/model"
)

func newTRoleMenu(db *gorm.DB) tRoleMenu {
	_tRoleMenu := tRoleMenu{}

	_tRoleMenu.tRoleMenuDo.UseDB(db)
	_tRoleMenu.tRoleMenuDo.UseModel(&model.TRoleMenu{})

	tableName := _tRoleMenu.tRoleMenuDo.TableName()
	_tRoleMenu.ALL = field.NewAsterisk(tableName)
	_tRoleMenu.ID = field.NewString(tableName, "id")
	_tRoleMenu.RoleID = field.NewString(tableName, "roleId")
	_tRoleMenu.MenuID = field.NewString(tableName, "menuId")
	_tRoleMenu.Bak1 = field.NewString(tableName, "bak1")
	_tRoleMenu.Bak2 = field.NewString(tableName, "bak2")
	_tRoleMenu.Bak3 = field.NewString(tableName, "bak3")
	_tRoleMenu.Bak4 = field.NewString(tableName, "bak4")
	_tRoleMenu.Bak5 = field.NewString(tableName, "bak5")
	_tRoleMenu.CreateTime = field.NewTime(tableName, "createTime")
	_tRoleMenu.CreateUserID = field.NewString(tableName, "createUserId")
	_tRoleMenu.UpdateTime = field.NewTime(tableName, "updateTime")
	_tRoleMenu.UpdateUserID = field.NewString(tableName, "updateUserId")

	_tRoleMenu.fillFieldMap()

	return _tRoleMenu
}

type tRoleMenu struct {
	tRoleMenuDo tRoleMenuDo

	ALL          field.Asterisk
	ID           field.String // 编号
	RoleID       field.String // 角色编号
	MenuID       field.String // 菜单编号
	Bak1         field.String
	Bak2         field.String
	Bak3         field.String
	Bak4         field.String
	Bak5         field.String
	CreateTime   field.Time
	CreateUserID field.String
	UpdateTime   field.Time
	UpdateUserID field.String

	fieldMap map[string]field.Expr
}

func (t tRoleMenu) Table(newTableName string) *tRoleMenu {
	t.tRoleMenuDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tRoleMenu) As(alias string) *tRoleMenu {
	t.tRoleMenuDo.DO = *(t.tRoleMenuDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tRoleMenu) updateTableName(table string) *tRoleMenu {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewString(table, "id")
	t.RoleID = field.NewString(table, "roleId")
	t.MenuID = field.NewString(table, "menuId")
	t.Bak1 = field.NewString(table, "bak1")
	t.Bak2 = field.NewString(table, "bak2")
	t.Bak3 = field.NewString(table, "bak3")
	t.Bak4 = field.NewString(table, "bak4")
	t.Bak5 = field.NewString(table, "bak5")
	t.CreateTime = field.NewTime(table, "createTime")
	t.CreateUserID = field.NewString(table, "createUserId")
	t.UpdateTime = field.NewTime(table, "updateTime")
	t.UpdateUserID = field.NewString(table, "updateUserId")

	t.fillFieldMap()

	return t
}

func (t *tRoleMenu) WithContext(ctx context.Context) ITRoleMenuDo {
	return t.tRoleMenuDo.WithContext(ctx)
}

func (t tRoleMenu) TableName() string { return t.tRoleMenuDo.TableName() }

func (t tRoleMenu) Alias() string { return t.tRoleMenuDo.Alias() }

func (t *tRoleMenu) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tRoleMenu) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 12)
	t.fieldMap["id"] = t.ID
	t.fieldMap["roleId"] = t.RoleID
	t.fieldMap["menuId"] = t.MenuID
	t.fieldMap["bak1"] = t.Bak1
	t.fieldMap["bak2"] = t.Bak2
	t.fieldMap["bak3"] = t.Bak3
	t.fieldMap["bak4"] = t.Bak4
	t.fieldMap["bak5"] = t.Bak5
	t.fieldMap["createTime"] = t.CreateTime
	t.fieldMap["createUserId"] = t.CreateUserID
	t.fieldMap["updateTime"] = t.UpdateTime
	t.fieldMap["updateUserId"] = t.UpdateUserID
}

func (t tRoleMenu) clone(db *gorm.DB) tRoleMenu {
	t.tRoleMenuDo.ReplaceDB(db)
	return t
}

type tRoleMenuDo struct{ gen.DO }

type ITRoleMenuDo interface {
	gen.SubQuery
	Debug() ITRoleMenuDo
	WithContext(ctx context.Context) ITRoleMenuDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITRoleMenuDo
	Not(conds ...gen.Condition) ITRoleMenuDo
	Or(conds ...gen.Condition) ITRoleMenuDo
	Select(conds ...field.Expr) ITRoleMenuDo
	Where(conds ...gen.Condition) ITRoleMenuDo
	Order(conds ...field.Expr) ITRoleMenuDo
	Distinct(cols ...field.Expr) ITRoleMenuDo
	Omit(cols ...field.Expr) ITRoleMenuDo
	Join(table schema.Tabler, on ...field.Expr) ITRoleMenuDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITRoleMenuDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITRoleMenuDo
	Group(cols ...field.Expr) ITRoleMenuDo
	Having(conds ...gen.Condition) ITRoleMenuDo
	Limit(limit int) ITRoleMenuDo
	Offset(offset int) ITRoleMenuDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITRoleMenuDo
	Unscoped() ITRoleMenuDo
	Create(values ...*model.TRoleMenu) error
	CreateInBatches(values []*model.TRoleMenu, batchSize int) error
	Save(values ...*model.TRoleMenu) error
	First() (*model.TRoleMenu, error)
	Take() (*model.TRoleMenu, error)
	Last() (*model.TRoleMenu, error)
	Find() ([]*model.TRoleMenu, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TRoleMenu, err error)
	FindInBatches(result *[]*model.TRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TRoleMenu) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITRoleMenuDo
	Assign(attrs ...field.AssignExpr) ITRoleMenuDo
	Joins(fields ...field.RelationField) ITRoleMenuDo
	Preload(fields ...field.RelationField) ITRoleMenuDo
	FirstOrInit() (*model.TRoleMenu, error)
	FirstOrCreate() (*model.TRoleMenu, error)
	FindByPage(offset int, limit int) (result []*model.TRoleMenu, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITRoleMenuDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tRoleMenuDo) Debug() ITRoleMenuDo {
	return t.withDO(t.DO.Debug())
}

func (t tRoleMenuDo) WithContext(ctx context.Context) ITRoleMenuDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tRoleMenuDo) ReadDB() ITRoleMenuDo {
	return t.Clauses(dbresolver.Read)
}

func (t tRoleMenuDo) WriteDB() ITRoleMenuDo {
	return t.Clauses(dbresolver.Write)
}

func (t tRoleMenuDo) Clauses(conds ...clause.Expression) ITRoleMenuDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tRoleMenuDo) Returning(value interface{}, columns ...string) ITRoleMenuDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tRoleMenuDo) Not(conds ...gen.Condition) ITRoleMenuDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tRoleMenuDo) Or(conds ...gen.Condition) ITRoleMenuDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tRoleMenuDo) Select(conds ...field.Expr) ITRoleMenuDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tRoleMenuDo) Where(conds ...gen.Condition) ITRoleMenuDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tRoleMenuDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITRoleMenuDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tRoleMenuDo) Order(conds ...field.Expr) ITRoleMenuDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tRoleMenuDo) Distinct(cols ...field.Expr) ITRoleMenuDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tRoleMenuDo) Omit(cols ...field.Expr) ITRoleMenuDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tRoleMenuDo) Join(table schema.Tabler, on ...field.Expr) ITRoleMenuDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tRoleMenuDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITRoleMenuDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tRoleMenuDo) RightJoin(table schema.Tabler, on ...field.Expr) ITRoleMenuDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tRoleMenuDo) Group(cols ...field.Expr) ITRoleMenuDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tRoleMenuDo) Having(conds ...gen.Condition) ITRoleMenuDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tRoleMenuDo) Limit(limit int) ITRoleMenuDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tRoleMenuDo) Offset(offset int) ITRoleMenuDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tRoleMenuDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITRoleMenuDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tRoleMenuDo) Unscoped() ITRoleMenuDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tRoleMenuDo) Create(values ...*model.TRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tRoleMenuDo) CreateInBatches(values []*model.TRoleMenu, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tRoleMenuDo) Save(values ...*model.TRoleMenu) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tRoleMenuDo) First() (*model.TRoleMenu, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRoleMenu), nil
	}
}

func (t tRoleMenuDo) Take() (*model.TRoleMenu, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRoleMenu), nil
	}
}

func (t tRoleMenuDo) Last() (*model.TRoleMenu, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRoleMenu), nil
	}
}

func (t tRoleMenuDo) Find() ([]*model.TRoleMenu, error) {
	result, err := t.DO.Find()
	return result.([]*model.TRoleMenu), err
}

func (t tRoleMenuDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TRoleMenu, err error) {
	buf := make([]*model.TRoleMenu, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tRoleMenuDo) FindInBatches(result *[]*model.TRoleMenu, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tRoleMenuDo) Attrs(attrs ...field.AssignExpr) ITRoleMenuDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tRoleMenuDo) Assign(attrs ...field.AssignExpr) ITRoleMenuDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tRoleMenuDo) Joins(fields ...field.RelationField) ITRoleMenuDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tRoleMenuDo) Preload(fields ...field.RelationField) ITRoleMenuDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tRoleMenuDo) FirstOrInit() (*model.TRoleMenu, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRoleMenu), nil
	}
}

func (t tRoleMenuDo) FirstOrCreate() (*model.TRoleMenu, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TRoleMenu), nil
	}
}

func (t tRoleMenuDo) FindByPage(offset int, limit int) (result []*model.TRoleMenu, count int64, err error) {
	result, err = t.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = t.Offset(-1).Limit(-1).Count()
	return
}

func (t tRoleMenuDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tRoleMenuDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tRoleMenuDo) Delete(models ...*model.TRoleMenu) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tRoleMenuDo) withDO(do gen.Dao) *tRoleMenuDo {
	t.DO = *do.(*gen.DO)
	return t
}