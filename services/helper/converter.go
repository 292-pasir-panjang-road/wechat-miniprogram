package converter

import (
  "encoding/json"

  recordStoreModels   "wechat-miniprogram/datastore/record/storeModels"
  recordServiceModels "wechat-miniprogram/services/record/storeModels"
)

func GenerateDetailBetweenUsers(records []*recordStoreModels.RecordRetrieveResult, hostID string, guestID string) (*recordServiceModels.DetailBetweenUsers, error) {
  userRecords := make([]*recordServiceModels.UserRecord, 0)
  for _, record := range records {

    totalAmount := record.Amount
    trans := totalAmount / len(record.Spliters)

    // Which means that host borrows money from guest. trans need to be negative
    if record.PayerID != hostID {
      trans = -1 * trans
    }
    userRecord := recordServiceModels.UserRecord{
      RecordID    record.RecordID
      GroupID     record.GroupID
      Date        record.Date
      Amount      trans
      Description record.Description
    }
    userRecords = append(userRecords, &userRecord)
  }
  return &recordServiceModels.DetailBetweenUsers{
    HostID:  hostID,
    GuestID: guestID,
    Records: userRecords,
  }, nil
}

func GenerateGroupDetails(records []*recordStoreModels.RecordRetrieveResult) (*recordServiceModels.GroupDetails, error) {
  return nil, nil
}

func ObjToString(object interface{}) string {
  objBytes, _ := json.Marshal(object)
  return string(objBytes)
}
