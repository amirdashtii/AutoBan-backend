package errors

// Auth service errors
var (
    ErrUserAlreadyExists            = NewWithCode("USER_ALREADY_EXISTS", "user already exists", "کاربر قبلا ثبت نام کرده است")
    ErrUserNotFound                 = NewWithCode("USER_NOT_FOUND", "User not found", "کاربر یافت نشد")
    TokenGenerationFailed           = NewWithCode("TOKEN_GENERATION_FAILED", "User Create successfully but failed to generate tokens", "کاربر ساخته شده اما توکن تولید نشد")
    ErrInvalidPhoneNumberOrPassword = NewWithCode("INVALID_PHONE_OR_PASSWORD", "invalid phone number or password", "شماره تلفن یا رمز عبور معتبر نیست")
    ErrInvalidTokenFormat           = NewWithCode("INVALID_TOKEN_FORMAT", "invalid token format", "فرمت توکن احراز هویت نامعتبر است")
    ErrInvalidTokenClaims           = NewWithCode("INVALID_TOKEN_CLAIMS", "invalid token claims", "اطلاعات توکن نامعتبر است")
)

// Verification sub-domain errors
var (
    ErrVerificationCodeNotFound = NewWithCode("VERIFICATION_CODE_NOT_FOUND", "verification code not found", "کد تایید یافت نشد")
    ErrInvalidVerificationCode  = NewWithCode("INVALID_VERIFICATION_CODE", "invalid verification code", "کد تایید نامعتبر است")
) 