package Utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"laundry/Constant"
	"os"
	"strconv"
)

func GenerateSecretTokenAndIdentifier(role string) (error, string) {
	var secretToken string

	switch role {
	case Constant.User:
		secretToken = os.Getenv("JWTUSERSECRETTOKEN")
	default:
		return BadRequestResponseWithMessage("role not found"), ""
	}

	if secretToken == "" {
		return BadRequestResponseWithMessage("secret token or identifier is empty"), ""
	}

	return nil, secretToken
}

func SetJwtClaims(c echo.Context, id string) {
	c.Set(Constant.UserClaimsId, id)
}

func GetJwtClaims(c echo.Context) uint {
	return c.Get(Constant.UserClaimsId).(uint)
}

func GenerateJwtToken(id uint, role string) (string, error) {
	err, secretToken := GenerateSecretTokenAndIdentifier(role)
	if err != nil {
		return "", err
	}

	token := jwt.New(jwt.SigningMethodHS512)
	claims := token.Claims.(jwt.MapClaims)
	claims[Constant.UserClaimsId] = strconv.Itoa(int(id))

	t, err := token.SignedString([]byte(secretToken))

	if err != nil {
		return "", err
	}
	return t, nil
}
