package sql

import (
	"database/sql"
	_interface "github.com/muhammadisa/barektest-tag/repository/interface"
	"github.com/muhammadisa/barektest-util/dbc"
	"go.opencensus.io/trace"
	"sync"
)

var mutex = &sync.RWMutex{}

type readWrite struct {
	tracer trace.Tracer
	db     *sql.DB
}

func NewSQL(config dbc.Config, tracer trace.Tracer) (_interface.ReadWrite, error) {
	sqlDB, err := dbc.OpenDB(config)
	if err != nil {
		return nil, err
	}
	return &readWrite{
		db:     sqlDB,
		tracer: tracer,
	}, nil
}
