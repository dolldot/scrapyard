package clients

import (
	"fmt"
	"os"

	"github.com/dolldot/scrapyard/bookmarkd/config"
	"github.com/dolldot/scrapyard/bookmarkd/generated/ent"
	"github.com/dolldot/scrapyard/bookmarkd/internal/logger"
	_ "github.com/lib/pq"
)

func NewEntclient() *ent.Client {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
		config.Load.DB.Host,
		config.Load.DB.Port,
		config.Load.DB.User,
		config.Load.DB.Name,
		config.Load.DB.Password,
	)

	client, err := ent.Open(config.Load.DB.Type, connectionString)
	if err != nil {
		logger.Errorf("failed to open db: %v", err)
		os.Exit(1)
	}

	return client
}
