package middlewares

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"github.com/AwesomeXjs/tma-server/internal/utils/response"
	"github.com/AwesomeXjs/tma-server/pkg/logger"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

// TelegramValidationMiddleware возвращает middleware для проверки подписи Telegram Mini App
func TelegramValidationMiddleware(botToken string) echo.MiddlewareFunc {
	const mark = "Middleware.TelegramValidation"
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Получаем тело запроса
			var dataUrl []string
			if err := c.Bind(&dataUrl); err != nil {
				// Логируем ошибку и возвращаем ошибку в ответ
				logger.Error("Error binding body", mark, zap.Error(err))
				return response.Response(c, http.StatusBadRequest, "bad request", err.Error())
			}

			// Ожидаем, что первый элемент — это строка данных, а второй — хэш
			if len(dataUrl) != 2 {
				// Логируем ошибку, если формат данных неверный
				return response.Response(c, http.StatusBadRequest, "bad request", "invalid data format")
			}

			// Извлекаем строку данных и хэш
			dataCheckString := dataUrl[0]
			receivedHash := dataUrl[1]
			// Проверяем хэш
			if !validateHash(dataCheckString, receivedHash, botToken) {
				return response.Response(c, http.StatusBadRequest, "bad request", "invalid hash")
			}

			// Если хэш валиден, продолжаем выполнение
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
