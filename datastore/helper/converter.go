package helper

import (
  storeErr        "mediocris/datastore/error"

  userStoreModels "mediocris/datastore/user/storeModels"
  userDBModels    "mediocris/utils/database/dbModels/user"
)

func serviceToStore(serviceModel interface{}) (interface{}, error) {
  switch serviceModel.(type) {
  case:
  default:
    return nil, storeErr.ErrUnrecognizedServiceModel
  }
}

func dbToStore(dbModel interface{}) (interface{}, error) {
  switch dbModel.(type) {
  case:
  default:
    return nil, storeErr.ErrUnrecognizedDBModel
  }
}
