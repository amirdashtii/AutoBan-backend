package errors

// General Errors
var (
	ErrBadRequest          = New("Invalid request body", "درخواست معتبر نیست")
	ErrInternalServerError = New("Internal Server Error", "خطای داخلی سرور")
	ErrInvalidToken        = New("invalid or expired token", "توکن نامعتبر یا منقضی شده است")
	ErrUserNotActive       = New("user not active", "کاربر فعال نیست")
	ErrInvalidPhoneNumber  = New("invalid phone number", "شماره تلفن معتبر نیست")
	ErrInvalidPassword     = New("invalid password", "رمز عبور نامعتبر است")
)

// Auth Errors
var (
	ErrUserAlreadyExists            = New("user already exists", "کاربر قبلا ثبت نام کرده است")
	ErrUserNotFound                 = New("User not found", "کاربر یافت نشد")
	TokenGenerationFailed           = New("User Create successfully but failed to generate tokens", "کاربر ساخته شده اما توکن تولید نشد")
	ErrInvalidPhoneNumberOrPassword = New("invalid phone number or password", "شماره تلفن یا رمز عبور معتبر نیست")
	ErrTokenNotFound                = New("authentication token not found", "توکن احراز هویت یافت نشد")
)

// Verification Errors
var (
	ErrVerificationCodeNotFound = New("verification code not found", "کد تایید یافت نشد")
	ErrInvalidVerificationCode  = New("invalid verification code", "کد تایید نامعتبر است")
)

var (
	ErrInvalidEmail               = New("invalid email", "ایمیل معتبر نیست")
	ErrEmailAlreadyExists         = New("email already exists", "این ایمیل قبلاً ثبت شده است")
	ErrInvalidBirthday            = New("invalid birthday format, expected YYYY-MM-DD", "فرمت تاریخ تولد نامعتبر است، فرمت صحیح: YYYY-MM-DD")
	ErrInvalidRole                = New("invalid role", "نقش معتبر نیست")
	ErrInvalidStatus              = New("invalid status", "وضعیت معتبر نیست")
	ErrLoadConfig                 = New("failed to load config", "خطای بارگذاری کانفیگ")
	ErrInvalidTokenFormat         = New("invalid token format", "فرمت توکن احراز هویت نامعتبر است")
	ErrInvalidTokenClaims         = New("invalid token claims", "اطلاعات توکن نامعتبر است")
	ErrAccessDenied               = New("access denied", "شما دسترسی لازم برای این عملیات را ندارید")
	ErrFailedToGetProfile         = New("failed to get profile", "خطای دریافت پروفایل")
	ErrFailedToUpdateProfile      = New("failed to update profile", "خطای به روز رسانی پروفایل")
	ErrFailedToUpdatePassword     = New("failed to update password", "خطای به روز رسانی رمز عبور")
	ErrFailedToDeleteUser         = New("failed to delete user", "خطای حذف کاربر")
	ErrFailedToListUsers          = New("failed to list users", "خطای فهرست کاربران")
	ErrFailedToGetUserById        = New("failed to get user by id", "خطای دریافت کاربر با id")
	ErrFailedToUpdateUser         = New("failed to update user", "خطای به روز رسانی کاربر")
	ErrFailedToHashPassword       = New("failed to hash password", "خطای هش کردن رمز عبور")
	ErrFailedToChangeUserRole     = New("failed to change user role", "خطای تغییر نقش کاربر")
	ErrFailedToChangeUserStatus   = New("failed to change user status", "خطای تغییر وضعیت کاربر")
	ErrFailedToChangeUserPassword = New("failed to change user password", "خطای تغییر رمز عبور کاربر")
	ErrInvalidUserID              = New("invalid user id", "شناسه کاربر نامعتبر است")
	ErrInvalidPurchaseDate        = New("invalid purchase date", "تاریخ خرید نامعتبر است")
	ErrUserVehicleNotOwned        = New("user vehicle not owned by user", "این وسیله نقلیه متعلق به شما نیست")
)

var (
	ErrInvalidVehicleTypeCreateRequest = New("invalid vehicle type create request", "درخواست ساخت نوع ماشین معتبر نیست")
	ErrInvalidVehicleTypeUpdateRequest = New("invalid vehicle type update request", "درخواست به روز رسانی نوع ماشین معتبر نیست")
	ErrFailedToListVehicleTypes        = New("failed to list vehicle types", "خطای فهرست نوع ماشین")
	ErrFailedToGetVehicleType          = New("failed to get vehicle type", "خطای دریافت نوع ماشین")
	ErrInvalidVehicleTypeID            = New("invalid vehicle type id", "شناسه نوع ماشین نامعتبر است")
	ErrFailedToCreateVehicleType       = New("failed to create vehicle type", "خطای ساخت نوع ماشین")
	ErrFailedToUpdateVehicleType       = New("failed to update vehicle type", "خطای به روز رسانی نوع ماشین")
	ErrFailedToDeleteVehicleType       = New("failed to delete vehicle type", "خطای حذف نوع ماشین")
)

var (
	ErrFailedToListVehicleBrands        = New("failed to list vehicle brands", "خطای فهرست برند ماشین")
	ErrFailedToGetVehicleBrand          = New("failed to get vehicle brand", "خطای دریافت برند ماشین")
	ErrInvalidVehicleBrandID            = New("invalid vehicle brand id", "شناسه برند ماشین نامعتبر است")
	ErrFailedToListVehicleBrandsByType  = New("failed to list vehicle brands by type", "خطای فهرست برند ماشین برای نوع ماشین")
	ErrInvalidVehicleBrandType          = New("invalid vehicle brand type", "نوع برند ماشین نامعتبر است")
	ErrInvalidVehicleBrandCreateRequest = New("invalid vehicle brand create request", "درخواست ساخت برند ماشین معتبر نیست")
	ErrInvalidVehicleBrandUpdateRequest = New("invalid vehicle brand update request", "درخواست به روز رسانی برند ماشین معتبر نیست")
	ErrFailedToCreateVehicleBrand       = New("failed to create vehicle brand", "خطای ساخت برند ماشین")
	ErrFailedToUpdateVehicleBrand       = New("failed to update vehicle brand", "خطای به روز رسانی برند ماشین")
	ErrFailedToDeleteVehicleBrand       = New("failed to delete vehicle brand", "خطای حذف برند ماشین")
)

var (
	ErrFailedToListVehicleModels        = New("failed to list vehicle models", "خطای فهرست مدل ماشین")
	ErrFailedToGetVehicleModel          = New("failed to get vehicle model", "خطای دریافت مدل ماشین")
	ErrInvalidVehicleModelID            = New("invalid vehicle model id", "شناسه مدل ماشین نامعتبر است")
	ErrInvalidVehicleModelCreateRequest = New("invalid vehicle model create request", "درخواست ساخت مدل ماشین معتبر نیست")
	ErrInvalidVehicleModelUpdateRequest = New("invalid vehicle model update request", "درخواست به روز رسانی مدل ماشین معتبر نیست")
	ErrFailedToCreateVehicleModel       = New("failed to create vehicle model", "خطای ساخت مدل ماشین")
	ErrFailedToUpdateVehicleModel       = New("failed to update vehicle model", "خطای به روز رسانی مدل ماشین")
	ErrFailedToDeleteVehicleModel       = New("failed to delete vehicle model", "خطای حذف مدل ماشین")
	ErrFailedToListVehicleModelsByBrand = New("failed to list vehicle models by brand", "خطای فهرست مدل ماشین برای برند ماشین")
)

var (
	ErrFailedToListVehicleGenerations        = New("failed to list vehicle generations", "خطای فهرست گنریشن ماشین")
	ErrFailedToListVehicleGenerationsByModel = New("failed to list vehicle generations by models", "خطای فهرست گنریشن ماشین برای مدل")
	ErrFailedToGetVehicleGeneration          = New("failed to get vehicle generation", "خطای دریافت گنریشن ماشین")
	ErrInvalidVehicleGenerationID            = New("invalid vehicle generation id", "شناسه گنریشن ماشین نامعتبر است")
	ErrInvalidVehicleGenerationCreateRequest = New("invalid vehicle generation create request", "درخواست ساخت گنریشن ماشین معتبر نیست")
	ErrInvalidVehicleGenerationUpdateRequest = New("invalid vehicle generation update request", "درخواست به روز رسانی گنریشن ماشین معتبر نیست")
	ErrFailedToCreateVehicleGeneration       = New("failed to create vehicle generation", "خطای ساخت گنریشن ماشین")
	ErrFailedToUpdateVehicleGeneration       = New("failed to update vehicle generation", "خطای به روز رسانی گنریشن ماشین")
	ErrFailedToDeleteVehicleGeneration       = New("failed to delete vehicle generation", "خطای حذف گنریشن ماشین")
)

var (
	ErrInvalidUserVehicleCreateRequest = New("invalid user vehicle create request", "درخواست ساخت وسیله نقلیه کاربر معتبر نیست")
	ErrInvalidUserVehicleUpdateRequest = New("invalid user vehicle update request", "درخواست به روز رسانی وسیله نقلیه کاربر معتبر نیست")
	ErrFailedToCreateUserVehicle       = New("failed to create user vehicle", "خطای ساخت وسیله نقلیه کاربر")
	ErrFailedToListUserVehicles        = New("failed to list user vehicles", "خطای فهرست وسیله نقلیه کاربر")
	ErrFailedToGetUserVehicle          = New("failed to get user vehicle", "خطای دریافت وسیله نقلیه کاربر")
	ErrInvalidUserVehicleID            = New("invalid vehicle user vehicle id", "شناسه وسیله نقلیه کاربر نامعتبر است")
	ErrFailedToUpdateUserVehicle       = New("failed to update user vehicle", "خطای به روز رسانی وسیلقه نقلیه کاربر")
	ErrFailedToDeleteUserVehicle       = New("failed to delete user vehicle", "خطای حذف وسیلثه نقلیه کاربر")
	ErrInvalidVehicleID                = New("invalid vehicle id", "شناسه وسیله نقلیه نامعتبر است")
)

// Oil Change Errors
var (
	ErrInvalidOilChangeCreateRequest = New("invalid oil change create request", "درخواست ساخت تعویض روغن معتبر نیست")
	ErrInvalidOilChangeUpdateRequest = New("invalid oil change update request", "درخواست به روز رسانی تعویض روغن معتبر نیست")
	ErrInvalidOilChangeID            = New("invalid oil change id", "شناسه تعویض روغن نامعتبر است")
	ErrFailedToCreateOilChange       = New("failed to create oil change", "خطای ساخت تعویض روغن")
	ErrFailedToGetOilChange          = New("failed to get oil change", "خطای دریافت تعویض روغن")
	ErrFailedToListOilChanges        = New("failed to list oil changes", "خطای فهرست تعویض روغن")
	ErrFailedToUpdateOilChange       = New("failed to update oil change", "خطای به روز رسانی تعویض روغن")
	ErrFailedToDeleteOilChange       = New("failed to delete oil change", "خطای حذف تعویض روغن")
	ErrInvalidDate                   = New("invalid date format", "فرمت تاریخ نامعتبر است")
	ErrUserVehicleIDRequired         = New("user vehicle id is required", "شناسه وسیله نقلیه کاربر الزامی است")
	ErrOilChangeNotOwned             = New("oil change not owned", "تعویض روغن متعلق به کاربر نیست")
)

// Oil Filter Errors
var (
	ErrInvalidOilFilterCreateRequest = New("invalid oil filter create request", "درخواست ساخت تعویض فیلتر روغن معتبر نیست")
	ErrInvalidOilFilterUpdateRequest = New("invalid oil filter update request", "درخواست به روز رسانی تعویض فیلتر روغن معتبر نیست")
	ErrInvalidOilFilterID            = New("invalid oil filter id", "شناسه تعویض فیلتر روغن نامعتبر است")
	ErrFailedToCreateOilFilter       = New("failed to create oil filter", "خطای ساخت تعویض فیلتر روغن")
	ErrFailedToGetOilFilter          = New("failed to get oil filter", "خطای دریافت تعویض فیلتر روغن")
	ErrOilFilterNotOwned             = New("oil filter not owned", "تعویض فیلتر روغن متعلق به کاربر نیست")
	ErrFailedToListOilFilters        = New("failed to list oil filters", "خطای فهرست تعویض فیلتر روغن")
	ErrFailedToUpdateOilFilter       = New("failed to update oil filter", "خطای به روز رسانی تعویض فیلتر روغن")
	ErrFailedToDeleteOilFilter       = New("failed to delete oil filter", "خطای حذف تعویض فیلتر روغن")
)

// Service Visit Errors
var (
	ErrInvalidServiceVisitCreateRequest = New("invalid service visit create request", "درخواست ساخت بازدید سرویس معتبر نیست")
	ErrInvalidServiceVisitUpdateRequest = New("invalid service visit update request", "درخواست به روز رسانی بازدید سرویس معتبر نیست")
	ErrInvalidServiceVisitID            = New("invalid service visit id", "شناسه بازدید سرویس نامعتبر است")
	ErrFailedToCreateServiceVisit       = New("failed to create service visit", "خطای ساخت بازدید سرویس")
	ErrFailedToGetServiceVisit          = New("failed to get service visit", "خطای دریافت بازدید سرویس")
	ErrFailedToListServiceVisits        = New("failed to list service visits", "خطای فهرست بازدیدهای سرویس")
	ErrFailedToUpdateServiceVisit       = New("failed to update service visit", "خطای به روز رسانی بازدید سرویس")
	ErrFailedToDeleteServiceVisit       = New("failed to delete service visit", "خطای حذف بازدید سرویس")
)
