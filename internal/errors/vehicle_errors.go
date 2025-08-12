package errors

// Vehicle - Types
var (
    ErrInvalidVehicleTypeCreateRequest = NewWithCode("INVALID_VEHICLE_TYPE_CREATE", "invalid vehicle type create request", "درخواست ساخت نوع ماشین معتبر نیست")
    ErrInvalidVehicleTypeUpdateRequest = NewWithCode("INVALID_VEHICLE_TYPE_UPDATE", "invalid vehicle type update request", "درخواست به روز رسانی نوع ماشین معتبر نیست")
    ErrFailedToListVehicleTypes        = NewWithCode("LIST_VEHICLE_TYPES_FAILED", "failed to list vehicle types", "خطای فهرست نوع ماشین")
    ErrFailedToGetVehicleType          = NewWithCode("GET_VEHICLE_TYPE_FAILED", "failed to get vehicle type", "خطای دریافت نوع ماشین")
    ErrInvalidVehicleTypeID            = NewWithCode("INVALID_VEHICLE_TYPE_ID", "invalid vehicle type id", "شناسه نوع ماشین نامعتبر است")
    ErrFailedToCreateVehicleType       = NewWithCode("CREATE_VEHICLE_TYPE_FAILED", "failed to create vehicle type", "خطای ساخت نوع ماشین")
    ErrFailedToUpdateVehicleType       = NewWithCode("UPDATE_VEHICLE_TYPE_FAILED", "failed to update vehicle type", "خطای به روز رسانی نوع ماشین")
    ErrFailedToDeleteVehicleType       = NewWithCode("DELETE_VEHICLE_TYPE_FAILED", "failed to delete vehicle type", "خطای حذف نوع ماشین")
)

// Vehicle - Brands
var (
    ErrFailedToListVehicleBrands        = NewWithCode("LIST_VEHICLE_BRANDS_FAILED", "failed to list vehicle brands", "خطای فهرست برند ماشین")
    ErrFailedToGetVehicleBrand          = NewWithCode("GET_VEHICLE_BRAND_FAILED", "failed to get vehicle brand", "خطای دریافت برند ماشین")
    ErrInvalidVehicleBrandID            = NewWithCode("INVALID_VEHICLE_BRAND_ID", "invalid vehicle brand id", "شناسه برند ماشین نامعتبر است")
    ErrFailedToListVehicleBrandsByType  = NewWithCode("LIST_VEHICLE_BRANDS_BY_TYPE_FAILED", "failed to list vehicle brands by type", "خطای فهرست برند ماشین برای نوع ماشین")
    ErrInvalidVehicleBrandType          = NewWithCode("INVALID_VEHICLE_BRAND_TYPE", "invalid vehicle brand type", "نوع برند ماشین نامعتبر است")
    ErrInvalidVehicleBrandCreateRequest = NewWithCode("INVALID_VEHICLE_BRAND_CREATE", "invalid vehicle brand create request", "درخواست ساخت برند ماشین معتبر نیست")
    ErrInvalidVehicleBrandUpdateRequest = NewWithCode("INVALID_VEHICLE_BRAND_UPDATE", "invalid vehicle brand update request", "درخواست به روز رسانی برند ماشین معتبر نیست")
    ErrFailedToCreateVehicleBrand       = NewWithCode("CREATE_VEHICLE_BRAND_FAILED", "failed to create vehicle brand", "خطای ساخت برند ماشین")
    ErrFailedToUpdateVehicleBrand       = NewWithCode("UPDATE_VEHICLE_BRAND_FAILED", "failed to update vehicle brand", "خطای به روز رسانی برند ماشین")
    ErrFailedToDeleteVehicleBrand       = NewWithCode("DELETE_VEHICLE_BRAND_FAILED", "failed to delete vehicle brand", "خطای حذف برند ماشین")
)

// Vehicle - Models
var (
    ErrFailedToListVehicleModels        = NewWithCode("LIST_VEHICLE_MODELS_FAILED", "failed to list vehicle models", "خطای فهرست مدل ماشین")
    ErrFailedToGetVehicleModel          = NewWithCode("GET_VEHICLE_MODEL_FAILED", "failed to get vehicle model", "خطای دریافت مدل ماشین")
    ErrInvalidVehicleModelID            = NewWithCode("INVALID_VEHICLE_MODEL_ID", "invalid vehicle model id", "شناسه مدل ماشین نامعتبر است")
    ErrInvalidVehicleModelCreateRequest = NewWithCode("INVALID_VEHICLE_MODEL_CREATE", "invalid vehicle model create request", "درخواست ساخت مدل ماشین معتبر نیست")
    ErrInvalidVehicleModelUpdateRequest = NewWithCode("INVALID_VEHICLE_MODEL_UPDATE", "invalid vehicle model update request", "درخواست به روز رسانی مدل ماشین معتبر نیست")
    ErrFailedToCreateVehicleModel       = NewWithCode("CREATE_VEHICLE_MODEL_FAILED", "failed to create vehicle model", "خطای ساخت مدل ماشین")
    ErrFailedToUpdateVehicleModel       = NewWithCode("UPDATE_VEHICLE_MODEL_FAILED", "failed to update vehicle model", "خطای به روز رسانی مدل ماشین")
    ErrFailedToDeleteVehicleModel       = NewWithCode("DELETE_VEHICLE_MODEL_FAILED", "failed to delete vehicle model", "خطای حذف مدل ماشین")
    ErrFailedToListVehicleModelsByBrand = NewWithCode("LIST_VEHICLE_MODELS_BY_BRAND_FAILED", "failed to list vehicle models by brand", "خطای فهرست مدل ماشین برای برند ماشین")
)

// Vehicle - Generations
var (
    ErrFailedToListVehicleGenerations        = NewWithCode("LIST_VEHICLE_GENERATIONS_FAILED", "failed to list vehicle generations", "خطای فهرست گنریشن ماشین")
    ErrFailedToListVehicleGenerationsByModel = NewWithCode("LIST_VEHICLE_GENERATIONS_BY_MODEL_FAILED", "failed to list vehicle generations by models", "خطای فهرست گنریشن ماشین برای مدل")
    ErrFailedToGetVehicleGeneration          = NewWithCode("GET_VEHICLE_GENERATION_FAILED", "failed to get vehicle generation", "خطای دریافت گنریشن ماشین")
    ErrInvalidVehicleGenerationID            = NewWithCode("INVALID_VEHICLE_GENERATION_ID", "invalid vehicle generation id", "شناسه گنریشن ماشین نامعتبر است")
    ErrInvalidVehicleGenerationCreateRequest = NewWithCode("INVALID_VEHICLE_GENERATION_CREATE", "invalid vehicle generation create request", "درخواست ساخت گنریشن ماشین معتبر نیست")
    ErrInvalidVehicleGenerationUpdateRequest = NewWithCode("INVALID_VEHICLE_GENERATION_UPDATE", "invalid vehicle generation update request", "درخواست به روز رسانی گنریشن ماشین معتبر نیست")
    ErrFailedToCreateVehicleGeneration       = NewWithCode("CREATE_VEHICLE_GENERATION_FAILED", "failed to create vehicle generation", "خطای ساخت گنریشن ماشین")
    ErrFailedToUpdateVehicleGeneration       = NewWithCode("UPDATE_VEHICLE_GENERATION_FAILED", "failed to update vehicle generation", "خطای به روز رسانی گنریشن ماشین")
    ErrFailedToDeleteVehicleGeneration       = NewWithCode("DELETE_VEHICLE_GENERATION_FAILED", "failed to delete vehicle generation", "خطای حذف گنریشن ماشین")
)

// Vehicle - User Vehicles
var (
    ErrInvalidUserVehicleCreateRequest = NewWithCode("INVALID_USER_VEHICLE_CREATE", "invalid user vehicle create request", "درخواست ساخت وسیله نقلیه کاربر معتبر نیست")
    ErrInvalidUserVehicleUpdateRequest = NewWithCode("INVALID_USER_VEHICLE_UPDATE", "invalid user vehicle update request", "درخواست به روز رسانی وسیله نقلیه کاربر معتبر نیست")
    ErrFailedToCreateUserVehicle       = NewWithCode("CREATE_USER_VEHICLE_FAILED", "failed to create user vehicle", "خطای ساخت وسیله نقلیه کاربر")
    ErrFailedToListUserVehicles        = NewWithCode("LIST_USER_VEHICLES_FAILED", "failed to list user vehicles", "خطای فهرست وسیله نقلیه کاربر")
    ErrFailedToGetUserVehicle          = NewWithCode("GET_USER_VEHICLE_FAILED", "failed to get user vehicle", "خطای دریافت وسیله نقلیه کاربر")
    ErrInvalidUserVehicleID            = NewWithCode("INVALID_USER_VEHICLE_ID", "invalid vehicle user vehicle id", "شناسه وسیله نقلیه کاربر نامعتبر است")
    ErrFailedToUpdateUserVehicle       = NewWithCode("UPDATE_USER_VEHICLE_FAILED", "failed to update user vehicle", "خطای به روز رسانی وسیلقه نقلیه کاربر")
    ErrFailedToDeleteUserVehicle       = NewWithCode("DELETE_USER_VEHICLE_FAILED", "failed to delete user vehicle", "خطای حذف وسیلثه نقلیه کاربر")
    ErrInvalidVehicleID                = NewWithCode("INVALID_VEHICLE_ID", "invalid vehicle id", "شناسه وسیله نقلیه نامعتبر است")
    ErrInvalidPurchaseDate             = NewWithCode("INVALID_PURCHASE_DATE", "invalid purchase date", "تاریخ خرید نامعتبر است")
    ErrUserVehicleNotOwned             = NewWithCode("USER_VEHICLE_NOT_OWNED", "user vehicle not owned by user", "این وسیله نقلیه متعلق به شما نیست")
) 