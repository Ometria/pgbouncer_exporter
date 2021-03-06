package cmd

import (
	"fmt"

	"github.com/jbub/pgbouncer_exporter/collector"
	"github.com/jbub/pgbouncer_exporter/config"
	"github.com/jbub/pgbouncer_exporter/server"
	"github.com/jbub/pgbouncer_exporter/store"

	"github.com/prometheus/common/log"
	"github.com/prometheus/common/version"
	"github.com/urfave/cli"
)

// Server is a cli command used for running exporter http server.
var Server = &cli.Command{
	Name:   "server",
	Usage:  "Starts exporter server.",
	Action: runServer,
}

func runServer(ctx *cli.Context) error {
	cfg := config.LoadFromCLI(ctx)
	st, err := store.NewSQLStore(cfg.DatabaseURL)
	if err != nil {
		return fmt.Errorf("unable to initialize store: %v", err)
	}
	defer st.Close()

	exp := collector.New(cfg, st)
	srv := server.New(cfg, exp)

	log.Infoln("Starting ", collector.Name, version.Info())
	log.Infoln("Server listening on", cfg.ListenAddress)
	log.Infoln("Metrics available at", cfg.TelemetryPath)
	log.Infoln("Build context", version.BuildContext())

	if err := srv.Run(); err != nil {
		return fmt.Errorf("unable to run server: %v", err)
	}
	return nil
}
