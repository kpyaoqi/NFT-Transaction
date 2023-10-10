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

func newTTableindex(db *gorm.DB) tTableindex {
	_tTableindex := tTableindex{}

	_tTableindex.tTableindexDo.UseDB(db)
	_tTableindex.tTableindexDo.UseModel(&model.TTableindex{})

	tableName := _tTableindex.tTableindexDo.TableName()
	_tTableindex.ALL = field.NewAsterisk(tableName)
	_tTableindex.ID = field.NewString(tableName, "id")
	_tTableindex.MaxIndex = field.NewInt64(tableName, "maxIndex")
	_tTableindex.Prefix = field.NewString(tableName, "prefix")
	_tTableindex.Bak1 = field.NewString(tableName, "bak1")
	_tTableindex.Bak2 = field.NewString(tableName, "bak2")
	_tTableindex.Bak3 = field.NewString(tableName, "bak3")
	_tTableindex.Bak4 = field.NewString(tableName, "bak4")
	_tTableindex.Bak5 = field.NewString(tableName, "bak5")

	_tTableindex.fillFieldMap()

	return _tTableindex
}

type tTableindex struct {
	tTableindexDo tTableindexDo

	ALL      field.Asterisk
	ID       field.String // 表名
	MaxIndex field.Int64  // 表记录最大的行,一直累加
	Prefix   field.String // 前缀 单个字母加 _
	Bak1     field.String
	Bak2     field.String
	Bak3     field.String
	Bak4     field.String
	Bak5     field.String

	fieldMap map[string]field.Expr
}

func (t tTableindex) Table(newTableName string) *tTableindex {
	t.tTableindexDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tTableindex) As(alias string) *tTableindex {
	t.tTableindexDo.DO = *(t.tTableindexDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tTableindex) updateTableName(table string) *tTableindex {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewString(table, "id")
	t.MaxIndex = field.NewInt64(table, "maxIndex")
	t.Prefix = field.NewString(table, "prefix")
	t.Bak1 = field.NewString(table, "bak1")
	t.Bak2 = field.NewString(table, "bak2")
	t.Bak3 = field.NewString(table, "bak3")
	t.Bak4 = field.NewString(table, "bak4")
	t.Bak5 = field.NewString(table, "bak5")

	t.fillFieldMap()

	return t
}

func (t *tTableindex) WithContext(ctx context.Context) ITTableindexDo {
	return t.tTableindexDo.WithContext(ctx)
}

func (t tTableindex) TableName() string { return t.tTableindexDo.TableName() }

func (t tTableindex) Alias() string { return t.tTableindexDo.Alias() }

func (t *tTableindex) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tTableindex) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 8)
	t.fieldMap["id"] = t.ID
	t.fieldMap["maxIndex"] = t.MaxIndex
	t.fieldMap["prefix"] = t.Prefix
	t.fieldMap["bak1"] = t.Bak1
	t.fieldMap["bak2"] = t.Bak2
	t.fieldMap["bak3"] = t.Bak3
	t.fieldMap["bak4"] = t.Bak4
	t.fieldMap["bak5"] = t.Bak5
}

func (t tTableindex) clone(db *gorm.DB) tTableindex {
	t.tTableindexDo.ReplaceDB(db)
	return t
}

type tTableindexDo struct{ gen.DO }

type ITTableindexDo interface {
	gen.SubQuery
	Debug() ITTableindexDo
	WithContext(ctx context.Context) ITTableindexDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITTableindexDo
	Not(conds ...gen.Condition) ITTableindexDo
	Or(conds ...gen.Condition) ITTableindexDo
	Select(conds ...field.Expr) ITTableindexDo
	Where(conds ...gen.Condition) ITTableindexDo
	Order(conds ...field.Expr) ITTableindexDo
	Distinct(cols ...field.Expr) ITTableindexDo
	Omit(cols ...field.Expr) ITTableindexDo
	Join(table schema.Tabler, on ...field.Expr) ITTableindexDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITTableindexDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITTableindexDo
	Group(cols ...field.Expr) ITTableindexDo
	Having(conds ...gen.Condition) ITTableindexDo
	Limit(limit int) ITTableindexDo
	Offset(offset int) ITTableindexDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITTableindexDo
	Unscoped() ITTableindexDo
	Create(values ...*model.TTableindex) error
	CreateInBatches(values []*model.TTableindex, batchSize int) error
	Save(values ...*model.TTableindex) error
	First() (*model.TTableindex, error)
	Take() (*model.TTableindex, error)
	Last() (*model.TTableindex, error)
	Find() ([]*model.TTableindex, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TTableindex, err error)
	FindInBatches(result *[]*model.TTableindex, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TTableindex) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITTableindexDo
	Assign(attrs ...field.AssignExpr) ITTableindexDo
	Joins(fields ...field.RelationField) ITTableindexDo
	Preload(fields ...field.RelationField) ITTableindexDo
	FirstOrInit() (*model.TTableindex, error)
	FirstOrCreate() (*model.TTableindex, error)
	FindByPage(offset int, limit int) (result []*model.TTableindex, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITTableindexDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tTableindexDo) Debug() ITTableindexDo {
	return t.withDO(t.DO.Debug())
}

func (t tTableindexDo) WithContext(ctx context.Context) ITTableindexDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tTableindexDo) ReadDB() ITTableindexDo {
	return t.Clauses(dbresolver.Read)
}

func (t tTableindexDo) WriteDB() ITTableindexDo {
	return t.Clauses(dbresolver.Write)
}

func (t tTableindexDo) Clauses(conds ...clause.Expression) ITTableindexDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tTableindexDo) Returning(value interface{}, columns ...string) ITTableindexDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tTableindexDo) Not(conds ...gen.Condition) ITTableindexDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tTableindexDo) Or(conds ...gen.Condition) ITTableindexDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tTableindexDo) Select(conds ...field.Expr) ITTableindexDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tTableindexDo) Where(conds ...gen.Condition) ITTableindexDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tTableindexDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITTableindexDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tTableindexDo) Order(conds ...field.Expr) ITTableindexDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tTableindexDo) Distinct(cols ...field.Expr) ITTableindexDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tTableindexDo) Omit(cols ...field.Expr) ITTableindexDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tTableindexDo) Join(table schema.Tabler, on ...field.Expr) ITTableindexDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tTableindexDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITTableindexDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tTableindexDo) RightJoin(table schema.Tabler, on ...field.Expr) ITTableindexDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tTableindexDo) Group(cols ...field.Expr) ITTableindexDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tTableindexDo) Having(conds ...gen.Condition) ITTableindexDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tTableindexDo) Limit(limit int) ITTableindexDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tTableindexDo) Offset(offset int) ITTableindexDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tTableindexDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITTableindexDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tTableindexDo) Unscoped() ITTableindexDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tTableindexDo) Create(values ...*model.TTableindex) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tTableindexDo) CreateInBatches(values []*model.TTableindex, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tTableindexDo) Save(values ...*model.TTableindex) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tTableindexDo) First() (*model.TTableindex, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTableindex), nil
	}
}

func (t tTableindexDo) Take() (*model.TTableindex, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTableindex), nil
	}
}

func (t tTableindexDo) Last() (*model.TTableindex, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTableindex), nil
	}
}

func (t tTableindexDo) Find() ([]*model.TTableindex, error) {
	result, err := t.DO.Find()
	return result.([]*model.TTableindex), err
}

func (t tTableindexDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TTableindex, err error) {
	buf := make([]*model.TTableindex, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tTableindexDo) FindInBatches(result *[]*model.TTableindex, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tTableindexDo) Attrs(attrs ...field.AssignExpr) ITTableindexDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tTableindexDo) Assign(attrs ...field.AssignExpr) ITTableindexDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tTableindexDo) Joins(fields ...field.RelationField) ITTableindexDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tTableindexDo) Preload(fields ...field.RelationField) ITTableindexDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tTableindexDo) FirstOrInit() (*model.TTableindex, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTableindex), nil
	}
}

func (t tTableindexDo) FirstOrCreate() (*model.TTableindex, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TTableindex), nil
	}
}

func (t tTableindexDo) FindByPage(offset int, limit int) (result []*model.TTableindex, count int64, err error) {
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

func (t tTableindexDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tTableindexDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tTableindexDo) Delete(models ...*model.TTableindex) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tTableindexDo) withDO(do gen.Dao) *tTableindexDo {
	t.DO = *do.(*gen.DO)
	return t
}
