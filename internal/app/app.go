package app

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/escoutdoor/spy-cat-agency-test/internal/config"
	catv1 "github.com/escoutdoor/spy-cat-agency-test/internal/controller/cat/v1"
	missionv1 "github.com/escoutdoor/spy-cat-agency-test/internal/controller/mission/v1"
	targetv1 "github.com/escoutdoor/spy-cat-agency-test/internal/controller/target/v1"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/closer"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/errwrap"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/logger"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	logger_middleware "github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

type App struct {
	di         *diContainer
	httpServer *fiber.App
}

func New(ctx context.Context) (*App, error) {
	app := &App{di: newDiContainer()}
	if err := app.initDeps(ctx); err != nil {
		return nil, err
	}

	if err := goose.SetDialect(string(goose.DialectPostgres)); err != nil {
		return nil, errwrap.Wrap("set migrations dialect", err)
	}

	db := stdlib.OpenDBFromPool(app.di.DBClient(ctx).DB().Pool())
	if err := goose.UpContext(ctx, db, config.Config().Postgres.MigrationsDir()); err != nil {
		return nil, errwrap.Wrap("migrate up", err)
	}

	if err := db.Close(); err != nil {
		return nil, errwrap.Wrap("close db after migrate up", err)
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	go func() {
		logger.Info(ctx, "http server is running")
		if err := a.runHttpServer(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Fatal(ctx, "run http server", err)
		}
	}()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(ctx context.Context) error{
		a.initHttpServer,
	}

	for _, d := range deps {
		if err := d(ctx); err != nil {
			return err
		}
	}

	return nil
}

type structValidator struct {
	validate *validator.Validate
}

func (v *structValidator) Validate(out any) error {
	return v.validate.Struct(out)
}

func (a *App) initHttpServer(ctx context.Context) error {
	app := fiber.New(fiber.Config{
		AppName:         config.Config().App.Name(),
		ReadTimeout:     time.Second * 5,
		StructValidator: &structValidator{validate: validator.New()},
		ErrorHandler: func(c fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			msg := "internal server error"

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
				msg = e.Message
			}

			return c.Status(code).JSON(fiber.Map{
				"error": msg,
			})
		},
	})

	app.Use(logger_middleware.New())

	catv1.Register(app, a.di.CatService(ctx))
	missionv1.Register(app, a.di.MissionService(ctx))
	targetv1.Register(app, a.di.TargetService(ctx))

	a.httpServer = app
	closer.Add(func(ctx context.Context) error {
		return app.ShutdownWithContext(ctx)
	})

	return nil
}

func (a *App) runHttpServer() error {
	if err := a.httpServer.Listen(config.Config().HttpServer.Address()); err != nil {
		return errwrap.Wrap("http server listen", err)
	}

	return nil
}
