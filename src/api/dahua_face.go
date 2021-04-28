package api

import (
	c "api-server/src/common"
	s "api-server/src/service"
)

func (r Request) VehicleFindFace() c.RequestDefine {
	return Get("/multiDataMiningService/rest/faceAndVehicle/vehicleFindFace", s.GetVehicleFindFace())
}

func (r Request) FaceFindVehicle() c.RequestDefine {
	return Get("/multiDataMiningService/rest/faceAndVehicle/faceFindVehicle", s.GetFaceFindVehicle())
}
