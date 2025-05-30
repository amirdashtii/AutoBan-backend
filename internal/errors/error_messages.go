package errors

// تعریف پیام‌های خطا به صورت ثابت
var (
	ErrInvalidRequestBody            = New("Invalid request body", "درخواست معتبر نیست")
	ErrUserNotFound                  = New("User not found", "کاربر یافت نشد")
	ErrInternalServerError           = New("Internal Server Error", "خطای داخلی سرور")
	ErrInvalidPhoneNumber            = New("Invalid phone number", "شماره تلفن معتبر نیست")
	ErrInvalidPassword               = New("password must be at least 8 characters long and include uppercase, lowercase, and numbers", "رمز عبور باید حداقل 8 کاراکتر باشد و شامل حروف بزرگ، کوچک و اعداد باشد")
	ErrLoadConfig                    = New("failed to load config", "خطای بارگذاری کانفیگ")
	ErrUserAlreadyExists             = New("user already exists", "کاربر قبلا ثبت نام کرده است")
	ErrInvalidPhoneNumberOrPassword = New("invalid phone number or password", "شماره تلفن یا رمز عبور معتبر نیست")
)
