package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"review-b/internal/conf"
	"review-b/internal/data/query"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewBusinessRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	query *query.Query
	log   *log.Helper
}

// NewData .
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	query.SetDefault(db)
	return &Data{
		query: query.Q,
		log:   log.NewHelper(logger),
	}, cleanup, nil
}

func NewDB(c *conf.Data) (db *gorm.DB, err error) {
	switch strings.ToLower(c.Database.Driver) {
	case "mysql":
		db, err = gorm.Open(mysql.Open(c.Database.Source))
		if err != nil {
			return nil, err
		}
		return db, nil
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(c.Database.Source))
		if err != nil {
			return nil, err
		}
		return db, nil
	}
	panic("the db open has panic")
}
