package main

import (
	"context"
	"flag"
	"fmt"
	"grpc-ecosystem-template/api"
	"grpc-ecosystem-template/config"
	"grpc-ecosystem-template/gateway"
	"grpc-ecosystem-template/internal/model"
	"grpc-ecosystem-template/service"
	"grpc-ecosystem-template/version"
	"net"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"github.com/xiaomLee/go-plugin/db"
	"google.golang.org/grpc"
)

//go:generate make proto

var (
	port       = flag.Int("Port", 10050, "grpc listen port, http port +1")
	configPath = flag.String("configPath", "./configs/config.yaml", "config path")
	debug      = flag.Bool("debug", false, "enable debug mode")
	pprof      = flag.Bool("pprof", false, "enable pprof on 6060")
	httpPrefix = flag.String("httpPrefix", "", "http prefix path. if set, url will be like: /{prefix}/api/v1/status")
	v          = flag.Bool("version", false, "version")
)

func main() {
	flag.Parse()
	if *v {
		version.PrintFullVersionInfo()
		return
	}

	if *debug || *pprof {
		startPprof()
	}

	if err := Init(); err != nil {
		logrus.Fatalf("init server err:%s", err)
	}

	// register shutdown signal
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)

	// run
	var (
		err     error
		grpcSrv *grpc.Server
		httpSrv *http.Server
	)
	if grpcSrv, err = startGrpc(fmt.Sprintf(":%d", *port)); err != nil {
		logrus.Fatal(err)
	}
	if httpSrv, err = startHttp(fmt.Sprintf("127.0.0.1:%d", *port), fmt.Sprintf(":%d", *port+1)); err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("start grpc service listen on %d, http on %d \n", *port, *port+1)

	// wait signal
	<-shutdown

	// stop and release resource
	grpcSrv.Stop()
	httpSrv.Shutdown(context.Background())
	db.ReleaseDBPool()
}

func startGrpc(addr string) (*grpc.Server, error) {
	srv := grpc.NewServer(
		grpc.MaxRecvMsgSize(1e7),
		grpc.MaxSendMsgSize(1e7),
		grpc.ChainUnaryInterceptor(
			grpc_recovery.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
		))

	api.RegisterUserServiceServer(srv, &service.UserService{})

	// listen
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	go func() {
		if err := srv.Serve(lis); err != nil {
			logrus.Errorf("grpc err:%s", err)
			return
		}
	}()
	return srv, nil
}

func startHttp(endpoint, addr string) (*http.Server, error) {
	mu := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(gateway.IncomingMatcher),
		runtime.WithOutgoingHeaderMatcher(gateway.OutgoingMatcher),
		runtime.WithErrorHandler(gateway.CustomHTTPError),
		//runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		//	MarshalOptions:   protojson.MarshalOptions{},
		//	UnmarshalOptions: protojson.UnmarshalOptions{},
		//},
		//),
	)
	api.RegisterUserServiceHandlerFromEndpoint(context.Background(), mu, endpoint, []grpc.DialOption{grpc.WithInsecure()})

	router := mux.NewRouter()
	if *httpPrefix != "" {
		router.PathPrefix(*httpPrefix).Subrouter().NewRoute().Handler(http.StripPrefix(*httpPrefix, mu))
	} else {
		router.NewRoute().Handler(mu)
	}

	srv := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			logrus.Errorf("http server err:%s", err.Error())
		}
	}()
	return srv, nil
}

func startPprof() {
	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			logrus.Errorf("pprof server err:%s \n", err)
		}
	}()
}

func Init() error {
	if err := config.Init(*configPath); err != nil {
		return err
	}
	// init db
	dbMap := config.Default.GetStringMap("db")
	for key := range dbMap {
		if err := db.AddDB(
			config.Default.GetString(fmt.Sprintf("db.%s.type", key)),
			key,
			config.Default.GetString(fmt.Sprintf("db.%s.dsn", key)),
			db.MaxConn(config.Default.GetInt(fmt.Sprintf("db.%s.max_conn", key))),
			db.IdleConn(config.Default.GetInt(fmt.Sprintf("db.%s.idle_conn", key))),
			db.MaxLeftTime(config.Default.GetInt64(fmt.Sprintf("db.%s.max_lefttime", key))),
		); err != nil {
			return err
		}
	}
	println("init db success")
	if err := model.AutoMigrate(db.MustGetDB("test").DB); err != nil {
		return err
	}
	logrus.Infoln("init db and auto migrate success")

	return nil
}
