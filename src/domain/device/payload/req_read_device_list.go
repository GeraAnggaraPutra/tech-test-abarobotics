package payload

import (
	"abarobotics-test/src/query"
	"abarobotics-test/src/util"
)

type ReadDeviceListRequest struct {
	util.PaginationPayload
	Status string `query:"status"`
}

func (req *ReadDeviceListRequest) ToParams() (params query.ReadListDeviceParams) {
	req.Init()

	params = query.ReadListDeviceParams{
		SetSearch: req.SetSearch,
		Search:    req.Search,
		Order:     req.Order,
		Offset:    req.Offset,
		Limit:     req.Limit,
	}

	if req.Status != "" {
		params.SetStatus = true
		params.Status = req.Status
	}

	return
}
