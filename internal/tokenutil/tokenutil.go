package tokenutil

import (
	"fmt"
	"forum/internal/domain"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	Name string `json:"name"`
	Email   string   `json:"email"`
	jwt.RegisteredClaims
}

func CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {

	//exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &JwtClaims{
		Name: user.Username,
		Email:   user.Email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

// func CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
// 	claimsRefresh := &domain.JwtRefreshClaims{
// 		ID: user.ID,
// 		RegisteredClaims: jwt.RegisteredClaims{
// 			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expiry))),
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
// 	rt, err := token.SignedString([]byte(secret))
// 	if err != nil {
// 		return "", err
// 	}

// 	return rt, err
// }


func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error){
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		} 
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}


func ExtractIDFromToken(requestToken string, secret string) (string, error) {
    token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(secret), nil
    })
    if err != nil {
        return "", err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok || !token.Valid {
        return "", fmt.Errorf("Invalid token")
    }

	// fmt.Printf("t1: %T\n", claims["id"])
    email, ok := claims["email"].(string)
	// fmt.Println("iddddd", idFloat)
    if !ok {
        return "", fmt.Errorf("ID is not a valid number")
    }

    // Convert the float64 ID to uint
    // userID := uint(idFloat)
    return email, nil
}

// func  ValidateRefreshToken(refreshToken string, secret string) (uint, error) {
//     token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
//         if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//             return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
//         }
//         return []byte(secret), nil
//     })
//     if err != nil {
//         return 0, err
//     }

//     claims, ok := token.Claims.(jwt.MapClaims)
//     if !ok || !token.Valid {
//         return 0, fmt.Errorf("Invalid refresh token")
//     }

//     idFloat, ok := claims["id"].(float64)
//     if !ok {
//         return 0, fmt.Errorf("User ID is not a valid number in the refresh token")
//     }

//     userID := uint(idFloat)
//     return userID, nil
// }
