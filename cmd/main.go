package main

import (
	"database/sql"
	"net/http"
	"time"

	graphql2 "github.com/nugrohoac/livestream/delivery/graphql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"

	"github.com/nugrohoac/livestream/delivery/graphql/mutation"
	"github.com/nugrohoac/livestream/delivery/graphql/schema"
	"github.com/nugrohoac/livestream/pkg/handler"
	"github.com/nugrohoac/livestream/repository/mysql"
	"github.com/nugrohoac/livestream/service"
)

func main() {
	// Init mysql
	mysqlDsn := "root1:password123@tcp(127.0.0.1:3307)/livestream_test?parseTime=true&loc=Asia%2FJakarta&charset=utf8mb4&collation=utf8mb4_unicode_ci"
	mysqlDB, err := sql.Open("mysql", mysqlDsn)
	if err != nil {
		logrus.Fatal("FAILED CONNECT TO MYSQL", err.Error())
	}

	if mysqlDB != nil {
		mysqlDB.SetConnMaxLifetime(time.Duration(5) * time.Second)
		mysqlDB.SetMaxIdleConns(3)
		mysqlDB.SetConnMaxLifetime(5)
	}

	livestreamMysqlrepo := mysql.NewLiveStreamMysql(mysqlDB)
	livestreamService := service.NewLivestreamService(livestreamMysqlrepo)

	livestreamMutation := mutation.NewLivestreamMutation(livestreamService)
	resolver := graphql2.NewRootResolver(livestreamMutation)

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.POST, echo.OPTIONS},
	}))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK!")
	})

	graphqlSchema := graphql.MustParseSchema(schema.String(), resolver)

	e.POST("/live-stream/graphql", handler.GraphQLHandler(&relay.Handler{Schema: graphqlSchema}))

	e.Logger.Fatal(e.Start(":8090"))
}
