// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"review-b/internal/data/model"
)

func newTodo(db *gorm.DB, opts ...gen.DOOption) todo {
	_todo := todo{}

	_todo.todoDo.UseDB(db, opts...)
	_todo.todoDo.UseModel(&model.Todo{})

	tableName := _todo.todoDo.TableName()
	_todo.ALL = field.NewAsterisk(tableName)
	_todo.ID = field.NewInt64(tableName, "id")
	_todo.Title = field.NewString(tableName, "title")
	_todo.Status = field.NewBool(tableName, "status")

	_todo.fillFieldMap()

	return _todo
}

type todo struct {
	todoDo todoDo

	ALL    field.Asterisk
	ID     field.Int64
	Title  field.String
	Status field.Bool

	fieldMap map[string]field.Expr
}

func (t todo) Table(newTableName string) *todo {
	t.todoDo.UseTable(newTableName)
	return t.updateTableName(newTableName)
}

func (t todo) As(alias string) *todo {
	t.todoDo.DO = *(t.todoDo.As(alias).(*gen.DO))
	return t.updateTableName(alias)
}

func (t *todo) updateTableName(table string) *todo {
	t.ALL = field.NewAsterisk(table)
	t.ID = field.NewInt64(table, "id")
	t.Title = field.NewString(table, "title")
	t.Status = field.NewBool(table, "status")

	t.fillFieldMap()

	return t
}

func (t *todo) WithContext(ctx context.Context) ITodoDo { return t.todoDo.WithContext(ctx) }

func (t todo) TableName() string { return t.todoDo.TableName() }

func (t todo) Alias() string { return t.todoDo.Alias() }

func (t todo) Columns(cols ...field.Expr) gen.Columns { return t.todoDo.Columns(cols...) }

func (t *todo) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := t.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (t *todo) fillFieldMap() {
	t.fieldMap = make(map[string]field.Expr, 3)
	t.fieldMap["id"] = t.ID
	t.fieldMap["title"] = t.Title
	t.fieldMap["status"] = t.Status
}

func (t todo) clone(db *gorm.DB) todo {
	t.todoDo.ReplaceConnPool(db.Statement.ConnPool)
	return t
}

func (t todo) replaceDB(db *gorm.DB) todo {
	t.todoDo.ReplaceDB(db)
	return t
}

type todoDo struct{ gen.DO }

type ITodoDo interface {
	gen.SubQuery
	Debug() ITodoDo
	WithContext(ctx context.Context) ITodoDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ITodoDo
	WriteDB() ITodoDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ITodoDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ITodoDo
	Not(conds ...gen.Condition) ITodoDo
	Or(conds ...gen.Condition) ITodoDo
	Select(conds ...field.Expr) ITodoDo
	Where(conds ...gen.Condition) ITodoDo
	Order(conds ...field.Expr) ITodoDo
	Distinct(cols ...field.Expr) ITodoDo
	Omit(cols ...field.Expr) ITodoDo
	Join(table schema.Tabler, on ...field.Expr) ITodoDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ITodoDo
	RightJoin(table schema.Tabler, on ...field.Expr) ITodoDo
	Group(cols ...field.Expr) ITodoDo
	Having(conds ...gen.Condition) ITodoDo
	Limit(limit int) ITodoDo
	Offset(offset int) ITodoDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ITodoDo
	Unscoped() ITodoDo
	Create(values ...*model.Todo) error
	CreateInBatches(values []*model.Todo, batchSize int) error
	Save(values ...*model.Todo) error
	First() (*model.Todo, error)
	Take() (*model.Todo, error)
	Last() (*model.Todo, error)
	Find() ([]*model.Todo, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Todo, err error)
	FindInBatches(result *[]*model.Todo, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Todo) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ITodoDo
	Assign(attrs ...field.AssignExpr) ITodoDo
	Joins(fields ...field.RelationField) ITodoDo
	Preload(fields ...field.RelationField) ITodoDo
	FirstOrInit() (*model.Todo, error)
	FirstOrCreate() (*model.Todo, error)
	FindByPage(offset int, limit int) (result []*model.Todo, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ITodoDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (t todoDo) Debug() ITodoDo {
	return t.withDO(t.DO.Debug())
}

func (t todoDo) WithContext(ctx context.Context) ITodoDo {
	return t.withDO(t.DO.WithContext(ctx))
}

func (t todoDo) ReadDB() ITodoDo {
	return t.Clauses(dbresolver.Read)
}

func (t todoDo) WriteDB() ITodoDo {
	return t.Clauses(dbresolver.Write)
}

func (t todoDo) Session(config *gorm.Session) ITodoDo {
	return t.withDO(t.DO.Session(config))
}

func (t todoDo) Clauses(conds ...clause.Expression) ITodoDo {
	return t.withDO(t.DO.Clauses(conds...))
}

func (t todoDo) Returning(value interface{}, columns ...string) ITodoDo {
	return t.withDO(t.DO.Returning(value, columns...))
}

func (t todoDo) Not(conds ...gen.Condition) ITodoDo {
	return t.withDO(t.DO.Not(conds...))
}

func (t todoDo) Or(conds ...gen.Condition) ITodoDo {
	return t.withDO(t.DO.Or(conds...))
}

func (t todoDo) Select(conds ...field.Expr) ITodoDo {
	return t.withDO(t.DO.Select(conds...))
}

func (t todoDo) Where(conds ...gen.Condition) ITodoDo {
	return t.withDO(t.DO.Where(conds...))
}

func (t todoDo) Order(conds ...field.Expr) ITodoDo {
	return t.withDO(t.DO.Order(conds...))
}

func (t todoDo) Distinct(cols ...field.Expr) ITodoDo {
	return t.withDO(t.DO.Distinct(cols...))
}

func (t todoDo) Omit(cols ...field.Expr) ITodoDo {
	return t.withDO(t.DO.Omit(cols...))
}

func (t todoDo) Join(table schema.Tabler, on ...field.Expr) ITodoDo {
	return t.withDO(t.DO.Join(table, on...))
}

func (t todoDo) LeftJoin(table schema.Tabler, on ...field.Expr) ITodoDo {
	return t.withDO(t.DO.LeftJoin(table, on...))
}

func (t todoDo) RightJoin(table schema.Tabler, on ...field.Expr) ITodoDo {
	return t.withDO(t.DO.RightJoin(table, on...))
}

func (t todoDo) Group(cols ...field.Expr) ITodoDo {
	return t.withDO(t.DO.Group(cols...))
}

func (t todoDo) Having(conds ...gen.Condition) ITodoDo {
	return t.withDO(t.DO.Having(conds...))
}

func (t todoDo) Limit(limit int) ITodoDo {
	return t.withDO(t.DO.Limit(limit))
}

func (t todoDo) Offset(offset int) ITodoDo {
	return t.withDO(t.DO.Offset(offset))
}

func (t todoDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ITodoDo {
	return t.withDO(t.DO.Scopes(funcs...))
}

func (t todoDo) Unscoped() ITodoDo {
	return t.withDO(t.DO.Unscoped())
}

func (t todoDo) Create(values ...*model.Todo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Create(values)
}

func (t todoDo) CreateInBatches(values []*model.Todo, batchSize int) error {
	return t.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (t todoDo) Save(values ...*model.Todo) error {
	if len(values) == 0 {
		return nil
	}
	return t.DO.Save(values)
}

func (t todoDo) First() (*model.Todo, error) {
	if result, err := t.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Todo), nil
	}
}

func (t todoDo) Take() (*model.Todo, error) {
	if result, err := t.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Todo), nil
	}
}

func (t todoDo) Last() (*model.Todo, error) {
	if result, err := t.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Todo), nil
	}
}

func (t todoDo) Find() ([]*model.Todo, error) {
	result, err := t.DO.Find()
	return result.([]*model.Todo), err
}

func (t todoDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Todo, err error) {
	buf := make([]*model.Todo, 0, batchSize)
	err = t.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (t todoDo) FindInBatches(result *[]*model.Todo, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return t.DO.FindInBatches(result, batchSize, fc)
}

func (t todoDo) Attrs(attrs ...field.AssignExpr) ITodoDo {
	return t.withDO(t.DO.Attrs(attrs...))
}

func (t todoDo) Assign(attrs ...field.AssignExpr) ITodoDo {
	return t.withDO(t.DO.Assign(attrs...))
}

func (t todoDo) Joins(fields ...field.RelationField) ITodoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Joins(_f))
	}
	return &t
}

func (t todoDo) Preload(fields ...field.RelationField) ITodoDo {
	for _, _f := range fields {
		t = *t.withDO(t.DO.Preload(_f))
	}
	return &t
}

func (t todoDo) FirstOrInit() (*model.Todo, error) {
	if result, err := t.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Todo), nil
	}
}

func (t todoDo) FirstOrCreate() (*model.Todo, error) {
	if result, err := t.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Todo), nil
	}
}

func (t todoDo) FindByPage(offset int, limit int) (result []*model.Todo, count int64, err error) {
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

func (t todoDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = t.Count()
	if err != nil {
		return
	}

	err = t.Offset(offset).Limit(limit).Scan(result)
	return
}

func (t todoDo) Scan(result interface{}) (err error) {
	return t.DO.Scan(result)
}

func (t todoDo) Delete(models ...*model.Todo) (result gen.ResultInfo, err error) {
	return t.DO.Delete(models)
}

func (t *todoDo) withDO(do gen.Dao) *todoDo {
	t.DO = *do.(*gen.DO)
	return t
}
