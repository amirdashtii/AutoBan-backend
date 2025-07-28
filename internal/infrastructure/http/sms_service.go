package http

import (
	"context"
	"net/http"
	"time"

	"github.com/amirdashtii/AutoBan/internal/dto"
	"github.com/amirdashtii/AutoBan/internal/errors"
)

// SMSService interface for SMS operations
type SMSService interface {
	SendVerificationCode(ctx context.Context, phoneNumber, code string) error
}

// smsService implements SMSService interface
type smsService struct {
	client HTTPClient
	apiKey string
}

// NewSMSService creates a new SMS service
func NewSMSService(baseURL, apiKey string) SMSService {
	client := NewClient(baseURL, 30*time.Second)
	return &smsService{
		client: client,
		apiKey: apiKey,
	}
}

// SendVerificationCode sends verification code via SMS
func (s *smsService) SendVerificationCode(ctx context.Context, phoneNumber, code string) error {
	request := dto.SmsIrRequest{
		Mobile:     phoneNumber,
		TemplateId: "123456",
		Parameters: []struct {
			Name  string `json:"name" validate:"required"`
			Value string `json:"value" validate:"required"`
		}{
			{
				Name:  "code",
				Value: code,
			},
		},
	}

	headers := map[string]string{
		"Content-Type": "application/json",
		"Accept":       "text/plain",
		"x-api-key":    s.apiKey,
	}

	resp, err := s.client.Post(ctx, "/v1/send/verify", request, headers)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return err
	}

	var response dto.SmsIrResponse
	if err := ParseResponse(resp, &response); err != nil {
		return err
	}

	if response.Status != 1 {
		return errors.ErrInternalServerError
	}
	return nil
}
