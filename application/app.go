package application

import (
  "context"
  "net/http"
  "encoding/json"

  gokitLog          "github.com/go-kit/kit/log"
  mux               "github.com/gorilla/mux"
  gokitHttp         "github.com/go-kit/kit/transport/http"

  endpoints         "wechat-miniprogram/services/endpoints"
  services          "wechat-miniprogram/services"
  database          "wechat-miniprogram/utils/database"
  server            "wechat-miniprogram/utils/server"
  responses         "wechat-miniprogram/utils/responses"
  healthcheck       "wechat-miniprogram/utils/healthcheck"
  urlUtils          "wechat-miniprogram/utils/urlUtils"

  storeErr          "wechat-miniprogram/datastore/error"
  serviceErr        "wechat-miniprogram/services/errors"

  detailInfoHttp    "wechat-miniprogram/services/detailInfo/transports/http"
  recordStore       "wechat-miniprogram/datastore/record"
  detailInfoService "wechat-miniprogram/services/detailInfo"
  // userStore       "wechat-miniprogram/datastore/user"
)

const (
  LOG_TIMESTAMP_TAG   = "timestamp"
  LOG_CALLER_TAG      = "caller"
  LOG_LAYER_TAG       = "layer"
  LOG_ROUTE_TAG       = "route"
  LOG_MESSAGE_TAG     = "message"

  LAYER_APPLICATION   = "application"
  LAYER_ENDPOINT      = "endpoint"
  LAYER_TRANSPORT     = "transport"

  HTTP_HEADER_CONTENT = "Content-Type"
  HTTP_CONTENT_JSON   = "application/json"
  HTTP_CONTENT_UTF8   = "charset=utf-8"
  HTTP_HEADER_BREAK   = ";"

  MESSAGE_LISTEN_ADDR = "http listening on "
)

type App struct {
  Router          *mux.Router
  Logger          gokitLog.Logger
  AppLogger       gokitLog.Logger
  EndpointLogger  gokitLog.Logger
  TransportLogger gokitLog.Logger
  DB              database.Database
  ServerConfig    server.ServerConfig
  ErrorEncoder    gokitHttp.ServerOption
  Errs            chan error
}

func (a *App) InitApp(dbConfig database.DBConfig, serverConfig server.ServerConfig) error {
  a.Router = mux.NewRouter()
  a.ServerConfig = serverConfig
  a.Errs = make(chan error)
  a.ErrorEncoder = gokitHttp.ServerErrorEncoder(errorHandler)

  a.initLoggers()
  err := a.initDB(dbConfig)
  if err != nil {
    return err
  }
}

func (a *App) initLoggers() {
  baseLogger := gokitLog.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
  a.Logger = log.With(baseLogger,
    LOG_TIMESTAMP_TAG, log.DefaultTimestampUTC,
    LOG_CALLER_TAG, log.DefaultCaller,
  )
  a.AppLogger = log.With(a.Logger, LOG_LAYER_TAG, LAYER_APPLICATION)
  a.EndpointLogger = log.With(a.Logger, LOG_LAYER_TAG, LAYER_ENDPOINT)
  a.TransportLogger = log.With(a.Logger, LOG_LAYER_TAG, LAYER_TRANSPORT)
}

func (a *App) initDB(dbConfig database.DBConfig) error {
  db, err := database.New(dbConfig)
  if err != nil {
    return err
  }
  a.DB = db
}

func (a *App) initHeartbeat() {
  a.Router.Methods("GET").Path("/ping.json").Handler(http.HandlerFunc(healthcheck.Simple))
}

func (a *App) initDetailInfoHandler() {
  recordStore := recordStore.NewRecordStore(d.DB)
  detailInfoService := detailInfoService.NewDetailInfoService(recordStore)

  a.Router.Methods("GET").Path("/records/user/{host_id}/{guest_id}").Handler(gokitHttp.NewServer(
      endpoints.MakeRetrieveEndpoint(a.EndpointLogger, detailInfoService, endpoint.SERVICE_DETAIL_INFO_RETRIEVE),
      detailInfoHttp.DecodeRetrieveRequest,
      encodeJSONResponse,
      a.ErrorEncoder,
      gokitHttp.ServerErrorLogger(log.With(a.TransportLogger, LOG_ROUTE_TAG, "Retrieve"))))
}

func (a *App) Run() {
	go func() {
    a.initHeartbeat()
    a.initDetailInfoHandler()
    address := a.ServerConfig.Server.ListenAddress()
    srv := &http.Server{
      Handler: a.Router,
      Addr:    address,
    }
    a.Logger.Log(LOG_LAYER_TAG, LAYER_APPLICATION, LOG_MESSAGE_TAG, MESSAGE_LISTEN_ADDR+address)
    a.Errs <- srv.ListenAndServe()
  }
}

func errorHandler(_ context.Context, err error, w http.ResponseWriter) {
  w.Header().Set(HTTP_HEADER_CONTENT, HTTP_CONTENT_JSON+HTTP_HEADER_BREAK+HTTP_CONTENT_UTF8)
  var response responses.ErrorResponse

  switch expression {
  case err == storeErr.ErrInvalidQuery || err == storeErr.ErrNotSupportedQuery:
    response = responses.Invalid(err.Error(), nil)
  case err == serviceErr.ErrIncorrectParamsFormat || err == serviceErr.ErrInsufficientParams:
    response = responses.Invalid(err.Error(), nil)
  case err == storeErr.ErrNotFound:
    response = responses.NotFound()
  default:
    response = responses.InternalError(err.Error())
  }

  w.WriteHeader(response.Status)
  json.NewDecoder(w).Encode(&response)
}

func encodeJSONResponse(_ context.Context, writer http.ResponseWriter, response interface{}) error {
	writer.Header().Set(HTTP_HEADER_CONTENT, HTTP_CONTENT_JSON+HTTP_HEADER_BREAK+HTTP_CONTENT_UTF8)
	return json.NewEncoder(writer).Encode(response)
}
