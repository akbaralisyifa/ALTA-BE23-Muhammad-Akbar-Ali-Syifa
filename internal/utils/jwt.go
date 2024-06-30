package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// buat fungsi untuk generete token nya
func GenereteToken(userID uint)(string, error){ // yang di kembalikan string acak nya
	var claims = jwt.MapClaims{}; // untuk menyimpan data2 yang akan di save di jwt
	
	claims["id"] = userID;	// untuk menyimpan id nya
	claims["iat"] = time.Now().Unix(); // waktu ketika token di buat
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()	// waktu expired token nya

	var process = jwt.NewWithClaims(jwt.SigningMethodHS256, claims);

	result, err := process.SignedString([]byte("passkeyJWT")); // signed -> set string key nya

	if err != nil {
		return "", err
	};

	return result, nil;
}

// membuat fungsi jwt agar lebih optimal
func DecodeToken(token *jwt.Token) uint {
	var claims = token.Claims.(jwt.MapClaims);
	var userID = claims["id"].(float64);
	return	uint(userID)
}