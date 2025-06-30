package usecase

import (
	"context"
	"strconv"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/repository"
	"github.com/amirdashtii/AutoBan/pkg/logger"
	"github.com/google/uuid"
)

type OilChangeUseCase interface {
	GetOilChange(ctx context.Context, userID string, vehicleID string, oilChangeID string) (*dto.OilChangeResponse, error)
	ListOilChanges(ctx context.Context, userID string, vehicleID string) (*dto.ListOilChangesResponse, error)
	GetLastOilChange(ctx context.Context, userID string, vehicleID string) (*dto.OilChangeResponse, error)
}

type oilChangeUseCase struct {
	oilChangeRepository repository.OilChangeRepository
	vehicleRepository   repository.VehicleRepository
}

func NewOilChangeUseCase() OilChangeUseCase {
	oilChangeRepository := repository.NewOilChangeRepository()
	vehicleRepository := repository.NewVehicleRepository()
	return &oilChangeUseCase{oilChangeRepository: oilChangeRepository, vehicleRepository: vehicleRepository}
}

func (uc *oilChangeUseCase) GetOilChange(ctx context.Context, userID string, vehicleID string, oilChangeID string) (*dto.OilChangeResponse, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}

	uintVehicleID, err := strconv.ParseUint(vehicleID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle id")
		return nil, errors.ErrInvalidVehicleID
	}

	uintOilChangeID, err := strconv.ParseUint(oilChangeID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse oil change id")
		return nil, errors.ErrInvalidOilChangeID
	}

	userVehicle := entity.UserVehicle{}
	err = uc.vehicleRepository.GetUserVehicle(ctx, uuidUserID, uintVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	oilChange := entity.OilChange{}
	oilChange.ID = uintOilChangeID

	err = uc.oilChangeRepository.GetOilChange(ctx, oilChange.ID, &oilChange)
	if err != nil {
		logger.Error(err, "Failed to get oil change")
		return nil, errors.ErrFailedToGetOilChange
	}

	return uc.mapOilChangeToResponse(&oilChange), nil
}

func (uc *oilChangeUseCase) ListOilChanges(ctx context.Context, userID string, vehicleID string) (*dto.ListOilChangesResponse, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}
	uintUserVehicleID, err := strconv.ParseUint(vehicleID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse user vehicle id")
		return nil, errors.ErrInvalidUserVehicleID
	}

	userVehicle := entity.UserVehicle{}
	err = uc.vehicleRepository.GetUserVehicle(ctx, uuidUserID, uintUserVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	oilChanges := []entity.OilChange{}
	err = uc.oilChangeRepository.ListOilChanges(ctx, uintUserVehicleID, &oilChanges)
	if err != nil {
		logger.Error(err, "Failed to list oil changes")
		return nil, errors.ErrFailedToListOilChanges
	}

	oilChangesResponse := []dto.OilChangeResponse{}
	for _, oilChange := range oilChanges {
		oilChangesResponse = append(oilChangesResponse, *uc.mapOilChangeToResponse(&oilChange))
	}

	return &dto.ListOilChangesResponse{
		OilChanges: oilChangesResponse,
	}, nil
}

func (uc *oilChangeUseCase) GetLastOilChange(ctx context.Context, userID string, vehicleID string) (*dto.OilChangeResponse, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}
	uintUserVehicleID, err := strconv.ParseUint(vehicleID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse user vehicle id")
		return nil, errors.ErrInvalidUserVehicleID
	}
	userVehicle := entity.UserVehicle{}
	err = uc.vehicleRepository.GetUserVehicle(ctx, uuidUserID, uintUserVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}
	oilChange := entity.OilChange{}
	err = uc.oilChangeRepository.GetLastOilChange(ctx, uintUserVehicleID, &oilChange)
	if err != nil {
		logger.Error(err, "Failed to get last oil change")
		return nil, errors.ErrFailedToGetOilChange
	}

	return uc.mapOilChangeToResponse(&oilChange), nil
}

func (uc *oilChangeUseCase) mapOilChangeToResponse(oilChange *entity.OilChange) *dto.OilChangeResponse {
	return &dto.OilChangeResponse{
		ID:                oilChange.ID,
		UserVehicleID:     oilChange.UserVehicleID,
		OilName:           oilChange.OilName,
		OilBrand:          oilChange.OilBrand,
		OilType:           oilChange.OilType,
		OilViscosity:      oilChange.OilViscosity,
		ChangeMileage:     oilChange.ChangeMileage,
		ChangeDate:        oilChange.ChangeDate.Format("2006-01-02"),
		OilCapacity:       oilChange.OilCapacity,
		NextChangeMileage: oilChange.NextChangeMileage,
		NextChangeDate:    oilChange.NextChangeDate.Format("2006-01-02"),
		ServiceCenter:     oilChange.ServiceCenter,
		Notes:             oilChange.Notes,
	}
}
