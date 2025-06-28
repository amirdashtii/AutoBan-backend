package validation

import (
	"errors"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateOilFilterCreateRequest(request dto.CreateOilFilterRequest) error {
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
				case "FilterName":
					if fieldError.Tag() == "required" {
						return errors.New("filter name is required")
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
				default:
					return errors.New("validation failed for field: " + fieldError.Field())
				}
			}
		}
		return errors.New("validation failed")
	}
	return nil
}

func ValidateOilFilterUpdateRequest(request dto.UpdateOilFilterRequest) error {
	// Check if at least one field has a value
	if request.FilterName == nil && request.FilterBrand == nil && request.FilterType == nil &&
		request.FilterPartNumber == nil && request.ChangeMileage == nil && request.ChangeDate == nil &&
		request.NextChangeMileage == nil && request.NextChangeDate == nil && request.ServiceCenter == nil &&
		request.Notes == nil {
		return errors.New("no fields to update")
	}

	validate := validator.New()
	validate.RegisterValidation("date", validateDate)

	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "ChangeMileage":
					if fieldError.Tag() == "min" {
						return errors.New("change mileage must be greater than 0")
					}
				case "ChangeDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid change date format")
					}
				case "NextChangeDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid next change date format")
					}
				default:
					return errors.New("validation failed for field: " + fieldError.Field())
				}
			}
		}
		return errors.New("validation failed")
	}

	return nil
}
