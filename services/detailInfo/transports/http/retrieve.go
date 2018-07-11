package http

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"wechat-miniprogram/services/detailInfo/serviceModels"
	"wechat-miniprogram/services/errors"
)

const (
	HOST_ID_KEY  = "host_id"
	GUEST_ID_KEY = "guest_id"
	GROUP_ID_KEY = "group_id"
	EMPTY_ID     = ""
)

// Decodes a raw incoming http request to readable object for services
func DecodeRetrieveRequest(_ conetext.Context, req *http.Request) (interface{}, error) {

	// extracts params from path
	vars := mux.Vars(req)
	hostID := vars[HOST_ID_KEY]
	guestID := vars[GUEST_ID_KEY]
	groupIDStr := vars[GROUP_ID_KEY]

	// do not allow empty host id
	if hostID == EMPTY_ID {
		return nil, errors.ErrInsufficientParams
	}

	// deal with group id. if no group id provided, sub with -1
	var groupID int
	if groupIDStr == EMPTY_ID {
		groupID = -1
	} else {
		groupID = strconv.Atoi(groupIDStr)
	}

	// returns retrieve params
	return serviceModels.DetailRetrieveParams{hostID, guestID, groupID}, nil
}
