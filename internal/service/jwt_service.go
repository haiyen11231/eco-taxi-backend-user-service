package service

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func GenerateToken(userId uint64, expiry time.Duration) (string, error) {
    expirationTime := time.Now().Add(expiry)
	claims := jwt.MapClaims{
        "user_id": userId,
        "exp":     expirationTime.Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func getKey(userID uint64) string {
    return "user_refresh_token:" + fmt.Sprint(userID)
}

func ValidateToken(encodedToken string, secret string) (*jwt.Token, error) {
    return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
        return []byte(secret), nil
    })
}


func ParseToken(tokenString, secret string) (int64, error) {
	// Parsing the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate that the signing method is HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		log.Println("Error parsing token:", err)
		return 0, err
	}

	// Extract claims and validate them
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		// Verify token expiration
		if !claims.VerifyExpiresAt(time.Now().Unix(), true) {
			return 0, errors.New("token expired")
		}

		// Convert the "id" claim to int64
		idFloat, ok := claims["id"].(float64)
		if !ok {
			return 0, errors.New("invalid ID format in token")
		}

		return int64(idFloat), nil
	}
	return 0, errors.New("Invalid token")
}