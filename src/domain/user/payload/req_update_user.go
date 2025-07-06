package payload

import "abarobotics-test/src/query"

type UpdateUserRequest struct {
	GUID string
	UserPayload
}

func (req *UpdateUserRequest) ToParams(userGUID, password string) (params query.UpdateUserParams) {
	params = query.UpdateUserParams{
		GUID:      req.GUID,
		Email:     req.Email,
		Password:  password,
		RoleGUID:  req.RoleGUID,
		UpdatedBy: userGUID,
	}

	return
}
