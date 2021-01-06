package db

import (
	"github.com/dolthub/go-mysql-server/sql"
	"github.com/lingsamuel/sqlserver/pkg/proxy"
	"github.com/lingsamuel/sqlserver/pkg/proxy/redis"
	"github.com/lingsamuel/sqlserver/pkg/util"
	"github.com/pkg/errors"
)

var _ TableCreator = NewRedisTable

// NewRedisTable creates a new sql.Table with the given name and schema.
func NewRedisTable(name string, schema sql.Schema, source string) (sql.Table, error) {
	if err := util.ValidateRedisTableSchema(schema); err != nil {
		return nil, err
	}
	if err := redis.PingRedisClient(source); err != nil {
		return nil, errors.Wrapf(err, "Ping redis %s failed", source)
	}
	return &ProxyTable{
		source:  source,
		name:    name,
		schema:  schema,
		fetcher: proxy.RedisFetch,
	}, nil
}
