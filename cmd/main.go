package main

import (
	"github.com/Richieid23/simple-forum/internal/configs"
	"github.com/Richieid23/simple-forum/internal/handlers/memberships"
	"github.com/Richieid23/simple-forum/pkg/internalsql"
	"github.com/gin-gonic/gin"
	"log"

	membershipRepository "github.com/Richieid23/simple-forum/internal/repositories/memberships"
	membershipService "github.com/Richieid23/simple-forum/internal/services/memberships"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.Get()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	membershipRepos := membershipRepository.NewRepository(db)
	membershipSvc := membershipService.NewService(membershipRepos)

	membershipHandler := memberships.NewHandler(r, membershipSvc)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
