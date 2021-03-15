package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	h "github.com/neel229/linktree-clone/internal/api"
	mr "github.com/neel229/linktree-clone/internal/repository/mongo"
	"github.com/neel229/linktree-clone/internal/shortener"
	"github.com/neel229/linktree-clone/internal/utils"
)

func main() {
	repoConfig, err := utils.LoadConfig("../../config")
	if err != nil {
		log.Fatal(err)
	}
	repo := chooseRepo(repoConfig)
	service := shortener.NewRedirectService(repo)
	handler := h.NewHandler(service)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/{code}", handler.Get)
	r.Post("/", handler.Post)

	errs := make(chan error, 2)
	go func() {
		fmt.Printf("starting a server on port: %s\n", repoConfig.Port)
		errs <- http.ListenAndServe(":"+repoConfig.Port, r)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func chooseRepo(repoConfig utils.Config) shortener.RedirectRepository {
	switch repoConfig.DB {
	case "redis":
		redisURL := repoConfig.RURL
		//repo, err := rr.NewRedisRepository(redisURL)
		//if err != nil {
		//log.Fatal(err)
		//}
		//return repo
		log.Print(redisURL)
	case "mongo":
		mongoURL := repoConfig.MURL
		mongodb := "linktree"
		mongoTimeout := int64(30)
		repo, err := mr.NewMongoRepo(mongoURL, mongodb, mongoTimeout)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
