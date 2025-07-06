package payload

import (
	"abarobotics-test/src/handler/jwt"
)

type SessionPayload struct {
	SessionGUID string `json:"session_guid"`
	UserGUID    string `json:"user_guid"`
	UserAgent   string `json:"user_agent"`
	IPAddress   string `json:"ip_address"`
}

func (request *SessionPayload) ToAccessTokenRequest() (
	params jwt.AccessTokenPayload,
) {
	params = jwt.AccessTokenPayload{
		GUID:     request.SessionGUID,
		UserGUID: request.UserGUID,
	}

	return
}

func (request *SessionPayload) ToRefreshTokenRequest() (
	params jwt.RefreshTokenPayload,
) {
	params = jwt.RefreshTokenPayload{
		GUID: request.SessionGUID,
	}

	return
}
