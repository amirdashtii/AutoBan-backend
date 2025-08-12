package errors

import (
	"reflect"
)

// ErrorMessage defines a structure for bilingual error messages
// ErrorMessage ساختاری برای پیام‌های خطای دو زبانه تعریف می‌کند

type ErrorMessage struct {
	English string `json:"english"`
	Persian string `json:"persian"`
}

// FieldError represents a validation error for a specific field
// خطای اعتبارسنجی برای یک فیلد مشخص

type FieldError struct {
	Field   string       `json:"field"`
	Message ErrorMessage `json:"message"`
}

// CustomError defines a custom error structure
// CustomError ساختار خطای کاستوم را تعریف می‌کند

type CustomError struct {
	// Machine-readable code (stable across languages)
	Code string `json:"code,omitempty"`
	// Bilingual message for human-readable display
	Message ErrorMessage `json:"message"`
	// Whether the operation can be retried by client later
	Retryable bool `json:"retryable,omitempty"`
	// Optional structured details for debugging (safe to expose)
	Details map[string]any `json:"details,omitempty"`
	// Optional list of field-level errors (validation)
	Fields []FieldError `json:"fields,omitempty"`
	// Not serialized by default; used internally to hint status mapping if needed
	statusHint int `json:"-"`
}

// Error implements the error interface
// متد Error اینترفیس error را پیاده‌سازی می‌کند

func (e *CustomError) Error() string { return e.Message.English }
func (e *CustomError) ErrorFa() string { return e.Message.Persian }

// New creates a new CustomError (backward-compatible)
// تابع New یک خطای کاستوم جدید ایجاد می‌کند

func New(messageEn, messageFa string) *CustomError {
	return &CustomError{
		Code: "GENERAL_ERROR",
		Message: ErrorMessage{English: messageEn, Persian: messageFa},
	}
}

// NewWithCode creates a new CustomError with an explicit code
func NewWithCode(code, messageEn, messageFa string) *CustomError {
	return &CustomError{
		Code: code,
		Message: ErrorMessage{English: messageEn, Persian: messageFa},
	}
}

// WithCode sets/overrides the error code
func (e *CustomError) WithCode(code string) *CustomError {
	e.Code = code
	return e
}

// WithRetryable marks the error as retryable/non-retryable
func (e *CustomError) WithRetryable(retryable bool) *CustomError {
	e.Retryable = retryable
	return e
}

// WithDetail adds a key/value to details
func (e *CustomError) WithDetail(key string, value any) *CustomError {
	if e.Details == nil {
		e.Details = map[string]any{}
	}
	e.Details[key] = value
	return e
}

// WithField appends a field-level validation error
func (e *CustomError) WithField(field, messageEn, messageFa string) *CustomError {
	e.Fields = append(e.Fields, FieldError{Field: field, Message: ErrorMessage{English: messageEn, Persian: messageFa}})
	return e
}

// WithStatusHint optionally hints an HTTP status for responders
func (e *CustomError) WithStatusHint(status int) *CustomError {
	e.statusHint = status
	return e
}

// Copy creates a shallow copy of a CustomError to avoid mutating shared singletons
func Copy(e *CustomError) *CustomError {
	if e == nil {
		return nil
	}
	dup := *e
	// start fresh maps/slices to avoid mutating shared instances after copy
	dup.Details = nil
	dup.Fields = nil
	dup.statusHint = 0
	return &dup
}

// Is compares errors in a robust way (preserved from original)
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