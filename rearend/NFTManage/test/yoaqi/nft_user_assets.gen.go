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

func newNftUserAsset(db *gorm.DB) nftUserAsset {
	_nftUserAsset := nftUserAsset{}

	_nftUserAsset.nftUserAssetDo.UseDB(db)
	_nftUserAsset.nftUserAssetDo.UseModel(&model.NftUserAsset{})

	tableName := _nftUserAsset.nftUserAssetDo.TableName()
	_nftUserAsset.ALL = field.NewAsterisk(tableName)
	_nftUserAsset.ID = field.NewString(tableName, "id")
	_nftUserAsset.UserID = field.NewString(tableName, "userId")
	_nftUserAsset.AssetsID = field.NewString(tableName, "assetsId")
	_nftUserAsset.Type = field.NewInt64(tableName, "type")
	_nftUserAsset.Origin = field.NewInt64(tableName, "origin")
	_nftUserAsset.SellState = field.NewInt64(tableName, "sellState")
	_nftUserAsset.CreateTime = field.NewTime(tableName, "createTime")
	_nftUserAsset.UpdateTime = field.NewTime(tableName, "updateTime")
	_nftUserAsset.Bak1 = field.NewString(tableName, "bak1")
	_nftUserAsset.Bak2 = field.NewString(tableName, "bak2")
	_nftUserAsset.Bak3 = field.NewString(tableName, "bak3")

	_nftUserAsset.fillFieldMap()

	return _nftUserAsset
}

type nftUserAsset struct {
	nftUserAssetDo nftUserAssetDo

	ALL        field.Asterisk
	ID         field.String // 唯一标识
	UserID     field.String // 用户标识
	AssetsID   field.String // 资源标识
	Type       field.Int64  // 资源类型0:合集,1:作品
	Origin     field.Int64  // 资产来源(0:自己创作,1购买其他作者)
	SellState  field.Int64  // 出售状态,0:不出售,1:售卖中,2:已生成订单,3:已删除
	CreateTime field.Time
	UpdateTime field.Time
	Bak1       field.String
	Bak2       field.String
	Bak3       field.String

	fieldMap map[string]field.Expr
}

func (n nftUserAsset) Table(newTableName string) *nftUserAsset {
	n.nftUserAssetDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n nftUserAsset) As(alias string) *nftUserAsset {
	n.nftUserAssetDo.DO = *(n.nftUserAssetDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *nftUserAsset) updateTableName(table string) *nftUserAsset {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewString(table, "id")
	n.UserID = field.NewString(table, "userId")
	n.AssetsID = field.NewString(table, "assetsId")
	n.Type = field.NewInt64(table, "type")
	n.Origin = field.NewInt64(table, "origin")
	n.SellState = field.NewInt64(table, "sellState")
	n.CreateTime = field.NewTime(table, "createTime")
	n.UpdateTime = field.NewTime(table, "updateTime")
	n.Bak1 = field.NewString(table, "bak1")
	n.Bak2 = field.NewString(table, "bak2")
	n.Bak3 = field.NewString(table, "bak3")

	n.fillFieldMap()

	return n
}

func (n *nftUserAsset) WithContext(ctx context.Context) INftUserAssetDo {
	return n.nftUserAssetDo.WithContext(ctx)
}

func (n nftUserAsset) TableName() string { return n.nftUserAssetDo.TableName() }

func (n nftUserAsset) Alias() string { return n.nftUserAssetDo.Alias() }

func (n *nftUserAsset) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *nftUserAsset) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 11)
	n.fieldMap["id"] = n.ID
	n.fieldMap["userId"] = n.UserID
	n.fieldMap["assetsId"] = n.AssetsID
	n.fieldMap["type"] = n.Type
	n.fieldMap["origin"] = n.Origin
	n.fieldMap["sellState"] = n.SellState
	n.fieldMap["createTime"] = n.CreateTime
	n.fieldMap["updateTime"] = n.UpdateTime
	n.fieldMap["bak1"] = n.Bak1
	n.fieldMap["bak2"] = n.Bak2
	n.fieldMap["bak3"] = n.Bak3
}

func (n nftUserAsset) clone(db *gorm.DB) nftUserAsset {
	n.nftUserAssetDo.ReplaceDB(db)
	return n
}

type nftUserAssetDo struct{ gen.DO }

type INftUserAssetDo interface {
	gen.SubQuery
	Debug() INftUserAssetDo
	WithContext(ctx context.Context) INftUserAssetDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INftUserAssetDo
	Not(conds ...gen.Condition) INftUserAssetDo
	Or(conds ...gen.Condition) INftUserAssetDo
	Select(conds ...field.Expr) INftUserAssetDo
	Where(conds ...gen.Condition) INftUserAssetDo
	Order(conds ...field.Expr) INftUserAssetDo
	Distinct(cols ...field.Expr) INftUserAssetDo
	Omit(cols ...field.Expr) INftUserAssetDo
	Join(table schema.Tabler, on ...field.Expr) INftUserAssetDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INftUserAssetDo
	RightJoin(table schema.Tabler, on ...field.Expr) INftUserAssetDo
	Group(cols ...field.Expr) INftUserAssetDo
	Having(conds ...gen.Condition) INftUserAssetDo
	Limit(limit int) INftUserAssetDo
	Offset(offset int) INftUserAssetDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INftUserAssetDo
	Unscoped() INftUserAssetDo
	Create(values ...*model.NftUserAsset) error
	CreateInBatches(values []*model.NftUserAsset, batchSize int) error
	Save(values ...*model.NftUserAsset) error
	First() (*model.NftUserAsset, error)
	Take() (*model.NftUserAsset, error)
	Last() (*model.NftUserAsset, error)
	Find() ([]*model.NftUserAsset, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NftUserAsset, err error)
	FindInBatches(result *[]*model.NftUserAsset, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NftUserAsset) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INftUserAssetDo
	Assign(attrs ...field.AssignExpr) INftUserAssetDo
	Joins(fields ...field.RelationField) INftUserAssetDo
	Preload(fields ...field.RelationField) INftUserAssetDo
	FirstOrInit() (*model.NftUserAsset, error)
	FirstOrCreate() (*model.NftUserAsset, error)
	FindByPage(offset int, limit int) (result []*model.NftUserAsset, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INftUserAssetDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n nftUserAssetDo) Debug() INftUserAssetDo {
	return n.withDO(n.DO.Debug())
}

func (n nftUserAssetDo) WithContext(ctx context.Context) INftUserAssetDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n nftUserAssetDo) ReadDB() INftUserAssetDo {
	return n.Clauses(dbresolver.Read)
}

func (n nftUserAssetDo) WriteDB() INftUserAssetDo {
	return n.Clauses(dbresolver.Write)
}

func (n nftUserAssetDo) Clauses(conds ...clause.Expression) INftUserAssetDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n nftUserAssetDo) Returning(value interface{}, columns ...string) INftUserAssetDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n nftUserAssetDo) Not(conds ...gen.Condition) INftUserAssetDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n nftUserAssetDo) Or(conds ...gen.Condition) INftUserAssetDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n nftUserAssetDo) Select(conds ...field.Expr) INftUserAssetDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n nftUserAssetDo) Where(conds ...gen.Condition) INftUserAssetDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n nftUserAssetDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) INftUserAssetDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n nftUserAssetDo) Order(conds ...field.Expr) INftUserAssetDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n nftUserAssetDo) Distinct(cols ...field.Expr) INftUserAssetDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n nftUserAssetDo) Omit(cols ...field.Expr) INftUserAssetDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n nftUserAssetDo) Join(table schema.Tabler, on ...field.Expr) INftUserAssetDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n nftUserAssetDo) LeftJoin(table schema.Tabler, on ...field.Expr) INftUserAssetDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n nftUserAssetDo) RightJoin(table schema.Tabler, on ...field.Expr) INftUserAssetDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n nftUserAssetDo) Group(cols ...field.Expr) INftUserAssetDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n nftUserAssetDo) Having(conds ...gen.Condition) INftUserAssetDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n nftUserAssetDo) Limit(limit int) INftUserAssetDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n nftUserAssetDo) Offset(offset int) INftUserAssetDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n nftUserAssetDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INftUserAssetDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n nftUserAssetDo) Unscoped() INftUserAssetDo {
	return n.withDO(n.DO.Unscoped())
}

func (n nftUserAssetDo) Create(values ...*model.NftUserAsset) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n nftUserAssetDo) CreateInBatches(values []*model.NftUserAsset, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n nftUserAssetDo) Save(values ...*model.NftUserAsset) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n nftUserAssetDo) First() (*model.NftUserAsset, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserAsset), nil
	}
}

func (n nftUserAssetDo) Take() (*model.NftUserAsset, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserAsset), nil
	}
}

func (n nftUserAssetDo) Last() (*model.NftUserAsset, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserAsset), nil
	}
}

func (n nftUserAssetDo) Find() ([]*model.NftUserAsset, error) {
	result, err := n.DO.Find()
	return result.([]*model.NftUserAsset), err
}

func (n nftUserAssetDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NftUserAsset, err error) {
	buf := make([]*model.NftUserAsset, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n nftUserAssetDo) FindInBatches(result *[]*model.NftUserAsset, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n nftUserAssetDo) Attrs(attrs ...field.AssignExpr) INftUserAssetDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n nftUserAssetDo) Assign(attrs ...field.AssignExpr) INftUserAssetDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n nftUserAssetDo) Joins(fields ...field.RelationField) INftUserAssetDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n nftUserAssetDo) Preload(fields ...field.RelationField) INftUserAssetDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n nftUserAssetDo) FirstOrInit() (*model.NftUserAsset, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserAsset), nil
	}
}

func (n nftUserAssetDo) FirstOrCreate() (*model.NftUserAsset, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserAsset), nil
	}
}

func (n nftUserAssetDo) FindByPage(offset int, limit int) (result []*model.NftUserAsset, count int64, err error) {
	result, err = n.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = n.Offset(-1).Limit(-1).Count()
	return
}

func (n nftUserAssetDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n nftUserAssetDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n nftUserAssetDo) Delete(models ...*model.NftUserAsset) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *nftUserAssetDo) withDO(do gen.Dao) *nftUserAssetDo {
	n.DO = *do.(*gen.DO)
	return n
}