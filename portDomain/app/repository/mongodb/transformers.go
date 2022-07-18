package mongodb

import "ekszuki/uploader/portDomain/app/models"

func (db *Port) fromDomain(dmPort *models.Port) error {
	db.Key = dmPort.Key
	db.Name = dmPort.Name
	db.City = dmPort.City
	db.Country = dmPort.Country
	db.Alias = dmPort.Alias
	db.Regions = dmPort.Regions
	db.Coordinates = dmPort.Coordinates
	db.Province = dmPort.Province
	db.Timezone = dmPort.Timezone
	db.Unlocs = dmPort.Unlocs
	db.Code = dmPort.Code

	return nil
}
