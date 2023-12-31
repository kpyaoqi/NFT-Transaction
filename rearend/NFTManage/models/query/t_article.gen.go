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

func newTArticle(db *gorm.DB) tArticle {
	_tArticle := tArticle{}

	_tArticle.tArticleDo.UseDB(db)
	_tArticle.tArticleDo.UseModel(&model.TArticle{})

	tableName := _tArticle.tArticleDo.TableName()
	_tArticle.ALL = field.NewAsterisk(tableName)
	_tArticle.ID = field.NewString(tableName, "id")
	_tArticle.Mintitle = field.NewString(tableName, "mintitle")
	_tArticle.Summary = field.NewString(tableName, "summary")
	_tArticle.Keywords = field.NewString(tableName, "keywords")
	_tArticle.Description = field.NewString(tableName, "description")
	_tArticle.Lookcount = field.NewInt64(tableName, "lookcount")
	_tArticle.CreatePerson = field.NewString(tableName, "createPerson")
	_tArticle.CreateDate = field.NewTime(tableName, "createDate")
	_tArticle.Status = field.NewInt64(tableName, "status")

	_tArticle.fillFieldMap()

	return _tArticle
}

type tArticle struct {
	tArticleDo tArticleDo

	ALL          field.Asterisk
	ID           field.String // ID
	Mintitle     field.String // 标题
	Summary      field.String // 摘要
	Keywords     field.String // 关键字
	Description  field.String // 描述
	Lookcount    field.Int64  // 打开次数
	CreatePerson field.String // 创建人
	CreateDate   field.Time   // 创建时间
	Status       field.Int64  // 状态  0未审核  1审核通过

	fieldMap map[string]field.Expr
}

func (t tArticle) Table(newTableName string) *tArticle {
	t.tArticleDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tArticle) As(alias string) *tArticle {
	t.tArticleDo.DO = *(t.tArticleDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tArticle) updateTableName(table string) *tArticle {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewString(table, "id")
	t.Mintitle = field.NewString(table, "mintitle")
	t.Summary = field.NewString(table, "summary")
	t.Keywords = field.NewString(table, "keywords")
	t.Description = field.NewString(table, "description")
	t.Lookcount = field.NewInt64(table, "lookcount")
	t.CreatePerson = field.NewString(table, "createPerson")
	t.CreateDate = field.NewTime(table, "createDate")
	t.Status = field.NewInt64(table, "status")

	t.fillFieldMap()

	return t
}

func (t *tArticle) WithContext(ctx context.Context) ITArticleDo { return t.tArticleDo.WithContext(ctx) }

func (t tArticle) TableName() string { return t.tArticleDo.TableName() }

func (t tArticle) Alias() string { return t.tArticleDo.Alias() }

func (t *tArticle) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tArticle) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 9)
	t.fieldMap["id"] = t.ID
	t.fieldMap["mintitle"] = t.Mintitle
	t.fieldMap["summary"] = t.Summary
	t.fieldMap["keywords"] = t.Keywords
	t.fieldMap["description"] = t.Description
	t.fieldMap["lookcount"] = t.Lookcount
	t.fieldMap["createPerson"] = t.CreatePerson
	t.fieldMap["createDate"] = t.CreateDate
	t.fieldMap["status"] = t.Status
}

func (t tArticle) clone(db *gorm.DB) tArticle {
	t.tArticleDo.ReplaceDB(db)
	return t
}

type tArticleDo struct{ gen.DO }

type ITArticleDo interface {
	gen.SubQuery
	Debug() ITArticleDo
	WithContext(ctx context.Context) ITArticleDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITArticleDo
	Not(conds ...gen.Condition) ITArticleDo
	Or(conds ...gen.Condition) ITArticleDo
	Select(conds ...field.Expr) ITArticleDo
	Where(conds ...gen.Condition) ITArticleDo
	Order(conds ...field.Expr) ITArticleDo
	Distinct(cols ...field.Expr) ITArticleDo
	Omit(cols ...field.Expr) ITArticleDo
	Join(table schema.Tabler, on ...field.Expr) ITArticleDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITArticleDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITArticleDo
	Group(cols ...field.Expr) ITArticleDo
	Having(conds ...gen.Condition) ITArticleDo
	Limit(limit int) ITArticleDo
	Offset(offset int) ITArticleDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITArticleDo
	Unscoped() ITArticleDo
	Create(values ...*model.TArticle) error
	CreateInBatches(values []*model.TArticle, batchSize int) error
	Save(values ...*model.TArticle) error
	First() (*model.TArticle, error)
	Take() (*model.TArticle, error)
	Last() (*model.TArticle, error)
	Find() ([]*model.TArticle, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TArticle, err error)
	FindInBatches(result *[]*model.TArticle, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TArticle) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITArticleDo
	Assign(attrs ...field.AssignExpr) ITArticleDo
	Joins(fields ...field.RelationField) ITArticleDo
	Preload(fields ...field.RelationField) ITArticleDo
	FirstOrInit() (*model.TArticle, error)
	FirstOrCreate() (*model.TArticle, error)
	FindByPage(offset int, limit int) (result []*model.TArticle, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITArticleDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tArticleDo) Debug() ITArticleDo {
	return t.withDO(t.DO.Debug())
}

func (t tArticleDo) WithContext(ctx context.Context) ITArticleDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tArticleDo) ReadDB() ITArticleDo {
	return t.Clauses(dbresolver.Read)
}

func (t tArticleDo) WriteDB() ITArticleDo {
	return t.Clauses(dbresolver.Write)
}

func (t tArticleDo) Clauses(conds ...clause.Expression) ITArticleDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tArticleDo) Returning(value interface{}, columns ...string) ITArticleDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tArticleDo) Not(conds ...gen.Condition) ITArticleDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tArticleDo) Or(conds ...gen.Condition) ITArticleDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tArticleDo) Select(conds ...field.Expr) ITArticleDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tArticleDo) Where(conds ...gen.Condition) ITArticleDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tArticleDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITArticleDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tArticleDo) Order(conds ...field.Expr) ITArticleDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tArticleDo) Distinct(cols ...field.Expr) ITArticleDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tArticleDo) Omit(cols ...field.Expr) ITArticleDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tArticleDo) Join(table schema.Tabler, on ...field.Expr) ITArticleDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tArticleDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITArticleDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tArticleDo) RightJoin(table schema.Tabler, on ...field.Expr) ITArticleDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tArticleDo) Group(cols ...field.Expr) ITArticleDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tArticleDo) Having(conds ...gen.Condition) ITArticleDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tArticleDo) Limit(limit int) ITArticleDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tArticleDo) Offset(offset int) ITArticleDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tArticleDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITArticleDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tArticleDo) Unscoped() ITArticleDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tArticleDo) Create(values ...*model.TArticle) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tArticleDo) CreateInBatches(values []*model.TArticle, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tArticleDo) Save(values ...*model.TArticle) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tArticleDo) First() (*model.TArticle, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TArticle), nil
	}
}

func (t tArticleDo) Take() (*model.TArticle, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TArticle), nil
	}
}

func (t tArticleDo) Last() (*model.TArticle, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TArticle), nil
	}
}

func (t tArticleDo) Find() ([]*model.TArticle, error) {
	result, err := t.DO.Find()
	return result.([]*model.TArticle), err
}

func (t tArticleDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TArticle, err error) {
	buf := make([]*model.TArticle, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tArticleDo) FindInBatches(result *[]*model.TArticle, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tArticleDo) Attrs(attrs ...field.AssignExpr) ITArticleDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tArticleDo) Assign(attrs ...field.AssignExpr) ITArticleDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tArticleDo) Joins(fields ...field.RelationField) ITArticleDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tArticleDo) Preload(fields ...field.RelationField) ITArticleDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tArticleDo) FirstOrInit() (*model.TArticle, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TArticle), nil
	}
}

func (t tArticleDo) FirstOrCreate() (*model.TArticle, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TArticle), nil
	}
}

func (t tArticleDo) FindByPage(offset int, limit int) (result []*model.TArticle, count int64, err error) {
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

func (t tArticleDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tArticleDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tArticleDo) Delete(models ...*model.TArticle) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tArticleDo) withDO(do gen.Dao) *tArticleDo {
	t.DO = *do.(*gen.DO)
	return t
}
