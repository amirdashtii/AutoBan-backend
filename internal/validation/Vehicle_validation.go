package validation

import (
	"errors"
	"reflect"
	"regexp"
	"strconv"
	"time"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/go-playground/validator/v10"
)

func validateYear(fl validator.FieldLevel) bool {
	var year int

	// Handle different field types
	switch fl.Field().Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		year = int(fl.Field().Int())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		year = int(fl.Field().Uint())
	case reflect.String:
		if yearStr, err := strconv.Atoi(fl.Field().String()); err == nil {
			year = yearStr
		} else {
			return false
		}
	default:
		return false
	}

	return year >= 1300 && year <= time.Now().Year()
}

func validateDate(fl validator.FieldLevel) bool {
	dateTime, err := time.Parse("2006-01-02", fl.Field().String())
	if err != nil {
		return false
	}
	return !dateTime.IsZero()
}

func validateIranianLicensePlate(fl validator.FieldLevel) bool {
	licensePlate := fl.Field().String()

	// پلاک موتورسیکلت
	motorcyclePlate := regexp.MustCompile(`^\d{3}-?\d{5}$`)
	// پلاک سفید عادی (فقط حروف مشخص شده)
	normalWhitePlate := regexp.MustCompile(`^\d{2}-[بجدسصطلقلمنوهی]-\d{3}-?\d{2}$`)
	// پلاک سفید با حرف گ (گذر موقت)
	temporaryPlate := regexp.MustCompile(`^\d{2}-گ-\d{3}-?\d{2}$`)
	// پلاک سفید با حرف ژ (معلولین و جانبازان)
	disabledPlate := regexp.MustCompile(`^\d{2}-ژ-\d{3}-?\d{2}$`)
	// پلاک قرمز تشریفات دولتی
	ceremonyPlate := regexp.MustCompile(`^\d{2}-تشریفات-\d{4}$`)
	// پلاک قرمز با حرف الف (دولتی)
	governmentPlate := regexp.MustCompile(`^\d{2}-الف-\d{3}-?\d{2}$`)
	// پلاک سبز با حرف پ (فراجا، راهور، پلیس مبارزه با مواد مخدر)
	policePlate := regexp.MustCompile(`^\d{2}-پ-\d{3}-?\d{2}$`)
	// پلاک سبز با حرف ث (سپاه پاسداران)
	irgcPlate := regexp.MustCompile(`^\d{2}-ث-\d{3}-?\d{2}$`)
	// پلاک زرد با حرف ک (وسایل کشاورزی)
	agriculturalPlate := regexp.MustCompile(`^\d{2}-ک-\d{3}-?\d{2}$`)
	// پلاک زرد با حرف ع (حمل‌ونقل عمومی)
	publicTransportPlate := regexp.MustCompile(`^\d{2}-ع-\d{3}-?\d{2}$`)
	// پلاک زرد با حرف ت (تاکسی)
	taxiPlate := regexp.MustCompile(`^\d{2}-ت-\d{3}-?\d{2}$`)
	// پلاک آبی با حرف D (دیپلمات)
	diplomatPlate := regexp.MustCompile(`^\d{2}-T-\d{3}-\d{2}$`)
	// پلاک آبی با حرف S (سفارت)
	embassyPlate := regexp.MustCompile(`^\d{2}-S-\d{3}-\d{2}$`)
	// پلاک آبی با حرف ز (وزارت دفاع)
	defensePlate := regexp.MustCompile(`^\d{2}-ز-\d{3}-?\d{2}$`)
	// پلاک آبی با حرف ف (ستاد کل نیروهای مسلح)
	militaryStaffPlate := regexp.MustCompile(`^\d{2}-ف-\d{3}-?\d{2}$`)
	// پلاک کرِم یا خاکی (ارتش)
	armyPlate := regexp.MustCompile(`^\d{2}-ش-\d{3}-?\d{2}$`)

	return motorcyclePlate.MatchString(licensePlate) ||
		normalWhitePlate.MatchString(licensePlate) ||
		temporaryPlate.MatchString(licensePlate) ||
		disabledPlate.MatchString(licensePlate) ||
		ceremonyPlate.MatchString(licensePlate) ||
		governmentPlate.MatchString(licensePlate) ||
		policePlate.MatchString(licensePlate) ||
		irgcPlate.MatchString(licensePlate) ||
		agriculturalPlate.MatchString(licensePlate) ||
		publicTransportPlate.MatchString(licensePlate) ||
		taxiPlate.MatchString(licensePlate) ||
		diplomatPlate.MatchString(licensePlate) ||
		embassyPlate.MatchString(licensePlate) ||
		defensePlate.MatchString(licensePlate) ||
		militaryStaffPlate.MatchString(licensePlate) ||
		armyPlate.MatchString(licensePlate)
}

func validateIranianVin(fl validator.FieldLevel) bool {
	vin := fl.Field().String()
	// باید دقیقاً ۱۷ کاراکتر باشد و فقط شامل حروف و عدد (بدون I, O, Q)
	re := regexp.MustCompile(`^[A-HJ-NPR-Z0-9]{17}$`)
	return re.MatchString(vin)
}

func ValidateVehicleTypeCreateRequest(request dto.CreateVehicleTypeRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "Name":
					if fieldError.Tag() == "required" {
						return errors.New("vehicle type name is required")
					}
				case "Description":
					if fieldError.Tag() == "max" {
						return errors.New("description is too long")
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

func ValidateVehicleTypeUpdateRequest(request dto.UpdateVehicleTypeRequest) error {
	if request.NameFa == nil && request.NameEn == nil && request.DescriptionFa == nil && request.DescriptionEn == nil {
		return errors.New("no fields to update")
	}

	if request.NameFa != nil && *request.NameFa == "" {
		return errors.New("name is required")
	}

	if request.NameEn != nil && *request.NameEn == "" {
		return errors.New("name is required")
	}

	return nil
}

func ValidateVehicleBrandCreateRequest(request dto.CreateVehicleBrandRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "VehicleTypeID":
					if fieldError.Tag() == "required" {
						return errors.New("vehicle type id is required")
					}
				case "Name":
					if fieldError.Tag() == "required" {
						return errors.New("brand name is required")
					}
				case "Description":
					if fieldError.Tag() == "max" {
						return errors.New("description is too long")
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

func ValidateVehicleBrandUpdateRequest(request dto.UpdateVehicleBrandRequest) error {
	// Check if at least one field has a value
	if request.VehicleTypeID == nil && request.NameFa == nil && request.NameEn == nil && request.DescriptionFa == nil && request.DescriptionEn == nil {
		return errors.New("no fields to update")
	}

	// If Name is provided, validate it's not empty
	if request.NameFa != nil && *request.NameFa == "" {
		return errors.New("name is required")
	}

	if request.NameEn != nil && *request.NameEn == "" {
		return errors.New("name is required")
	}

	return nil
}

func ValidateVehicleModelCreateRequest(request dto.CreateVehicleModelRequest) error {
	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "BrandID":
					if fieldError.Tag() == "required" {
						return errors.New("brand id is required")
					}
				case "Name":
					if fieldError.Tag() == "required" {
						return errors.New("model name is required")
					}
				case "StartYear":
					if fieldError.Tag() == "year" {
						return errors.New("invalid start year format")
					}
				case "EndYear":
					if fieldError.Tag() == "year" {
						return errors.New("invalid end year format")
					}
				case "Description":
					if fieldError.Tag() == "max" {
						return errors.New("description is too long")
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

func ValidateVehicleModelUpdateRequest(request dto.UpdateVehicleModelRequest) error {
	// Check if at least one field has a value
	if request.BrandID == nil && request.NameFa == nil && request.NameEn == nil && request.DescriptionFa == nil && request.DescriptionEn == nil {
		return errors.New("no fields to update")
	}

	// If Name is provided, validate it's not empty
	if request.NameFa != nil && *request.NameFa == "" {
		return errors.New("name is required")
	}

	if request.NameEn != nil && *request.NameEn == "" {
		return errors.New("name is required")
	}

	return nil
}

func ValidateVehicleGenerationCreateRequest(request dto.CreateVehicleGenerationRequest) error {
	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "ModelID":
					return errors.New("model id is required")
				case "Name":
					return errors.New("generation name is required")
				case "StartYear":
					return errors.New("invalid start year format")
				case "EndYear":
					return errors.New("invalid end year format")
				}
			}
		}
		return errors.New("validation failed")
	}
	return nil
}

func ValidateVehicleGenerationUpdateRequest(request dto.UpdateVehicleGenerationRequest) error {
	// Check if at least one field has a value
	if request.ModelID == nil && request.NameFa == nil && request.NameEn == nil && request.DescriptionFa == nil && request.DescriptionEn == nil &&
		request.StartYear == nil && request.EndYear == nil && request.Engine == nil &&
		request.EngineVolume == nil && request.Cylinders == nil && request.DrivetrainFa == nil &&
		request.DrivetrainEn == nil && request.Gearbox == nil && request.FuelType == nil &&
		request.Battery == nil && request.Seller == nil && request.AssemblyType == nil &&
		request.Assembler == nil {
		return errors.New("no fields to update")
	}

	// If Name is provided, validate it's not empty
	if request.NameFa != nil && *request.NameFa == "" {
		return errors.New("name is required")
	}

	if request.NameEn != nil && *request.NameEn == "" {
		return errors.New("name is required")
	}

	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "StartYear":
					return errors.New("invalid start year format")
				case "EndYear":
					return errors.New("invalid end year format")
				}
			}
		}
		return errors.New("validation failed")
	}
	return nil
}

func ValidateUserVehicleCreateRequest(request dto.CreateUserVehicleRequest) error {
	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	validate.RegisterValidation("date", validateDate)
	validate.RegisterValidation("iranian_license_plate", validateIranianLicensePlate)
	validate.RegisterValidation("iranian_vin", validateIranianVin)

	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "Name":
					if fieldError.Tag() == "required" {
						return errors.New("name is required")
					}
				case "GenerationID":
					if fieldError.Tag() == "required" {
						return errors.New("generation id is required")
					}
				case "ProductionYear":
					if fieldError.Tag() == "year" {
						return errors.New("invalid production year format")
					}
				case "LicensePlate":
					if fieldError.Tag() == "iranian_license_plate" {
						return errors.New("invalid Iranian license plate format")
					}
				case "VIN":
					if fieldError.Tag() == "iranian_vin" {
						return errors.New("invalid Iranian VIN format")
					}
				case "CurrentMileage":
					if fieldError.Tag() == "min" {
						return errors.New("current mileage must be greater than 0")
					}
				case "PurchaseDate":
					if fieldError.Tag() == "date" {
						return errors.New("invalid purchase date format")
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

func ValidateUserVehicleUpdateRequest(request dto.UpdateUserVehicleRequest) error {
	// Check if at least one field has a value
	if request.Name == nil && request.GenerationID == nil && request.ProductionYear == nil &&
		request.Color == nil && request.LicensePlate == nil && request.VIN == nil &&
		request.CurrentMileage == nil && request.PurchaseDate == nil {
		return errors.New("no fields to update")
	}

	// If Name is provided, validate it's not empty
	if request.Name != nil && *request.Name == "" {
		return errors.New("name is required")
	}

	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	validate.RegisterValidation("date", validateDate)
	validate.RegisterValidation("iranian_license_plate", validateIranianLicensePlate)
	validate.RegisterValidation("iranian_vin", validateIranianVin)

	err := validate.Struct(request)
	if err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range validationErrors {
				switch fieldError.Field() {
				case "StartYear":
					return errors.New("invalid start year format")
				case "EndYear":
					return errors.New("invalid end year format")
				}
			}
		}
		return errors.New("validation failed")
	}

	return nil
}
