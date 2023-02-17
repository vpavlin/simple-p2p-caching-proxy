package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/vpavlin/simple-p2p-caching-proxy/internal/api"
	"github.com/vpavlin/simple-p2p-caching-proxy/internal/store/leveldb"
)

func main() {
	e := echo.New()

	db, err := leveldb.NewDB("./level.db")
	if err != nil {
		log.Fatal(err)
	}

	_ = api.NewAPI(e, db)
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	e.Start(fmt.Sprintf("localhost:%d", port))
}
