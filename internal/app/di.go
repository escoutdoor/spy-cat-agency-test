package app

import (
	"context"

	"github.com/escoutdoor/spy-cat-agency-test/internal/client"
	"github.com/escoutdoor/spy-cat-agency-test/internal/config"
	"github.com/escoutdoor/spy-cat-agency-test/internal/repository"
	cat_repository "github.com/escoutdoor/spy-cat-agency-test/internal/repository/cat"
	mission_repository "github.com/escoutdoor/spy-cat-agency-test/internal/repository/mission"
	target_repository "github.com/escoutdoor/spy-cat-agency-test/internal/repository/target"
	"github.com/escoutdoor/spy-cat-agency-test/internal/service"
	cat_service "github.com/escoutdoor/spy-cat-agency-test/internal/service/cat"
	mission_service "github.com/escoutdoor/spy-cat-agency-test/internal/service/mission"
	target_service "github.com/escoutdoor/spy-cat-agency-test/internal/service/target"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/closer"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/database"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/database/pg"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/database/txmanager"
	"github.com/escoutdoor/spy-cat-agency-test/pkg/logger"
)

type diContainer struct {
	dbClient  database.Client
	txManager database.TxManager

	catApiClient client.CatClient

	catService     service.CatService
	missionService service.MissionService
	targetService  service.TargetService

	catRepository     repository.CatRepisitory
	missionRepository repository.MissionRepository
	targetRepository  repository.TargetRepository
}

func newDiContainer() *diContainer {
	return &diContainer{}
}

func (d *diContainer) DBClient(ctx context.Context) database.Client {
	if d.dbClient == nil {
		client, err := pg.NewClient(ctx, config.Config().Postgres.Dsn())
		if err != nil {
			logger.Fatal(ctx, "new database client", err)
		}

		if err := client.DB().Ping(ctx); err != nil {
			logger.Fatal(ctx, "ping database: %s", err)
		}

		d.dbClient = client
		closer.Add(func(ctx context.Context) error {
			client.Close()
			return nil
		})
	}

	return d.dbClient
}

func (d *diContainer) TxManager(ctx context.Context) database.TxManager {
	if d.txManager == nil {
		d.txManager = txmanager.NewTransactionManager(d.DBClient(ctx).DB())
	}

	return d.txManager
}

func (d *diContainer) CatService(ctx context.Context) service.CatService {
	if d.catService == nil {
		d.catService = cat_service.New(d.CatRepisitory(ctx), d.CatApiClient())
	}

	return d.catService
}

func (d *diContainer) TargetService(ctx context.Context) service.TargetService {
	if d.targetService == nil {
		d.targetService = target_service.New(
			d.MissionRepository(ctx),
			d.TargetRepository(ctx),
			d.TxManager(ctx),
		)
	}

	return d.targetService
}

func (d *diContainer) MissionService(ctx context.Context) service.MissionService {
	if d.missionService == nil {
		d.missionService = mission_service.New(
			d.MissionRepository(ctx),
			d.TargetRepository(ctx),
			d.TxManager(ctx),
		)
	}

	return d.missionService
}

func (d *diContainer) CatRepisitory(ctx context.Context) repository.CatRepisitory {
	if d.catRepository == nil {
		d.catRepository = cat_repository.New(d.DBClient(ctx))
	}

	return d.catRepository
}

func (d *diContainer) MissionRepository(ctx context.Context) repository.MissionRepository {
	if d.missionRepository == nil {
		d.missionRepository = mission_repository.New(d.DBClient(ctx))
	}

	return d.missionRepository
}

func (d *diContainer) TargetRepository(ctx context.Context) repository.TargetRepository {
	if d.targetRepository == nil {
		d.targetRepository = target_repository.New(d.DBClient(ctx))
	}

	return d.targetRepository
}

func (d *diContainer) CatApiClient() client.CatClient {
	if d.catApiClient == nil {
		d.catApiClient = client.New(config.Config().CatClient.ApiKey())
	}

	return d.catApiClient
}
