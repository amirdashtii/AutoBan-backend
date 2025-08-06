package errors
import "reflect"

// ErrorMessage defines a structure for bilingual error messages
// ErrorMessage ساختاری برای پیام‌های خطای دو زبانه تعریف می‌کند

type ErrorMessage struct {
	English string `json:"english"`
	Persian string `json:"persian"`
}

// CustomError defines a custom error structure
// CustomError ساختار خطای کاستوم را تعریف می‌کند

type CustomError struct {
	Message ErrorMessage `json:"message"`
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


func Is(err, target error) bool {
	if err == nil || target == nil {
		return err == target
	}

	isComparable := reflect.TypeOf(target).Comparable()
	return is(err, target, isComparable)
}

func is(err, target error, targetComparable bool) bool {
	for {
		if targetComparable && err == target {
			return true
		}
		if x, ok := err.(interface{ Is(error) bool }); ok && x.Is(target) {
			return true
		}
		switch x := err.(type) {
		case interface{ Unwrap() error }:
			err = x.Unwrap()
			if err == nil {
				return false
			}
		case interface{ Unwrap() []error }:
			for _, err := range x.Unwrap() {
				if is(err, target, targetComparable) {
					return true
				}
			}
			return false
		default:
			return false
		}
	}
}