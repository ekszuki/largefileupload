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
	dmPort.Alias = req.Alias
	dmPort.Regions = req.Regions
	dmPort.Coordinates = req.Coordinates
	dmPort.Province = req.GetProvince()
	dmPort.Timezone = req.GetTimezone()
	dmPort.Unlocs = req.Unlocs
	dmPort.Code = req.Code

	return dmPort
}

func FromDomainToFindByKeyResponse(dmPort *models.Port) *protoport.FindByKeyResponse {
	resp := new(protoport.FindByKeyResponse)

	resp.Key = dmPort.Key
	resp.Name = dmPort.Name
	resp.City = dmPort.City
	resp.Country = dmPort.Country
	resp.Alias = dmPort.Alias
	resp.Regions = dmPort.Regions
	resp.Coordinates = dmPort.Coordinates
	resp.Province = dmPort.Province
	resp.Timezone = dmPort.Timezone
	resp.Unlocs = dmPort.Unlocs
	resp.Code = dmPort.Code

	return resp
}
