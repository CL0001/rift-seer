package auth

import (
	"errors"
	"github.com/CL0001/rift-seer/pkg/db"
	"github.com/CL0001/rift-seer/pkg/models"
	"github.com/CL0001/rift-seer/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

var TokenKey = []byte(os.Getenv("TOKEN_KEY"))

func RegisterUser(context echo.Context) error {
	user, err := models.NewUser(context.FormValue("username"), context.FormValue("summoner-name"), context.FormValue("email"), context.FormValue("password"))
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to register user"})
	}

	err = db.AddUser(user)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to register user"})
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate authentication token"})
	}

	return context.JSON(http.StatusOK, map[string]string{
		"message": "register successful",
		"token":   token,
	})
}

func LoginUser(context echo.Context) error {
	email := context.FormValue("email")
	password := context.FormValue("password")

	user, err := db.FetchUser(email)
	if err != nil {
		return context.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid email or password"})
	}

	if err := utils.ComparePasswords(password, user.Password); err != nil {
		return context.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid email or password"})
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return context.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to generate authentication token"})
	}

	return context.JSON(http.StatusOK, map[string]string{
		"message": "register successful",
		"token":   token,
	})
}

func GenerateToken(userID uuid.UUID) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)

	claims := &jwt.RegisteredClaims{
	    ExpiresAt: jwt.NewNumericDate(expirationTime),
	    Subject:   userID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(TokenKey)
}

func IsAuthenticated(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		tokenString := context.Request().Header.Get("Authorization")

		if len(tokenString) < 7 || tokenString[:7] != "Bearer " {
			return context.JSON(http.StatusUnauthorized, map[string]string{"error": "missing or invalid authorization header"})
		}

		tokenString = tokenString[7:]

		claims, err := ValidateToken(tokenString)
		if err != nil {
			return context.JSON(http.StatusUnauthorized, map[string]string{"error": "invalid or expired token"})
		}

		context.Set("user-id", claims.Subject)

		return next(context)
	}
}

func ValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return TokenKey, nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, jwt.ErrTokenExpired
		} else {
			return nil, jwt.ErrTokenSignatureInvalid
		}
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
