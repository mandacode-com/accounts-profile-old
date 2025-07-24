package dto

import (
	"github.com/google/uuid"
	dbrepo "mandacode.com/accounts/profile/internal/repository/database"
)

type CreateProfileData struct {
	UserID uuid.UUID `json:"user_id"`
	Email  string    `json:"email,omitempty"`
}

func (data CreateProfileData) ToRepoModel(nickname string) *dbrepo.CreateProfileModel {
	return &dbrepo.CreateProfileModel{
		UserID:   data.UserID,
		Email:    data.Email,
		Nickname: nickname,
	}
}
