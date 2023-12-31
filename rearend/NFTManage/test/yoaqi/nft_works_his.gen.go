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

func newNftWorksHi(db *gorm.DB) nftWorksHi {
	_nftWorksHi := nftWorksHi{}

	_nftWorksHi.nftWorksHiDo.UseDB(db)
	_nftWorksHi.nftWorksHiDo.UseModel(&model.NftWorksHi{})

	tableName := _nftWorksHi.nftWorksHiDo.TableName()
	_nftWorksHi.ALL = field.NewAsterisk(tableName)
	_nftWorksHi.ID = field.NewString(tableName, "id")
	_nftWorksHi.WorksID = field.NewString(tableName, "worksId")
	_nftWorksHi.Type = field.NewInt64(tableName, "type")
	_nftWorksHi.ChainPlatCert = field.NewString(tableName, "chainPlatCert")
	_nftWorksHi.Num = field.NewInt64(tableName, "num")
	_nftWorksHi.Price = field.NewFloat64(tableName, "price")
	_nftWorksHi.FromUser = field.NewString(tableName, "fromUser")
	_nftWorksHi.ToUser = field.NewString(tableName, "toUser")
	_nftWorksHi.TradeTime = field.NewTime(tableName, "tradeTime")
	_nftWorksHi.CreateTime = field.NewTime(tableName, "createTime")
	_nftWorksHi.UpdateTime = field.NewTime(tableName, "updateTime")
	_nftWorksHi.Bak1 = field.NewString(tableName, "bak1")
	_nftWorksHi.Bak2 = field.NewString(tableName, "bak2")
	_nftWorksHi.Bak3 = field.NewString(tableName, "bak3")

	_nftWorksHi.fillFieldMap()

	return _nftWorksHi
}

type nftWorksHi struct {
	nftWorksHiDo nftWorksHiDo

	ALL           field.Asterisk
	ID            field.String  // 唯一标识
	WorksID       field.String  // 作品标识
	Type          field.Int64   // 类型(0:售卖,1:转移)
	ChainPlatCert field.String  // 在不同链平台上链,认证后的数据json(使用区块链浏览器查询)
	Num           field.Int64   // 作品数量
	Price         field.Float64 // 价格
	FromUser      field.String  // 卖方
	ToUser        field.String  // 买方
	TradeTime     field.Time    // 交易时间
	CreateTime    field.Time
	UpdateTime    field.Time
	Bak1          field.String
	Bak2          field.String
	Bak3          field.String

	fieldMap map[string]field.Expr
}

func (n nftWorksHi) Table(newTableName string) *nftWorksHi {
	n.nftWorksHiDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n nftWorksHi) As(alias string) *nftWorksHi {
	n.nftWorksHiDo.DO = *(n.nftWorksHiDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *nftWorksHi) updateTableName(table string) *nftWorksHi {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewString(table, "id")
	n.WorksID = field.NewString(table, "worksId")
	n.Type = field.NewInt64(table, "type")
	n.ChainPlatCert = field.NewString(table, "chainPlatCert")
	n.Num = field.NewInt64(table, "num")
	n.Price = field.NewFloat64(table, "price")
	n.FromUser = field.NewString(table, "fromUser")
	n.ToUser = field.NewString(table, "toUser")
	n.TradeTime = field.NewTime(table, "tradeTime")
	n.CreateTime = field.NewTime(table, "createTime")
	n.UpdateTime = field.NewTime(table, "updateTime")
	n.Bak1 = field.NewString(table, "bak1")
	n.Bak2 = field.NewString(table, "bak2")
	n.Bak3 = field.NewString(table, "bak3")

	n.fillFieldMap()

	return n
}

func (n *nftWorksHi) WithContext(ctx context.Context) INftWorksHiDo {
	return n.nftWorksHiDo.WithContext(ctx)
}

func (n nftWorksHi) TableName() string { return n.nftWorksHiDo.TableName() }

func (n nftWorksHi) Alias() string { return n.nftWorksHiDo.Alias() }

func (n *nftWorksHi) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *nftWorksHi) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 14)
	n.fieldMap["id"] = n.ID
	n.fieldMap["worksId"] = n.WorksID
	n.fieldMap["type"] = n.Type
	n.fieldMap["chainPlatCert"] = n.ChainPlatCert
	n.fieldMap["num"] = n.Num
	n.fieldMap["price"] = n.Price
	n.fieldMap["fromUser"] = n.FromUser
	n.fieldMap["toUser"] = n.ToUser
	n.fieldMap["tradeTime"] = n.TradeTime
	n.fieldMap["createTime"] = n.CreateTime
	n.fieldMap["updateTime"] = n.UpdateTime
	n.fieldMap["bak1"] = n.Bak1
	n.fieldMap["bak2"] = n.Bak2
	n.fieldMap["bak3"] = n.Bak3
}

func (n nftWorksHi) clone(db *gorm.DB) nftWorksHi {
	n.nftWorksHiDo.ReplaceDB(db)
	return n
}

type nftWorksHiDo struct{ gen.DO }

type INftWorksHiDo interface {
	gen.SubQuery
	Debug() INftWorksHiDo
	WithContext(ctx context.Context) INftWorksHiDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INftWorksHiDo
	Not(conds ...gen.Condition) INftWorksHiDo
	Or(conds ...gen.Condition) INftWorksHiDo
	Select(conds ...field.Expr) INftWorksHiDo
	Where(conds ...gen.Condition) INftWorksHiDo
	Order(conds ...field.Expr) INftWorksHiDo
	Distinct(cols ...field.Expr) INftWorksHiDo
	Omit(cols ...field.Expr) INftWorksHiDo
	Join(table schema.Tabler, on ...field.Expr) INftWorksHiDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INftWorksHiDo
	RightJoin(table schema.Tabler, on ...field.Expr) INftWorksHiDo
	Group(cols ...field.Expr) INftWorksHiDo
	Having(conds ...gen.Condition) INftWorksHiDo
	Limit(limit int) INftWorksHiDo
	Offset(offset int) INftWorksHiDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INftWorksHiDo
	Unscoped() INftWorksHiDo
	Create(values ...*model.NftWorksHi) error
	CreateInBatches(values []*model.NftWorksHi, batchSize int) error
	Save(values ...*model.NftWorksHi) error
	First() (*model.NftWorksHi, error)
	Take() (*model.NftWorksHi, error)
	Last() (*model.NftWorksHi, error)
	Find() ([]*model.NftWorksHi, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NftWorksHi, err error)
	FindInBatches(result *[]*model.NftWorksHi, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NftWorksHi) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INftWorksHiDo
	Assign(attrs ...field.AssignExpr) INftWorksHiDo
	Joins(fields ...field.RelationField) INftWorksHiDo
	Preload(fields ...field.RelationField) INftWorksHiDo
	FirstOrInit() (*model.NftWorksHi, error)
	FirstOrCreate() (*model.NftWorksHi, error)
	FindByPage(offset int, limit int) (result []*model.NftWorksHi, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INftWorksHiDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n nftWorksHiDo) Debug() INftWorksHiDo {
	return n.withDO(n.DO.Debug())
}

func (n nftWorksHiDo) WithContext(ctx context.Context) INftWorksHiDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n nftWorksHiDo) ReadDB() INftWorksHiDo {
	return n.Clauses(dbresolver.Read)
}

func (n nftWorksHiDo) WriteDB() INftWorksHiDo {
	return n.Clauses(dbresolver.Write)
}

func (n nftWorksHiDo) Clauses(conds ...clause.Expression) INftWorksHiDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n nftWorksHiDo) Returning(value interface{}, columns ...string) INftWorksHiDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n nftWorksHiDo) Not(conds ...gen.Condition) INftWorksHiDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n nftWorksHiDo) Or(conds ...gen.Condition) INftWorksHiDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n nftWorksHiDo) Select(conds ...field.Expr) INftWorksHiDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n nftWorksHiDo) Where(conds ...gen.Condition) INftWorksHiDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n nftWorksHiDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) INftWorksHiDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n nftWorksHiDo) Order(conds ...field.Expr) INftWorksHiDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n nftWorksHiDo) Distinct(cols ...field.Expr) INftWorksHiDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n nftWorksHiDo) Omit(cols ...field.Expr) INftWorksHiDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n nftWorksHiDo) Join(table schema.Tabler, on ...field.Expr) INftWorksHiDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n nftWorksHiDo) LeftJoin(table schema.Tabler, on ...field.Expr) INftWorksHiDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n nftWorksHiDo) RightJoin(table schema.Tabler, on ...field.Expr) INftWorksHiDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n nftWorksHiDo) Group(cols ...field.Expr) INftWorksHiDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n nftWorksHiDo) Having(conds ...gen.Condition) INftWorksHiDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n nftWorksHiDo) Limit(limit int) INftWorksHiDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n nftWorksHiDo) Offset(offset int) INftWorksHiDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n nftWorksHiDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INftWorksHiDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n nftWorksHiDo) Unscoped() INftWorksHiDo {
	return n.withDO(n.DO.Unscoped())
}

func (n nftWorksHiDo) Create(values ...*model.NftWorksHi) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n nftWorksHiDo) CreateInBatches(values []*model.NftWorksHi, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n nftWorksHiDo) Save(values ...*model.NftWorksHi) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n nftWorksHiDo) First() (*model.NftWorksHi, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftWorksHi), nil
	}
}

func (n nftWorksHiDo) Take() (*model.NftWorksHi, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftWorksHi), nil
	}
}

func (n nftWorksHiDo) Last() (*model.NftWorksHi, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftWorksHi), nil
	}
}

func (n nftWorksHiDo) Find() ([]*model.NftWorksHi, error) {
	result, err := n.DO.Find()
	return result.([]*model.NftWorksHi), err
}

func (n nftWorksHiDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NftWorksHi, err error) {
	buf := make([]*model.NftWorksHi, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n nftWorksHiDo) FindInBatches(result *[]*model.NftWorksHi, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n nftWorksHiDo) Attrs(attrs ...field.AssignExpr) INftWorksHiDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n nftWorksHiDo) Assign(attrs ...field.AssignExpr) INftWorksHiDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n nftWorksHiDo) Joins(fields ...field.RelationField) INftWorksHiDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n nftWorksHiDo) Preload(fields ...field.RelationField) INftWorksHiDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n nftWorksHiDo) FirstOrInit() (*model.NftWorksHi, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftWorksHi), nil
	}
}

func (n nftWorksHiDo) FirstOrCreate() (*model.NftWorksHi, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftWorksHi), nil
	}
}

func (n nftWorksHiDo) FindByPage(offset int, limit int) (result []*model.NftWorksHi, count int64, err error) {
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

func (n nftWorksHiDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n nftWorksHiDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n nftWorksHiDo) Delete(models ...*model.NftWorksHi) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *nftWorksHiDo) withDO(do gen.Dao) *nftWorksHiDo {
	n.DO = *do.(*gen.DO)
	return n
}
