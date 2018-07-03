package record

import (
  mDB               "mediocris/utils/database"
  mStore            "mediocris/datastore"
  converter         "mediocris/datastore/helper"
  recordStoreModels "mediocris/datastore/record/storeModels"
  recordDBModels    "mediocris/utils/database/dbModels"
  storeErr          "mediocris/datastore/error"
)

type RecordStore struct {
  DB mDB.Database
}

func NewRecordStore(db mDB.Database) mStore.Store {
  return RecordStore{db}
}

// Creates a record. all arguments are included in args
func (s RecordStore) Create(args interface{}) (interface{}, error) {
  
}

// Retrieves record(s).
func (s RecordStore) Retrieve(args interface{}) (interface{}, error) {

}

// Updates a record. id should be specified
func (s RecordStore) Update(id string, args interface{}) (interface{}, error) {

}

// Deletes a record. id should be specified
func (s RecordStore) Delete(id string, args interface{}) (interface{}, error) {

}
