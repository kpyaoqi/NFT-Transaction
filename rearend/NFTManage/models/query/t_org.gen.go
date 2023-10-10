// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"yqnft/NFTManage/models/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newTOrg(db *gorm.DB) tOrg {
	_tOrg := tOrg{}

	_tOrg.tOrgDo.UseDB(db)
	_tOrg.tOrgDo.UseModel(&model.TOrg{})

	tableName := _tOrg.tOrgDo.TableName()
	_tOrg.ALL = field.NewAsterisk(tableName)
	_tOrg.ID = field.NewString(tableName, "id")
	_tOrg.Name = field.NewString(tableName, "name")
	_tOrg.Comcode = field.NewString(tableName, "comcode")
	_tOrg.Pid = field.NewString(tableName, "pid")
	_tOrg.OrgType = field.NewInt64(tableName, "orgType")
	_tOrg.Sortno = field.NewInt64(tableName, "sortno")
	_tOrg.Remark = field.NewString(tableName, "remark")
	_tOrg.CreateTime = field.NewTime(tableName, "createTime")
	_tOrg.CreateUserID = field.NewString(tableName, "createUserId")
	_tOrg.UpdateTime = field.NewTime(tableName, "updateTime")
	_tOrg.UpdateUserID = field.NewString(tableName, "updateUserId")
	_tOrg.Active = field.NewInt64(tableName, "active")
	_tOrg.Bak1 = field.NewString(tableName, "bak1")
	_tOrg.Bak2 = field.NewString(tableName, "bak2")
	_tOrg.Bak3 = field.NewString(tableName, "bak3")
	_tOrg.Bak4 = field.NewString(tableName, "bak4")
	_tOrg.Bak5 = field.NewString(tableName, "bak5")

	_tOrg.fillFieldMap()

	return _tOrg
}

type tOrg struct {
	tOrgDo tOrgDo

	ALL          field.Asterisk
	ID           field.String // 编号
	Name         field.String // 名称
	Comcode      field.String // 代码
	Pid          field.String // 上级部门ID
	OrgType      field.Int64  // 0-99门店,100-199部门,200-299,分公司,300-399集团公司,900-999总平台
	Sortno       field.Int64  // 排序,查询时倒叙排列
	Remark       field.String // 备注
	CreateTime   field.Time
	CreateUserID field.String
	UpdateTime   field.Time
	UpdateUserID field.String
	Active       field.Int64 // 是否有效(0否,1是)
	Bak1         field.String
	Bak2         field.String
	Bak3         field.String
	Bak4         field.String
	Bak5         field.String

	fieldMap map[string]field.Expr
}

func (t tOrg) Table(newTableName string) *tOrg {
	t.tOrgDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tOrg) As(alias string) *tOrg {
	t.tOrgDo.DO = *(t.tOrgDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tOrg) updateTableName(table string) *tOrg {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewString(table, "id")
	t.Name = field.NewString(table, "name")
	t.Comcode = field.NewString(table, "comcode")
	t.Pid = field.NewString(table, "pid")
	t.OrgType = field.NewInt64(table, "orgType")
	t.Sortno = field.NewInt64(table, "sortno")
	t.Remark = field.NewString(table, "remark")
	t.CreateTime = field.NewTime(table, "createTime")
	t.CreateUserID = field.NewString(table, "createUserId")
	t.UpdateTime = field.NewTime(table, "updateTime")
	t.UpdateUserID = field.NewString(table, "updateUserId")
	t.Active = field.NewInt64(table, "active")
	t.Bak1 = field.NewString(table, "bak1")
	t.Bak2 = field.NewString(table, "bak2")
	t.Bak3 = field.NewString(table, "bak3")
	t.Bak4 = field.NewString(table, "bak4")
	t.Bak5 = field.NewString(table, "bak5")

	t.fillFieldMap()

	return t
}

func (t *tOrg) WithContext(ctx context.Context) ITOrgDo { return t.tOrgDo.WithContext(ctx) }

func (t tOrg) TableName() string { return t.tOrgDo.TableName() }

func (t tOrg) Alias() string { return t.tOrgDo.Alias() }

func (t *tOrg) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tOrg) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 17)
	t.fieldMap["id"] = t.ID
	t.fieldMap["name"] = t.Name
	t.fieldMap["comcode"] = t.Comcode
	t.fieldMap["pid"] = t.Pid
	t.fieldMap["orgType"] = t.OrgType
	t.fieldMap["sortno"] = t.Sortno
	t.fieldMap["remark"] = t.Remark
	t.fieldMap["createTime"] = t.CreateTime
	t.fieldMap["createUserId"] = t.CreateUserID
	t.fieldMap["updateTime"] = t.UpdateTime
	t.fieldMap["updateUserId"] = t.UpdateUserID
	t.fieldMap["active"] = t.Active
	t.fieldMap["bak1"] = t.Bak1
	t.fieldMap["bak2"] = t.Bak2
	t.fieldMap["bak3"] = t.Bak3
	t.fieldMap["bak4"] = t.Bak4
	t.fieldMap["bak5"] = t.Bak5
}

func (t tOrg) clone(db *gorm.DB) tOrg {
	t.tOrgDo.ReplaceDB(db)
	return t
}

type tOrgDo struct{ gen.DO }

type ITOrgDo interface {
	gen.SubQuery
	Debug() ITOrgDo
	WithContext(ctx context.Context) ITOrgDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITOrgDo
	Not(conds ...gen.Condition) ITOrgDo
	Or(conds ...gen.Condition) ITOrgDo
	Select(conds ...field.Expr) ITOrgDo
	Where(conds ...gen.Condition) ITOrgDo
	Order(conds ...field.Expr) ITOrgDo
	Distinct(cols ...field.Expr) ITOrgDo
	Omit(cols ...field.Expr) ITOrgDo
	Join(table schema.Tabler, on ...field.Expr) ITOrgDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITOrgDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITOrgDo
	Group(cols ...field.Expr) ITOrgDo
	Having(conds ...gen.Condition) ITOrgDo
	Limit(limit int) ITOrgDo
	Offset(offset int) ITOrgDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITOrgDo
	Unscoped() ITOrgDo
	Create(values ...*model.TOrg) error
	CreateInBatches(values []*model.TOrg, batchSize int) error
	Save(values ...*model.TOrg) error
	First() (*model.TOrg, error)
	Take() (*model.TOrg, error)
	Last() (*model.TOrg, error)
	Find() ([]*model.TOrg, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOrg, err error)
	FindInBatches(result *[]*model.TOrg, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TOrg) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITOrgDo
	Assign(attrs ...field.AssignExpr) ITOrgDo
	Joins(fields ...field.RelationField) ITOrgDo
	Preload(fields ...field.RelationField) ITOrgDo
	FirstOrInit() (*model.TOrg, error)
	FirstOrCreate() (*model.TOrg, error)
	FindByPage(offset int, limit int) (result []*model.TOrg, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITOrgDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tOrgDo) Debug() ITOrgDo {
	return t.withDO(t.DO.Debug())
}

func (t tOrgDo) WithContext(ctx context.Context) ITOrgDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tOrgDo) ReadDB() ITOrgDo {
	return t.Clauses(dbresolver.Read)
}

func (t tOrgDo) WriteDB() ITOrgDo {
	return t.Clauses(dbresolver.Write)
}

func (t tOrgDo) Clauses(conds ...clause.Expression) ITOrgDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tOrgDo) Returning(value interface{}, columns ...string) ITOrgDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tOrgDo) Not(conds ...gen.Condition) ITOrgDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tOrgDo) Or(conds ...gen.Condition) ITOrgDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tOrgDo) Select(conds ...field.Expr) ITOrgDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tOrgDo) Where(conds ...gen.Condition) ITOrgDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tOrgDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITOrgDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tOrgDo) Order(conds ...field.Expr) ITOrgDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tOrgDo) Distinct(cols ...field.Expr) ITOrgDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tOrgDo) Omit(cols ...field.Expr) ITOrgDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tOrgDo) Join(table schema.Tabler, on ...field.Expr) ITOrgDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tOrgDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITOrgDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tOrgDo) RightJoin(table schema.Tabler, on ...field.Expr) ITOrgDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tOrgDo) Group(cols ...field.Expr) ITOrgDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tOrgDo) Having(conds ...gen.Condition) ITOrgDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tOrgDo) Limit(limit int) ITOrgDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tOrgDo) Offset(offset int) ITOrgDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tOrgDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITOrgDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tOrgDo) Unscoped() ITOrgDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tOrgDo) Create(values ...*model.TOrg) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tOrgDo) CreateInBatches(values []*model.TOrg, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tOrgDo) Save(values ...*model.TOrg) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tOrgDo) First() (*model.TOrg, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOrg), nil
	}
}

func (t tOrgDo) Take() (*model.TOrg, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOrg), nil
	}
}

func (t tOrgDo) Last() (*model.TOrg, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOrg), nil
	}
}

func (t tOrgDo) Find() ([]*model.TOrg, error) {
	result, err := t.DO.Find()
	return result.([]*model.TOrg), err
}

func (t tOrgDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TOrg, err error) {
	buf := make([]*model.TOrg, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tOrgDo) FindInBatches(result *[]*model.TOrg, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tOrgDo) Attrs(attrs ...field.AssignExpr) ITOrgDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tOrgDo) Assign(attrs ...field.AssignExpr) ITOrgDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tOrgDo) Joins(fields ...field.RelationField) ITOrgDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tOrgDo) Preload(fields ...field.RelationField) ITOrgDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tOrgDo) FirstOrInit() (*model.TOrg, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOrg), nil
	}
}

func (t tOrgDo) FirstOrCreate() (*model.TOrg, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TOrg), nil
	}
}

func (t tOrgDo) FindByPage(offset int, limit int) (result []*model.TOrg, count int64, err error) {
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

func (t tOrgDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tOrgDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tOrgDo) Delete(models ...*model.TOrg) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tOrgDo) withDO(do gen.Dao) *tOrgDo {
	t.DO = *do.(*gen.DO)
	return t
}