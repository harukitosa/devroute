package main

import (
	"context"
	"devroute/ent"
	"devroute/ent/migrate"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// connect db and migration
	client, err := ent.Open("mysql", "root:root@tcp(localhost:3306)/devroute")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	err = client.Schema.Create(
		ctx,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true),
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.POST("/login", login)
	router.GET("/", topHandler)

	r := router.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))
	r.GET("", restricted)
	router.Start(":8080")
}

func restricted(c echo.Context) error {
	user, ok := c.Get("user").(*jwt.Token)
	if !ok {
		return fmt.Errorf("failed to parse")
	}
	claims, ok := user.Claims.(jwt.MapClaims)
	if !ok {
		return fmt.Errorf("failed to parse")
	}
	name := claims["name"].(string)
	return c.String(http.StatusOK, "Welcome "+name+"!")
}

func topHandler(c echo.Context) error {
	return c.String(http.StatusOK, "OK")
}

type UserDto struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func login(c echo.Context) error {
	user := UserDto{}
	err := c.Bind(&user)
	if err != nil {
		return err
	}

	if user.UserName != "taro" || user.Password != "shhh!" {
		return echo.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = "Taro"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

// curl -X POST -H "Content-Type: application/json" -d '{"user_name":"taro", "Password": "shhh!"}' localhost:8080/login
