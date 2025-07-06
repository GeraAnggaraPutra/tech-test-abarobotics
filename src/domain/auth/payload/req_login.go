package payload

import "abarobotics-test/src/util"

type LoginRequest struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required"`
}

func (request *LoginRequest) ToSessionPayload(userGUID, userAgent, iPAddress string) (
	params SessionPayload,
) {
	params = SessionPayload{
		SessionGUID: util.GenerateUUID(),
		UserGUID:    userGUID,
		UserAgent:   userAgent,
		IPAddress:   iPAddress,
	}

	return
}
