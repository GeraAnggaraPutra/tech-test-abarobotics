package payload

import (
	"abarobotics-test/src/query"
	"abarobotics-test/src/util"
)

type CreateUserRequest struct {
	UserPayload
}

func (req *CreateUserRequest) ToParams(userGUID, password string) (params query.CreateUserParams) {
	params = query.CreateUserParams{
		GUID:      util.GenerateUUID(),
		Email:     req.Email,
		Password:  password,
		RoleGUID:  req.RoleGUID,
		CreatedBy: userGUID,
	}

	return
}
