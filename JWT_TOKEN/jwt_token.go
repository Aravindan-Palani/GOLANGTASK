package main

import (
        "github.com/golang-jwt/jwt"
        "time"
        "log"
        "fmt"
)

/* security Key */

var mySigningKey = []byte("test")

func main() {
        var jwtToken string
        var valid bool
        token, err := CreateToken("S497177")
        if err != nil{
                log.Println("Error in creating token",err)
                return
        }
        log.Println("TOKEN", token)
        jwtToken = token
        valid = ValidToken(jwtToken)
        if(!valid){
                log.Println("Not valid token")
                return
        }
        log.Println("valid token")
}

/* function to create Token */

func CreateToken(userid string) (string, error) {
        claims := jwt.MapClaims{}

        claims["authorized"] = true
        claims["user"] = userid
        claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
        token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

        tokenString, err := token.SignedString(mySigningKey)

        if err != nil {
                return "", err
        }

        return tokenString, nil
}

/* function to check the token is valid or not */

func ValidToken(tokenString string) (bool) {

        token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
                return []byte("test"), nil
        })

        if token.Valid {
                fmt.Println("Inside Valid Token")
                return true
        } else {
                fmt.Println("Inside inValid Token:", err)
                return false
        }
}

