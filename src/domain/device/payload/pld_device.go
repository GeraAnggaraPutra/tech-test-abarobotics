package payload

type DevicePayload struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Status   string `json:"status" validate:"required,oneof=online offline"`
}
