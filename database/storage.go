package database

type CRUDInterface interface {
	ListFromDB() (*[]User, error)
	GetByID(id uint) (*User, error)
	GetByName(id string) (*User, error)
	GetByApiKey(apiKey string) (*User, error)
}

func (h *DBHandler) ListDatabase(i CRUDInterface) (any, error) {
	queries, err := i.ListFromDB()
	if err != nil {
		return nil, err
	} else {
		return queries, nil
	}
}

func (h *DBHandler) GetDatabaseByName(i CRUDInterface, name string) (any, error) {
	query, err := i.GetByName(name)
	if err != nil {
		return nil, err
	} else {
		return query, nil
	}
}

func (h *DBHandler) GetDatabaseByApiKey(i CRUDInterface, apiKey string) (any, error) {
	query, err := i.GetByApiKey(apiKey)
	if err != nil {
		return nil, err
	} else {
		return query, nil
	}
}
