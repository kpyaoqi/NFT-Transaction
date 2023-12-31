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

func newTFwlogHistory2019(db *gorm.DB) tFwlogHistory2019 {
	_tFwlogHistory2019 := tFwlogHistory2019{}

	_tFwlogHistory2019.tFwlogHistory2019Do.UseDB(db)
	_tFwlogHistory2019.tFwlogHistory2019Do.UseModel(&model.TFwlogHistory2019{})

	tableName := _tFwlogHistory2019.tFwlogHistory2019Do.TableName()
	_tFwlogHistory2019.ALL = field.NewAsterisk(tableName)
	_tFwlogHistory2019.ID = field.NewString(tableName, "id")
	_tFwlogHistory2019.StartDate = field.NewTime(tableName, "startDate")
	_tFwlogHistory2019.StrDate = field.NewString(tableName, "strDate")
	_tFwlogHistory2019.Tomcat = field.NewString(tableName, "tomcat")
	_tFwlogHistory2019.UserCode = field.NewString(tableName, "userCode")
	_tFwlogHistory2019.UserName = field.NewString(tableName, "userName")
	_tFwlogHistory2019.SessionID = field.NewString(tableName, "sessionId")
	_tFwlogHistory2019.IP = field.NewString(tableName, "ip")
	_tFwlogHistory2019.FwURL = field.NewString(tableName, "fwUrl")
	_tFwlogHistory2019.MenuName = field.NewString(tableName, "menuName")
	_tFwlogHistory2019.Isqx = field.NewString(tableName, "isqx")
	_tFwlogHistory2019.Bak1 = field.NewString(tableName, "bak1")
	_tFwlogHistory2019.Bak2 = field.NewString(tableName, "bak2")
	_tFwlogHistory2019.Bak3 = field.NewString(tableName, "bak3")
	_tFwlogHistory2019.Bak4 = field.NewString(tableName, "bak4")
	_tFwlogHistory2019.Bak5 = field.NewString(tableName, "bak5")

	_tFwlogHistory2019.fillFieldMap()

	return _tFwlogHistory2019
}

type tFwlogHistory2019 struct {
	tFwlogHistory2019Do tFwlogHistory2019Do

	ALL       field.Asterisk
	ID        field.String // ID
	StartDate field.Time   // 访问时间
	StrDate   field.String // 访问时间(毫秒)
	Tomcat    field.String // Tomcat
	UserCode  field.String // 登录账号
	UserName  field.String // 姓名
	SessionID field.String // sessionId
	IP        field.String // IP
	FwURL     field.String // 访问菜单
	MenuName  field.String // 菜单名称
	Isqx      field.String // 是否有权限访问
	Bak1      field.String
	Bak2      field.String
	Bak3      field.String
	Bak4      field.String
	Bak5      field.String

	fieldMap map[string]field.Expr
}

func (t tFwlogHistory2019) Table(newTableName string) *tFwlogHistory2019 {
	t.tFwlogHistory2019Do.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tFwlogHistory2019) As(alias string) *tFwlogHistory2019 {
	t.tFwlogHistory2019Do.DO = *(t.tFwlogHistory2019Do.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tFwlogHistory2019) updateTableName(table string) *tFwlogHistory2019 {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewString(table, "id")
	t.StartDate = field.NewTime(table, "startDate")
	t.StrDate = field.NewString(table, "strDate")
	t.Tomcat = field.NewString(table, "tomcat")
	t.UserCode = field.NewString(table, "userCode")
	t.UserName = field.NewString(table, "userName")
	t.SessionID = field.NewString(table, "sessionId")
	t.IP = field.NewString(table, "ip")
	t.FwURL = field.NewString(table, "fwUrl")
	t.MenuName = field.NewString(table, "menuName")
	t.Isqx = field.NewString(table, "isqx")
	t.Bak1 = field.NewString(table, "bak1")
	t.Bak2 = field.NewString(table, "bak2")
	t.Bak3 = field.NewString(table, "bak3")
	t.Bak4 = field.NewString(table, "bak4")
	t.Bak5 = field.NewString(table, "bak5")

	t.fillFieldMap()

	return t
}

func (t *tFwlogHistory2019) WithContext(ctx context.Context) ITFwlogHistory2019Do {
	return t.tFwlogHistory2019Do.WithContext(ctx)
}

func (t tFwlogHistory2019) TableName() string { return t.tFwlogHistory2019Do.TableName() }

func (t tFwlogHistory2019) Alias() string { return t.tFwlogHistory2019Do.Alias() }

func (t *tFwlogHistory2019) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tFwlogHistory2019) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 16)
	t.fieldMap["id"] = t.ID
	t.fieldMap["startDate"] = t.StartDate
	t.fieldMap["strDate"] = t.StrDate
	t.fieldMap["tomcat"] = t.Tomcat
	t.fieldMap["userCode"] = t.UserCode
	t.fieldMap["userName"] = t.UserName
	t.fieldMap["sessionId"] = t.SessionID
	t.fieldMap["ip"] = t.IP
	t.fieldMap["fwUrl"] = t.FwURL
	t.fieldMap["menuName"] = t.MenuName
	t.fieldMap["isqx"] = t.Isqx
	t.fieldMap["bak1"] = t.Bak1
	t.fieldMap["bak2"] = t.Bak2
	t.fieldMap["bak3"] = t.Bak3
	t.fieldMap["bak4"] = t.Bak4
	t.fieldMap["bak5"] = t.Bak5
}

func (t tFwlogHistory2019) clone(db *gorm.DB) tFwlogHistory2019 {
	t.tFwlogHistory2019Do.ReplaceDB(db)
	return t
}

type tFwlogHistory2019Do struct{ gen.DO }

type ITFwlogHistory2019Do interface {
	gen.SubQuery
	Debug() ITFwlogHistory2019Do
	WithContext(ctx context.Context) ITFwlogHistory2019Do
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITFwlogHistory2019Do
	Not(conds ...gen.Condition) ITFwlogHistory2019Do
	Or(conds ...gen.Condition) ITFwlogHistory2019Do
	Select(conds ...field.Expr) ITFwlogHistory2019Do
	Where(conds ...gen.Condition) ITFwlogHistory2019Do
	Order(conds ...field.Expr) ITFwlogHistory2019Do
	Distinct(cols ...field.Expr) ITFwlogHistory2019Do
	Omit(cols ...field.Expr) ITFwlogHistory2019Do
	Join(table schema.Tabler, on ...field.Expr) ITFwlogHistory2019Do
	LeftJoin(table schema.Tabler, on ...field.Expr) ITFwlogHistory2019Do
	RightJoin(table schema.Tabler, on ...field.Expr) ITFwlogHistory2019Do
	Group(cols ...field.Expr) ITFwlogHistory2019Do
	Having(conds ...gen.Condition) ITFwlogHistory2019Do
	Limit(limit int) ITFwlogHistory2019Do
	Offset(offset int) ITFwlogHistory2019Do
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITFwlogHistory2019Do
	Unscoped() ITFwlogHistory2019Do
	Create(values ...*model.TFwlogHistory2019) error
	CreateInBatches(values []*model.TFwlogHistory2019, batchSize int) error
	Save(values ...*model.TFwlogHistory2019) error
	First() (*model.TFwlogHistory2019, error)
	Take() (*model.TFwlogHistory2019, error)
	Last() (*model.TFwlogHistory2019, error)
	Find() ([]*model.TFwlogHistory2019, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TFwlogHistory2019, err error)
	FindInBatches(result *[]*model.TFwlogHistory2019, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TFwlogHistory2019) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITFwlogHistory2019Do
	Assign(attrs ...field.AssignExpr) ITFwlogHistory2019Do
	Joins(fields ...field.RelationField) ITFwlogHistory2019Do
	Preload(fields ...field.RelationField) ITFwlogHistory2019Do
	FirstOrInit() (*model.TFwlogHistory2019, error)
	FirstOrCreate() (*model.TFwlogHistory2019, error)
	FindByPage(offset int, limit int) (result []*model.TFwlogHistory2019, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITFwlogHistory2019Do
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tFwlogHistory2019Do) Debug() ITFwlogHistory2019Do {
	return t.withDO(t.DO.Debug())
}

func (t tFwlogHistory2019Do) WithContext(ctx context.Context) ITFwlogHistory2019Do {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tFwlogHistory2019Do) ReadDB() ITFwlogHistory2019Do {
	return t.Clauses(dbresolver.Read)
}

func (t tFwlogHistory2019Do) WriteDB() ITFwlogHistory2019Do {
	return t.Clauses(dbresolver.Write)
}

func (t tFwlogHistory2019Do) Clauses(conds ...clause.Expression) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tFwlogHistory2019Do) Returning(value interface{}, columns ...string) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tFwlogHistory2019Do) Not(conds ...gen.Condition) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Not(conds...))
}

func (t tFwlogHistory2019Do) Or(conds ...gen.Condition) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Or(conds...))
}

func (t tFwlogHistory2019Do) Select(conds ...field.Expr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Select(conds...))
}

func (t tFwlogHistory2019Do) Where(conds ...gen.Condition) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Where(conds...))
}

func (t tFwlogHistory2019Do) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITFwlogHistory2019Do {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tFwlogHistory2019Do) Order(conds ...field.Expr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Order(conds...))
}

func (t tFwlogHistory2019Do) Distinct(cols ...field.Expr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tFwlogHistory2019Do) Omit(cols ...field.Expr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tFwlogHistory2019Do) Join(table schema.Tabler, on ...field.Expr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tFwlogHistory2019Do) LeftJoin(table schema.Tabler, on ...field.Expr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tFwlogHistory2019Do) RightJoin(table schema.Tabler, on ...field.Expr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tFwlogHistory2019Do) Group(cols ...field.Expr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Group(cols...))
}

func (t tFwlogHistory2019Do) Having(conds ...gen.Condition) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Having(conds...))
}

func (t tFwlogHistory2019Do) Limit(limit int) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Limit(limit))
}

func (t tFwlogHistory2019Do) Offset(offset int) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Offset(offset))
}

func (t tFwlogHistory2019Do) Scopes(funcs ...func(gen.Dao) gen.Dao) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tFwlogHistory2019Do) Unscoped() ITFwlogHistory2019Do {
	return t.withDO(t.DO.Unscoped())
}

func (t tFwlogHistory2019Do) Create(values ...*model.TFwlogHistory2019) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tFwlogHistory2019Do) CreateInBatches(values []*model.TFwlogHistory2019, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tFwlogHistory2019Do) Save(values ...*model.TFwlogHistory2019) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tFwlogHistory2019Do) First() (*model.TFwlogHistory2019, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFwlogHistory2019), nil
	}
}

func (t tFwlogHistory2019Do) Take() (*model.TFwlogHistory2019, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFwlogHistory2019), nil
	}
}

func (t tFwlogHistory2019Do) Last() (*model.TFwlogHistory2019, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFwlogHistory2019), nil
	}
}

func (t tFwlogHistory2019Do) Find() ([]*model.TFwlogHistory2019, error) {
	result, err := t.DO.Find()
	return result.([]*model.TFwlogHistory2019), err
}

func (t tFwlogHistory2019Do) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TFwlogHistory2019, err error) {
	buf := make([]*model.TFwlogHistory2019, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tFwlogHistory2019Do) FindInBatches(result *[]*model.TFwlogHistory2019, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tFwlogHistory2019Do) Attrs(attrs ...field.AssignExpr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tFwlogHistory2019Do) Assign(attrs ...field.AssignExpr) ITFwlogHistory2019Do {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tFwlogHistory2019Do) Joins(fields ...field.RelationField) ITFwlogHistory2019Do {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tFwlogHistory2019Do) Preload(fields ...field.RelationField) ITFwlogHistory2019Do {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tFwlogHistory2019Do) FirstOrInit() (*model.TFwlogHistory2019, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFwlogHistory2019), nil
	}
}

func (t tFwlogHistory2019Do) FirstOrCreate() (*model.TFwlogHistory2019, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TFwlogHistory2019), nil
	}
}

func (t tFwlogHistory2019Do) FindByPage(offset int, limit int) (result []*model.TFwlogHistory2019, count int64, err error) {
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

func (t tFwlogHistory2019Do) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tFwlogHistory2019Do) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tFwlogHistory2019Do) Delete(models ...*model.TFwlogHistory2019) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tFwlogHistory2019Do) withDO(do gen.Dao) *tFwlogHistory2019Do {
	t.DO = *do.(*gen.DO)
	return t
}
