package payload

import (
	"time"

	"abarobotics-test/src/model"
)

type ReadListUserResponse struct {
	GUID      string       `json:"guid"`
	Email     string       `json:"email"`
	Role      RoleResponse `json:"role"`
	CreatedAt time.Time    `json:"created_at"`
	CreatedBy *string      `json:"created_by"`
	UpdatedAt *time.Time   `json:"updated_at"`
	UpdatedBy *string      `json:"updated_by"`
}

type RoleResponse struct {
	GUID string `json:"guid"`
	Name string `json:"name"`
}

func ToReadListUserResponse(entity model.User) (response ReadListUserResponse) {
	response.GUID = entity.GUID
	response.Email = entity.Email
	response.CreatedAt = entity.CreatedAt
	response.Role.GUID = entity.RoleGUID
	response.Role.Name = entity.RoleName.String

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

func ToReadListUserResponses(entities []model.User) (response []ReadListUserResponse) {
	response = make([]ReadListUserResponse, len(entities))

	for i := range entities {
		response[i] = ToReadListUserResponse(entities[i])
	}

	return
}
