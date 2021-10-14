package repository

import (
	"context"
	_interface "github.com/muhammadisa/barektest-tag/repository/interface"
	"github.com/muhammadisa/barektest-tag/repository/sql"
	"github.com/muhammadisa/barektest-util/dbc"
	"go.opencensus.io/trace"
)

type Repository struct {
	ReadWriter _interface.ReadWrite
}

type RepoConf struct {
	SQL dbc.Config
}

func NewRepositories(ctx context.Context, rc RepoConf, tracer trace.Tracer) (*Repository, error) {
	readWriter, err := sql.NewSQL(rc.SQL, tracer)
	if err != nil {
		return nil, err
	}
	return &Repository{
		ReadWriter: readWriter,
	}, nil
}
