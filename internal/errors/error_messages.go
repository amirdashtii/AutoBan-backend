package errors

// تعریف پیام‌های خطا به صورت ثابت
var (
	ErrUserNotFound                  = New("User not found", "کاربر یافت نشد")
	ErrInternalServerError           = New("Internal Server Error", "خطای داخلی سرور")
	ErrPhoneNumberRequired           = New("phone number is required", "شماره تلفن الزامی است")
	ErrPasswordRequired              = New("password is required", "رمز عبور الزامی است")
	ErrPhoneNumberOrPasswordRequired = New("phone number and password are required", "شماره تلفن و رمز عبور الزامی است")
)
