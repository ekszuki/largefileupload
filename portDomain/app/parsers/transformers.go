package parsers

import (
	"ekszuki/uploader/portDomain/app/models"
	protoport "ekszuki/uploader/portDomain/protos/port"
)

func FromUpLoadPortRequestToDomain(req *protoport.UpLoadPortRequest) *models.Port {
	dmPort := new(models.Port)

	dmPort.Key = req.GetKey()
	dmPort.Name = req.GetName()
	dmPort.City = req.GetCity()
	dmPort.Country = req.GetCountry()
	dmPort.Alias = append(dmPort.Alias, req.Alias...)
	dmPort.Regions = append(dmPort.Regions, req.Regions...)
	dmPort.Coordinates = append(dmPort.Coordinates, req.Coordinates...)
	dmPort.Province = req.GetProvince()
	dmPort.Timezone = req.GetTimezone()
	dmPort.Unlocs = append(dmPort.Unlocs, req.Unlocs...)
	dmPort.Code = req.Code

	return dmPort
}
