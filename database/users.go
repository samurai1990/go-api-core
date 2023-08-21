package database

func (user *User) ListFromDB() (*[]User, error) {
	users := []User{
		{
			BaseModel: &BaseModel{
				ID: 22222,
			},
			Username: "samurai",
			Email:    "samurai@com.com",
			ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
			IsAdmin:  true,
		},
		{
			BaseModel: &BaseModel{
				ID: 123123,
			},
			Username: "bob",
			Email:    "bob@com.com",
			ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
			IsAdmin:  false,
		},
	}
	return &users, nil
}

func (user *User) GetByID(id uint) (*User, error) {
	user = &User{
		BaseModel: &BaseModel{
			ID: id,
		},
		Username: "samurai",
		Email:    "samurai@com.com",
		IsAdmin:  true,
	}
	return user, nil
}

func (user *User) GetByName(name string) (*User, error) {
	user = &User{
		BaseModel: &BaseModel{
			ID: 13213,
		},
		Username: "samurai",
		Password: "858a39db1542daacc92d9bb4fb8b563d35e13833cfc35e4ec2106c2043aa4bfc5a4373350267278dbcb8ee34c214898dfe27ce286b615b74bcb23642a9a067b5",
		Email:    "samurai@com.com",
		ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
		IsAdmin:  true,
	}
	return user, nil
}

func (user *User) GetByApiKey(apiKey string) (*User, error) {
	user = &User{
		BaseModel: &BaseModel{
			ID: 13213,
		},
		Username: "samurai",
		Password: "858a39db1542daacc92d9bb4fb8b563d35e13833cfc35e4ec2106c2043aa4bfc5a4373350267278dbcb8ee34c214898dfe27ce286b615b74bcb23642a9a067b5",
		Email:    "samurai@com.com",
		ApiKey:   "KjClltV/gjuAwKBVqbpEoLJ9YIfsvQoC9d/csOkhPeLPm5aI9UkLAzgmBNbgvRRb+7DoJx5KIevCbi0FiMzDoQ==",
		IsAdmin:  true,
	}
	return user, nil
}
