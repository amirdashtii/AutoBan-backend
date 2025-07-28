# HTTP Client Infrastructure

This package contains the necessary tools for communicating with external services.

## File Structure

- `client.go` - General HTTP client for sending requests
- `sms_service.go` - SMS service for sending verification messages

## Usage

### 1. General HTTP Client

```go
import "github.com/amirdashtii/AutoBan/internal/infrastructure/http"

// Create new client
client := http.NewClient("https://api.example.com", 30*time.Second)

// Send GET request
resp, err := client.Get(ctx, "/api/users", map[string]string{
    "Authorization": "Bearer token",
})

// Send POST request
requestBody := map[string]interface{}{
    "name": "John",
    "email": "john@example.com",
}
resp, err := client.Post(ctx, "/api/users", requestBody, nil)

// Process response
var result UserResponse
err = http.ParseResponse(resp, &result)
```

### 2. SMS Service

```go
import "github.com/amirdashtii/AutoBan/internal/infrastructure/http"

// Create SMS service
smsService := http.NewSMSService(
    "https://api.sms.ir",
    "your-api-key",
)

// Send verification code
err := smsService.SendVerificationCode(ctx, "09123456789", "123456")
```

## Configuration

To use the SMS service, add the following settings to your `.env` file:

```env
SMS_BASE_URL=https://api.sms.ir
SMS_X_API_KEY=your-api-key
```

## Best Practices

1. **Timeout**: Always set appropriate timeout for requests
2. **Error Handling**: Handle errors properly
3. **Logging**: Use logger to record requests and responses
4. **Context**: Use context to cancel requests
5. **Headers**: Set appropriate headers for each service

## Complete Example

```go
type ExternalService struct {
    client http.HTTPClient
}

func NewExternalService() *ExternalService {
    return &ExternalService{
        client: http.NewClient("https://api.external.com", 30*time.Second),
    }
}

func (s *ExternalService) GetUserData(ctx context.Context, userID string) (*UserData, error) {
    headers := map[string]string{
        "Authorization": "Bearer " + s.apiKey,
        "Content-Type": "application/json",
    }
    
    resp, err := s.client.Get(ctx, "/users/"+userID, headers)
    if err != nil {
        return nil, fmt.Errorf("failed to get user data: %w", err)
    }
    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("external service returned status %d", resp.StatusCode)
    }
    
    var userData UserData
    if err := http.ParseResponse(resp, &userData); err != nil {
        return nil, fmt.Errorf("failed to parse response: %w", err)
    }
    
    return &userData, nil
}
```

## HTTP Client Interface

```go
type HTTPClient interface {
    Get(ctx context.Context, url string, headers map[string]string) (*http.Response, error)
    Post(ctx context.Context, url string, body interface{}, headers map[string]string) (*http.Response, error)
    Put(ctx context.Context, url string, body interface{}, headers map[string]string) (*http.Response, error)
    Delete(ctx context.Context, url string, headers map[string]string) (*http.Response, error)
}
```

## SMS Service Interface

```go
type SMSService interface {
    SendVerificationCode(ctx context.Context, phoneNumber, code string) error
}
```

## Error Handling

The HTTP client provides comprehensive error handling:

```go
resp, err := client.Post(ctx, "/api/endpoint", data, headers)
if err != nil {
    // Handle network errors, timeouts, etc.
    return fmt.Errorf("request failed: %w", err)
}
defer resp.Body.Close()

if resp.StatusCode != http.StatusOK {
    // Handle HTTP status errors
    return fmt.Errorf("unexpected status: %d", resp.StatusCode)
}
```

## Response Parsing

Use the `ParseResponse` function to easily parse JSON responses:

```go
var response MyResponse
if err := http.ParseResponse(resp, &response); err != nil {
    return fmt.Errorf("failed to parse response: %w", err)
}
```

## Custom Headers

Set custom headers for different services:

```go
headers := map[string]string{
    "Authorization": "Bearer " + token,
    "X-API-Key": apiKey,
    "Content-Type": "application/json",
    "Accept": "application/json",
}

resp, err := client.Post(ctx, "/api/endpoint", data, headers)
```

## Timeout Configuration

Set appropriate timeouts based on service requirements:

```go
// Short timeout for fast services
fastClient := http.NewClient("https://fast-api.com", 5*time.Second)

// Longer timeout for slow services
slowClient := http.NewClient("https://slow-api.com", 60*time.Second)
```

## Logging

All HTTP operations are logged using the application logger:

```go
// Request logging
logger.Info("Making HTTP request", "method", "POST", "url", "/api/endpoint")

// Response logging
logger.Info("HTTP response received", "status", resp.StatusCode)
```

## Testing

### Mock HTTP Client

For testing, you can create a mock HTTP client:

```go
type MockHTTPClient struct {
    responses map[string]*http.Response
}

func (m *MockHTTPClient) Get(ctx context.Context, url string, headers map[string]string) (*http.Response, error) {
    if resp, exists := m.responses[url]; exists {
        return resp, nil
    }
    return nil, fmt.Errorf("not found")
}
```

### Integration Tests

Test the actual HTTP client with real services:

```go
func TestSMSService_Integration(t *testing.T) {
    smsService := http.NewSMSService("https://test-api.com", "test-key")
    
    err := smsService.SendVerificationCode(context.Background(), "09123456789", "123456")
    assert.NoError(t, err)
}
```

## Security Considerations

1. **API Keys**: Never hardcode API keys in source code
2. **HTTPS**: Always use HTTPS for external communications
3. **Timeouts**: Set reasonable timeouts to prevent hanging requests
4. **Validation**: Validate all responses before processing
5. **Rate Limiting**: Implement rate limiting for external services

## Performance Optimization

1. **Connection Pooling**: The HTTP client uses Go's default connection pooling
2. **Keep-Alive**: Connections are reused when possible
3. **Compression**: Enable gzip compression for large responses
4. **Caching**: Implement caching for frequently accessed data

## Monitoring

Monitor HTTP client performance:

```go
// Track request duration
start := time.Now()
resp, err := client.Get(ctx, "/api/endpoint", nil)
duration := time.Since(start)

// Log metrics
logger.Info("HTTP request completed", 
    "duration", duration,
    "status", resp.StatusCode,
    "url", "/api/endpoint")
``` 