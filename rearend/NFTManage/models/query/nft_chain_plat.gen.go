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

func newNftChainPlat(db *gorm.DB) nftChainPlat {
	_nftChainPlat := nftChainPlat{}

	_nftChainPlat.nftChainPlatDo.UseDB(db)
	_nftChainPlat.nftChainPlatDo.UseModel(&model.NftChainPlat{})

	tableName := _nftChainPlat.nftChainPlatDo.TableName()
	_nftChainPlat.ALL = field.NewAsterisk(tableName)
	_nftChainPlat.ID = field.NewString(tableName, "id")
	_nftChainPlat.ChainName = field.NewString(tableName, "chainName")
	_nftChainPlat.EnglishName = field.NewString(tableName, "englishName")
	_nftChainPlat.Name = field.NewString(tableName, "name")
	_nftChainPlat.Details = field.NewString(tableName, "details")
	_nftChainPlat.NodeHost = field.NewString(tableName, "nodeHost")
	_nftChainPlat.AddressPath = field.NewString(tableName, "addressPath")
	_nftChainPlat.PrivatePath = field.NewString(tableName, "privatePath")
	_nftChainPlat.Passwd = field.NewString(tableName, "passwd")
	_nftChainPlat.PublicPath = field.NewString(tableName, "publicPath")
	_nftChainPlat.ContractAccount = field.NewString(tableName, "contractAccount")
	_nftChainPlat.ContractAddress = field.NewString(tableName, "contractAddress")
	_nftChainPlat.BrowserAddress = field.NewString(tableName, "browserAddress")
	_nftChainPlat.CurrentState = field.NewInt64(tableName, "currentState")
	_nftChainPlat.Balance = field.NewFloat64(tableName, "balance")
	_nftChainPlat.CreateTime = field.NewTime(tableName, "createTime")
	_nftChainPlat.UpdateTime = field.NewTime(tableName, "updateTime")
	_nftChainPlat.Bak1 = field.NewString(tableName, "bak1")
	_nftChainPlat.Bak2 = field.NewString(tableName, "bak2")
	_nftChainPlat.Bak3 = field.NewString(tableName, "bak3")

	_nftChainPlat.fillFieldMap()

	return _nftChainPlat
}

type nftChainPlat struct {
	nftChainPlatDo nftChainPlatDo

	ALL             field.Asterisk
	ID              field.String  // 唯一标识
	ChainName       field.String  // 链名
	EnglishName     field.String  // 链平台英文名称
	Name            field.String  // 链平台名称
	Details         field.String  // 平台描述
	NodeHost        field.String  // 节点链接地址host=ip+port
	AddressPath     field.String  // ak路径(classpath路径下)
	PrivatePath     field.String  // 私钥路径(classpath路径下)
	Passwd          field.String  // 密码,使用SecUtils.encoderByRSAPrivateKey()加密后的数据
	PublicPath      field.String  // 共钥路径(classpath路径下)非必填
	ContractAccount field.String  // 合约账户
	ContractAddress field.String  // 合约地址
	BrowserAddress  field.String  // 区块链浏览器地址
	CurrentState    field.Int64   // 当前适配状态 0:适配中,1:适配完成
	Balance         field.Float64 // root账户余额
	CreateTime      field.Time    // 创建时间
	UpdateTime      field.Time    // 更新时间
	Bak1            field.String  // bak
	Bak2            field.String
	Bak3            field.String

	fieldMap map[string]field.Expr
}

func (n nftChainPlat) Table(newTableName string) *nftChainPlat {
	n.nftChainPlatDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n nftChainPlat) As(alias string) *nftChainPlat {
	n.nftChainPlatDo.DO = *(n.nftChainPlatDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *nftChainPlat) updateTableName(table string) *nftChainPlat {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewString(table, "id")
	n.ChainName = field.NewString(table, "chainName")
	n.EnglishName = field.NewString(table, "englishName")
	n.Name = field.NewString(table, "name")
	n.Details = field.NewString(table, "details")
	n.NodeHost = field.NewString(table, "nodeHost")
	n.AddressPath = field.NewString(table, "addressPath")
	n.PrivatePath = field.NewString(table, "privatePath")
	n.Passwd = field.NewString(table, "passwd")
	n.PublicPath = field.NewString(table, "publicPath")
	n.ContractAccount = field.NewString(table, "contractAccount")
	n.ContractAddress = field.NewString(table, "contractAddress")
	n.BrowserAddress = field.NewString(table, "browserAddress")
	n.CurrentState = field.NewInt64(table, "currentState")
	n.Balance = field.NewFloat64(table, "balance")
	n.CreateTime = field.NewTime(table, "createTime")
	n.UpdateTime = field.NewTime(table, "updateTime")
	n.Bak1 = field.NewString(table, "bak1")
	n.Bak2 = field.NewString(table, "bak2")
	n.Bak3 = field.NewString(table, "bak3")

	n.fillFieldMap()

	return n
}

func (n *nftChainPlat) WithContext(ctx context.Context) INftChainPlatDo {
	return n.nftChainPlatDo.WithContext(ctx)
}

func (n nftChainPlat) TableName() string { return n.nftChainPlatDo.TableName() }

func (n nftChainPlat) Alias() string { return n.nftChainPlatDo.Alias() }

func (n *nftChainPlat) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *nftChainPlat) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 20)
	n.fieldMap["id"] = n.ID
	n.fieldMap["chainName"] = n.ChainName
	n.fieldMap["englishName"] = n.EnglishName
	n.fieldMap["name"] = n.Name
	n.fieldMap["details"] = n.Details
	n.fieldMap["nodeHost"] = n.NodeHost
	n.fieldMap["addressPath"] = n.AddressPath
	n.fieldMap["privatePath"] = n.PrivatePath
	n.fieldMap["passwd"] = n.Passwd
	n.fieldMap["publicPath"] = n.PublicPath
	n.fieldMap["contractAccount"] = n.ContractAccount
	n.fieldMap["contractAddress"] = n.ContractAddress
	n.fieldMap["browserAddress"] = n.BrowserAddress
	n.fieldMap["currentState"] = n.CurrentState
	n.fieldMap["balance"] = n.Balance
	n.fieldMap["createTime"] = n.CreateTime
	n.fieldMap["updateTime"] = n.UpdateTime
	n.fieldMap["bak1"] = n.Bak1
	n.fieldMap["bak2"] = n.Bak2
	n.fieldMap["bak3"] = n.Bak3
}

func (n nftChainPlat) clone(db *gorm.DB) nftChainPlat {
	n.nftChainPlatDo.ReplaceDB(db)
	return n
}

type nftChainPlatDo struct{ gen.DO }

type INftChainPlatDo interface {
	gen.SubQuery
	Debug() INftChainPlatDo
	WithContext(ctx context.Context) INftChainPlatDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INftChainPlatDo
	Not(conds ...gen.Condition) INftChainPlatDo
	Or(conds ...gen.Condition) INftChainPlatDo
	Select(conds ...field.Expr) INftChainPlatDo
	Where(conds ...gen.Condition) INftChainPlatDo
	Order(conds ...field.Expr) INftChainPlatDo
	Distinct(cols ...field.Expr) INftChainPlatDo
	Omit(cols ...field.Expr) INftChainPlatDo
	Join(table schema.Tabler, on ...field.Expr) INftChainPlatDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INftChainPlatDo
	RightJoin(table schema.Tabler, on ...field.Expr) INftChainPlatDo
	Group(cols ...field.Expr) INftChainPlatDo
	Having(conds ...gen.Condition) INftChainPlatDo
	Limit(limit int) INftChainPlatDo
	Offset(offset int) INftChainPlatDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INftChainPlatDo
	Unscoped() INftChainPlatDo
	Create(values ...*model.NftChainPlat) error
	CreateInBatches(values []*model.NftChainPlat, batchSize int) error
	Save(values ...*model.NftChainPlat) error
	First() (*model.NftChainPlat, error)
	Take() (*model.NftChainPlat, error)
	Last() (*model.NftChainPlat, error)
	Find() ([]*model.NftChainPlat, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NftChainPlat, err error)
	FindInBatches(result *[]*model.NftChainPlat, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NftChainPlat) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INftChainPlatDo
	Assign(attrs ...field.AssignExpr) INftChainPlatDo
	Joins(fields ...field.RelationField) INftChainPlatDo
	Preload(fields ...field.RelationField) INftChainPlatDo
	FirstOrInit() (*model.NftChainPlat, error)
	FirstOrCreate() (*model.NftChainPlat, error)
	FindByPage(offset int, limit int) (result []*model.NftChainPlat, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INftChainPlatDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n nftChainPlatDo) Debug() INftChainPlatDo {
	return n.withDO(n.DO.Debug())
}

func (n nftChainPlatDo) WithContext(ctx context.Context) INftChainPlatDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n nftChainPlatDo) ReadDB() INftChainPlatDo {
	return n.Clauses(dbresolver.Read)
}

func (n nftChainPlatDo) WriteDB() INftChainPlatDo {
	return n.Clauses(dbresolver.Write)
}

func (n nftChainPlatDo) Clauses(conds ...clause.Expression) INftChainPlatDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n nftChainPlatDo) Returning(value interface{}, columns ...string) INftChainPlatDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n nftChainPlatDo) Not(conds ...gen.Condition) INftChainPlatDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n nftChainPlatDo) Or(conds ...gen.Condition) INftChainPlatDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n nftChainPlatDo) Select(conds ...field.Expr) INftChainPlatDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n nftChainPlatDo) Where(conds ...gen.Condition) INftChainPlatDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n nftChainPlatDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) INftChainPlatDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n nftChainPlatDo) Order(conds ...field.Expr) INftChainPlatDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n nftChainPlatDo) Distinct(cols ...field.Expr) INftChainPlatDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n nftChainPlatDo) Omit(cols ...field.Expr) INftChainPlatDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n nftChainPlatDo) Join(table schema.Tabler, on ...field.Expr) INftChainPlatDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n nftChainPlatDo) LeftJoin(table schema.Tabler, on ...field.Expr) INftChainPlatDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n nftChainPlatDo) RightJoin(table schema.Tabler, on ...field.Expr) INftChainPlatDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n nftChainPlatDo) Group(cols ...field.Expr) INftChainPlatDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n nftChainPlatDo) Having(conds ...gen.Condition) INftChainPlatDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n nftChainPlatDo) Limit(limit int) INftChainPlatDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n nftChainPlatDo) Offset(offset int) INftChainPlatDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n nftChainPlatDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INftChainPlatDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n nftChainPlatDo) Unscoped() INftChainPlatDo {
	return n.withDO(n.DO.Unscoped())
}

func (n nftChainPlatDo) Create(values ...*model.NftChainPlat) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n nftChainPlatDo) CreateInBatches(values []*model.NftChainPlat, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n nftChainPlatDo) Save(values ...*model.NftChainPlat) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n nftChainPlatDo) First() (*model.NftChainPlat, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftChainPlat), nil
	}
}

func (n nftChainPlatDo) Take() (*model.NftChainPlat, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftChainPlat), nil
	}
}

func (n nftChainPlatDo) Last() (*model.NftChainPlat, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftChainPlat), nil
	}
}

func (n nftChainPlatDo) Find() ([]*model.NftChainPlat, error) {
	result, err := n.DO.Find()
	return result.([]*model.NftChainPlat), err
}

func (n nftChainPlatDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NftChainPlat, err error) {
	buf := make([]*model.NftChainPlat, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n nftChainPlatDo) FindInBatches(result *[]*model.NftChainPlat, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n nftChainPlatDo) Attrs(attrs ...field.AssignExpr) INftChainPlatDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n nftChainPlatDo) Assign(attrs ...field.AssignExpr) INftChainPlatDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n nftChainPlatDo) Joins(fields ...field.RelationField) INftChainPlatDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n nftChainPlatDo) Preload(fields ...field.RelationField) INftChainPlatDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n nftChainPlatDo) FirstOrInit() (*model.NftChainPlat, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftChainPlat), nil
	}
}

func (n nftChainPlatDo) FirstOrCreate() (*model.NftChainPlat, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftChainPlat), nil
	}
}

func (n nftChainPlatDo) FindByPage(offset int, limit int) (result []*model.NftChainPlat, count int64, err error) {
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

func (n nftChainPlatDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n nftChainPlatDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n nftChainPlatDo) Delete(models ...*model.NftChainPlat) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *nftChainPlatDo) withDO(do gen.Dao) *nftChainPlatDo {
	n.DO = *do.(*gen.DO)
	return n
}
