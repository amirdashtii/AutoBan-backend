package validation

import (
	"errors"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateOilChangeCreateRequest(request dto.CreateOilChangeRequest) error {
	validate := validator.New()
	validate.RegisterValidation("date", validateDate)

	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "UserVehicleID":
					if fieldError.Tag() == "required" {
						return errors.New("user vehicle id is required")
					}
				case "OilName":
					if fieldError.Tag() == "required" {
						return errors.New("oil name is required")
					}
				case "ChangeMileage":
					if fieldError.Tag() == "required" {
						return errors.New("change mileage is required")
					}
					if fieldError.Tag() == "min" {
						return errors.New("change mileage must be greater than 0")
					}
				case "ChangeDate":
					if fieldError.Tag() == "required" {
						return errors.New("change date is required")
					}
					if fieldError.Tag() == "date" {
						return errors.New("invalid change date format")
					}
				case "NextChangeDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid next change date format")
					}
				case "NextChangeMileage":
					if fieldError.Tag() == "min" {
						return errors.New("next change mileage must be greater than 0")
					}
				default:
					return errors.New("validation failed for field: " + fieldError.Field())
				}
			}
		}
	}
	return nil
}

func ValidateOilChangeUpdateRequest(request dto.UpdateOilChangeRequest) error {
	// Check if at least one field has a value
	if request.OilName == nil && request.OilBrand == nil && request.OilType == nil &&
		request.OilViscosity == nil && request.ChangeMileage == nil && request.ChangeDate == nil &&
		request.OilCapacity == nil && request.NextChangeMileage == nil &&
		request.NextChangeDate == nil && request.ServiceCenter == nil && request.Notes == nil {
		return errors.New("no fields to update")
	}

	// If OilName is provided, validate it's not empty
	if request.OilName != nil && *request.OilName == "" {
		return errors.New("invalid oil name")
	}

	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	validate.RegisterValidation("date", validateDate)

	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "ChangeDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid change date format")
					}
				case "ChangeMileage":
					if fieldError.Tag() == "min" {
						return errors.New("change mileage must be greater than 0")
					}
				case "NextChangeDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid next change date format")
					}
				case "NextChangeMileage":
					if fieldError.Tag() == "min" {
						return errors.New("next change mileage must be greater than 0")
					}
				default:
					return errors.New("validation failed for field: " + fieldError.Field())
				}
			}
		}
	}
	return nil
}
