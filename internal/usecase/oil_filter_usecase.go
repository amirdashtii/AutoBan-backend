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

type OilFilterUseCase interface {
	ListOilFilters(ctx context.Context, userID, userVehicleID string) (*dto.ListOilFiltersResponse, error)
	GetLastOilFilter(ctx context.Context, userID, userVehicleID string) (*dto.OilFilterResponse, error)
	GetOilFilter(ctx context.Context, userID, userVehicleID, oilFilterID string) (*dto.OilFilterResponse, error)
}

type oilFilterUseCase struct {
	oilFilterRepository repository.OilFilterRepository
	vehicleRepository   repository.VehicleRepository
}

func NewOilFilterUseCase() OilFilterUseCase {
	oilFilterRepository := repository.NewOilFilterRepository()
	vehicleRepository := repository.NewVehicleRepository()
	return &oilFilterUseCase{oilFilterRepository: oilFilterRepository, vehicleRepository: vehicleRepository}
}

func (uc *oilFilterUseCase) ListOilFilters(ctx context.Context, userID, userVehicleID string) (*dto.ListOilFiltersResponse, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}
	uintUserVehicleID, err := strconv.ParseUint(userVehicleID, 10, 64)
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

	oilFilters := []entity.OilFilter{}
	err = uc.oilFilterRepository.ListOilFilters(ctx, uintUserVehicleID, &oilFilters)
	if err != nil {
		logger.Error(err, "Failed to list oil filters")
		return nil, errors.ErrFailedToListOilFilters
	}

	oilFiltersResponse := []dto.OilFilterResponse{}
	for _, oilFilter := range oilFilters {
		oilFiltersResponse = append(oilFiltersResponse, *uc.mapOilFilterToResponse(&oilFilter))
	}

	return &dto.ListOilFiltersResponse{
		OilFilters: oilFiltersResponse,
	}, nil
}

func (uc *oilFilterUseCase) GetLastOilFilter(ctx context.Context, userID, userVehicleID string) (*dto.OilFilterResponse, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}
	uintUserVehicleID, err := strconv.ParseUint(userVehicleID, 10, 64)
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

	oilFilter := entity.OilFilter{}
	err = uc.oilFilterRepository.GetLastOilFilter(ctx, uintUserVehicleID, &oilFilter)
	if err != nil {
		logger.Error(err, "Failed to get last oil filter")
		return nil, errors.ErrFailedToGetOilFilter
	}

	return uc.mapOilFilterToResponse(&oilFilter), nil
}

func (uc *oilFilterUseCase) GetOilFilter(ctx context.Context, userID, userVehicleID, oilFilterID string) (*dto.OilFilterResponse, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}
	uintOilFilterID, err := strconv.ParseUint(oilFilterID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse oil filter id")
		return nil, errors.ErrInvalidOilFilterID
	}

	oilFilter := entity.OilFilter{}
	err = uc.oilFilterRepository.GetOilFilter(ctx, uintOilFilterID, &oilFilter)
	if err != nil {
		logger.Error(err, "Failed to get oil filter")
		return nil, errors.ErrFailedToGetOilFilter
	}

	if oilFilter.UserID != uuidUserID {
		logger.Error(err, "Oil filter not owned by user")
		return nil, errors.ErrOilFilterNotOwned
	}

	return uc.mapOilFilterToResponse(&oilFilter), nil
}

func (uc *oilFilterUseCase) mapOilFilterToResponse(oilFilter *entity.OilFilter) *dto.OilFilterResponse {
	return &dto.OilFilterResponse{
		ID:                oilFilter.ID,
		UserVehicleID:     oilFilter.UserVehicleID,
		FilterName:        oilFilter.FilterName,
		FilterBrand:       oilFilter.FilterBrand,
		FilterType:        oilFilter.FilterType,
		FilterPartNumber:  oilFilter.FilterPartNumber,
		ChangeMileage:     oilFilter.ChangeMileage,
		ChangeDate:        oilFilter.ChangeDate.Format("2006-01-02"),
		NextChangeMileage: oilFilter.NextChangeMileage,
		NextChangeDate:    oilFilter.NextChangeDate.Format("2006-01-02"),
		ServiceCenter:     oilFilter.ServiceCenter,
		Notes:             oilFilter.Notes,
	}
}
