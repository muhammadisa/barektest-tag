package service

import (
	_repointerface "github.com/muhammadisa/barektest-tag/repository"
	_interface "github.com/muhammadisa/barektest-tag/service/interface"
	"github.com/muhammadisa/barektest-util/vlt"
	"go.opencensus.io/trace"
)

type service struct {
	vault  vlt.VLT
	tracer trace.Tracer
	repo   _repointerface.Repository
}

func NewUsecases(repo _repointerface.Repository, tracer trace.Tracer, vault vlt.VLT) _interface.Service {
	return &service{
		tracer: tracer,
		repo:   repo,
		vault:  vault,
	}
}
