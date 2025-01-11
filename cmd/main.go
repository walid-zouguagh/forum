package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"forum/models"
	"forum/router"
	"forum/utils"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	/* port handling */
	port := ":8001" // use env variable

	/* init database tables */
	var err error
	utils.DB, err = sql.Open("sqlite3", "./database.sqlite")
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer utils.DB.Close()
	models.InitTables(utils.DB)

	/* --------------------------- main mux --------------------------- */
	mainMux := http.NewServeMux()
	mainMux.Handle("/", router.WebRouter())
	mainMux.Handle("/api/", router.APIRouter())

	/* --------------------------- run server --------------------------- */
	fmt.Printf("server running on http://localhost%s\n", port)
	server_err := http.ListenAndServe(port, mainMux)
	if server_err != nil {
		fmt.Printf("server runnig error! %v", server_err.Error())
		os.Exit(1)
	}
}
