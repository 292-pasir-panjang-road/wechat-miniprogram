package helper

import (
	"encoding/json"

	recordStoreModels "wechat-miniprogram/datastore/record/storeModels"
	recordServiceModels "wechat-miniprogram/services/detailInfo/serviceModels"
)

func GenerateDetailBetweenUsers(records []*recordStoreModels.RecordRetrieveResult, hostID string, guestID string) (*recordServiceModels.DetailBetweenUsers, error) {
	userRecords := make([]*recordServiceModels.UserRecord, 0)
	for _, record := range records {

		totalAmount := record.Amount
		trans := totalAmount / float32(len(record.Spliters))

		// Which means that host borrows money from guest. trans need to be negative
		if record.Payer != hostID {
			trans = -1 * trans
		}
		userRecord := recordServiceModels.UserRecord{
			RecordID:    record.RecordID,
			GroupID:     record.GroupID,
			Date:        record.Date,
			Amount:      trans,
			Description: record.Description,
		}
		userRecords = append(userRecords, &userRecord)
	}
	return &recordServiceModels.DetailBetweenUsers{hostID, guestID, userRecords}, nil
}

func GenerateGroupDetails(records []*recordStoreModels.RecordRetrieveResult) (*recordServiceModels.GroupDetails, error) {
	return nil, nil
}

func ObjToString(object interface{}) string {
	objBytes, _ := json.Marshal(object)
	return string(objBytes)
}
