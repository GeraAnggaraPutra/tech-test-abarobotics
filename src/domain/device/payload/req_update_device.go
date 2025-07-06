package payload

import "abarobotics-test/src/query"

type UpdateDeviceRequest struct {
	GUID string
	DevicePayload
}

func (req *UpdateDeviceRequest) ToParams(userGUID string) (params query.UpdateDeviceParams) {
	params = query.UpdateDeviceParams{
		GUID:      req.GUID,
		Name:      req.Name,
		Location:  req.Location,
		Status:    req.Status,
		UpdatedBy: userGUID,
	}

	return
}
