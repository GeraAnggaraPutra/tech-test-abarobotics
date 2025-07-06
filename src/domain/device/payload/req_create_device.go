package payload

import (
	"abarobotics-test/src/query"
	"abarobotics-test/src/util"
)

type CreateDeviceRequest struct {
	DevicePayload
}

func (req *CreateDeviceRequest) ToParams(userGUID string) (params query.CreateDeviceParams) {
	params = query.CreateDeviceParams{
		GUID:      util.GenerateUUID(),
		Name:      req.Name,
		Location:  req.Location,
		Status:    req.Status,
		CreatedBy: userGUID,
	}

	return
}
