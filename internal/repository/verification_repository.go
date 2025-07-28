package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/amirdashtii/AutoBan/internal/infrastructure/database"

	"github.com/redis/go-redis/v9"
)

type VerificationRepository interface {
	SaveVerificationCode(ctx context.Context, phoneNumber, code string) error
	GetVerificationCode(ctx context.Context, phoneNumber string) (string, error)
	DeleteVerificationCode(ctx context.Context, phoneNumber string) error
	IsVerificationCodeValid(ctx context.Context, phoneNumber, code string) bool
}

type verificationRepository struct {
	client *redis.Client
}

func NewVerificationRepository() VerificationRepository {
	return &verificationRepository{
		client: database.ConnectRedis(),
	}
}

// ساخت کلید برای ذخیره کد تایید در Redis
func makeVerificationKey(phoneNumber string) string {
	return fmt.Sprintf("verification:%s", phoneNumber)
}

// SaveVerificationCode ذخیره کد تایید در Redis با زمان انقضا
func (r *verificationRepository) SaveVerificationCode(ctx context.Context, phoneNumber, code string) error {
	key := makeVerificationKey(phoneNumber)

	// کد تایید به مدت ۲ دقیقه معتبر است
	err := r.client.Set(ctx, key, code, 2*time.Minute).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetVerificationCode دریافت کد تایید از Redis
func (r *verificationRepository) GetVerificationCode(ctx context.Context, phoneNumber string) (string, error) {
	key := makeVerificationKey(phoneNumber)

	code, err := r.client.Get(ctx, key).Result()

	if err != nil {
		return "", err
	}

	return code, nil
}

// DeleteVerificationCode حذف کد تایید از Redis
func (r *verificationRepository) DeleteVerificationCode(ctx context.Context, phoneNumber string) error {
	key := makeVerificationKey(phoneNumber)

	err := r.client.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}

// IsVerificationCodeValid بررسی اعتبار کد تایید
func (r *verificationRepository) IsVerificationCodeValid(ctx context.Context, phoneNumber, code string) bool {
	storedCode, err := r.GetVerificationCode(ctx, phoneNumber)
	if err != nil {
		return false
	}

	return storedCode == code
}
