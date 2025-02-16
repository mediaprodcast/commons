package storage

import (
	"fmt"

	"github.com/golang-jwt/jwt"
	pb "github.com/mediaprodcast/proto/genproto/go/storage/v1"
	"google.golang.org/protobuf/encoding/protojson"
)

var jwtSecretKey = []byte("2861c20e1d54ee422b33d6fafec308976eb93a768cada05ffd82280d123efc55")

// EncodeStorageConfig converts a StorageConfig struct into a JWT token string signed with the provided secret key.
func EncodeStorageConfig(cfg *pb.StorageConfig) (string, error) {
	// Marshal the struct into a JSON byte slice using protojson.
	jsonData, err := protojson.MarshalOptions{EmitUnpopulated: true}.Marshal(cfg)
	if err != nil {
		return "", fmt.Errorf("failed to marshal config: %w", err)
	}

	// Create JWT claims with the JSON data as a "storage_config" claim.
	claims := jwt.MapClaims{
		"storage_config": string(jsonData),
	}

	// Create a new token with the claims and sign it using HS256.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return tokenString, nil
}

// DecodeStorageConfig converts a JWT token string back into a StorageConfig struct after verifying the signature with the provided secret key.
func DecodeStorageConfig(tokenString string) (*pb.StorageConfig, error) {
	// Parse the token and verify the signature using the secret key.
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecretKey, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Check if the token is valid.
	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	// Extract claims from the token.
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token claims format")
	}

	// Retrieve the "storage_config" claim.
	jsonData, ok := claims["storage_config"].(string)
	if !ok {
		return nil, fmt.Errorf("storage_config claim is missing or invalid")
	}

	// Unmarshal the JSON data into the StorageConfig struct.
	var cfg pb.StorageConfig
	if err := protojson.Unmarshal([]byte(jsonData), &cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
