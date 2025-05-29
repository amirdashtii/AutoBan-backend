package errors

// تعریف پیام‌های خطا به صورت ثابت
var (
	ErrUserNotFound                  = New("User not found", "کاربر یافت نشد")
	ErrInternalServerError           = New("Internal Server Error", "خطای داخلی سرور")
	ErrInvalidPhoneNumber            = New("Invalid phone number", "شماره تلفن معتبر نیست")
	ErrInvalidPassword               = New("password must be at least 8 characters long and include uppercase, lowercase, and numbers", "رمز عبور باید حداقل 8 کاراکتر باشد و شامل حروف بزرگ، کوچک و اعداد باشد")
	ErrPhoneNumberOrPasswordRequired = New("phone number and password are required", "شماره تلفن و رمز عبور الزامی است")
	ErrLoadConfig                    = New("failed to load config", "خطای بارگذاری کانفیگ")
)
