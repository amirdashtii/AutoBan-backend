package usecase

import (
	"context"
	"strconv"
	"time"

	"github.com/amirdashtii/AutoBan/internal/domain/entity"
	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/amirdashtii/AutoBan/internal/repository"
	"github.com/amirdashtii/AutoBan/internal/validation"
	"github.com/amirdashtii/AutoBan/pkg/logger"
)

type OilFilterUseCase interface {
	CreateOilFilter(ctx context.Context, request dto.CreateOilFilterRequest) (*dto.OilFilterResponse, error)
	GetOilFilter(ctx context.Context, id string) (*dto.OilFilterResponse, error)
	ListOilFilters(ctx context.Context, userVehicleID string) (*dto.ListOilFiltersResponse, error)
	UpdateOilFilter(ctx context.Context, id string, request dto.UpdateOilFilterRequest) (*dto.OilFilterResponse, error)
	DeleteOilFilter(ctx context.Context, id string) error
	GetLastOilFilter(ctx context.Context, userVehicleID string) (*dto.OilFilterResponse, error)
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

func (uc *oilFilterUseCase) CreateOilFilter(ctx context.Context, request dto.CreateOilFilterRequest) (*dto.OilFilterResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, strconv.Itoa(int(request.UserVehicleID)), &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	err = validation.ValidateOilFilterCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate oil filter create request")
		return nil, errors.ErrInvalidOilFilterCreateRequest
	}

	changeDate, err := time.Parse("2006-01-02", request.ChangeDate)
	if err != nil {
		logger.Error(err, "Failed to parse change date")
		return nil, errors.ErrInvalidDate
	}

	var nextChangeDate time.Time
	if request.NextChangeDate != "" {
		nextChangeDate, err = time.Parse("2006-01-02", request.NextChangeDate)
		if err != nil {
			logger.Error(err, "Failed to parse next change date")
			return nil, errors.ErrInvalidDate
		}
	}

	oilFilter := entity.OilFilter{
		UserVehicleID:     request.UserVehicleID,
		FilterName:        request.FilterName,
		FilterBrand:       request.FilterBrand,
		FilterType:        request.FilterType,
		FilterPartNumber:  request.FilterPartNumber,
		ChangeMileage:     request.ChangeMileage,
		ChangeDate:        changeDate,
		NextChangeMileage: request.NextChangeMileage,
		NextChangeDate:    nextChangeDate,
		ServiceCenter:     request.ServiceCenter,
		Notes:             request.Notes,
	}

	err = uc.oilFilterRepository.CreateOilFilter(ctx, &oilFilter)
	if err != nil {
		logger.Error(err, "Failed to create oil filter")
		return nil, errors.ErrFailedToCreateOilFilter
	}

	return &dto.OilFilterResponse{
		ID:                oilFilter.ID,
		UserVehicleID:     oilFilter.UserVehicleID,
		FilterName:        oilFilter.FilterName,
		FilterBrand:       oilFilter.FilterBrand,
		FilterType:        oilFilter.FilterType,
		FilterPartNumber:  oilFilter.FilterPartNumber,
		ChangeMileage:     oilFilter.ChangeMileage,
		ChangeDate:        oilFilter.ChangeDate,
		NextChangeMileage: oilFilter.NextChangeMileage,
		NextChangeDate:    oilFilter.NextChangeDate,
		ServiceCenter:     oilFilter.ServiceCenter,
		Notes:             oilFilter.Notes,
	}, nil
}

func (uc *oilFilterUseCase) GetOilFilter(ctx context.Context, id string) (*dto.OilFilterResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, id, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}	

	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse oil filter id")
		return nil, errors.ErrInvalidOilFilterID
	}

	oilFilter := entity.OilFilter{}
	oilFilter.ID = uint(uintID)

	err = uc.oilFilterRepository.GetOilFilter(ctx, oilFilter.ID, &oilFilter)
	if err != nil {
		logger.Error(err, "Failed to get oil filter")
		return nil, errors.ErrFailedToGetOilFilter
	}

	return &dto.OilFilterResponse{
		ID:                oilFilter.ID,
		UserVehicleID:     oilFilter.UserVehicleID,
		FilterName:        oilFilter.FilterName,
		FilterBrand:       oilFilter.FilterBrand,
		FilterType:        oilFilter.FilterType,
		FilterPartNumber:  oilFilter.FilterPartNumber,
		ChangeMileage:     oilFilter.ChangeMileage,
		ChangeDate:        oilFilter.ChangeDate,
		NextChangeMileage: oilFilter.NextChangeMileage,
		NextChangeDate:    oilFilter.NextChangeDate,
		ServiceCenter:     oilFilter.ServiceCenter,
		Notes:             oilFilter.Notes,
	}, nil
}

func (uc *oilFilterUseCase) ListOilFilters(ctx context.Context, userVehicleID string) (*dto.ListOilFiltersResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, userVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	oilFilters := []entity.OilFilter{}
	err = uc.oilFilterRepository.ListOilFilters(ctx, userVehicleID, &oilFilters)
	if err != nil {
		logger.Error(err, "Failed to list oil filters")
		return nil, errors.ErrFailedToListOilFilters
	}

	oilFiltersResponse := []dto.OilFilterResponse{}
	for _, oilFilter := range oilFilters {
		oilFiltersResponse = append(oilFiltersResponse, dto.OilFilterResponse{
			ID:                oilFilter.ID,
			UserVehicleID:     oilFilter.UserVehicleID,
			FilterName:        oilFilter.FilterName,
			FilterBrand:       oilFilter.FilterBrand,
			FilterType:        oilFilter.FilterType,
			FilterPartNumber:  oilFilter.FilterPartNumber,
			ChangeMileage:     oilFilter.ChangeMileage,
			ChangeDate:        oilFilter.ChangeDate,
			NextChangeMileage: oilFilter.NextChangeMileage,
			NextChangeDate:    oilFilter.NextChangeDate,
			ServiceCenter:     oilFilter.ServiceCenter,
			Notes:             oilFilter.Notes,
		})
	}

	return &dto.ListOilFiltersResponse{
		OilFilters: oilFiltersResponse,
	}, nil
}

func (uc *oilFilterUseCase) UpdateOilFilter(ctx context.Context, id string, request dto.UpdateOilFilterRequest) (*dto.OilFilterResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, id, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	err = validation.ValidateOilFilterUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate oil filter update request")
		return nil, errors.ErrInvalidOilFilterUpdateRequest
	}

	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse oil filter id")
		return nil, errors.ErrInvalidOilFilterID
	}

	oilFilter := entity.OilFilter{}
	oilFilter.ID = uint(uintID)

	// Get existing oil filter
	err = uc.oilFilterRepository.GetOilFilter(ctx, oilFilter.ID, &oilFilter)
	if err != nil {
		logger.Error(err, "Failed to get oil filter for update")
		return nil, errors.ErrFailedToGetOilFilter
	}

	// Update fields if provided
	if request.FilterName != nil {
		oilFilter.FilterName = *request.FilterName
	}
	if request.FilterBrand != nil {
		oilFilter.FilterBrand = *request.FilterBrand
	}
	if request.FilterType != nil {
		oilFilter.FilterType = *request.FilterType
	}
	if request.FilterPartNumber != nil {
		oilFilter.FilterPartNumber = *request.FilterPartNumber
	}
	if request.ChangeMileage != nil {
		oilFilter.ChangeMileage = *request.ChangeMileage
	}
	if request.ChangeDate != nil {
		changeDate, err := time.Parse("2006-01-02", *request.ChangeDate)
		if err != nil {
			logger.Error(err, "Failed to parse change date")
			return nil, errors.ErrInvalidDate
		}
		oilFilter.ChangeDate = changeDate
	}
	if request.NextChangeMileage != nil {
		oilFilter.NextChangeMileage = *request.NextChangeMileage
	}
	if request.NextChangeDate != nil {
		nextChangeDate, err := time.Parse("2006-01-02", *request.NextChangeDate)
		if err != nil {
			logger.Error(err, "Failed to parse next change date")
			return nil, errors.ErrInvalidDate
		}
		oilFilter.NextChangeDate = nextChangeDate
	}
	if request.ServiceCenter != nil {
		oilFilter.ServiceCenter = *request.ServiceCenter
	}
	if request.Notes != nil {
		oilFilter.Notes = *request.Notes
	}

	err = uc.oilFilterRepository.UpdateOilFilter(ctx, &oilFilter)
	if err != nil {
		logger.Error(err, "Failed to update oil filter")
		return nil, errors.ErrFailedToUpdateOilFilter
	}

	return &dto.OilFilterResponse{
		ID:                oilFilter.ID,
		UserVehicleID:     oilFilter.UserVehicleID,
		FilterName:        oilFilter.FilterName,
		FilterBrand:       oilFilter.FilterBrand,
		FilterType:        oilFilter.FilterType,
		FilterPartNumber:  oilFilter.FilterPartNumber,
		ChangeMileage:     oilFilter.ChangeMileage,
		ChangeDate:        oilFilter.ChangeDate,
		NextChangeMileage: oilFilter.NextChangeMileage,
		NextChangeDate:    oilFilter.NextChangeDate,
		ServiceCenter:     oilFilter.ServiceCenter,
		Notes:             oilFilter.Notes,
	}, nil
}

func (uc *oilFilterUseCase) DeleteOilFilter(ctx context.Context, id string) error {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, id, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return errors.ErrUserVehicleNotOwned
	}

	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse oil filter id")
		return errors.ErrInvalidOilFilterID
	}

	oilFilter := entity.OilFilter{}
	oilFilter.ID = uint(uintID)

	err = uc.oilFilterRepository.DeleteOilFilter(ctx, &oilFilter)
	if err != nil {
		logger.Error(err, "Failed to delete oil filter")
		return errors.ErrFailedToDeleteOilFilter
	}

	return nil
}

func (uc *oilFilterUseCase) GetLastOilFilter(ctx context.Context, userVehicleID string) (*dto.OilFilterResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, userVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	oilFilter := entity.OilFilter{}
	err = uc.oilFilterRepository.GetLastOilFilter(ctx, userVehicleID, &oilFilter)
	if err != nil {
		logger.Error(err, "Failed to get last oil filter")
		return nil, errors.ErrFailedToGetOilFilter
	}

	return &dto.OilFilterResponse{
		ID:                oilFilter.ID,
		UserVehicleID:     oilFilter.UserVehicleID,
		FilterName:        oilFilter.FilterName,
		FilterBrand:       oilFilter.FilterBrand,
		FilterType:        oilFilter.FilterType,
		FilterPartNumber:  oilFilter.FilterPartNumber,
		ChangeMileage:     oilFilter.ChangeMileage,
		ChangeDate:        oilFilter.ChangeDate,
		NextChangeMileage: oilFilter.NextChangeMileage,
		NextChangeDate:    oilFilter.NextChangeDate,
		ServiceCenter:     oilFilter.ServiceCenter,
		Notes:             oilFilter.Notes,
	}, nil
}
