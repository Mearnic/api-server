package api

import (
	c "api-server/src/common"
	s "api-server/src/service"
)

func (r Request) StormAlarmList() c.RequestDefine {
	return Get("/v7/tropical/storm-list", s.GetStormAlarmList())
}

func (r Request) StormAlarm() c.RequestDefine {
	return Get("/v7/tropical/storm-forecast", s.GetStormAlarm())
}

func (r Request) StormAlarmHistory() c.RequestDefine {
	return Get("/v7/tropical/storm-track", s.GetStormAlarmHistory())
}
