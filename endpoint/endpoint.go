package endpoint

import (
	"github.com/go-kit/kit/endpoint"
	logger "github.com/go-kit/kit/log"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	"github.com/muhammadisa/barektest-tag/constant"
	_interface "github.com/muhammadisa/barektest-tag/service/interface"
	"github.com/muhammadisa/barektest-util/mw"
	"github.com/muhammadisa/barektest-util/vlt"
)

type TagEndpoint struct {
	AddTagEndpoint    endpoint.Endpoint
	EditTagEndpoint   endpoint.Endpoint
	DeleteTagEndpoint endpoint.Endpoint
	GetTagsEndpoint   endpoint.Endpoint
}

func NewTagEndpoint(tagSvc _interface.Service, logger logger.Logger, vault vlt.VLT) (TagEndpoint, error) {
	var addTagEp endpoint.Endpoint
	{
		const name = `AddTag`
		addTagEp = makeAddTagEndpoint(tagSvc)
		addTagEp = mw.LoggingMiddleware(logger)(addTagEp)
		addTagEp = mw.CircuitBreakerMiddleware(constant.ServiceName)(addTagEp)
		addTagEp = kitoc.TraceEndpoint(name)(addTagEp)
	}

	var editTagEp endpoint.Endpoint
	{
		const name = `EditTag`
		editTagEp = makeEditTagEndpoint(tagSvc)
		editTagEp = mw.LoggingMiddleware(logger)(editTagEp)
		editTagEp = mw.CircuitBreakerMiddleware(constant.ServiceName)(editTagEp)
		editTagEp = kitoc.TraceEndpoint(name)(editTagEp)
	}

	var deleteTagEp endpoint.Endpoint
	{
		const name = `DeleteTag`
		deleteTagEp = makeDeleteTagEndpoint(tagSvc)
		deleteTagEp = mw.LoggingMiddleware(logger)(deleteTagEp)
		deleteTagEp = mw.CircuitBreakerMiddleware(constant.ServiceName)(deleteTagEp)
		deleteTagEp = kitoc.TraceEndpoint(name)(deleteTagEp)
	}

	var getTagsEp endpoint.Endpoint
	{
		const name = `DeleteTag`
		getTagsEp = makeGetTagsEndpoint(tagSvc)
		getTagsEp = mw.LoggingMiddleware(logger)(getTagsEp)
		getTagsEp = mw.CircuitBreakerMiddleware(constant.ServiceName)(getTagsEp)
		getTagsEp = kitoc.TraceEndpoint(name)(getTagsEp)
	}

	return TagEndpoint{
		AddTagEndpoint:    addTagEp,
		EditTagEndpoint:   editTagEp,
		DeleteTagEndpoint: deleteTagEp,
		GetTagsEndpoint:   getTagsEp,
	}, nil
}
