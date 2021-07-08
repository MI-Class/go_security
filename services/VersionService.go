package services

import (
	"testaa/db"
	"testaa/dto"
	"testaa/util"
)

type VersionService struct {
}

func (main VersionService) GetAll() ([]*dto.Version, error) {

	rows, err := db.DB.Query("SELECT * FROM version")
	if err != nil {
		return dto.VersionList, err
	}

	for rows.Next() {
		f := new(dto.Version)
		rows.Scan(&f.Id, &f.Version, &f.LastVersion, &f.CreatTime)
		f.Version, _ = util.Decrypt(f.Version)
		f.LastVersion, _ = util.Decrypt(f.LastVersion)
		f.CreatTime, _ = util.Decrypt(f.CreatTime)
		dto.VersionList = append(dto.VersionList, f)
	}
	return dto.VersionList, nil
}

func (v VersionService) GetPW() (string, error) {
	var pw string
	rows := db.DB.QueryRow("select id,lastVersion from version where id = 1")
	rows.Scan(&pw)
	return pw, nil
}
