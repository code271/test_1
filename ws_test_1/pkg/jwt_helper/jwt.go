package jwt_helper

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"sync"
	"time"
)

var privateKey *rsa.PrivateKey
var publicKey *rsa.PublicKey
var once sync.Once

type CustomClaims struct {
	UserID      int64  `json:"user_id,string,omitempty"`
	AccountName string `json:"account_name,omitempty"`
	RoleID      int32  `json:"role_id,omitempty"`
	jwt.StandardClaims
}

type JwtRsaKey struct {
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}

func Generate(usrID int64, accountName string, roleId int32) (tokenString string, err error) {
	now := time.Now()
	exp := now.Add(30 * 24 * time.Hour)
	var claims CustomClaims
	claims.UserID = usrID
	claims.AccountName = accountName
	claims.RoleID = roleId
	claims.StandardClaims.ExpiresAt = exp.Unix()
	claims.StandardClaims.IssuedAt = now.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
	tokenString, err = token.SignedString(privateKey)
	return
}

func Verify(tokenString string) (claims *CustomClaims, ok bool) {
	parseToken, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return publicKey, nil
	})
	if err != nil {
		return
	}
	claims, ok = parseToken.Claims.(*CustomClaims)

	return
}

func IsExpired(claims *CustomClaims) bool {
	return claims.ExpiresAt <= time.Now().Unix()
}

func Init(private, public []byte) (err error) {
	once.Do(func() {
		privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(private)
		if err != nil {
			return
		}
		publicKey, err = jwt.ParseRSAPublicKeyFromPEM(public)
		if err != nil {
			return
		}
	})
	return
}
