package endpoints

import (
  "context"
	"time"
  "github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"

  "wechat-miniprogram/services"
  "wechat-miniprogram/services/helper"
)

const (
  SERVICE_DETAIL_INFO_RETRIEVE = "detail_info_retrieve"

  LOG_ERROR_TAG    = "error"
  LOG_TIME_TAG     = "took"
  LOG_ENDPOINT_TAG = "endpoint"
  LOG_PARAMS_TAG   = "params"
)

func MakeRetrieveEndpoint(logger log.Logger, service services.Service, serviceType string) endpoint.Endpoint {

  // Returns an endpoint (basically, an enpoint is a place to deal with request)
  // For here we just pass it to services to do it.
  return func(ctx context.Context, request interface{}) (interface{}, error) {

    // Before return, log status
    defer func(start time.Time) {
      logger.Log(
        LOG_ERROR_TAG,    err,
        LOG_TIME_TAG,     time.Since(start),
        LOG_ENDPOINT_TAG, serviceType,
        LOG_PARAMS_TAG,   helper.ObjToString(request)
      )
    }(time.Now())

    // Pass job to services
    return service.Retrieve(request)
  }
}
