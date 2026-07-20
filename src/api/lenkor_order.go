package api

import (
	c "api-server/src/common"
	s "api-server/src/service"
)

func (r Request) LenkorOrder() c.RequestDefine {
	return Get("/v7/lenkor/order", s.GetLenkorOrder())
}



func (r Request) LenkorOrderCancel() c.RequestDefine {
	return Post("/openapi/tongtool/orderCancel", s.GetOrderCancel())
}
