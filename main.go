package main

import (
	"context"
	"devroute/ent"
	"devroute/ent/migrate"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"google.golang.org/api/option"
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

var app *firebase.App

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

	opt := option.WithCredentialsFile("./devroute-firebase-adminsdk.json")
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	router := echo.New()
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	router.GET("/", topHandler)
	router.GET("/private", echoAuthMiddleware(privateHandler))
	router.Start(":8080")
}

func topHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK")
}

func privateHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "OK private")
}

func echoAuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Println("before action")
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
		}
		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			fmt.Printf("error verifying ID token: %v\n", err)
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		}
		log.Println(token)
		if err := next(c); err != nil {
			c.Error(err)
		}
		log.Println("after action")
		return nil
	}
}
