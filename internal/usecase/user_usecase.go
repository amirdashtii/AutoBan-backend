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

	return &dto.GetProfileResponse{
		ID:          user.ID,
		PhoneNumber: user.PhoneNumber,
		FirstName:   user.FirstName,
		LastName:    user.LastName,
		Email:       user.Email,
		Birthday:    user.Birthday.Format("2006-01-02"),
		Status:      user.Status.String(),
		Role:        user.Role.String(),
		CreatedAt:   user.CreatedAt.Format("2006-01-02"),
		UpdatedAt:   user.UpdatedAt.Format("2006-01-02"),
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

	user := entity.User{}
	user.ID = userUUID

	if request.FirstName != nil {
		user.FirstName = *request.FirstName
	}
	if request.LastName != nil {
		user.LastName = *request.LastName
	}
	if request.Email != nil {
		user.Email = *request.Email
	}
	if request.Birthday != nil {
		user.Birthday = *birthday
	}

	err = u.userRepository.UpdateProfile(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to update profile")
		if err == errors.ErrEmailAlreadyExists {
			return nil, err
		}
		return nil, errors.ErrFailedToUpdateProfile
	}

	// Get the updated user with all fields
	updatedUser := entity.User{}
	updatedUser.ID = userUUID
	err = u.userRepository.GetProfile(ctx, &updatedUser)
	if err != nil {
		logger.Error(err, "Failed to get updated user")
		return nil, errors.ErrFailedToGetProfile
	}

	return &dto.UpdateProfileResponse{
		ID:          updatedUser.ID,
		PhoneNumber: updatedUser.PhoneNumber,
		FirstName:   updatedUser.FirstName,
		LastName:    updatedUser.LastName,
		Email:       updatedUser.Email,
		Birthday:    updatedUser.Birthday.Format("2006-01-02"),
		Status:      updatedUser.Status.String(),
		Role:        updatedUser.Role.String(),
		CreatedAt:   updatedUser.CreatedAt.Format("2006-01-02"),
		UpdatedAt:   updatedUser.UpdatedAt.Format("2006-01-02"),
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
