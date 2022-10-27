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

func GetStormAlarmList() string {
	return io.ReadJson("storm_alarm_list.json")
}

func GetStormAlarm() string {
	return io.ReadJson("storm_alarm.json")
}

func GetStormAlarmHistory() string {
	return io.ReadJson("storm_alarm_history.json")
}
