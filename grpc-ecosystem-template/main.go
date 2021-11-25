package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	log "github.com/sirupsen/logrus"
	"github.com/xiaomLee/go-plugin/db"
	"google.golang.org/grpc"
	"grpc-ecosystem-template/api"
	"grpc-ecosystem-template/gateway"
	"grpc-ecosystem-template/internal/model"
	"grpc-ecosystem-template/service"
	"grpc-ecosystem-template/version"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

//go:generate make proto

var (
	port = flag.Int("Port", 10050, "grpc listen port, http port +1")
	v          = flag.Bool("version", false, "version")
)

func main()  {
	flag.Parse()
	if *v {
		version.PrintFullVersionInfo()
		return
	}

	if err:=Init(); err!=nil {
		log.Fatalf("init server err:%s", err)
	}

	// register shutdown signal
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, os.Kill, syscall.SIGUSR1, syscall.SIGUSR2)

	// run
	var (
		err error
		grpcSrv *grpc.Server
		httpSrv *http.Server
	)
	if grpcSrv, err =startGrpc(fmt.Sprintf(":%d", *port)); err !=nil {
		log.Fatal(err)
	}
	if httpSrv, err = startHttp(fmt.Sprintf("127.0.0.1:%d", *port), fmt.Sprintf(":%d", *port+1)); err!=nil {
		log.Fatal(err)
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
	lis, err:=net.Listen("tcp", addr)
	if err!=nil {
		return nil, err
	}
	go func() {
		if err := srv.Serve(lis); err !=nil {
			log.Errorf("grpc err:%s", err)
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

	prefix := "/test"
	router := mux.NewRouter()
	externalRouter := router.PathPrefix(prefix).Subrouter()
	externalRouter.NewRoute().Handler(http.StripPrefix(prefix, mu))

	srv := &http.Server{
		Addr:              addr,
		Handler:           router,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Errorf("http server err:%s", err.Error())
		}
	}()
	return srv, nil
}

func Init() error {
	if err:=db.AddDB(db.SQLITE, "test", "file:test.db?_auth&_auth_user=admin&_auth_pass=admin&_auth_crypt=sha1"); err!=nil {
		return err
	}
	if err := model.AutoMigrate(db.MustGetDB("test").DB);err!=nil {
		return err
	}
	log.Infoln("init db and auto migrate success")

	return nil
}