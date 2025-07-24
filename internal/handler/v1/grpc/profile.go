package grpchandlerv1

import (
	"context"
	"github.com/google/uuid"
	profilev1 "github.com/mandacode-com/accounts-proto/go/profile/v1"
	"github.com/mandacode-com/golib/errors"
	"github.com/mandacode-com/golib/errors/errcode"
	"go.uber.org/zap"
	"google.golang.org/protobuf/types/known/timestamppb"
	"mandacode.com/accounts/profile/internal/usecase/dto"
	"mandacode.com/accounts/profile/internal/usecase/system"
)

type ProfileHandler struct {
	profilev1.UnimplementedProfileServiceServer
	profile *system.ProfileUsecase
	logger  *zap.Logger
}

// UpdateEmail implements profilev1.ProfileServiceServer.
func (u *ProfileHandler) UpdateEmail(ctx context.Context, req *profilev1.UpdateEmailRequest) (*profilev1.UpdateEmailResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Upgrade(err, "Failed to validate UpdateEmailRequest", errcode.ErrInvalidFormat)
	}
	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errors.Upgrade(err, "Failed to parse user ID", errcode.ErrInvalidFormat)
	}
	profile, err := u.profile.UpdateProfile(ctx, &dto.UpdateProfileData{
		UserID: userID,
		Email:  &req.NewEmail,
	})
	if err != nil {
		return nil, err
	}

	return &profilev1.UpdateEmailResponse{
		UserId:       profile.UserID.String(),
		UpdatedEmail: profile.Email,
		UpdatedAt:    timestamppb.Now(),
	}, nil
}

// DeleteUser implements userv1.UserServiceServer.
func (u *ProfileHandler) DeleteUser(ctx context.Context, req *profilev1.DeleteUserRequest) (*profilev1.DeleteUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Upgrade(err, "Failed to validate DeleteUserRequest", errcode.ErrInvalidFormat)
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errors.Upgrade(err, "Failed to parse user ID", errcode.ErrInvalidFormat)
	}

	err = u.profile.DeleteProfile(ctx, userID)
	if err != nil {
		return nil, err
	}

	return &profilev1.DeleteUserResponse{
		UserId:    req.UserId,
		DeletedAt: timestamppb.Now(),
	}, nil
}

// InitUser implements userv1.UserServiceServer.
func (u *ProfileHandler) InitUser(ctx context.Context, req *profilev1.InitUserRequest) (*profilev1.InitUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, errors.Upgrade(err, "Failed to validate InitUserRequest", errcode.ErrInvalidFormat)
	}

	userID, err := uuid.Parse(req.UserId)
	if err != nil {
		return nil, errors.Upgrade(err, "Failed to parse user ID", errcode.ErrInvalidFormat)
	}

	user, err := u.profile.CreateProfile(ctx, &dto.CreateProfileData{
		UserID: userID,
		Email:  req.Email,
	})
	if err != nil {
		return nil, err
	}

	return &profilev1.InitUserResponse{
		UserId:        user.UserID.String(),
		InitializedAt: timestamppb.Now(),
	}, nil

}

// NewProfileHandler creates a new UserSystemHandler with the provided use case.
func NewProfileHandler(profile *system.ProfileUsecase, logger *zap.Logger) profilev1.ProfileServiceServer {
	return &ProfileHandler{
		profile: profile,
		logger:  logger,
	}
}
