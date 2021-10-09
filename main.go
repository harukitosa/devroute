package main

import (
	"context"
	"devroute/ent"
	"devroute/ent/migrate"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func initDB() (*ent.Client, error) {
	var (
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME")
		dbName                 = mustGetenv("DB_NAME")
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	dns := fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)
	db, err := ent.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func initLocalDB() (*ent.Client, error) {
	dns := "root:root@tcp(127.0.0.1:3306)/devroute?charset=utf8mb4&parseTime=True&loc=Local"
	client, err := ent.Open("mysql", dns)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	v := os.Getenv("RUNENV")

	var client *ent.Client
	var err error
	if v == "production" {
		client, err = initDB()
	} else {
		client, err = initLocalDB()
	}

	// connect db and migration
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
