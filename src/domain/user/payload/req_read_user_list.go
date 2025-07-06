package payload

import (
	"abarobotics-test/src/query"
	"abarobotics-test/src/util"
)

type ReadUserListRequest struct {
	util.PaginationPayload
}

func (req *ReadUserListRequest) ToParams() (params query.ReadListUserParams) {
	req.Init()

	params = query.ReadListUserParams{
		SetSearch: req.SetSearch,
		Search:    req.Search,
		Order:     req.Order,
		Offset:    req.Offset,
		Limit:     req.Limit,
	}

	return
}
