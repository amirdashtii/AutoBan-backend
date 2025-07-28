# SMS Verification Code System

This system is designed for sending and verifying verification codes via SMS.

## System Architecture

### 1. Infrastructure Layer
- **`internal/infrastructure/http/client.go`** - General HTTP client
- **`internal/infrastructure/http/sms_service.go`** - SMS service

### 2. Repository Layer
- **`internal/repository/verification_repository.go`** - Verification code management in Redis

### 3. UseCase Layer
- **`internal/usecase/auth_usecase.go`** - Verification code business logic

### 4. Controller Layer
- **`internal/interface/controller/auth_controller.go`** - API endpoints

## Workflow

### 1. Send Verification Code
```
POST /api/v1/auth/send-verifycode
```

**Request:**
```json
{
  "phone_number": "09123456789"
}
```

**Response:**
```json
{
  "message": "verify code sent successfully"
}
```

**Steps:**
1. Phone number validation
2. Check user existence
3. Generate 6-digit random code
4. Store code in Redis (2 minutes validity)
5. Send code via SMS

### 2. Verify Code
```
POST /api/v1/auth/verify-code
```

**Request:**
```json
{
  "phone_number": "09123456789",
  "code": "123456"
}
```

**Response:**
```json
{
  "message": "code verified successfully"
}
```

**Steps:**
1. Request validation
2. Check code in Redis
3. Compare sent code with stored code
4. Delete code after successful verification

## Configuration

### `.env` File
```env
# SMS Service Configuration
SMS_BASE_URL=https://api.sms.ir
SMS_X_API_KEY=your-api-key
```

### `config/config.go` File
```go
SMS struct {
    BaseURL string `mapstructure:"base_url"`
    XAPIKey string `mapstructure:"x_api_key"`
} `mapstructure:"sms"`
```

## Security Features

### 1. Time Limitation
- Verification code is valid for **2 minutes**
- After expiration, code is automatically removed from Redis

### 2. Auto Deletion
- After successful verification, code is deleted from Redis
- Prevents code reuse

### 3. Validation
- Phone number must have Iranian format (09XXXXXXXXX)
- Code must be exactly 6 characters
- Only numbers are allowed

## Errors

### Common Errors
- `ErrInvalidPhoneNumber` - Invalid phone number
- `ErrInvalidVerificationCode` - Invalid verification code
- `ErrVerificationCodeNotFound` - Verification code not found
- `ErrVerificationCodeExpired` - Verification code expired

### HTTP Status Codes
- `200` - Successful verification
- `400` - Invalid request
- `404` - Verification code not found
- `500` - Internal server error

## Usage Examples

### 1. Send Verification Code
```bash
curl -X POST http://localhost:8080/api/v1/auth/send-verifycode \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "09123456789"}'
```

### 2. Verify Code
```bash
curl -X POST http://localhost:8080/api/v1/auth/verify-code \
  -H "Content-Type: application/json" \
  -d '{"phone_number": "09123456789", "code": "123456"}'
```

## Important Notes

1. **Redis**: System depends on Redis, ensure Redis is running
2. **SMS Service**: SMS service configuration must be correct
3. **Logging**: All operations are logged
4. **Error Handling**: Errors are properly managed
5. **Validation**: All inputs are validated

## Future Development

- Rate limiting for requests in specific time intervals
- Support for variable-length verification codes
- Email code delivery
- QR code support
- Two-factor authentication (2FA)

## Technical Details

### Code Generation
```go
func generateCode() string {
    return fmt.Sprintf("%d", rand.Intn(1000000))
}
```

### Redis Key Structure
```go
func makeVerificationKey(phoneNumber string) string {
    return fmt.Sprintf("verification:%s", phoneNumber)
}
```

### SMS Service Integration
```go
type SMSService interface {
    SendVerificationCode(ctx context.Context, phoneNumber, code string) error
}
```

## Testing

### Unit Tests
- Repository layer tests
- UseCase layer tests
- Validation tests

### Integration Tests
- SMS service integration
- Redis integration
- API endpoint tests

## Monitoring

### Metrics
- SMS delivery success rate
- Code verification success rate
- Redis operation performance
- API response times

### Alerts
- SMS service failures
- Redis connection issues
- High error rates
- Code expiration warnings 