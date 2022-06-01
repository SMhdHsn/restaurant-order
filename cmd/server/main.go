package main

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/smhdhsn/restaurant-order/internal/config"
	"github.com/smhdhsn/restaurant-order/internal/db"
	"github.com/smhdhsn/restaurant-order/internal/model"
	"github.com/smhdhsn/restaurant-order/internal/repository/mysql"
	"github.com/smhdhsn/restaurant-order/internal/service"

	log "github.com/smhdhsn/restaurant-order/internal/logger"
	eipb "github.com/smhdhsn/restaurant-order/internal/protos/edible/inventory"
	remoteRepository "github.com/smhdhsn/restaurant-order/internal/repository/remote"
)

// ctx holds application's context.
var ctx context.Context

// init will be called when this package is imported.
func init() {
	ctx = context.Background()
}

// main is the application's kernel.
func main() {
	// read configurations.
	conf, err := config.LoadConf()
	if err != nil {
		log.Fatal(err)
	}

	// create a database connection.
	dbConn, err := db.Connect(&conf.DB)
	if err != nil {
		log.Fatal(err)
	}

	// initialize auto migration.
	if err := db.InitMigrations(dbConn); err != nil {
		log.Fatal(err)
	}

	// make connection with external services.
	eConn, err := grpc.Dial(
		conf.Services["edible"].Address,
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	// instantiate gRPC clients.
	eiClient := eipb.NewEdibleInventoryServiceClient(eConn)

	// instantiate models.
	oModel := new(model.Order)

	// instantiate repositories.
	oRepo := mysql.NewOrderRepository(dbConn, *oModel)
	eiRepo := remoteRepository.NewEdibleInventoryRepository(&ctx, eiClient)

	// instantiate services.
	osServ := service.NewOrderSubmitService(eiRepo, oRepo)

	_ = osServ
}
