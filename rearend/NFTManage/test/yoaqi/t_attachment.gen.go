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

func newTAttachment(db *gorm.DB) tAttachment {
	_tAttachment := tAttachment{}

	_tAttachment.tAttachmentDo.UseDB(db)
	_tAttachment.tAttachmentDo.UseModel(&model.TAttachment{})

	tableName := _tAttachment.tAttachmentDo.TableName()
	_tAttachment.ALL = field.NewAsterisk(tableName)
	_tAttachment.ID = field.NewString(tableName, "id")
	_tAttachment.OrgID = field.NewString(tableName, "orgId")
	_tAttachment.OrgTypePathKey = field.NewString(tableName, "orgTypePathKey")
	_tAttachment.BusinessID = field.NewString(tableName, "businessId")
	_tAttachment.AttachmentType = field.NewInt64(tableName, "attachmentType")
	_tAttachment.FileName = field.NewString(tableName, "fileName")
	_tAttachment.FileURL = field.NewString(tableName, "fileURL")
	_tAttachment.Suffix = field.NewString(tableName, "suffix")
	_tAttachment.FileSize = field.NewInt64(tableName, "fileSize")
	_tAttachment.LastDownTime = field.NewTime(tableName, "lastDownTime")
	_tAttachment.Sortno = field.NewInt64(tableName, "sortno")
	_tAttachment.Active = field.NewInt64(tableName, "active")
	_tAttachment.CreateUser = field.NewString(tableName, "createUser")
	_tAttachment.CreateTime = field.NewTime(tableName, "createTime")
	_tAttachment.UpdateUser = field.NewString(tableName, "updateUser")
	_tAttachment.UpdateTime = field.NewTime(tableName, "updateTime")
	_tAttachment.Md5Code = field.NewString(tableName, "md5Code")

	_tAttachment.fillFieldMap()

	return _tAttachment
}

type tAttachment struct {
	tAttachmentDo tAttachmentDo

	ALL            field.Asterisk
	ID             field.String // id主键
	OrgID          field.String // 部门id
	OrgTypePathKey field.String // URL路径中的部门类型,例如 URL路径中的 kjj
	BusinessID     field.String // 业务ID,用于业务关联查询
	AttachmentType field.Int64  // 附件类型,1.政策附件.2.企业认证文件3.专家认证文件.4.企业个人认证文件.0.其他文件
	FileName       field.String // 附件名称
	FileURL        field.String // 路径
	Suffix         field.String // 文件后缀
	FileSize       field.Int64  // 文件大小,单位K
	LastDownTime   field.Time   // 最后下载时间
	Sortno         field.Int64  // 排序,查询时倒叙排列
	Active         field.Int64  // 是否有效(0否,1是)
	CreateUser     field.String // 创建者
	CreateTime     field.Time   // 上传时间
	UpdateUser     field.String // 更新者
	UpdateTime     field.Time   // 更新时间
	Md5Code        field.String // 文件的md5值

	fieldMap map[string]field.Expr
}

func (t tAttachment) Table(newTableName string) *tAttachment {
	t.tAttachmentDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t tAttachment) As(alias string) *tAttachment {
	t.tAttachmentDo.DO = *(t.tAttachmentDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *tAttachment) updateTableName(table string) *tAttachment {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewString(table, "id")
	t.OrgID = field.NewString(table, "orgId")
	t.OrgTypePathKey = field.NewString(table, "orgTypePathKey")
	t.BusinessID = field.NewString(table, "businessId")
	t.AttachmentType = field.NewInt64(table, "attachmentType")
	t.FileName = field.NewString(table, "fileName")
	t.FileURL = field.NewString(table, "fileURL")
	t.Suffix = field.NewString(table, "suffix")
	t.FileSize = field.NewInt64(table, "fileSize")
	t.LastDownTime = field.NewTime(table, "lastDownTime")
	t.Sortno = field.NewInt64(table, "sortno")
	t.Active = field.NewInt64(table, "active")
	t.CreateUser = field.NewString(table, "createUser")
	t.CreateTime = field.NewTime(table, "createTime")
	t.UpdateUser = field.NewString(table, "updateUser")
	t.UpdateTime = field.NewTime(table, "updateTime")
	t.Md5Code = field.NewString(table, "md5Code")

	t.fillFieldMap()

	return t
}

func (t *tAttachment) WithContext(ctx context.Context) ITAttachmentDo {
	return t.tAttachmentDo.WithContext(ctx)
}

func (t tAttachment) TableName() string { return t.tAttachmentDo.TableName() }

func (t tAttachment) Alias() string { return t.tAttachmentDo.Alias() }

func (t *tAttachment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *tAttachment) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 17)
	t.fieldMap["id"] = t.ID
	t.fieldMap["orgId"] = t.OrgID
	t.fieldMap["orgTypePathKey"] = t.OrgTypePathKey
	t.fieldMap["businessId"] = t.BusinessID
	t.fieldMap["attachmentType"] = t.AttachmentType
	t.fieldMap["fileName"] = t.FileName
	t.fieldMap["fileURL"] = t.FileURL
	t.fieldMap["suffix"] = t.Suffix
	t.fieldMap["fileSize"] = t.FileSize
	t.fieldMap["lastDownTime"] = t.LastDownTime
	t.fieldMap["sortno"] = t.Sortno
	t.fieldMap["active"] = t.Active
	t.fieldMap["createUser"] = t.CreateUser
	t.fieldMap["createTime"] = t.CreateTime
	t.fieldMap["updateUser"] = t.UpdateUser
	t.fieldMap["updateTime"] = t.UpdateTime
	t.fieldMap["md5Code"] = t.Md5Code
}

func (t tAttachment) clone(db *gorm.DB) tAttachment {
	t.tAttachmentDo.ReplaceDB(db)
	return t
}

type tAttachmentDo struct{ gen.DO }

type ITAttachmentDo interface {
	gen.SubQuery
	Debug() ITAttachmentDo
	WithContext(ctx context.Context) ITAttachmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITAttachmentDo
	Not(conds ...gen.Condition) ITAttachmentDo
	Or(conds ...gen.Condition) ITAttachmentDo
	Select(conds ...field.Expr) ITAttachmentDo
	Where(conds ...gen.Condition) ITAttachmentDo
	Order(conds ...field.Expr) ITAttachmentDo
	Distinct(cols ...field.Expr) ITAttachmentDo
	Omit(cols ...field.Expr) ITAttachmentDo
	Join(table schema.Tabler, on ...field.Expr) ITAttachmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITAttachmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITAttachmentDo
	Group(cols ...field.Expr) ITAttachmentDo
	Having(conds ...gen.Condition) ITAttachmentDo
	Limit(limit int) ITAttachmentDo
	Offset(offset int) ITAttachmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITAttachmentDo
	Unscoped() ITAttachmentDo
	Create(values ...*model.TAttachment) error
	CreateInBatches(values []*model.TAttachment, batchSize int) error
	Save(values ...*model.TAttachment) error
	First() (*model.TAttachment, error)
	Take() (*model.TAttachment, error)
	Last() (*model.TAttachment, error)
	Find() ([]*model.TAttachment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TAttachment, err error)
	FindInBatches(result *[]*model.TAttachment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.TAttachment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITAttachmentDo
	Assign(attrs ...field.AssignExpr) ITAttachmentDo
	Joins(fields ...field.RelationField) ITAttachmentDo
	Preload(fields ...field.RelationField) ITAttachmentDo
	FirstOrInit() (*model.TAttachment, error)
	FirstOrCreate() (*model.TAttachment, error)
	FindByPage(offset int, limit int) (result []*model.TAttachment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITAttachmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t tAttachmentDo) Debug() ITAttachmentDo {
	return t.withDO(t.DO.Debug())
}

func (t tAttachmentDo) WithContext(ctx context.Context) ITAttachmentDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t tAttachmentDo) ReadDB() ITAttachmentDo {
	return t.Clauses(dbresolver.Read)
}

func (t tAttachmentDo) WriteDB() ITAttachmentDo {
	return t.Clauses(dbresolver.Write)
}

func (t tAttachmentDo) Clauses(conds ...clause.Expression) ITAttachmentDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t tAttachmentDo) Returning(value interface{}, columns ...string) ITAttachmentDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t tAttachmentDo) Not(conds ...gen.Condition) ITAttachmentDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t tAttachmentDo) Or(conds ...gen.Condition) ITAttachmentDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t tAttachmentDo) Select(conds ...field.Expr) ITAttachmentDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t tAttachmentDo) Where(conds ...gen.Condition) ITAttachmentDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t tAttachmentDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) ITAttachmentDo {
	return t.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (t tAttachmentDo) Order(conds ...field.Expr) ITAttachmentDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t tAttachmentDo) Distinct(cols ...field.Expr) ITAttachmentDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t tAttachmentDo) Omit(cols ...field.Expr) ITAttachmentDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t tAttachmentDo) Join(table schema.Tabler, on ...field.Expr) ITAttachmentDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t tAttachmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITAttachmentDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t tAttachmentDo) RightJoin(table schema.Tabler, on ...field.Expr) ITAttachmentDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t tAttachmentDo) Group(cols ...field.Expr) ITAttachmentDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t tAttachmentDo) Having(conds ...gen.Condition) ITAttachmentDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t tAttachmentDo) Limit(limit int) ITAttachmentDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t tAttachmentDo) Offset(offset int) ITAttachmentDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t tAttachmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITAttachmentDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t tAttachmentDo) Unscoped() ITAttachmentDo {
	return t.withDO(t.DO.Unscoped())
}

func (t tAttachmentDo) Create(values ...*model.TAttachment) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t tAttachmentDo) CreateInBatches(values []*model.TAttachment, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t tAttachmentDo) Save(values ...*model.TAttachment) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t tAttachmentDo) First() (*model.TAttachment, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAttachment), nil
	}
}

func (t tAttachmentDo) Take() (*model.TAttachment, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAttachment), nil
	}
}

func (t tAttachmentDo) Last() (*model.TAttachment, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAttachment), nil
	}
}

func (t tAttachmentDo) Find() ([]*model.TAttachment, error) {
	result, err := t.DO.Find()
	return result.([]*model.TAttachment), err
}

func (t tAttachmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.TAttachment, err error) {
	buf := make([]*model.TAttachment, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t tAttachmentDo) FindInBatches(result *[]*model.TAttachment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t tAttachmentDo) Attrs(attrs ...field.AssignExpr) ITAttachmentDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t tAttachmentDo) Assign(attrs ...field.AssignExpr) ITAttachmentDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t tAttachmentDo) Joins(fields ...field.RelationField) ITAttachmentDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t tAttachmentDo) Preload(fields ...field.RelationField) ITAttachmentDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t tAttachmentDo) FirstOrInit() (*model.TAttachment, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAttachment), nil
	}
}

func (t tAttachmentDo) FirstOrCreate() (*model.TAttachment, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.TAttachment), nil
	}
}

func (t tAttachmentDo) FindByPage(offset int, limit int) (result []*model.TAttachment, count int64, err error) {
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

func (t tAttachmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t tAttachmentDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t tAttachmentDo) Delete(models ...*model.TAttachment) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *tAttachmentDo) withDO(do gen.Dao) *tAttachmentDo {
	t.DO = *do.(*gen.DO)
	return t
}