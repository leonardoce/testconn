package main

import (
	"fmt"
	"os"

	"github.com/leonardoce/testconn/cmd/testconn"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := testconn.Run(); err != nil {
		fmt.Printf("error: %+v", err)
		os.Exit(1)
	}
}
