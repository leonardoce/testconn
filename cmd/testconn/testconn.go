// Package testconn implements the main command
package testconn

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/pkg/errors"
)

// pingTimeout is used as a timeout when pinging the database
const pingTimeout = 1 * time.Second

// db is the connection pool used by this application
var db *sql.DB

// Run implements the "testconn" command
func Run() error {
	listenAddressDefault := os.Getenv("LISTEN_ADDRESSES")
	if listenAddressDefault == "" {
		listenAddressDefault = ":8000"
	}
	dsnDefault := os.Getenv("DSN")
	if dsnDefault == "" {
		dsnDefault = "dbname=postgres"
	}

	listenAddresses := flag.String(
		"listenAddresses",
		listenAddressDefault,
		"the IP addresses where we should listen to",
	)
	dsn := flag.String(
		"dsn",
		dsnDefault,
		"the DSN where we should connect to",
	)
	flag.Parse()

	var err error
	if db, err = sql.Open("pgx", *dsn); err != nil {
		return errors.Wrap(err, "creating the connection pool")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/readyz", readyzHandler)
	mux.HandleFunc("/livez", livezHandler)
	mux.HandleFunc("/ping", pingHandler)

	server := &http.Server{
		Addr:              *listenAddresses,
		Handler:           NewLoggingDecorator(mux),
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	log.Printf("Starting web server: %s", *listenAddresses)
	return errors.Wrapf(
		server.ListenAndServe(),
		"listening to %s",
		*listenAddresses,
	)
}

// readyzHandler handles the readiness probe for this application
func readyzHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), pingTimeout)
	defer cancel()

	// In this example, we declare the application ready when
	// it can connect to the database
	if err := db.PingContext(ctx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	fmt.Fprintf(w, "OK")
}

// livezHandler handles the liveness probe for this application
func livezHandler(w http.ResponseWriter, r *http.Request) {
	// We simply return ok when we're able to listen for an
	// HTTP request. This is reasonable: if we are not
	// able to that, we need to be restarted
	fmt.Fprintf(w, "OK")
}

// pingHandler is the only entrypoint of this go application
func pingHandler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), pingTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	fmt.Fprintf(w, "OK")
}
