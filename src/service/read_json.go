package service

import (
	"api-server/src/io"
)

func GetVehicleFindFace() string {
	return io.ReadJson("vehicleFindFace.json")
}

func GetFaceFindVehicle() string {
	return io.ReadJson("faceFindVehicle.json")
}
