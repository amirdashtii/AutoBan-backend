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
	"github.com/google/uuid"
)

type ServiceVisitUseCase interface {
	CreateServiceVisit(ctx context.Context, userID, vehicleID string, request dto.CreateServiceVisitRequest) (*dto.ServiceVisitResponse, error)
	GetServiceVisit(ctx context.Context, userID, vehicleID, visitID string) (*dto.ServiceVisitResponse, error)
	ListServiceVisits(ctx context.Context, userID, vehicleID string) (*dto.ListServiceVisitsResponse, error)
	UpdateServiceVisit(ctx context.Context, userID, vehicleID, visitID string, request dto.UpdateServiceVisitRequest) (*dto.ServiceVisitResponse, error)
	DeleteServiceVisit(ctx context.Context, userID, vehicleID, visitID string) error
	GetLastServiceVisit(ctx context.Context, userID, vehicleID string) (*dto.ServiceVisitResponse, error)
}

type serviceVisitUseCase struct {
	serviceVisitRepository repository.ServiceVisitRepository
	oilChangeRepository    repository.OilChangeRepository
	oilFilterRepository    repository.OilFilterRepository
	vehicleRepository      repository.VehicleRepository
}

func NewServiceVisitUseCase() ServiceVisitUseCase {
	serviceVisitRepository := repository.NewServiceVisitRepository()
	oilChangeRepository := repository.NewOilChangeRepository()
	oilFilterRepository := repository.NewOilFilterRepository()
	vehicleRepository := repository.NewVehicleRepository()
	return &serviceVisitUseCase{
		serviceVisitRepository: serviceVisitRepository,
		oilChangeRepository:    oilChangeRepository,
		oilFilterRepository:    oilFilterRepository,
		vehicleRepository:      vehicleRepository,
	}
}

func (uc *serviceVisitUseCase) CreateServiceVisit(ctx context.Context, userID, vehicleID string, request dto.CreateServiceVisitRequest) (*dto.ServiceVisitResponse, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}
	uintVehicleID, err := strconv.ParseUint(vehicleID, 10, 64)
	if err != nil {
		logger.Error(err, "Failed to parse vehicle id")
		return nil, errors.ErrInvalidUserVehicleID
	}
	userVehicle := entity.UserVehicle{}
	err = uc.vehicleRepository.GetUserVehicle(ctx, uuidUserID, uintVehicleID, &userVehicle)
	if err != nil {
		logger.Error(err, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	err = validation.ValidateServiceVisitCreateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate service visit create request")
		return nil, errors.ErrInvalidServiceVisitCreateRequest
	}

	serviceDate, err := time.Parse("2006-01-02", request.ServiceDate)
	if err != nil {
		logger.Error(err, "Failed to parse service date")
		return nil, errors.ErrInvalidDate
	}

	serviceVisit := entity.ServiceVisit{
		UserID:         uuidUserID,
		UserVehicleID:  uintVehicleID,
		ServiceMileage: request.ServiceMileage,
		ServiceDate:    serviceDate,
		ServiceCenter:  request.ServiceCenter,
		Notes:          request.Notes,
	}
	serviceVisit.ID = uuid.New()

	if request.OilChange != nil {
		serviceVisit.OilChange = entity.OilChange{
			UserID:            uuidUserID,
			UserVehicleID:     uintVehicleID,
			ServiceVisitID:    serviceVisit.ID,
			OilName:           request.OilChange.OilName,
			OilBrand:          request.OilChange.OilBrand,
			OilType:           request.OilChange.OilType,
			OilViscosity:      request.OilChange.OilViscosity,
			ChangeMileage:     request.ServiceMileage,
			ChangeDate:        serviceDate,
			OilCapacity:       request.OilChange.OilCapacity,
			NextChangeMileage: request.OilChange.NextChangeMileage,
			ServiceCenter:     request.ServiceCenter,
			Notes:             request.OilChange.Notes,
		}

		if request.OilChange.NextChangeDate != "" {
			nextChangeDate, err := time.Parse("2006-01-02", request.OilChange.NextChangeDate)
			if err != nil {
				logger.Error(err, "Failed to parse oil change next change date")
				return nil, errors.ErrInvalidDate
			}
			serviceVisit.OilChange.NextChangeDate = nextChangeDate
		}
	}

	if request.OilFilter != nil {
		serviceVisit.OilFilter = entity.OilFilter{
			UserID:            uuidUserID,
			UserVehicleID:     uintVehicleID,
			ServiceVisitID:    serviceVisit.ID,
			FilterName:        request.OilFilter.FilterName,
			FilterBrand:       request.OilFilter.FilterBrand,
			FilterType:        request.OilFilter.FilterType,
			FilterPartNumber:  request.OilFilter.FilterPartNumber,
			ChangeMileage:     request.ServiceMileage,
			ChangeDate:        serviceDate,
			NextChangeMileage: request.OilFilter.NextChangeMileage,
			ServiceCenter:     request.ServiceCenter,
			Notes:             request.OilFilter.Notes,
		}

		if request.OilFilter.NextChangeDate != "" {
			nextChangeDate, err := time.Parse("2006-01-02", request.OilFilter.NextChangeDate)
			if err != nil {
				logger.Error(err, "Failed to parse oil filter next change date")
				return nil, errors.ErrInvalidDate
			}
			serviceVisit.OilFilter.NextChangeDate = nextChangeDate
		}

		err = uc.serviceVisitRepository.CreateServiceVisit(ctx, &serviceVisit)
		if err != nil {
			logger.Error(err, "Failed to create service visit")
			return nil, errors.ErrFailedToCreateServiceVisit
		}
	}

	return uc.mapServiceVisitToResponse(&serviceVisit), nil
}

func (uc *serviceVisitUseCase) GetServiceVisit(ctx context.Context, userID, vehicleID, visitID string) (*dto.ServiceVisitResponse, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}
	uuidServiceVisitID, err := uuid.Parse(visitID)
	if err != nil {
		logger.Error(err, "Failed to parse service visit id")
		return nil, errors.ErrInvalidServiceVisitID
	}

	serviceVisit := entity.ServiceVisit{}
	serviceVisit.ID = uuidServiceVisitID

	err = uc.serviceVisitRepository.GetServiceVisit(ctx, &serviceVisit)
	if err != nil {
		logger.Error(err, "Failed to get service visit")
		return nil, errors.ErrFailedToGetServiceVisit
	}
	if serviceVisit.UserID != uuidUserID {
		logger.Error(errors.ErrUserVehicleNotOwned, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	return uc.mapServiceVisitToResponse(&serviceVisit), nil
}

func (uc *serviceVisitUseCase) ListServiceVisits(ctx context.Context, userID, vehicleID string) (*dto.ListServiceVisitsResponse, error) {
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

	serviceVisits := []entity.ServiceVisit{}
	err = uc.serviceVisitRepository.ListServiceVisits(ctx, vehicleID, &serviceVisits)
	if err != nil {
		logger.Error(err, "Failed to list service visits")
		return nil, errors.ErrFailedToListServiceVisits
	}

	serviceVisitsResponse := []dto.ServiceVisitResponse{}
	for _, serviceVisit := range serviceVisits {
		serviceVisitsResponse = append(serviceVisitsResponse, *uc.mapServiceVisitToResponse(&serviceVisit))
	}

	return &dto.ListServiceVisitsResponse{
		ServiceVisits: serviceVisitsResponse,
	}, nil
}

func (uc *serviceVisitUseCase) UpdateServiceVisit(ctx context.Context, userID, vehicleID, visitID string, request dto.UpdateServiceVisitRequest) (*dto.ServiceVisitResponse, error) {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return nil, errors.ErrInvalidUserID
	}

	uuidServiceVisitID, err := uuid.Parse(visitID)
	if err != nil {
		logger.Error(err, "Failed to parse service visit id")
		return nil, errors.ErrInvalidServiceVisitID
	}

	serviceVisit := entity.ServiceVisit{}
	serviceVisit.ID = uuidServiceVisitID

	err = uc.serviceVisitRepository.GetServiceVisit(ctx, &serviceVisit)
	if err != nil {
		logger.Error(err, "Failed to get service visit")
		return nil, errors.ErrFailedToGetServiceVisit
	}
	if serviceVisit.UserID != uuidUserID {
		logger.Error(errors.ErrUserVehicleNotOwned, "User vehicle not owned by user")
		return nil, errors.ErrUserVehicleNotOwned
	}

	err = validation.ValidateServiceVisitUpdateRequest(request)
	if err != nil {
		logger.Error(err, "Failed to validate service visit update request")
		return nil, errors.ErrInvalidServiceVisitUpdateRequest
	}

	// Update fields if provided
	if request.ServiceMileage != nil {
		serviceVisit.ServiceMileage = *request.ServiceMileage
	}
	if request.ServiceDate != nil {
		serviceDate, err := time.Parse("2006-01-02", *request.ServiceDate)
		if err != nil {
			logger.Error(err, "Failed to parse service date")
			return nil, errors.ErrInvalidDate
		}
		serviceVisit.ServiceDate = serviceDate
	}
	if request.ServiceCenter != nil {
		serviceVisit.ServiceCenter = *request.ServiceCenter
	}
	if request.Notes != nil {
		serviceVisit.Notes = *request.Notes
	}

	// Update oil change if provided
	if request.OilChange != nil {
		oilChange := entity.OilChange{}
		
		// Update oil change fields if provided
		if request.OilChange.OilName != nil {
			oilChange.OilName = *request.OilChange.OilName
		}
		if request.OilChange.OilBrand != nil {
			oilChange.OilBrand = *request.OilChange.OilBrand
		}
		if request.OilChange.OilType != nil {
			oilChange.OilType = *request.OilChange.OilType
		}
		if request.OilChange.OilViscosity != nil {
			oilChange.OilViscosity = *request.OilChange.OilViscosity
		}
		if request.OilChange.OilCapacity != nil {
			oilChange.OilCapacity = *request.OilChange.OilCapacity
		}
		if request.OilChange.NextChangeMileage != nil {
			oilChange.NextChangeMileage = *request.OilChange.NextChangeMileage
		}
		if request.OilChange.NextChangeDate != nil {
			nextChangeDate, err := time.Parse("2006-01-02", *request.OilChange.NextChangeDate)
			if err != nil {
				logger.Error(err, "Failed to parse oil change next change date")
				return nil, errors.ErrInvalidDate
			}
			oilChange.NextChangeDate = nextChangeDate
		}
		if request.OilChange.Notes != nil {
			oilChange.Notes = *request.OilChange.Notes
		}

		serviceVisit.OilChange = oilChange
	}

	// Update oil filter if provided
	if request.OilFilter != nil {
		oilFilter := entity.OilFilter{}

		// Update oil filter fields if provided
		if request.OilFilter.FilterName != nil {
			oilFilter.FilterName = *request.OilFilter.FilterName
		}
		if request.OilFilter.FilterBrand != nil {
			oilFilter.FilterBrand = *request.OilFilter.FilterBrand
		}
		if request.OilFilter.FilterType != nil {
			oilFilter.FilterType = *request.OilFilter.FilterType
		}
		if request.OilFilter.FilterPartNumber != nil {
			oilFilter.FilterPartNumber = *request.OilFilter.FilterPartNumber
		}
		if request.OilFilter.NextChangeMileage != nil {
			oilFilter.NextChangeMileage = *request.OilFilter.NextChangeMileage
		}
		if request.OilFilter.NextChangeDate != nil {
			nextChangeDate, err := time.Parse("2006-01-02", *request.OilFilter.NextChangeDate)
			if err != nil {
				logger.Error(err, "Failed to parse oil filter next change date")
				return nil, errors.ErrInvalidDate
			}
			oilFilter.NextChangeDate = nextChangeDate
		}
		if request.OilFilter.Notes != nil {
			oilFilter.Notes = *request.OilFilter.Notes
		}

		serviceVisit.OilFilter = oilFilter
	}

	err = uc.serviceVisitRepository.UpdateServiceVisit(ctx, &serviceVisit)
	if err != nil {
		logger.Error(err, "Failed to update service visit")
		return nil, errors.ErrFailedToUpdateServiceVisit
	}

	return uc.mapServiceVisitToResponse(&serviceVisit), nil
}

func (uc *serviceVisitUseCase) DeleteServiceVisit(ctx context.Context, userID, vehicleID, visitID string) error {
	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		logger.Error(err, "Failed to parse user id")
		return errors.ErrInvalidUserID
	}
	uuidServiceVisitID, err := uuid.Parse(visitID)
	if err != nil {
		logger.Error(err, "Failed to parse service visit id")
		return errors.ErrInvalidServiceVisitID
	}

	serviceVisit := entity.ServiceVisit{}
	serviceVisit.ID = uuidServiceVisitID

	err = uc.serviceVisitRepository.GetServiceVisit(ctx, &serviceVisit)
	if err != nil {
		logger.Error(err, "Failed to get service visit")
		return errors.ErrFailedToGetServiceVisit
	}
	if serviceVisit.UserID != uuidUserID {
		logger.Error(errors.ErrUserVehicleNotOwned, "User vehicle not owned by user")
		return errors.ErrUserVehicleNotOwned
	}

	err = uc.serviceVisitRepository.DeleteServiceVisit(ctx, &serviceVisit)
	if err != nil {
		logger.Error(err, "Failed to delete service visit")
		return errors.ErrFailedToDeleteServiceVisit
	}

	return nil
}

func (uc *serviceVisitUseCase) GetLastServiceVisit(ctx context.Context, userID, vehicleID string) (*dto.ServiceVisitResponse, error) {
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

	serviceVisit := entity.ServiceVisit{}

	serviceVisit.UserVehicleID = uintUserVehicleID

	err = uc.serviceVisitRepository.GetLastServiceVisit(ctx, &serviceVisit)
	if err != nil {
		logger.Error(err, "Failed to get last service visit")
		return nil, errors.ErrFailedToGetServiceVisit
	}

	return uc.mapServiceVisitToResponse(&serviceVisit), nil
}

func (uc *serviceVisitUseCase) mapServiceVisitToResponse(serviceVisit *entity.ServiceVisit) *dto.ServiceVisitResponse {
	response := &dto.ServiceVisitResponse{
		ID:             serviceVisit.ID,
		UserVehicleID:  serviceVisit.UserVehicleID,
		ServiceMileage: serviceVisit.ServiceMileage,
		ServiceDate:    serviceVisit.ServiceDate.Format("2006-01-02"),
		ServiceCenter:  serviceVisit.ServiceCenter,
		Notes:          serviceVisit.Notes,
	}

	// Map oil change if exists
	if serviceVisit.OilChange != (entity.OilChange{}) {
		response.OilChange = &dto.ServiceVisitOilChangeResponse{
			ID:                serviceVisit.OilChange.ID,
			OilName:           serviceVisit.OilChange.OilName,
			OilBrand:          serviceVisit.OilChange.OilBrand,
			OilType:           serviceVisit.OilChange.OilType,
			OilViscosity:      serviceVisit.OilChange.OilViscosity,
			OilCapacity:       serviceVisit.OilChange.OilCapacity,
			NextChangeMileage: serviceVisit.OilChange.NextChangeMileage,
			NextChangeDate:    serviceVisit.OilChange.NextChangeDate.Format("2006-01-02"),
			Notes:             serviceVisit.OilChange.Notes,
		}
	}

	// Map oil filter if exists
	if serviceVisit.OilFilter != (entity.OilFilter{}) {
		response.OilFilter = &dto.ServiceVisitOilFilterResponse{
			ID:                serviceVisit.OilFilter.ID,
			FilterName:        serviceVisit.OilFilter.FilterName,
			FilterBrand:       serviceVisit.OilFilter.FilterBrand,
			FilterType:        serviceVisit.OilFilter.FilterType,
			FilterPartNumber:  serviceVisit.OilFilter.FilterPartNumber,
			NextChangeMileage: serviceVisit.OilFilter.NextChangeMileage,
			NextChangeDate:    serviceVisit.OilFilter.NextChangeDate.Format("2006-01-02"),
			Notes:             serviceVisit.OilFilter.Notes,
		}
	}

	return response
}
