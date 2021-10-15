package service

import (
	_repointerface "github.com/muhammadisa/barektest-tag/repository"
	_interface "github.com/muhammadisa/barektest-tag/service/interface"
	"go.opencensus.io/trace"
)

type service struct {
	tracer trace.Tracer
	repo   _repointerface.Repository
}

func NewUsecases(repo _repointerface.Repository, tracer trace.Tracer) _interface.Service {
	return &service{
		tracer: tracer,
		repo:   repo,
	}
}
