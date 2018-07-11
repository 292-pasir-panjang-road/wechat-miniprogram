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
)

func MakeRetrieveEndpoint(logger log.Logger, service services.Service, serviceType string) endpoint.Endpoint {
  return func(ctx context.Context, request interface{}) (interface{}, error) {
    defer func(start time.Time) {
      logger.Log(
        "error",    err,
        "took",     time.Since(start),
        "endpoint", serviceType,
        "params",   helper.ObjToString(request)
      )
    }(time.Now())
    return service.Retrieve(request)
  }
}
