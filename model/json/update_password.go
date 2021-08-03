package json

type UpdatePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"password"`
}
