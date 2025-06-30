package validation

import (
	"errors"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/go-playground/validator/v10"
)

func ValidateServiceVisitCreateRequest(request dto.CreateServiceVisitRequest) error {
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
				case "ServiceMileage":
					if fieldError.Tag() == "required" {
						return errors.New("service mileage is required")
					}
					if fieldError.Tag() == "min" {
						return errors.New("service mileage must be greater than 0")
					}
				case "ServiceDate":
					if fieldError.Tag() == "required" {
						return errors.New("service date is required")
					}
					if fieldError.Tag() == "date" {
						return errors.New("invalid service date format")
					}
				default:
					return errors.New("validation failed for field: " + fieldError.Field())
				}
			}
		}
		return errors.New("validation failed")
	}

	// Validate nested oil change if provided
	if request.OilChange != nil {
		if err := validateServiceVisitOilChange(*request.OilChange); err != nil {
			return err
		}
	}

	// Validate nested oil filter if provided
	if request.OilFilter != nil {
		if err := validateServiceVisitOilFilter(*request.OilFilter); err != nil {
			return err
		}
	}

	return nil
}

func ValidateServiceVisitUpdateRequest(request dto.UpdateServiceVisitRequest) error {
	// Check if at least one field has a value
	if request.ServiceMileage == nil && request.ServiceDate == nil && request.ServiceCenter == nil &&
		request.Notes == nil && request.OilChange == nil && request.OilFilter == nil {
		return errors.New("no fields to update")
	}

	validate := validator.New()
	validate.RegisterValidation("date", validateDate)

	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "ServiceMileage":
					if fieldError.Tag() == "min" {
						return errors.New("service mileage must be greater than 0")
					}
				case "ServiceDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid service date format")
					}
				default:
					return errors.New("validation failed for field: " + fieldError.Field())
				}
			}
		}
		return errors.New("validation failed")
	}

	// Validate nested oil change if provided
	if request.OilChange != nil {
		if err := validateUpdateServiceVisitOilChange(*request.OilChange); err != nil {
			return err
		}
	}

	// Validate nested oil filter if provided
	if request.OilFilter != nil {
		if err := validateUpdateServiceVisitOilFilter(*request.OilFilter); err != nil {
			return err
		}
	}

	return nil
}

func validateServiceVisitOilChange(oilChange dto.ServiceVisitOilChange) error {
	validate := validator.New()
	validate.RegisterValidation("date", validateDate)

	err := validate.Struct(oilChange)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "OilName":
					if fieldError.Tag() == "required" {
						return errors.New("oil name is required")
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
					return errors.New("validation failed for oil change field: " + fieldError.Field())
				}
			}
		}
		return errors.New("oil change validation failed")
	}
	return nil
}

func validateServiceVisitOilFilter(oilFilter dto.ServiceVisitOilFilter) error {
	validate := validator.New()
	validate.RegisterValidation("date", validateDate)

	err := validate.Struct(oilFilter)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "FilterName":
					if fieldError.Tag() == "required" {
						return errors.New("filter name is required")
					}
				case "NextChangeMileage":
					if fieldError.Tag() == "min" {
						return errors.New("next change mileage must be greater than 0")
					}
				case "NextChangeDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid next change date format")
					}
				default:
					return errors.New("validation failed for oil filter field: " + fieldError.Field())
				}
			}
		}
		return errors.New("oil filter validation failed")
	}
	return nil
}

func validateUpdateServiceVisitOilChange(oilChange dto.UpdateServiceVisitOilChange) error {
	validate := validator.New()
	validate.RegisterValidation("date", validateDate)

	err := validate.Struct(oilChange)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "NextChangeDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid next change date format")
					}
				case "NextChangeMileage":
					if fieldError.Tag() == "min" {
						return errors.New("next change mileage must be greater than 0")
					}
				default:
					return errors.New("validation failed for oil change field: " + fieldError.Field())
				}
			}
		}
		return errors.New("oil change validation failed")
	}
	return nil
}

func validateUpdateServiceVisitOilFilter(oilFilter dto.UpdateServiceVisitOilFilter) error {
	validate := validator.New()
	validate.RegisterValidation("date", validateDate)

	err := validate.Struct(oilFilter)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "NextChangeMileage":
					if fieldError.Tag() == "min" {
						return errors.New("next change mileage must be greater than 0")
					}
				case "NextChangeDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid next change date format")
					}
				default:
					return errors.New("validation failed for oil filter field: " + fieldError.Field())
				}
			}
		}
		return errors.New("oil filter validation failed")
	}
	return nil
}
