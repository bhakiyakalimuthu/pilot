package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net/http"
	"strings"
	"github.com/mattn/go-colorable"
	"github.com/bhakiyakalimuthu/pilot/internal/app"
)

func main()  {

	r := chi.NewRouter()
	l := loggerSetup()
	// store:= app.NewMemoryStore(l)
	store := app.NewPostgresStore(l,postgresSetup())
	svc := app.NewService(l,store)
	_ = app.NewController(l,svc).SetupRouter(r)
	http.ListenAndServe(":8080",r)
}

func postgresSetup() *sqlx.DB {
	connStrParts := []string{
		fmt.Sprintf("host=%s", "localhost"),
		fmt.Sprintf("port=%d", 5432),
		fmt.Sprintf("user=%s", "dbuser"),
		fmt.Sprintf("password=%s", "dbpassword"),
		fmt.Sprintf("dbname=%s", "users"),
	}
	connector, _ := pq.NewConnector(strings.Join(connStrParts, " "))
	driver := sql.OpenDB(connector)
	return sqlx.NewDb(driver, "postgres")
}
func loggerSetup() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create zap logger : %v", err)
	}
	return logger
	aa := zap.NewDevelopmentEncoderConfig()
	aa.EncodeLevel = zapcore.CapitalColorLevelEncoder
	bb := zap.New(zapcore.NewCore(
		zapcore.NewConsoleEncoder(aa),
		zapcore.AddSync(colorable.NewColorableStdout()),
		zapcore.DebugLevel,
	))
	bb.Warn("logger setup done")
	return bb
}

// main - > controller
// GetUser (/GetUser)
// CreateUser
// UpdateProfile
// DeleteProfile