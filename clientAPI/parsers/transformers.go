package parsers

import (
	"ekszuki/uploader/portDomain/app/models"
	"ekszuki/uploader/portDomain/protos/port"
)

func ToUpLoadPortRequest(data *PortFile) *port.UpLoadPortRequest {
	loadPortRequest := new(port.UpLoadPortRequest)
	loadPortRequest.Key = data.Key
	loadPortRequest.Name = data.Name
	loadPortRequest.City = data.City
	loadPortRequest.Country = data.Country
	loadPortRequest.Alias = append(loadPortRequest.Alias, data.Alias...)
	loadPortRequest.Regions = append(loadPortRequest.Regions, data.Regions...)
	loadPortRequest.Coordinates = append(loadPortRequest.Coordinates, data.Coordinates...)
	loadPortRequest.Province = data.Province
	loadPortRequest.Timezone = data.Timezone
	loadPortRequest.Unlocs = append(loadPortRequest.Unlocs, data.Unlocs...)
	loadPortRequest.Code = data.Code

	return loadPortRequest
}

func ToPortDomain(data *PortFile) *models.Port {
	dmPort := new(models.Port)
	dmPort.Key = data.Key
	dmPort.Name = data.Name
	dmPort.City = data.City
	dmPort.Country = data.Country
	dmPort.Alias = append(dmPort.Alias, data.Alias...)
	dmPort.Regions = append(dmPort.Regions, data.Regions...)
	dmPort.Coordinates = append(dmPort.Coordinates, data.Coordinates...)
	dmPort.Province = data.Province
	dmPort.Timezone = data.Timezone
	dmPort.Unlocs = append(dmPort.Unlocs, data.Unlocs...)
	dmPort.Code = data.Code

	return dmPort
}
