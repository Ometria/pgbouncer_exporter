package cmd

import (
	"fmt"

	"github.com/jbub/pgbouncer_exporter/config"
	"github.com/jbub/pgbouncer_exporter/store"

	"github.com/urfave/cli"
)

// Health is a cli command used for checking the health of the system.
var Health = &cli.Command{
	Name:   "health",
	Usage:  "Checks the health of the system.",
	Action: checkHealth,
}

func checkHealth(ctx *cli.Context) error {
	cfg := config.LoadFromCLI(ctx)
	st, err := store.NewSQLStore(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("unable to initialize store: %v", err)
	}
	defer st.Close()

	if err := st.Check(); err != nil {
		return fmt.Errorf("store health check failed: %v", err)
	}
	return nil
}
