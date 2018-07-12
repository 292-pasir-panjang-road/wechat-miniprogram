package application

import (
	"os"
	"context"
	"encoding/json"
	"net/http"

	gokitLog "github.com/go-kit/kit/log"
	gokitHttp "github.com/go-kit/kit/transport/http"
	mux "github.com/gorilla/mux"

	endpoints "wechat-miniprogram/services/endpoints"
	database "wechat-miniprogram/utils/database"
	healthcheck "wechat-miniprogram/utils/healthcheck"
	responses "wechat-miniprogram/utils/responses"
	server "wechat-miniprogram/utils/server"

	storeErr "wechat-miniprogram/datastore/error"
	serviceErr "wechat-miniprogram/services/errors"

	recordStore "wechat-miniprogram/datastore/record"
	detailInfoService "wechat-miniprogram/services/detailInfo"
	detailInfoHttp "wechat-miniprogram/services/detailInfo/transports/http"
	// userStore       "wechat-miniprogram/datastore/user"
)

const (
	LOG_TIMESTAMP_TAG = "timestamp"
	LOG_CALLER_TAG    = "caller"
	LOG_LAYER_TAG     = "layer"
	LOG_ROUTE_TAG     = "route"
	LOG_MESSAGE_TAG   = "message"
	LOG_ERROR_TAG			= "error"

	LAYER_APPLICATION = "application"
	LAYER_ENDPOINT    = "endpoint"
	LAYER_TRANSPORT   = "transport"

	HTTP_HEADER_CONTENT = "Content-Type"
	HTTP_CONTENT_JSON   = "application/json"
	HTTP_CONTENT_UTF8   = "charset=utf-8"
	HTTP_HEADER_BREAK   = ";"

	MESSAGE_LISTEN_ADDR = "http listening on "
	MESSAGE_HALTING     = "halting!"
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
	return nil
}

func (a *App) initLoggers() {
	baseLogger := gokitLog.NewLogfmtLogger(gokitLog.NewSyncWriter(os.Stderr))
	a.Logger = gokitLog.With(baseLogger,
		LOG_TIMESTAMP_TAG, gokitLog.DefaultTimestampUTC,
		LOG_CALLER_TAG, gokitLog.DefaultCaller,
	)
	a.AppLogger = gokitLog.With(a.Logger, LOG_LAYER_TAG, LAYER_APPLICATION)
	a.EndpointLogger = gokitLog.With(a.Logger, LOG_LAYER_TAG, LAYER_ENDPOINT)
	a.TransportLogger = gokitLog.With(a.Logger, LOG_LAYER_TAG, LAYER_TRANSPORT)
}

func (a *App) initDB(dbConfig database.DBConfig) error {
	db, err := database.New(dbConfig)
	if err != nil {
		return err
	}
	a.DB = db
	return nil
}

func (a *App) initHeartbeat() {
	a.Router.Methods("GET").Path("/ping.json").Handler(http.HandlerFunc(healthcheck.Simple))
}

func (a *App) initDetailInfoHandler() {
	recordStore := recordStore.NewRecordStore(a.DB)
	detailInfoService := detailInfoService.NewDetailInfoService(recordStore)

	a.Router.Methods("GET").Path("/records/user/{host_id}/{guest_id}").Handler(gokitHttp.NewServer(
		endpoints.MakeRetrieveEndpoint(a.EndpointLogger, detailInfoService, endpoints.SERVICE_DETAIL_INFO_RETRIEVE),
		detailInfoHttp.DecodeRetrieveRequest,
		encodeJSONResponse,
		a.ErrorEncoder,
		gokitHttp.ServerErrorLogger(gokitLog.With(a.TransportLogger, LOG_ROUTE_TAG, "Retrieve"))))
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
	}()
}

func errorHandler(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set(HTTP_HEADER_CONTENT, HTTP_CONTENT_JSON+HTTP_HEADER_BREAK+HTTP_CONTENT_UTF8)
	var response responses.ErrorResponse

	switch {
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
	json.NewEncoder(w).Encode(&response)
}

func encodeJSONResponse(_ context.Context, writer http.ResponseWriter, response interface{}) error {
	writer.Header().Set(HTTP_HEADER_CONTENT, HTTP_CONTENT_JSON+HTTP_HEADER_BREAK+HTTP_CONTENT_UTF8)
	return json.NewEncoder(writer).Encode(response)
}
