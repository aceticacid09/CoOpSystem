package interfaces

import (
	"context"
)

// AuthRepository handles logging and API key validation
type AuthRepository interface {
	LogRequest(ctx context.Context, clientID *int, userID *string, endpoint, method string) error
	ValidateJWTToken(ctx context.Context, token string) (map[string]interface{}, error)
	ValidateAPIKey(ctx context.Context, apiKey string) (int, error)
}
