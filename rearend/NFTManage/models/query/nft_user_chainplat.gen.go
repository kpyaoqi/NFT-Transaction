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

func newNftUserChainplat(db *gorm.DB) nftUserChainplat {
	_nftUserChainplat := nftUserChainplat{}

	_nftUserChainplat.nftUserChainplatDo.UseDB(db)
	_nftUserChainplat.nftUserChainplatDo.UseModel(&model.NftUserChainplat{})

	tableName := _nftUserChainplat.nftUserChainplatDo.TableName()
	_nftUserChainplat.ALL = field.NewAsterisk(tableName)
	_nftUserChainplat.ID = field.NewString(tableName, "id")
	_nftUserChainplat.UserID = field.NewString(tableName, "userId")
	_nftUserChainplat.ChainPlatName = field.NewString(tableName, "chainPlatName")
	_nftUserChainplat.PrivatePath = field.NewString(tableName, "privatePath")
	_nftUserChainplat.Passwd = field.NewString(tableName, "passwd")
	_nftUserChainplat.PublicPath = field.NewString(tableName, "publicPath")
	_nftUserChainplat.AddressPath = field.NewString(tableName, "addressPath")
	_nftUserChainplat.Address = field.NewString(tableName, "address")
	_nftUserChainplat.EVMAddress = field.NewString(tableName, "EVMAddress")
	_nftUserChainplat.Mnemonic = field.NewString(tableName, "mnemonic")
	_nftUserChainplat.PrevTime = field.NewTime(tableName, "prevTime")
	_nftUserChainplat.CreateTime = field.NewTime(tableName, "createTime")
	_nftUserChainplat.UpdateTime = field.NewTime(tableName, "updateTime")
	_nftUserChainplat.Bak1 = field.NewString(tableName, "bak1")
	_nftUserChainplat.Bak2 = field.NewString(tableName, "bak2")
	_nftUserChainplat.Bak3 = field.NewString(tableName, "bak3")

	_nftUserChainplat.fillFieldMap()

	return _nftUserChainplat
}

type nftUserChainplat struct {
	nftUserChainplatDo nftUserChainplatDo

	ALL           field.Asterisk
	ID            field.String // 唯一标识
	UserID        field.String // 用户标识
	ChainPlatName field.String // 链平台英文名称(标识)
	PrivatePath   field.String // 私钥地址
	Passwd        field.String // 密码,使用SecUtils.encoderByRSAPrivateKey()加密后的数据
	PublicPath    field.String // 公钥地址
	AddressPath   field.String // ak地址文件路径
	Address       field.String // ak地址
	EVMAddress    field.String // EVMaddress(由address转换得来的Account.xchainAKToEVMAddress())
	Mnemonic      field.String // 助记词
	PrevTime      field.Time   // 上次链接时间
	CreateTime    field.Time
	UpdateTime    field.Time
	Bak1          field.String
	Bak2          field.String
	Bak3          field.String

	fieldMap map[string]field.Expr
}

func (n nftUserChainplat) Table(newTableName string) *nftUserChainplat {
	n.nftUserChainplatDo.UseTable(newTableName)
	return n.updateTableName(newTableName)
}

func (n nftUserChainplat) As(alias string) *nftUserChainplat {
	n.nftUserChainplatDo.DO = *(n.nftUserChainplatDo.As(alias).(*gen.DO))
	return n.updateTableName(alias)
}

func (n *nftUserChainplat) updateTableName(table string) *nftUserChainplat {
	n.ALL = field.NewAsterisk(table)
	n.ID = field.NewString(table, "id")
	n.UserID = field.NewString(table, "userId")
	n.ChainPlatName = field.NewString(table, "chainPlatName")
	n.PrivatePath = field.NewString(table, "privatePath")
	n.Passwd = field.NewString(table, "passwd")
	n.PublicPath = field.NewString(table, "publicPath")
	n.AddressPath = field.NewString(table, "addressPath")
	n.Address = field.NewString(table, "address")
	n.EVMAddress = field.NewString(table, "EVMAddress")
	n.Mnemonic = field.NewString(table, "mnemonic")
	n.PrevTime = field.NewTime(table, "prevTime")
	n.CreateTime = field.NewTime(table, "createTime")
	n.UpdateTime = field.NewTime(table, "updateTime")
	n.Bak1 = field.NewString(table, "bak1")
	n.Bak2 = field.NewString(table, "bak2")
	n.Bak3 = field.NewString(table, "bak3")

	n.fillFieldMap()

	return n
}

func (n *nftUserChainplat) WithContext(ctx context.Context) INftUserChainplatDo {
	return n.nftUserChainplatDo.WithContext(ctx)
}

func (n nftUserChainplat) TableName() string { return n.nftUserChainplatDo.TableName() }

func (n nftUserChainplat) Alias() string { return n.nftUserChainplatDo.Alias() }

func (n *nftUserChainplat) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := n.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (n *nftUserChainplat) fillFieldMap() {
	n.fieldMap = make(map[string]field.Expr, 16)
	n.fieldMap["id"] = n.ID
	n.fieldMap["userId"] = n.UserID
	n.fieldMap["chainPlatName"] = n.ChainPlatName
	n.fieldMap["privatePath"] = n.PrivatePath
	n.fieldMap["passwd"] = n.Passwd
	n.fieldMap["publicPath"] = n.PublicPath
	n.fieldMap["addressPath"] = n.AddressPath
	n.fieldMap["address"] = n.Address
	n.fieldMap["EVMAddress"] = n.EVMAddress
	n.fieldMap["mnemonic"] = n.Mnemonic
	n.fieldMap["prevTime"] = n.PrevTime
	n.fieldMap["createTime"] = n.CreateTime
	n.fieldMap["updateTime"] = n.UpdateTime
	n.fieldMap["bak1"] = n.Bak1
	n.fieldMap["bak2"] = n.Bak2
	n.fieldMap["bak3"] = n.Bak3
}

func (n nftUserChainplat) clone(db *gorm.DB) nftUserChainplat {
	n.nftUserChainplatDo.ReplaceDB(db)
	return n
}

type nftUserChainplatDo struct{ gen.DO }

type INftUserChainplatDo interface {
	gen.SubQuery
	Debug() INftUserChainplatDo
	WithContext(ctx context.Context) INftUserChainplatDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	As(alias string) gen.Dao
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) INftUserChainplatDo
	Not(conds ...gen.Condition) INftUserChainplatDo
	Or(conds ...gen.Condition) INftUserChainplatDo
	Select(conds ...field.Expr) INftUserChainplatDo
	Where(conds ...gen.Condition) INftUserChainplatDo
	Order(conds ...field.Expr) INftUserChainplatDo
	Distinct(cols ...field.Expr) INftUserChainplatDo
	Omit(cols ...field.Expr) INftUserChainplatDo
	Join(table schema.Tabler, on ...field.Expr) INftUserChainplatDo
	LeftJoin(table schema.Tabler, on ...field.Expr) INftUserChainplatDo
	RightJoin(table schema.Tabler, on ...field.Expr) INftUserChainplatDo
	Group(cols ...field.Expr) INftUserChainplatDo
	Having(conds ...gen.Condition) INftUserChainplatDo
	Limit(limit int) INftUserChainplatDo
	Offset(offset int) INftUserChainplatDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) INftUserChainplatDo
	Unscoped() INftUserChainplatDo
	Create(values ...*model.NftUserChainplat) error
	CreateInBatches(values []*model.NftUserChainplat, batchSize int) error
	Save(values ...*model.NftUserChainplat) error
	First() (*model.NftUserChainplat, error)
	Take() (*model.NftUserChainplat, error)
	Last() (*model.NftUserChainplat, error)
	Find() ([]*model.NftUserChainplat, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NftUserChainplat, err error)
	FindInBatches(result *[]*model.NftUserChainplat, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.NftUserChainplat) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) INftUserChainplatDo
	Assign(attrs ...field.AssignExpr) INftUserChainplatDo
	Joins(fields ...field.RelationField) INftUserChainplatDo
	Preload(fields ...field.RelationField) INftUserChainplatDo
	FirstOrInit() (*model.NftUserChainplat, error)
	FirstOrCreate() (*model.NftUserChainplat, error)
	FindByPage(offset int, limit int) (result []*model.NftUserChainplat, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) INftUserChainplatDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (n nftUserChainplatDo) Debug() INftUserChainplatDo {
	return n.withDO(n.DO.Debug())
}

func (n nftUserChainplatDo) WithContext(ctx context.Context) INftUserChainplatDo {
	return n.withDO(n.DO.WithContext(ctx))
}

func (n nftUserChainplatDo) ReadDB() INftUserChainplatDo {
	return n.Clauses(dbresolver.Read)
}

func (n nftUserChainplatDo) WriteDB() INftUserChainplatDo {
	return n.Clauses(dbresolver.Write)
}

func (n nftUserChainplatDo) Clauses(conds ...clause.Expression) INftUserChainplatDo {
	return n.withDO(n.DO.Clauses(conds...))
}

func (n nftUserChainplatDo) Returning(value interface{}, columns ...string) INftUserChainplatDo {
	return n.withDO(n.DO.Returning(value, columns...))
}

func (n nftUserChainplatDo) Not(conds ...gen.Condition) INftUserChainplatDo {
	return n.withDO(n.DO.Not(conds...))
}

func (n nftUserChainplatDo) Or(conds ...gen.Condition) INftUserChainplatDo {
	return n.withDO(n.DO.Or(conds...))
}

func (n nftUserChainplatDo) Select(conds ...field.Expr) INftUserChainplatDo {
	return n.withDO(n.DO.Select(conds...))
}

func (n nftUserChainplatDo) Where(conds ...gen.Condition) INftUserChainplatDo {
	return n.withDO(n.DO.Where(conds...))
}

func (n nftUserChainplatDo) Exists(subquery interface{ UnderlyingDB() *gorm.DB }) INftUserChainplatDo {
	return n.Where(field.CompareSubQuery(field.ExistsOp, nil, subquery.UnderlyingDB()))
}

func (n nftUserChainplatDo) Order(conds ...field.Expr) INftUserChainplatDo {
	return n.withDO(n.DO.Order(conds...))
}

func (n nftUserChainplatDo) Distinct(cols ...field.Expr) INftUserChainplatDo {
	return n.withDO(n.DO.Distinct(cols...))
}

func (n nftUserChainplatDo) Omit(cols ...field.Expr) INftUserChainplatDo {
	return n.withDO(n.DO.Omit(cols...))
}

func (n nftUserChainplatDo) Join(table schema.Tabler, on ...field.Expr) INftUserChainplatDo {
	return n.withDO(n.DO.Join(table, on...))
}

func (n nftUserChainplatDo) LeftJoin(table schema.Tabler, on ...field.Expr) INftUserChainplatDo {
	return n.withDO(n.DO.LeftJoin(table, on...))
}

func (n nftUserChainplatDo) RightJoin(table schema.Tabler, on ...field.Expr) INftUserChainplatDo {
	return n.withDO(n.DO.RightJoin(table, on...))
}

func (n nftUserChainplatDo) Group(cols ...field.Expr) INftUserChainplatDo {
	return n.withDO(n.DO.Group(cols...))
}

func (n nftUserChainplatDo) Having(conds ...gen.Condition) INftUserChainplatDo {
	return n.withDO(n.DO.Having(conds...))
}

func (n nftUserChainplatDo) Limit(limit int) INftUserChainplatDo {
	return n.withDO(n.DO.Limit(limit))
}

func (n nftUserChainplatDo) Offset(offset int) INftUserChainplatDo {
	return n.withDO(n.DO.Offset(offset))
}

func (n nftUserChainplatDo) Scopes(funcs ...func(gen.Dao) gen.Dao) INftUserChainplatDo {
	return n.withDO(n.DO.Scopes(funcs...))
}

func (n nftUserChainplatDo) Unscoped() INftUserChainplatDo {
	return n.withDO(n.DO.Unscoped())
}

func (n nftUserChainplatDo) Create(values ...*model.NftUserChainplat) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Create(values)
}

func (n nftUserChainplatDo) CreateInBatches(values []*model.NftUserChainplat, batchSize int) error {
	return n.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (n nftUserChainplatDo) Save(values ...*model.NftUserChainplat) error {
	if len(values) == 0 {
		return nil
	}
	return n.DO.Save(values)
}

func (n nftUserChainplatDo) First() (*model.NftUserChainplat, error) {
	if result, err := n.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserChainplat), nil
	}
}

func (n nftUserChainplatDo) Take() (*model.NftUserChainplat, error) {
	if result, err := n.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserChainplat), nil
	}
}

func (n nftUserChainplatDo) Last() (*model.NftUserChainplat, error) {
	if result, err := n.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserChainplat), nil
	}
}

func (n nftUserChainplatDo) Find() ([]*model.NftUserChainplat, error) {
	result, err := n.DO.Find()
	return result.([]*model.NftUserChainplat), err
}

func (n nftUserChainplatDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.NftUserChainplat, err error) {
	buf := make([]*model.NftUserChainplat, 0, batchSize)
	err = n.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (n nftUserChainplatDo) FindInBatches(result *[]*model.NftUserChainplat, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return n.DO.FindInBatches(result, batchSize, fc)
}

func (n nftUserChainplatDo) Attrs(attrs ...field.AssignExpr) INftUserChainplatDo {
	return n.withDO(n.DO.Attrs(attrs...))
}

func (n nftUserChainplatDo) Assign(attrs ...field.AssignExpr) INftUserChainplatDo {
	return n.withDO(n.DO.Assign(attrs...))
}

func (n nftUserChainplatDo) Joins(fields ...field.RelationField) INftUserChainplatDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Joins(_f))
	}
	return &n
}

func (n nftUserChainplatDo) Preload(fields ...field.RelationField) INftUserChainplatDo {
	for _, _f := range fields {
		n = *n.withDO(n.DO.Preload(_f))
	}
	return &n
}

func (n nftUserChainplatDo) FirstOrInit() (*model.NftUserChainplat, error) {
	if result, err := n.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserChainplat), nil
	}
}

func (n nftUserChainplatDo) FirstOrCreate() (*model.NftUserChainplat, error) {
	if result, err := n.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.NftUserChainplat), nil
	}
}

func (n nftUserChainplatDo) FindByPage(offset int, limit int) (result []*model.NftUserChainplat, count int64, err error) {
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

func (n nftUserChainplatDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = n.Count()
	if err != nil {
		return
	}

	err = n.Offset(offset).Limit(limit).Scan(result)
	return
}

func (n nftUserChainplatDo) Scan(result interface{}) (err error) {
	return n.DO.Scan(result)
}

func (n nftUserChainplatDo) Delete(models ...*model.NftUserChainplat) (result gen.ResultInfo, err error) {
	return n.DO.Delete(models)
}

func (n *nftUserChainplatDo) withDO(do gen.Dao) *nftUserChainplatDo {
	n.DO = *do.(*gen.DO)
	return n
}