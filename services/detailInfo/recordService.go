package records

import (
  "context"

  "wechat-miniprogram/services"
  "wechat-miniprogram/datastore"

  "wechat-miniprogram/services/detailInfo/serviceModels"

  "wechat-miniprogram/services/errors"
  "wechat-miniprogram/services/helper"
)

type DetailInfoService struct {
  GroupStore  datastore.Store
  RecordStore datastore.Store
}

func NewDetailInfoService(groupStore datastore.Store, recordStore datastore.Store) services.Service {
  return DetailInfoService{groupStore, recordStore}
}

func (s DetailInfoService) Retrieve(_ context.Context, args interface{}) (interface{}, error) {
  infoRetrieveParams, ok := args.(serviceModels.detailInfoServiceModels)
  if !ok {
    return nil, errors.ErrIncorrectParamsFormat
  }

  // If has group id, need to get group info
  // Else first
  records, err := RecordStore.Retrieve(infoRetrieveParams)
  if err != nil {
    return nil, err
  }

  return helper.GenerateDetailBetweenUsers(records, infoRetrieveParams.HostID, infoRetrieveParams.GuestID), nil
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
