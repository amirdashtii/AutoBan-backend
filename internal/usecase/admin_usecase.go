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
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase interface {
	ListUsers(ctx context.Context) (*dto.ListUserResponse, error)
	GetUserById(ctx context.Context, userID string) (*dto.User, error)
	UpdateUser(ctx context.Context, userID string, request dto.UpdateUserRequest) error
	ChangeUserRole(ctx context.Context, userID string, request dto.ChangeUserRoleRequest) error
	ChangeUserStatus(ctx context.Context, userID string, request dto.ChangeUserStatusRequest) error
	ChangeUserPassword(ctx context.Context, userID string, request dto.ChangeUserPasswordRequest) error
	DeleteUser(ctx context.Context, userID string) error
}

type adminUseCase struct {
	adminRepository repository.AdminRepository
}

func NewAdminUseCase() AdminUseCase {
	adminRepository := repository.NewAdminRepository()
	return &adminUseCase{
		adminRepository: adminRepository,
	}
}

func (u *adminUseCase) ListUsers(ctx context.Context) (*dto.ListUserResponse, error) {
	users := []entity.User{}
	err := u.adminRepository.ListUsers(ctx, &users)
	if err != nil {
		logger.Error(err, "Failed to list users")
		return nil, errors.ErrFailedToListUsers
	}

	userResponse := []dto.User{}
	for _, user := range users {
		userResponse = append(userResponse, dto.User{
			ID:        user.ID.String(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Phone:     user.PhoneNumber,
			Role:      user.Role.String(),
			Status:    user.Status.String(),
			Birthday:  user.Birthday.Format("2006-01-02"),
		})
	}

	return &dto.ListUserResponse{Users: userResponse}, nil
}

func (u *adminUseCase) GetUserById(ctx context.Context, userID string) (*dto.User, error) {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return nil, errors.ErrInvalidUserID
	}

	var user entity.User
	user.ID = userUUID

	err = u.adminRepository.GetUserById(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to get user by id")
		return nil, errors.ErrFailedToGetUserById
	}

	return &dto.User{
		ID:        user.ID.String(),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
		Phone:     user.PhoneNumber,
		Role:      user.Role.String(),
		Status:    user.Status.String(),
		Birthday:  user.Birthday.Format("2006-01-02"),
	}, nil
}

func (u *adminUseCase) UpdateUser(ctx context.Context, userID string, request dto.UpdateUserRequest) error {
	err := validation.AdminValidateUpdateProfileRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate update user request")
		return err
	}
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return errors.ErrInvalidUserID
	}

	var birthday *time.Time
	if request.Birthday != nil {
		parsedTime, err := time.Parse("2006-01-02", *request.Birthday)
		if err != nil {
			logger.Error(err, "Failed to parse birthday")
			return errors.ErrInvalidBirthday
		}
		birthday = &parsedTime
	}

	var user entity.User
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
	if request.Phone != nil {
		user.PhoneNumber = *request.Phone
	}
	if birthday != nil {
		user.Birthday = *birthday
	}

	err = u.adminRepository.UpdateUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to update user")
		if err == errors.ErrEmailAlreadyExists {
			return err
		}
		return errors.ErrFailedToUpdateUser
	}
	return nil
}

func (u *adminUseCase) ChangeUserRole(ctx context.Context, userID string, request dto.ChangeUserRoleRequest) error {
	err := validation.AdminValidateChangeUserRoleRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate change user role request")
		return err
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return errors.ErrInvalidUserID
	}

	var user entity.User
	user.ID = userUUID
	user.Role = entity.ParseRoleType(request.Role)

	err = u.adminRepository.UpdateUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to change user role")
		return errors.ErrFailedToChangeUserRole
	}
	return nil
}

func (u *adminUseCase) ChangeUserStatus(ctx context.Context, userID string, request dto.ChangeUserStatusRequest) error {
	err := validation.AdminValidateChangeUserStatusRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate change user status request")
		return err
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return errors.ErrInvalidUserID
	}

	var user entity.User
	user.ID = userUUID
	user.Status = entity.ParseStatusType(request.Status)

	err = u.adminRepository.UpdateUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to change user status")
		return errors.ErrFailedToChangeUserStatus
	}
	return nil
}

func (u *adminUseCase) ChangeUserPassword(ctx context.Context, userID string, request dto.ChangeUserPasswordRequest) error {
	err := validation.AdminValidateChangeUserPasswordRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate change user password request")
		return err
	}

	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return errors.ErrInvalidUserID
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err, "Failed to hash password")
		return errors.ErrFailedToHashPassword
	}

	var user entity.User
	user.ID = userUUID
	user.Password = string(hashedPassword)

	err = u.adminRepository.UpdateUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to change user password")
		return errors.ErrFailedToChangeUserPassword
	}
	return nil
}

func (u *adminUseCase) DeleteUser(ctx context.Context, userID string) error {
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user ID")
		return errors.ErrInvalidUserID
	}

	var user entity.User
	user.ID = userUUID
	err = u.adminRepository.DeleteUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to delete user")
		return errors.ErrFailedToDeleteUser
	}
	return nil
}
