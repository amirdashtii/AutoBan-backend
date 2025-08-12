package errors

// General/Common Errors
var (
    ErrBadRequest          = NewWithCode("BAD_REQUEST", "Invalid request body", "درخواست معتبر نیست")
    ErrInternalServerError = NewWithCode("INTERNAL_SERVER_ERROR", "Internal Server Error", "خطای داخلی سرور")
    ErrInvalidToken        = NewWithCode("INVALID_TOKEN", "invalid or expired token", "توکن نامعتبر یا منقضی شده است")
    ErrUserNotActive       = NewWithCode("USER_NOT_ACTIVE", "user not active", "کاربر فعال نیست")
    ErrInvalidPhoneNumber  = NewWithCode("INVALID_PHONE_NUMBER", "invalid phone number", "شماره تلفن معتبر نیست")
    ErrInvalidPassword     = NewWithCode("INVALID_PASSWORD", "invalid password", "رمز عبور نامعتبر است")
    ErrTokenNotFound       = NewWithCode("TOKEN_NOT_FOUND", "authentication token not found", "توکن احراز هویت یافت نشد")
    ErrAccessDenied        = NewWithCode("ACCESS_DENIED", "access denied", "شما دسترسی لازم برای این عملیات را ندارید")
)
