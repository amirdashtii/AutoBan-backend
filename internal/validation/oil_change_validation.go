package validation

import (
	"errors"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateOilChangeCreateRequest(request dto.CreateOilChangeRequest) error {
	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	validate.RegisterValidation("date", validateDate)
	err := validate.Struct(request)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "year":
				return errors.New("invalid year format")
			case "date":
				return errors.New("invalid date format")
			default:
				return errors.New("invalid oil change create request")
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
		request.NextChangeDate == nil && request.ServiceCenter == nil && request.Cost == nil && request.Notes == nil {
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
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "year":
				return errors.New("invalid year format")
			case "date":
				return errors.New("invalid date format")
			default:
				return errors.New("invalid oil change update request")
			}
		}
	}
	return nil
}
