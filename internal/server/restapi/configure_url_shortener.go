// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/sirupsen/logrus"
	"github.com/theoptz/url-shortener/internal/infrastructure/etcd"
	"github.com/theoptz/url-shortener/internal/infrastructure/mongodb"
	"github.com/theoptz/url-shortener/internal/repositories"
	"github.com/theoptz/url-shortener/internal/server/handlers"
	"github.com/theoptz/url-shortener/internal/server/restapi/operations"
	"github.com/theoptz/url-shortener/internal/services/b62"
	"github.com/theoptz/url-shortener/internal/services/counter"
	"github.com/theoptz/url-shortener/internal/services/links"
	"github.com/theoptz/url-shortener/internal/services/validators"
	"github.com/theoptz/url-shortener/internal/utils"
	"github.com/urfave/negroni"
)

//go:generate swagger generate server --target ../../server --name URLShortener --spec ../../../api/api.yaml --principal interface{} --exclude-main

func configureFlags(api *operations.URLShortenerAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.URLShortenerAPI) http.Handler {
	nodeName := os.Getenv("name")

	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.Logger = logrus.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	coder := b62.NewCode()
	database := mongodb.GetDatabase()
	etcdClient := etcd.GetClient()

	linksRepository := repositories.NewLinksRepository(database)
	counterRepository := repositories.NewCounterRepository(database)
	rangeRepository := repositories.NewRangRepo(etcdClient, nodeName)

	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()

	item, err := rangeRepository.Get(ctx)
	if err != nil {
		logrus.Fatalln(err)
	}

	startIndex, err := counterRepository.Get(ctx, item)
	if err != nil {
		logrus.Fatalln(err)
	}

	v, _ := json.Marshal(item)
	logrus.Info("starting node", nodeName, utils.GetString(v))

	counterService := counter.NewCounter(rangeRepository, startIndex, *item)
	linksService := links.NewLinks(linksRepository, counterService, coder)
	urlValidator := validators.NewURLValidator()

	api.GetLinksCodeHandler = handlers.NewGetLinksHandler(linksService)
	api.PostLinksHandler = handlers.NewPostLinksHandler(linksService, urlValidator)

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return addLogging(handler)
}

func addLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now().UTC()

		lrw := negroni.NewResponseWriter(w)

		next.ServeHTTP(lrw, r)

		duration := time.Since(startTime)
		logrus.WithFields(logrus.Fields{
			"time":         startTime.Format("2006-01-02T15:04:05Z"),
			"ip":           strings.Split(r.RemoteAddr, ":")[0],
			"host":         r.Host,
			"method":       r.Method,
			"path":         r.URL.Path,
			"status":       lrw.Status(),
			"resp_time_ms": duration.Milliseconds(),
		}).Info("HTTP")
	})
}
