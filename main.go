package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/dgrijalva/jwt-go"
	"strings"


)

type User struct {
	 Username string `json:"username"`
	 DOB string
	 Age int
	 Email string
	 Phonenumber string
}

var user = User{
	Username: "mayank",
	DOB:"10102020",
	Age:23,
	Email:"test@test.com",
	Phonenumber:"1234567890",
}


var mySigningKey =[]byte("thiscannotbebreak")

func main(){
	route:=gin.Default()
	route.POST("/auth",AuthCheck)
	route.GET("/user/profile",UserProfile)
	route.Run(":8080")
}


func AuthCheck(c *gin.Context){
	var u User
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	 if user.Username != u.Username {
		c.JSON(http.StatusUnauthorized, "Please provide valid login details")
		return
	}
	token, err := CreateToken(user.Username)
	if err != nil {
	   c.JSON(http.StatusUnprocessableEntity, err.Error())
	   return
	}
	c.JSON(http.StatusOK, token)

}

func CreateToken(username string) (string,error){
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_name"] = username
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	tokenString,err:= at.SignedString(mySigningKey)
	if err != nil {
		return "", err
	 }
	 return tokenString, nil
}

func ExtractToken(c *gin.Context) string{
	bearToken := c.GetHeader("Authorization")

	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
	   return strArr[1]
	}
	return ""

}

func UserProfile(c *gin.Context){
	token:=ExtractToken(c)
	if token!=""{
		c.JSON(200,gin.H{
			"DOB":"10102020",
			"age":23,
			"email":"test@test.com",
			"Phonenumber":"1234567890",
		})
		return 
	}else{
		c.JSON(200,gin.H{
			"message":"userseriv-name",
		})
		return
	}

}
