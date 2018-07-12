package records

import (
	"context"

	"wechat-miniprogram/datastore"
	"wechat-miniprogram/services"

	"wechat-miniprogram/services/detailInfo/serviceModels"
	recordStoreModels "wechat-miniprogram/datastore/record/storeModels"

	"wechat-miniprogram/services/helper"
	serviceErr "wechat-miniprogram/services/errors"
)

// Service for Detailed info
// Used for two situations:
// 1. records page between users
// 2. records page inside a group
type DetailInfoService struct {
	// GroupStore  datastore.Store
	RecordStore datastore.Store
}

// Constructor
func NewDetailInfoService(recordStore datastore.Store) services.Service {
	return DetailInfoService{recordStore}
}

// Retrieves detailed infos
// compulsary param: host_id
// possible params:
// - guest_id
// - group_id
func (s DetailInfoService) Retrieve(_ context.Context, args interface{}) (interface{}, error) {
	infoRetrieveParams, ok := args.(serviceModels.DetailRetrieveParams)
	if !ok {
		return nil, serviceErr.ErrIncorrectParamsFormat
	}

	// If has group id, need to get group info
	// Else first
	records, err := s.RecordStore.Retrieve(infoRetrieveParams)
	if err != nil {
		return nil, err
	}
	castedRecords := records.([]*recordStoreModels.RecordRetrieveResult)
	return helper.GenerateDetailBetweenUsers(castedRecords, infoRetrieveParams.HostID, infoRetrieveParams.GuestID)
}

func (s DetailInfoService) Create(ctx context.Context, args interface{}) (interface{}, error) {
	return nil, nil
}

func (s DetailInfoService) Update(ctx context.Context, args interface{}) (interface{}, error) {
	return nil, nil
}

func (s DetailInfoService) Delete(ctx context.Context, args interface{}) (interface{}, error) {
	return nil, nil
}
