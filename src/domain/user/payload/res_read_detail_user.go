package payload

import (
	"encoding/json"
	"time"

	"abarobotics-test/src/model"
)

type ReadDetailUserResponse struct {
	GUID        string           `json:"guid"`
	Email       string           `json:"email"`
	Role        RoleResponse     `json:"role"`
	CreatedAt   time.Time        `json:"created_at"`
	CreatedBy   *string          `json:"created_by"`
	UpdatedAt   *time.Time       `json:"updated_at"`
	UpdatedBy   *string          `json:"updated_by"`
	Permissions *json.RawMessage `json:"permissions"`
}

func ToReadDetailUserResponse(entity model.User) (response ReadDetailUserResponse, err error) {
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

	if err = json.Unmarshal(entity.Permissions, &response.Permissions); err != nil {
		return
	}

	return
}
