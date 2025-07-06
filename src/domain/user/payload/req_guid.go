package payload

type GUIDRequest struct {
	GUID string `param:"guid" validate:"required"`
}
