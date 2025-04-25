package test

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/require"
	"github.com/viniosilva/go-boilerplateapi/config"
	"github.com/viniosilva/go-boilerplateapi/internal/application/customer/dto"
	"github.com/viniosilva/go-boilerplateapi/internal/container"
	"github.com/viniosilva/go-boilerplateapi/internal/infrastructure/api"
	"github.com/viniosilva/go-boilerplateapi/internal/infrastructure/db"
	"github.com/viniosilva/go-boilerplateapi/pkg/pagination"
	"github.com/viniosilva/go-boilerplateapi/test/seed"
	"gorm.io/gorm"
)

var (
	server  *http.Server
	dbConn  *gorm.DB
	client  *resty.Client
	baseUrl string
)

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	teardown()
	os.Exit(code)
}

func setup() {
	os.Setenv("IB_ENV", "test")
	cfg, err := config.LoadConfig(config.WithPath("../"))
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	dbConn, err = db.NewGorm(cfg.DB.Host, cfg.DB.Port, cfg.DB.DBName,
		cfg.DB.User, cfg.DB.Password, cfg.DB.SslMode,
	)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	di := container.New(dbConn)
	server = api.NewServer(di, cfg.App.Host, cfg.App.Port, time.Second*time.Duration(cfg.App.TimeoutSec))
	baseUrl = fmt.Sprintf("http://%s/api", server.Addr)

	go func() {
		_ = server.ListenAndServe()
	}()

	client = resty.New()
	client.SetTimeout(10 * time.Second)
}

func teardown() {
	err := dbConn.Exec("TRUNCATE TABLE customers RESTART IDENTITY CASCADE").Error
	if err != nil {
		log.Fatalf("failed to clean database: %v", err)
	}

	db.Close(dbConn)
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Printf("failed to shutdown server: %v\n", err)
	}
	client.GetClient().CloseIdleConnections()
}

func TestE2E(t *testing.T) {
	var customerCurrent *dto.Customer

	t.Run("POST /api/customers", func(t *testing.T) {
		payload := seed.CreateCustomer

		res, err := client.R().
			SetBody(payload).
			SetResult(&customerCurrent).
			Post(fmt.Sprintf("%s/customers", baseUrl))

		require.NoError(t, err)
		require.NotNil(t, customerCurrent)

		require.Equal(t, http.StatusCreated, res.StatusCode())
	})

	t.Run("GET /api/customers", func(t *testing.T) {
		var body pagination.Pagination[dto.Customer]
		res, err := client.R().
			SetResult(&body).
			Get(fmt.Sprintf("%s/customers", baseUrl))

		require.NoError(t, err)

		require.Equal(t, http.StatusOK, res.StatusCode())
		require.Len(t, body.Data, 1)
	})
}
