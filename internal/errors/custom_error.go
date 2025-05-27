package errors

// ErrorMessage defines a structure for bilingual error messages
// ErrorMessage ساختاری برای پیام‌های خطای دو زبانه تعریف می‌کند

type ErrorMessage struct {
	English string
	Persian string
}

// CustomError defines a custom error structure
// CustomError ساختار خطای کاستوم را تعریف می‌کند

type CustomError struct {
	Message ErrorMessage
}

// Error implements the error interface
// متد Error اینترفیس error را پیاده‌سازی می‌کند

func (e *CustomError) Error() string {
	return e.Message.English
}

func (e *CustomError) ErrorFa() string {
	return e.Message.Persian
}

// New creates a new CustomError
// تابع New یک خطای کاستوم جدید ایجاد می‌کند

func New(messageEn, messageFa string) *CustomError {
	return &CustomError{
		Message: ErrorMessage{
			English: messageEn,
			Persian: messageFa,
		},
	}
}
