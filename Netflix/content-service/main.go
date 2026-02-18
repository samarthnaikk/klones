package main

import (
	"log"
	"net/http"

	"content-service/availability"
	"content-service/cache"
	"content-service/cast"
	"content-service/config"
	"content-service/database"
	"content-service/episode"
	"content-service/genre"
	"content-service/indexing"
	"content-service/localization"
	"content-service/metadata"
	"content-service/movie"
	"content-service/router"
	"content-service/season"
	"content-service/series"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	redisClient := cache.NewClient(cfg)
	defer redisClient.Close()

	r := router.New()

	movieHandler := movie.NewHandler(movie.NewService(movie.NewRepository(db)))
	movieHandler.RegisterRoutes(r)

	seriesHandler := series.NewHandler(series.NewService(series.NewRepository(db)))
	seriesHandler.RegisterRoutes(r)

	seasonHandler := season.NewHandler(season.NewService(season.NewRepository(db)))
	seasonHandler.RegisterRoutes(r)

	episodeHandler := episode.NewHandler(episode.NewService(episode.NewRepository(db)))
	episodeHandler.RegisterRoutes(r)

	genreHandler := genre.NewHandler(genre.NewService(genre.NewRepository(db)))
	genreHandler.RegisterRoutes(r)

	castHandler := cast.NewHandler(cast.NewService(cast.NewRepository(db)))
	castHandler.RegisterRoutes(r)

	metadataHandler := metadata.NewHandler(metadata.NewService(metadata.NewRepository(db)))
	metadataHandler.RegisterRoutes(r)

	availabilityHandler := availability.NewHandler(availability.NewService(availability.NewRepository(db)))
	availabilityHandler.RegisterRoutes(r)

	localizationHandler := localization.NewHandler(localization.NewService(localization.NewRepository(db)))
	localizationHandler.RegisterRoutes(r)

	indexingHandler := indexing.NewHandler(indexing.NewService(indexing.NewRepository(db)))
	indexingHandler.RegisterRoutes(r)

	log.Printf("Starting server on port %s", cfg.ServerPort)
	if err := http.ListenAndServe(":"+cfg.ServerPort, r); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
