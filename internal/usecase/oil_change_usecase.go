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

type OilChangeUseCase interface {
	CreateOilChange(ctx context.Context, request dto.CreateOilChangeRequest) (*dto.OilChangeResponse, error)
	GetOilChange(ctx context.Context, id string) (*dto.OilChangeResponse, error)
	ListOilChanges(ctx context.Context, userVehicleID string) (*dto.ListOilChangesResponse, error)
	UpdateOilChange(ctx context.Context, id string, request dto.UpdateOilChangeRequest) (*dto.OilChangeResponse, error)
	DeleteOilChange(ctx context.Context, id string) error
	GetLastOilChange(ctx context.Context, userVehicleID string) (*dto.OilChangeResponse, error)
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

func (uc *oilChangeUseCase) CreateOilChange(ctx context.Context, request dto.CreateOilChangeRequest) (*dto.OilChangeResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, strconv.Itoa(int(request.UserVehicleID)), &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	err = validation.ValidateOilChangeCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate oil change create request")
		return nil, errors.ErrInvalidOilChangeCreateRequest
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

	oilChange := entity.OilChange{
		UserVehicleID:     request.UserVehicleID,
		OilName:           request.OilName,
		OilBrand:          request.OilBrand,
		OilType:           request.OilType,
		OilViscosity:      request.OilViscosity,
		ChangeMileage:     request.ChangeMileage,
		ChangeDate:        changeDate,
		OilCapacity:       request.OilCapacity,
		NextChangeMileage: request.NextChangeMileage,
		NextChangeDate:    nextChangeDate,
		ServiceCenter:     request.ServiceCenter,
		Cost:              request.Cost,
		Notes:             request.Notes,
	}

	err = uc.oilChangeRepository.CreateOilChange(ctx, &oilChange)
	if err != nil {
		logger.Error(err, "Failed to create oil change")
		return nil, errors.ErrFailedToCreateOilChange
	}

	return &dto.OilChangeResponse{
		ID:                oilChange.ID,
		UserVehicleID:     oilChange.UserVehicleID,
		OilName:           oilChange.OilName,
		OilBrand:          oilChange.OilBrand,
		OilType:           oilChange.OilType,
		OilViscosity:      oilChange.OilViscosity,
		ChangeMileage:     oilChange.ChangeMileage,
		ChangeDate:        oilChange.ChangeDate,
		OilCapacity:       oilChange.OilCapacity,
		NextChangeMileage: oilChange.NextChangeMileage,
		NextChangeDate:    oilChange.NextChangeDate,
		ServiceCenter:     oilChange.ServiceCenter,
		Cost:              oilChange.Cost,
		Notes:             oilChange.Notes,
	}, nil
}

func (uc *oilChangeUseCase) GetOilChange(ctx context.Context, vehicleID string) (*dto.OilChangeResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, vehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	uintID, err := strconv.ParseUint(vehicleID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse oil change id")
		return nil, errors.ErrInvalidOilChangeID
	}

	oilChange := entity.OilChange{}
	oilChange.ID = uint(uintID)

	err = uc.oilChangeRepository.GetOilChange(ctx, oilChange.ID, &oilChange)
	if err != nil {
		logger.Error(err, "Failed to get oil change")
		return nil, errors.ErrFailedToGetOilChange
	}

	return &dto.OilChangeResponse{
		ID:                oilChange.ID,
		UserVehicleID:     oilChange.UserVehicleID,
		OilName:           oilChange.OilName,
		OilBrand:          oilChange.OilBrand,
		OilType:           oilChange.OilType,
		OilViscosity:      oilChange.OilViscosity,
		ChangeMileage:     oilChange.ChangeMileage,
		ChangeDate:        oilChange.ChangeDate,
		OilCapacity:       oilChange.OilCapacity,
		NextChangeMileage: oilChange.NextChangeMileage,
		NextChangeDate:    oilChange.NextChangeDate,
		ServiceCenter:     oilChange.ServiceCenter,
		Cost:              oilChange.Cost,
		Notes:             oilChange.Notes,
	}, nil
}

func (uc *oilChangeUseCase) ListOilChanges(ctx context.Context, userVehicleID string) (*dto.ListOilChangesResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, userVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	oilChanges := []entity.OilChange{}
	err = uc.oilChangeRepository.ListOilChanges(ctx, userVehicleID, &oilChanges)
	if err != nil {
		logger.Error(err, "Failed to list oil changes")
		return nil, errors.ErrFailedToListOilChanges
	}

	oilChangesResponse := []dto.OilChangeResponse{}
	for _, oilChange := range oilChanges {
		oilChangesResponse = append(oilChangesResponse, dto.OilChangeResponse{
			ID:                oilChange.ID,
			UserVehicleID:     oilChange.UserVehicleID,
			OilName:           oilChange.OilName,
			OilBrand:          oilChange.OilBrand,
			OilType:           oilChange.OilType,
			OilViscosity:      oilChange.OilViscosity,
			ChangeMileage:     oilChange.ChangeMileage,
			ChangeDate:        oilChange.ChangeDate,
			OilCapacity:       oilChange.OilCapacity,
			NextChangeMileage: oilChange.NextChangeMileage,
			NextChangeDate:    oilChange.NextChangeDate,
			ServiceCenter:     oilChange.ServiceCenter,
			Cost:              oilChange.Cost,
			Notes:             oilChange.Notes,
		})
	}

	return &dto.ListOilChangesResponse{
		OilChanges: oilChangesResponse,
	}, nil
}

func (uc *oilChangeUseCase) UpdateOilChange(ctx context.Context, id string, request dto.UpdateOilChangeRequest) (*dto.OilChangeResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, id, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}
	err = validation.ValidateOilChangeUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate oil change update request")
		return nil, errors.ErrInvalidOilChangeUpdateRequest
	}

	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse oil change id")
		return nil, errors.ErrInvalidOilChangeID
	}

	oilChange := entity.OilChange{}
	oilChange.ID = uint(uintID)

	// Get existing oil change
	err = uc.oilChangeRepository.GetOilChange(ctx, oilChange.ID, &oilChange)
	if err != nil {
		logger.Error(err, "Failed to get oil change for update")
		return nil, errors.ErrFailedToGetOilChange
	}

	// Update fields if provided
	if request.OilName != nil {
		oilChange.OilName = *request.OilName
	}
	if request.OilBrand != nil {
		oilChange.OilBrand = *request.OilBrand
	}
	if request.OilType != nil {
		oilChange.OilType = *request.OilType
	}
	if request.OilViscosity != nil {
		oilChange.OilViscosity = *request.OilViscosity
	}
	if request.ChangeMileage != nil {
		oilChange.ChangeMileage = *request.ChangeMileage
	}
	if request.ChangeDate != nil {
		changeDate, err := time.Parse("2006-01-02", *request.ChangeDate)
		if err != nil {
			logger.Error(err, "Failed to parse change date")
			return nil, errors.ErrInvalidDate
		}
		oilChange.ChangeDate = changeDate
	}
	if request.OilCapacity != nil {
		oilChange.OilCapacity = *request.OilCapacity
	}
	if request.NextChangeMileage != nil {
		oilChange.NextChangeMileage = *request.NextChangeMileage
	}
	if request.NextChangeDate != nil {
		nextChangeDate, err := time.Parse("2006-01-02", *request.NextChangeDate)
		if err != nil {
			logger.Error(err, "Failed to parse next change date")
			return nil, errors.ErrInvalidDate
		}
		oilChange.NextChangeDate = nextChangeDate
	}
	if request.ServiceCenter != nil {
		oilChange.ServiceCenter = *request.ServiceCenter
	}
	if request.Cost != nil {
		oilChange.Cost = *request.Cost
	}
	if request.Notes != nil {
		oilChange.Notes = *request.Notes
	}

	err = uc.oilChangeRepository.UpdateOilChange(ctx, &oilChange)
	if err != nil {
		logger.Error(err, "Failed to update oil change")
		return nil, errors.ErrFailedToUpdateOilChange
	}

	return &dto.OilChangeResponse{
		ID:                oilChange.ID,
		UserVehicleID:     oilChange.UserVehicleID,
		OilName:           oilChange.OilName,
		OilBrand:          oilChange.OilBrand,
		OilType:           oilChange.OilType,
		OilViscosity:      oilChange.OilViscosity,
		ChangeMileage:     oilChange.ChangeMileage,
		ChangeDate:        oilChange.ChangeDate,
		OilCapacity:       oilChange.OilCapacity,
		NextChangeMileage: oilChange.NextChangeMileage,
		NextChangeDate:    oilChange.NextChangeDate,
		ServiceCenter:     oilChange.ServiceCenter,
		Cost:              oilChange.Cost,
		Notes:             oilChange.Notes,
	}, nil
}

func (uc *oilChangeUseCase) DeleteOilChange(ctx context.Context, id string) error {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, id, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return errors.ErrUserVehicleNotOwned
	}
	uintID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse oil change id")
		return errors.ErrInvalidOilChangeID
	}

	oilChange := entity.OilChange{}
	oilChange.ID = uint(uintID)

	err = uc.oilChangeRepository.DeleteOilChange(ctx, &oilChange)
	if err != nil {
		logger.Error(err, "Failed to delete oil change")
		return errors.ErrFailedToDeleteOilChange
	}

	return nil
}

func (uc *oilChangeUseCase) GetLastOilChange(ctx context.Context, userVehicleID string) (*dto.OilChangeResponse, error) {
	userID := ctx.Value("user_id").(string)
	userVehicle := entity.UserVehicle{}
	err := uc.vehicleRepository.GetUserVehicle(ctx, userID, userVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}
	oilChange := entity.OilChange{}
	err = uc.oilChangeRepository.GetLastOilChange(ctx, userVehicleID, &oilChange)
	if err != nil {
		logger.Error(err, "Failed to get last oil change")
		return nil, errors.ErrFailedToGetOilChange
	}

	return &dto.OilChangeResponse{
		ID:                oilChange.ID,
		UserVehicleID:     oilChange.UserVehicleID,
		OilName:           oilChange.OilName,
		OilBrand:          oilChange.OilBrand,
		OilType:           oilChange.OilType,
		OilViscosity:      oilChange.OilViscosity,
		ChangeMileage:     oilChange.ChangeMileage,
		ChangeDate:        oilChange.ChangeDate,
		OilCapacity:       oilChange.OilCapacity,
		NextChangeMileage: oilChange.NextChangeMileage,
		NextChangeDate:    oilChange.NextChangeDate,
		ServiceCenter:     oilChange.ServiceCenter,
		Cost:              oilChange.Cost,
		Notes:             oilChange.Notes,
	}, nil
}
