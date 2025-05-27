package errors

import "fmt"

// ErrorMessage defines a structure for bilingual error messages
// ErrorMessage ساختاری برای پیام‌های خطای دو زبانه تعریف می‌کند

type ErrorMessage struct {
	English string
	Persian string
}

// CustomError defines a custom error structure
// CustomError ساختار خطای کاستوم را تعریف می‌کند

type CustomError struct {
	Code    int
	Message ErrorMessage
}

// Error implements the error interface
// متد Error اینترفیس error را پیاده‌سازی می‌کند

func (e *CustomError) Error(lang string) string {
	switch lang {
	case "fa":
		return fmt.Sprintf("خطا %d: %s", e.Code, e.Message.Persian)
	default:
		return fmt.Sprintf("Error %d: %s", e.Code, e.Message.English)
	}
}

// New creates a new CustomError
// تابع New یک خطای کاستوم جدید ایجاد می‌کند

func New(code int, messageEn, messageFa string) *CustomError {
	return &CustomError{
		Code: code,
		Message: ErrorMessage{
			English: messageEn,
			Persian: messageFa,
		},
	}
}
