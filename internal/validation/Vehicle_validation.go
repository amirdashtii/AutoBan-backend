package validation

import (
	"regexp"
	"strconv"
	"time"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
	"github.com/go-playground/validator/v10"
)

func validateYear(fl validator.FieldLevel) bool {
	year, err := strconv.Atoi(fl.Field().String())
	if err != nil {
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

	// پلاک سواری و تاکسی و دولتی و انتظامی و مناطق آزاد
	normalPlate := regexp.MustCompile(`^\d{2}[آ-ی]\d{3}-?\d{2}$`)
	// پلاک موتورسیکلت (۸ رقم)
	motorPlate := regexp.MustCompile(`^\d{8}$`)
	// پلاک گذر موقت (مثلاً: موقت12-345-67 یا موقت1234567)
	temporaryPlate := regexp.MustCompile(`^موقت\d{2,3}-?\d{3}-?\d{2}$|^موقت\d{7,8}$`)
	// پلاک دیپلمات (مثلاً: D12-345-67)
	diplomatPlate := regexp.MustCompile(`^[DS]\d{2}-\d{3}-\d{2}$`)
	// پلاک معلولین و جانبازان (حرف ژ)
	disabledPlate := regexp.MustCompile(`^\d{2}ژ\d{3}-?\d{2}$`)

	return normalPlate.MatchString(licensePlate) ||
		motorPlate.MatchString(licensePlate) ||
		temporaryPlate.MatchString(licensePlate) ||
		diplomatPlate.MatchString(licensePlate) ||
		disabledPlate.MatchString(licensePlate)
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
		return errors.ErrInvalidVehicleTypeCreateRequest
	}
	return nil
}

func ValidateVehicleTypeUpdateRequest(request dto.UpdateVehicleTypeRequest) error {
	if request.Name == nil && request.Description == nil {
		return errors.ErrInvalidVehicleTypeUpdateRequest
	}

	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidVehicleTypeUpdateRequest
	}

	return nil
}

func ValidateVehicleBrandCreateRequest(request dto.CreateVehicleBrandRequest) error {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidVehicleBrandCreateRequest
	}
	return nil
}

func ValidateVehicleBrandUpdateRequest(request dto.UpdateVehicleBrandRequest) error {
	// Check if at least one field has a value
	if request.VehicleTypeID == nil && request.Name == nil && request.Description == nil {
		return errors.ErrInvalidVehicleBrandUpdateRequest
	}

	// If Name is provided, validate it's not empty
	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidVehicleBrandUpdateRequest
	}

	return nil
}

func ValidateVehicleModelCreateRequest(request dto.CreateVehicleModelRequest) error {
	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidVehicleModelCreateRequest
	}
	return nil
}

func ValidateVehicleModelUpdateRequest(request dto.UpdateVehicleModelRequest) error {
	// Check if at least one field has a value
	if request.BrandID == nil && request.Name == nil && request.Description == nil && request.StartYear == nil && request.EndYear == nil {
		return errors.ErrInvalidVehicleModelUpdateRequest
	}

	// If Name is provided, validate it's not empty
	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidVehicleModelUpdateRequest
	}

	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidVehicleModelUpdateRequest
	}
	return nil
}

func ValidateVehicleGenerationCreateRequest(request dto.CreateVehicleGenerationRequest) error {
	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidVehicleGenerationCreateRequest
	}
	return nil
}

func ValidateVehicleGenerationUpdateRequest(request dto.UpdateVehicleGenerationRequest) error {
	// Check if at least one field has a value
	if request.Name == nil && request.Description == nil && request.ModelID == nil &&
		request.StartYear == nil && request.EndYear == nil && request.EngineType == nil &&
		request.AssemblyType == nil && request.Assembler == nil && request.Transmission == nil &&
		request.EngineSize == nil && request.BodyStyle == nil && request.SpecialFeatures == nil {
		return errors.ErrInvalidVehicleGenerationUpdateRequest
	}

	// If Name is provided, validate it's not empty
	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidVehicleGenerationUpdateRequest
	}

	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidVehicleGenerationUpdateRequest
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
		return errors.ErrInvalidUserVehicleCreateRequest
	}
	return nil
}

func ValidateUserVehicleUpdateRequest(request dto.UpdateUserVehicleRequest) error {
	// Check if at least one field has a value
	if request.Name == nil && request.GenerationID == nil && request.ProductionYear == nil &&
		request.Color == nil && request.LicensePlate == nil && request.VIN == nil &&
		request.CurrentMileage == nil && request.PurchaseDate == nil {
		return errors.ErrInvalidUserVehicleUpdateRequest
	}

	// If Name is provided, validate it's not empty
	if request.Name != nil && *request.Name == "" {
		return errors.ErrInvalidUserVehicleUpdateRequest
	}

	validate := validator.New()
	validate.RegisterValidation("year", validateYear)
	validate.RegisterValidation("date", validateDate)
	validate.RegisterValidation("iranian_license_plate", validateIranianLicensePlate)
	validate.RegisterValidation("iranian_vin", validateIranianVin)

	err := validate.Struct(request)
	if err != nil {
		return errors.ErrInvalidUserVehicleUpdateRequest
	}

	return nil
}
