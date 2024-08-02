package main

import (
	"database/sql"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	"github.com/mori-dv/RSS/internal/database"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {

	fmt.Println("Start Program...")
	baseDir := os.Getenv("BASE_DIR")
	if baseDir == "" {
		log.Fatal("BASE DIR variable not assigned!")
	}
	fmt.Println(baseDir)
	err := godotenv.Load(filepath.Join(baseDir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("database url environment variable not set")
	}

	connection, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}

	db := database.New(connection)
	apicfg := apiConfig{
		DB: db,
	}
	go startScraping(db, 10, time.Minute)

	router := chi.NewRouter()
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)

	v1Router.Get("/error", handlerError)

	v1Router.Get("/users", apicfg.handlerGetAllUsers)
	v1Router.Post("/user/new", apicfg.handlerCreateUser)
	v1Router.Get("/user/get", apicfg.middlewareAuth(apicfg.handlerGetUser))

	v1Router.Post("/feed/new", apicfg.middlewareAuth(apicfg.handlerCreateFeed))
	v1Router.Get("/feeds", apicfg.handlerGetFeeds)

	v1Router.Post("/feed/follow/new", apicfg.middlewareAuth(apicfg.handlerCreateFeedFollows))
	v1Router.Get("/feed/follows", apicfg.middlewareAuth(apicfg.handlerGetAllFeedFollows))
	v1Router.Delete("/feed/follow/delete/{feedFollowsID}", apicfg.middlewareAuth(apicfg.handlerDeleteFeedFollows))
	v1Router.Get("/posts", apicfg.middlewareAuth(apicfg.handlerGetPostsForUser))
	v1Router.Post("/webhook", webhookHandler)

	router.Mount("/v1", v1Router)

	server := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}
	log.Printf("Server is starting on port %s\n", portString)

	httpErr := server.ListenAndServe()
	if httpErr != nil {
		log.Fatal(httpErr)
	}
	closeErr := server.Close()

	if closeErr != nil {
		log.Fatal(closeErr)
	}
}
