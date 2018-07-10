package helper

import (
  storeErr          "mediocris/datastore/error"

  serviceModels     "mediocris/services/serviceModels"

  userStoreModels   "mediocris/datastore/user/storeModels"
  recordStoreModels "mediocris/datastore/record/storeModels"
  userDBModels      "mediocris/utils/database/dbModels/user"
  recordDBModels    "mediocris/utils/database/dbModels/record"
)

func ServiceToStore(serviceModel interface{}) (interface{}, error) {
  switch serviceModel.(type) {
  case serviceModels.RecordRetrieveParams:
    recordRetrieval := serviceModel.(serviceModels.RecordRetrieveParams)
    return recordRetrieveServiceToStore(recordRetrieval), nil
  default:
    return nil, storeErr.ErrUnrecognizedServiceModel
  }
}

func recordRetrieveServiceToStore(serviceModel serviceModels.RecordRetrieveParams) recordStoreModels.RecordRetrieveParams {
  return recordStoreModels.RecordRetrieveParams{
    HostID:  serviceModel.HostID,
    GuestID: serviceModel.GuestID,
    GroupID: serviceModel.GroupID,
  }
}

func DBRecordsToStore(dbRecords []recordDBModels.TansRecord) []*recordStoreModels.RecordRetrieveResult {
  result := make([]*recordStoreModels.RecordRetrieveResult, 0)
  for _, dbRecord := range dbRecords {
    temp := recordStoreModels.RecordRetrieveResult{
      RecordID:    dbRecord.RecordID,
      GroupID:     dbRecord.GroupID,
      Date:        dbRecord.Date,
      Payer:       dbRecord.Payer,
      Spliters:    dbRecord.Spliters,
      Amount:      dbRecord.Amount,
      Description: dbRecord.Description,
    }
    result = append(result, &temp)
  }
  return result
}
