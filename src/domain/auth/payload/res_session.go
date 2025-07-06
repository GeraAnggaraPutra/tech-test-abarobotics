package payload

import (
	"time"

	"abarobotics-test/src/model"
)

type UserResponse struct {
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

type SessionResponse struct {
	AccessToken           string       `json:"access_token"`
	AccessTokenExpiredAt  time.Time    `json:"access_token_expired_at"`
	RefreshToken          string       `json:"refresh_token"`
	RefreshTokenExpiredAt time.Time    `json:"refresh_token_expired_at"`
	User                  UserResponse `json:"user"`
}

func ToSessionResponse(entity model.Session, user model.User) (response SessionResponse) {
	response.AccessToken = entity.AccessToken
	response.AccessTokenExpiredAt = entity.AccessTokenExpiredAt
	response.RefreshToken = entity.RefreshToken
	response.RefreshTokenExpiredAt = entity.RefreshTokenExpiredAt
	response.User.GUID = user.GUID
	response.User.Email = user.Email
	response.User.CreatedAt = user.CreatedAt
	response.User.Role.GUID = user.RoleGUID
	response.User.Role.Name = user.RoleName.String

	if user.CreatedBy.Valid {
		response.User.CreatedBy = &user.CreatedBy.String
	}

	if user.UpdatedAt.Valid {
		response.User.UpdatedAt = &user.UpdatedAt.Time
	}

	if user.UpdatedBy.Valid {
		response.User.UpdatedBy = &user.UpdatedBy.String
	}

	return
}
