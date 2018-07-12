package helper

import (
	storeErr "wechat-miniprogram/datastore/error"

	detailInfoServiceModels "wechat-miniprogram/services/detailInfo/serviceModels"

	recordStoreModels "wechat-miniprogram/datastore/record/storeModels"
	// userStoreModels "wechat-miniprogram/datastore/user/storeModels"
	recordDBModels "wechat-miniprogram/utils/database/dbModels/record"
	// userDBModels "wechat-miniprogram/utils/database/dbModels/user"
)

func ServiceToStore(serviceModel interface{}) (interface{}, error) {
	switch serviceModel.(type) {
	case detailInfoServiceModels.DetailRetrieveParams:
		recordRetrieval := serviceModel.(detailInfoServiceModels.DetailRetrieveParams)
		return recordRetrieveServiceToStore(recordRetrieval), nil
	default:
		return nil, storeErr.ErrUnrecognizedServiceModel
	}
}

func recordRetrieveServiceToStore(serviceModel detailInfoServiceModels.DetailRetrieveParams) recordStoreModels.RecordRetrieveParams {
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
