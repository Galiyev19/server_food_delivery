package helpers

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte("secret-key")

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("Hashed password error - %v", err)
	}
	return string(hashedPassword), nil
}

func GenerateId() uuid.UUID {
	return uuid.New()
}

func GenerateToken(email string) (string, error) {
	expTime := time.Now().Add(time.Hour * 24)
	expNumeric := jwt.NewNumericDate(expTime)

	claims := jwt.RegisteredClaims{
		ExpiresAt: expNumeric,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		Subject:   email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(secretKey)
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	// Парсим токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Здесь возвращаем секретный ключ для верификации подписи токена
		return []byte(secretKey), nil
	})
	// Проверяем на ошибки при парсинге токена
	if err != nil {
		return nil, fmt.Errorf("Error parsing token: %v", err)
	}

	// Проверяем валидность токена
	if !token.Valid {
		return nil, fmt.Errorf("Invalid token")
	}

	// Извлекаем утверждения (claims) из токена
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("Error decoding claims")
	}

	return claims, nil
}

func ConvertToDate(numericTime float64) (time.Time, error) {
	// numericTime, ok := data["exp"].(float64)
	// if !ok {
	// 	return time.Time{}, fmt.Errorf("Failed to convert exp to float64")
	// }
	seconds := int64(numericTime)
	date := time.Unix(seconds, 0)

	return date, nil
}
