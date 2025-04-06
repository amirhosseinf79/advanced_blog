package dto

type RefreshTokenRequest struct {
	RefreshToken string `form:"refresh_token" json:"refresh_token" validate:"required"`
}
