package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-openapi/loads"
	"github.com/jessevdk/go-flags"
	"github.com/sirupsen/logrus"
	"github.com/theoptz/url-shortener/internal/infrastructure/etcd"
	"github.com/theoptz/url-shortener/internal/infrastructure/mongodb"
	"github.com/theoptz/url-shortener/internal/server/restapi"
	"github.com/theoptz/url-shortener/internal/server/restapi/operations"
)

type App struct{}

func (a *App) connectToEtcd(addr string) error {
	_, err := etcd.Connect(addr)
	return err
}

func (a *App) connectToMongo(ctx context.Context, addr string) error {
	// connect to mongodb
	_, err := mongodb.Connect(ctx, addr, "url_shortener")
	return err
}

func (a *App) startRestAPIServer() (*restapi.Server, error) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return nil, err
	}

	api := operations.NewURLShortenerAPI(swaggerSpec)
	server := restapi.NewServer(api)

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "URL Shortener"
	parser.LongDescription = "Distributed URL Shortener"
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err = parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			return nil, err
		}
	}

	if _, err = parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}

		if code != 0 {
			return nil, err
		}
	}

	server.ConfigureAPI()

	return server, nil
}

func (a *App) Run() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// make default app context
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	signal.NotifyContext(ctx, syscall.SIGTERM, syscall.SIGINT)

	// connect to mongodb
	err := a.connectToMongo(ctx, os.Getenv("MONGO_DB_URI"))
	if err != nil {
		logrus.Fatalln(err)
	}

	// connect to etcd
	err = a.connectToEtcd(os.Getenv("ETCD_URI"))
	if err != nil {
		logrus.Fatalln(err)
	}

	server, err := a.startRestAPIServer()
	if err != nil {
		logrus.Fatalln(err)
	}
	defer server.Shutdown()

	if err = server.Serve(); err != nil {
		logrus.Fatalln(err)
	}
}

func NewApp() *App {
	return &App{}
}
