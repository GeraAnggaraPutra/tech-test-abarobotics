package payload

import (
	"time"

	"abarobotics-test/src/model"

)

type ReadDeviceResponse struct {
	GUID      string     `json:"guid"`
	Name      string     `json:"name"`
	Location  string     `json:"location"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	CreatedBy *string    `json:"created_by"`
	UpdatedAt *time.Time `json:"updated_at"`
	UpdatedBy *string    `json:"updated_by"`
}

func ToReadDeviceResponse(entity model.Device) (response ReadDeviceResponse) {
	response.GUID = entity.GUID
	response.Name = entity.Name
	response.Location = entity.Location
	response.Status = entity.Status
	response.CreatedAt = entity.CreatedAt

	if entity.CreatedBy.Valid {
		response.CreatedBy = &entity.CreatedBy.String
	}

	if entity.UpdatedAt.Valid {
		response.UpdatedAt = &entity.UpdatedAt.Time
	}

	if entity.UpdatedBy.Valid {
		response.UpdatedBy = &entity.UpdatedBy.String
	}

	return
}

func ToReadDeviceResponses(entities []model.Device) (response []ReadDeviceResponse) {
	response = make([]ReadDeviceResponse, len(entities))

	for i := range entities {
		response[i] = ToReadDeviceResponse(entities[i])
	}

	return
}
