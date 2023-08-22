package database

import "github.com/google/uuid"

type CRUDInterface interface {
	ListRecord() (*[]User, error)
	GetRecordByID(id uuid.UUID) (*User, error)
	GetRecordByName(name string) (*User, error)
	GetRecordByApiKey(apiKey string) (*User, error)
	CreateRecord(model any) (*User, error)
	DeleteRecord(id string) error
	UpdateRecord(model any) (*User, error)
}


type SwapVariable interface{
	User| BaseModel
}

func (h *DBHandler) ListDatabase(i CRUDInterface) (any, error) {
	queries, err := i.ListRecord()
	if err != nil {
		return nil, err
	} else {
		return queries, nil
	}
}

func (h *DBHandler) GetRecordDatabaseByName(i CRUDInterface, name string) (any, error) {
	query, err := i.GetRecordByName(name)
	if err != nil {
		return nil, err
	} else {
		return query, nil
	}
}

func (h *DBHandler) GetRecordDatabaseByApiKey(i CRUDInterface, apiKey string) (any, error) {
	query, err := i.GetRecordByApiKey(apiKey)
	if err != nil {
		return nil, err
	} else {
		return query, nil
	}
}

func (h *DBHandler) GetRecordDatabaseByID(i CRUDInterface, id uuid.UUID) (any, error) {
	query, err := i.GetRecordByID(id)
	if err != nil {
		return nil, err
	} else {
		return query, nil
	}
}

func (h *DBHandler) CreateRecordToDatabase(i CRUDInterface) (any, error) {
	query, err := i.CreateRecord(h.Model)
	if err != nil {
		return nil, err
	} else {
		return query, nil
	}
}

func (h *DBHandler) UpdateRecordToDatabase(i CRUDInterface) (any, error) {
	query, err := i.UpdateRecord(h.Model)
	if err != nil {
		return nil, err
	} else {
		return query, nil
	}
}

func (h *DBHandler) DeleteRecordFromDatabase(i CRUDInterface, id string) error {
	return i.DeleteRecord(id)
}
