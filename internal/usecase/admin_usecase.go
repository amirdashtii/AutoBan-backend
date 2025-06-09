package usecase

import (
	"context"
	"time"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/repository"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AdminUseCase interface {
	ListUsers(ctx context.Context) ([]dto.User, error)
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

func (u *adminUseCase) ListUsers(ctx context.Context) ([]dto.User, error) {
	users, err := u.adminRepository.ListUsers(ctx)
	if err != nil {
		logger.Error(err, "Failed to list users")
		return nil, errors.ErrFailedToListUsers
	}

	userDtos := []dto.User{}
	for _, user := range users {
		firstName := ""
		if user.FirstName != nil {
			firstName = *user.FirstName
		}

		lastName := ""
		if user.LastName != nil {
			lastName = *user.LastName
		}

		email := ""
		if user.Email != nil {
			email = *user.Email
		}

		birthday := ""
		if user.Birthday != nil {
			birthday = user.Birthday.Format("2006-01-02")
		}

		userDtos = append(userDtos, dto.User{
			ID:        user.ID.String(),
			FirstName: firstName,
			LastName:  lastName,
			Email:     email,
			Phone:     user.PhoneNumber,
			Role:      user.Role.String(),
			Status:    user.Status.String(),
			Birthday:  birthday,
		})
	}

	return userDtos, nil
}

func (u *adminUseCase) GetUserById(ctx context.Context, userID string) (*dto.User, error) {
	var user entity.User
	user.ID = uuid.MustParse(userID)
	err := u.adminRepository.GetUserById(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to get user by id")
		return nil, errors.ErrFailedToGetUserById
	}

	firstName := ""
	if user.FirstName != nil {
		firstName = *user.FirstName
	}

	lastName := ""
	if user.LastName != nil {
		lastName = *user.LastName
	}

	email := ""
	if user.Email != nil {
		email = *user.Email
	}

	birthday := ""
	if user.Birthday != nil {
		birthday = user.Birthday.Format("2006-01-02")
	}

	return &dto.User{
		ID:        user.ID.String(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     user.PhoneNumber,
		Role:      user.Role.String(),
		Status:    user.Status.String(),
		Birthday:  birthday,
	}, nil
}

func (u *adminUseCase) UpdateUser(ctx context.Context, userID string, request dto.UpdateUserRequest) error {
	var birthday *time.Time
	if request.Birthday != "" {
		parsedTime, err := time.Parse("2006-01-02", request.Birthday)
		if err != nil {
			logger.Error(err, "Failed to parse birthday")
			return errors.ErrInvalidBirthday
		}
		birthday = &parsedTime
	}

	var user entity.User
	user.ID = uuid.MustParse(userID)
	user.FirstName = &request.FirstName
	user.LastName = &request.LastName
	user.Email = &request.Email
	user.PhoneNumber = request.Phone
	user.Birthday = birthday
	
	err := u.adminRepository.UpdateUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to update user")
		return errors.ErrFailedToUpdateUser
	}
	return nil
}

func (u *adminUseCase) ChangeUserRole(ctx context.Context, userID string, request dto.ChangeUserRoleRequest) error {
	var user entity.User
	user.ID = uuid.MustParse(userID)
	user.Role = entity.ParseRoleType(request.Role)
	err := u.adminRepository.UpdateUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to change user role")
		return errors.ErrFailedToChangeUserRole
	}
	return nil
}

func (u *adminUseCase) ChangeUserStatus(ctx context.Context, userID string, request dto.ChangeUserStatusRequest) error {
	var user entity.User
	user.ID = uuid.MustParse(userID)
	user.Status = entity.ParseStatusType(request.Status)
	err := u.adminRepository.UpdateUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to change user status")
		return errors.ErrFailedToChangeUserStatus
	}
	return nil
}

func (u *adminUseCase) ChangeUserPassword(ctx context.Context, userID string, request dto.ChangeUserPasswordRequest) error {
	var user entity.User
	user.ID = uuid.MustParse(userID)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		logger.Error(err, "Failed to hash password")
		return errors.ErrFailedToHashPassword
	}
	user.Password = string(hashedPassword)

	err = u.adminRepository.UpdateUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to change user password")
		return errors.ErrFailedToChangeUserPassword
	}
	return nil
}

func (u *adminUseCase) DeleteUser(ctx context.Context, userID string) error {
	var user entity.User
	user.ID = uuid.MustParse(userID)
	err := u.adminRepository.DeleteUser(ctx, &user)
	if err != nil {
		logger.Error(err, "Failed to delete user")
		return errors.ErrFailedToDeleteUser
	}
	return nil
}
