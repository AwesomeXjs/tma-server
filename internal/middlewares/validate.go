package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/http"
	"strings"

	"github.com/AwesomeXjs/tma-server/internal/utils"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// TelegramValidationMiddleware возвращает middleware для проверки подписи Telegram Mini App
func TelegramValidationMiddleware(botToken string) echo.MiddlewareFunc {
	const mark = "Middleware.TelegramValidation"
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("TGWebAppToken")
			if authHeader == "" {
				logger.Error("Missing 'TGWebAppToken' header", mark)
				return utils.Response(c, http.StatusUnauthorized, "unauthorized", "missing authorization header")
			}

			parts := strings.SplitN(authHeader, " ", 2)
			if len(parts) != 2 || parts[0] != "TGWebApp" {
				logger.Error("Invalid Authorization format", mark)
				return utils.Response(c, http.StatusUnauthorized, "unauthorized", "invalid authorization format")
			}

			decoded, err := base64.StdEncoding.DecodeString(parts[1])
			if err != nil {
				logger.Error("Base64 decode error", mark, zap.Error(err))
				return utils.Response(c, http.StatusBadRequest, "bad request", "invalid base64 encoding")
			}

			// Исправленное разделение данных
			lastColonIndex := strings.LastIndex(string(decoded), ":")
			if lastColonIndex == -1 {
				logger.Error("Invalid auth data format", mark)
				return utils.Response(c, http.StatusBadRequest, "bad request", "invalid auth data format")
			}

			dataCheckString := string(decoded)[:lastColonIndex]
			receivedHash := string(decoded)[lastColonIndex+1:]
			
			// Проверяем хэш
			if !validateHash(dataCheckString, receivedHash, botToken) {
				logger.Warn("Invalid hash received", mark)
				return utils.Response(c, http.StatusForbidden, "forbidden", "invalid hash")
			}

			return next(c)
		}
	}
}

// validateHash проверяет хэш с использованием токена бота
func validateHash(dataCheckString, receivedHash, botToken string) bool {
	// Генерация секретного ключа на основе botToken
	secretKey := hmacSHA256([]byte(botToken), []byte("WebAppData"))

	// Генерация хэша с использованием секретного ключа
	calculatedHash := hmacSHA256([]byte(dataCheckString), secretKey)

	// Преобразуем вычисленный хэш в строку
	calculatedHashHex := hex.EncodeToString(calculatedHash)

	// Сравнение вычисленного хэша с полученным
	return strings.EqualFold(calculatedHashHex, receivedHash)
}

// hmacSHA256 генерирует HMAC-SHA256
func hmacSHA256(data, key []byte) []byte {
	h := hmac.New(sha256.New, key)
	h.Write(data)
	return h.Sum(nil)
}
