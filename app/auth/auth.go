package auth

import (
	"crypto/ed25519"
	"encoding/hex"
	"log"
	"os"
	"teamhub-backend/app/model"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/pascaldekloe/jwt"
)

func GenerateToken(user model.User) ([]byte, error) {
	var claims jwt.Claims
	keyString := os.Getenv("PRIV_KEY")
	hkey, err := hex.DecodeString(keyString)
	if err != nil {
		return nil, err
	}
	key := ed25519.PrivateKey(hkey)
	claims.Subject = user.Username
	claims.Issued = jwt.NewNumericTime(time.Now().Round(time.Second))
	claims.Set = map[string]interface{}{"full_name": user.Name, "uuid": user.UUID}
	token, err := claims.EdDSASign(key)
	if err != nil {
		return nil, err
	}
	log.Println("token created for ", user.Username)
	return token, nil
}

func VerifyToken(token []byte) (bool, string) {
	keyString := os.Getenv("PUB_KEY")
	hkey, err := hex.DecodeString(keyString)
	if err != nil {
		log.Println("key decode failed ", err)
		return false, ""
	}
	key := ed25519.PublicKey(hkey)
	claims, err := jwt.EdDSACheck(token, key)
	if err != nil {
		log.Println("credentials denied with ", err)
		return false, ""
	}
	if !claims.Valid(time.Now()) {
		log.Println("credential time constraints exceeded")
		return false, ""
	}
	log.Println("token verified for ", claims.Subject)
	return true, claims.Subject
}
