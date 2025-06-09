package usecase

import (
	"context"
	"time"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/repository"
	"github.com/amirdashtii/AutoBan/internal/validation"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"github.com/google/uuid"
)

type UserUseCase interface {
	GetProfile(ctx context.Context, userID string) (*dto.GetProfileResponse, error)
	UpdateProfile(ctx context.Context, userID string, request dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error)
	ChangePassword(ctx context.Context, userID string, request dto.UpdatePasswordRequest) error
	DeleteUser(ctx context.Context, userID string) error
}

type userUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase() UserUseCase {
	userRepository := repository.NewUserRepository()
	return &userUseCase{userRepository: userRepository}
}

func (u *userUseCase) GetProfile(ctx context.Context, userID string) (*dto.GetProfileResponse, error) {
	var user entity.User
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return nil, errors.ErrInvalidUserID
	}
	user.ID = userUUID
	err = u.userRepository.GetProfile(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to get profile")
		return nil, errors.ErrFailedToGetProfile
	}
	birthday := ""
	if user.Birthday != nil {
		birthday = user.Birthday.Format("2006-01-02")
	}
	return &dto.GetProfileResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Birthday:  &birthday,
	}, nil
}

func (u *userUseCase) UpdateProfile(ctx context.Context, userID string, request dto.UpdateProfileRequest) (*dto.UpdateProfileResponse, error) {
	err := validation.ValidateUpdateProfileRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate update profile request")
		return nil, err
	}

	var birthday *time.Time
	if request.Birthday != nil {
		parsedTime, err := time.Parse("2006-01-02", *request.Birthday)
		if err != nil {
			logger.Error(err, "Failed to parse birthday")
			return nil, errors.ErrInvalidBirthday
		}
		birthday = &parsedTime
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return nil, errors.ErrInvalidUserID
	}

	user := entity.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		Birthday:  birthday,
	}
	user.ID = userUUID

	err = u.userRepository.UpdateProfile(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to update profile")
		return nil, errors.ErrFailedToUpdateProfile
	}

	responseBirthday := ""
	if user.Birthday != nil {
		responseBirthday = user.Birthday.Format("2006-01-02")
	}
	return &dto.UpdateProfileResponse{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Birthday:  &responseBirthday,
	}, nil
}

func (u *userUseCase) ChangePassword(ctx context.Context, userID string, request dto.UpdatePasswordRequest) error {
	err := validation.ValidateUpdatePasswordRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate update password request")
		return err
	}
	
	user := entity.User{
		Password: request.Password,
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return errors.ErrInvalidUserID
	}
	user.ID = userUUID

	err = u.userRepository.ChangePassword(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to update password")
		return errors.ErrFailedToUpdatePassword
	}
	return nil
}

func (u *userUseCase) DeleteUser(ctx context.Context, userID string) error {
	var user entity.User
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return errors.ErrInvalidUserID
	}
	user.ID = userUUID
	
	err = u.userRepository.DeleteUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to delete user")
		return errors.ErrFailedToDeleteUser
	}
	return nil
}
