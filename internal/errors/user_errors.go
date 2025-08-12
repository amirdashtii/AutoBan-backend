package errors

// User service errors
var (
    ErrInvalidEmail               = NewWithCode("INVALID_EMAIL", "invalid email", "ایمیل معتبر نیست")
    ErrEmailAlreadyExists         = NewWithCode("EMAIL_ALREADY_EXISTS", "email already exists", "این ایمیل قبلاً ثبت شده است")
    ErrInvalidBirthday            = NewWithCode("INVALID_BIRTHDAY", "invalid birthday format, expected YYYY-MM-DD", "فرمت تاریخ تولد نامعتبر است، فرمت صحیح: YYYY-MM-DD")
    ErrInvalidRole                = NewWithCode("INVALID_ROLE", "invalid role", "نقش معتبر نیست")
    ErrInvalidStatus              = NewWithCode("INVALID_STATUS", "invalid status", "وضعیت معتبر نیست")
    ErrLoadConfig                 = NewWithCode("LOAD_CONFIG_FAILED", "failed to load config", "خطای بارگذاری کانفیگ")
    ErrFailedToGetProfile         = NewWithCode("GET_PROFILE_FAILED", "failed to get profile", "خطای دریافت پروفایل")
    ErrFailedToUpdateProfile      = NewWithCode("UPDATE_PROFILE_FAILED", "failed to update profile", "خطای به روز رسانی پروفایل")
    ErrFailedToUpdatePassword     = NewWithCode("UPDATE_PASSWORD_FAILED", "failed to update password", "خطای به روز رسانی رمز عبور")
    ErrFailedToDeleteUser         = NewWithCode("DELETE_USER_FAILED", "failed to delete user", "خطای حذف کاربر")
    ErrFailedToListUsers          = NewWithCode("LIST_USERS_FAILED", "failed to list users", "خطای فهرست کاربران")
    ErrFailedToGetUserById        = NewWithCode("GET_USER_BY_ID_FAILED", "failed to get user by id", "خطای دریافت کاربر با id")
    ErrFailedToUpdateUser         = NewWithCode("UPDATE_USER_FAILED", "failed to update user", "خطای به روز رسانی کاربر")
    ErrFailedToHashPassword       = NewWithCode("HASH_PASSWORD_FAILED", "failed to hash password", "خطای هش کردن رمز عبور")
    ErrFailedToChangeUserRole     = NewWithCode("CHANGE_USER_ROLE_FAILED", "failed to change user role", "خطای تغییر نقش کاربر")
    ErrFailedToChangeUserStatus   = NewWithCode("CHANGE_USER_STATUS_FAILED", "failed to change user status", "خطای تغییر وضعیت کاربر")
    ErrFailedToChangeUserPassword = NewWithCode("CHANGE_USER_PASSWORD_FAILED", "failed to change user password", "خطای تغییر رمز عبور کاربر")
    ErrInvalidUserID              = NewWithCode("INVALID_USER_ID", "invalid user id", "شناسه کاربر نامعتبر است")
) 