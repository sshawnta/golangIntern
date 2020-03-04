package main


import (
	"flag"
	"fmt"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	fasthttpprometheus "github.com/flf2ko/fasthttp-prometheus"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/kelseyhightower/envconfig"
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"

	"github.com/sshawnta/TestTransport/pkg/service"
	"github.com/sshawnta/TestTransport/pkg/service/httpserver"

)

const (
	serviceVersion = "dev"
)

type configuration struct {
	Port               string `envconfig:"PORT" required:"true" default:"8080"`
	MaxRequestBodySize int    `envconfig:"MAX_REQUEST_BODY_SIZE" default:"10485760"` // 10 MB
	Debug              bool   `envconfig:"DEBUG" default:"false"`

	ReadTimeout time.Duration `envconnfig:"READ_TIMEOUT" default:"1s"`

	MetricsNamespace    string `envconfig:"METRICS_NAMESPACE" default:"wb"`
	MetricsSubsystem    string `envconfig:"METRICS_SUBSYSTEM" default:"barcode_service"`
	MetricsNameCount    string `envconfig:"METRICS_NAME_COUNT" default:"request_count"`
	MetricsNameDuration string `envconfig:"METRICS_NAME_DURATION" default:"request_duration"`
	MetricsHelpCount    string `envconfig:"METRICS_HELP_COUNT" default:"Request count"`
	MetricsHelpDuration string `envconfig:"METRICS_HELP_DURATION" default:"Request duration"`
}

func main() {
	printVersion := flag.Bool("version", false, "print version and exit")
	flag.Parse()

	if *printVersion {
		fmt.Println(serviceVersion)
		os.Exit(0)
	}

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC, "caller", log.DefaultCaller)
	_ = level.Info(logger).Log("msg", "initializing", "version", serviceVersion)

	var cfg configuration
	if err := envconfig.Process("", &cfg); err != nil {
		_ = level.Error(logger).Log("msg", "failed to load configuration", "err", err)
		os.Exit(1)
	}
	if !cfg.Debug {
		logger = level.NewFilter(logger, level.AllowInfo())
	}

	svc := service.NewService()

	router := httpserver.NewPreparedServer(svc)
	router.Handle("GET", "/debug/pprof/", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Index))
	router.Handle("GET", "/debug/pprof/profile", fasthttpadaptor.NewFastHTTPHandlerFunc(pprof.Profile))
	p := fasthttpprometheus.NewPrometheus(cfg.MetricsSubsystem)
	fasthttpServer := &fasthttp.Server{
		Handler:            p.WrapHandler(router),
		MaxRequestBodySize: cfg.MaxRequestBodySize,
		ReadTimeout:        cfg.ReadTimeout,
	}

	go func() {
		_ = level.Info(logger).Log("msg", "starting http server", "port", cfg.Port)
		if err := fasthttpServer.ListenAndServe(":" + cfg.Port); err != nil {
			_ = level.Error(logger).Log("msg", "server run failure", "err", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {
		_ = level.Info(logger).Log("msg", "received signal, exiting", "signal", sig)

		if err := fasthttpServer.Shutdown(); err != nil {
			_ = level.Error(logger).Log("msg", "server shutdown failure", "err", err)
		}

		_ = level.Info(logger).Log("msg", "goodbye")
	}(<-c)
}