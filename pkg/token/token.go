package token

import (
	log "odilbekqazaqov4657/my_blog_backend/pkg/logger"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/saidamir98/udevs_pkg/logger"
)

type JWTHandler struct {
	Sub        string // useer_id == sub
	Exp        string
	Iat        string // token yaratilgan vaqt
	Aud        []string
	Role       string
	SighnedKey string
	Log        log.Log
	Token      string
	Timeout    int
}

type CustomClaims struct {
	*jwt.Token
	Sub  string   `json:"sub"`
	Exp  string   `json:"exp"`
	Iat  string   `json:"iat"`
	Aud  []string `json:"aud"`
	Role string   `json:"role"`
}

func (j *JWTHandler) GenerateToken() (accessToken string, err error) {

	var token *jwt.Token
	var claims jwt.MapClaims

	token = jwt.New(jwt.SigningMethodHS256)

	claims = token.Claims.(jwt.MapClaims)

	claims["sub"] = j.Sub
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(j.Timeout)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["aud"] = j.Aud
	claims["user_role"] = j.Role

	accessToken, err = token.SignedString([]byte(j.SighnedKey))

	if err != nil {

		j.Log.Error("error on generating token !", logger.Error(err))
		return
	}

	return
}
