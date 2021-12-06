package middleware

import (
	"io/ioutil"
	"os"
	"strconv"
	"time"
	"user-login-service/pkg/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type userInfo struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UserClaimsInfo struct {
	userInfo
	*jwt.StandardClaims
}

func GenerateToken(ctx *fiber.Ctx) error {
	var userDetails userInfo
	userDetails.ID = ctx.GetRespHeader("id")

	if userDetails.ID == "" {
		return &fiber.Error{
			Code:    fiber.ErrBadRequest.Code,
			Message: "unable to load userId",
		}
	}

	secretKey, err := ioutil.ReadFile(helpers.PrivateKeyFileName)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}
	var JWTSecretKeyExpirationTime = os.Getenv("JWT_SECRET_KEY_EXP")

	expiryTime, err := strconv.Atoi(JWTSecretKeyExpirationTime)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	claims := UserClaimsInfo{
		userDetails,
		&jwt.StandardClaims{
			Subject:   "User Creation and Authentication Service",
			Issuer:    "DalvinCodes",
			Audience:  "github.com/DalvinCodes/user-login-service",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(expiryTime)).Unix(),
		},
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(secretKey)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return &fiber.Error{
			Code:    fiber.ErrInternalServerError.Code,
			Message: err.Error(),
		}
	}

	ctx.Set("Authorization", signedToken)
	return nil
}

func ValidateToken(ctx *fiber.Ctx) error {
	token , err := verifyToken(ctx)
	if err != nil {
		return err
	}
	ctx.Set("Authorization", token.Raw)
	return ctx.Next()
}

func extractToken(ctx *fiber.Ctx) (string, error) {
	bearerToken := ctx.Get("Authorization")
	token, err := helpers.ValidateBearerToken(bearerToken)
	if err != nil {
		return "", err
	}

	return token, nil
}

func verifyToken(ctx *fiber.Ctx) (*jwt.Token, error) {

	tokenString, err := extractToken(ctx)
	if err != nil {
		return nil, err
	}

	publicKey, err := ioutil.ReadFile(helpers.PublicKeyFileName)
	if err != nil {
		return nil, err
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(publicKey)
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) { return key, nil })

	if err := helpers.TokenErrorValidationResponseHandler(err); err != nil {
		return nil, err
	}

	return token, nil
}
