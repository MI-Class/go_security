package services

import (
	"testaa/db"
	"testaa/dto"
)

type MainDataService struct {
}

func (main MainDataService) GetAll() ([]*dto.Main, error) {
	dto.MainList=
	rows, err := db.DB.Query("SELECT * FROM main")
	if err != nil {
		return "", err
	}
	for rows.Next() {

	}
	return "", nil
}

func (main MainDataService)GetOne (id string)  {
	db.DB.Query()
}
