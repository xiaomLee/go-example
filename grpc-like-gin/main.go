package main

//go:generate make proto

import (
	"github.com/google/uuid"
	"math"
	"os"
	"os/signal"

	"github.com/xiaomLee/go-example/grpc-like-gin/engine"
	"github.com/xiaomLee/go-example/grpc-like-gin/middleware"

	"google.golang.org/grpc"
)

func main() {
	ser, err := newGRpcEngine().Run(":1234", grpc.MaxRecvMsgSize(math.MaxInt32))
	if err != nil {
		panic(err.Error())
	}
	println("server start success")

	// hold here and deal request...
	// time.Sleep(time.Second * 100)
	exit := make(chan os.Signal)
	signal.Notify(exit, os.Interrupt, os.Kill)
	select {
	case <-exit:
		ser.GracefulStop()
		println("server stop success")
	}

}

func newGRpcEngine() *engine.GRpcEngine {
	e := engine.NewGRpcEngine("MyAppName")

	// register handler here. implement HandleFunc func(c *GRpcContext)
	e.RegisterFunc("hello", "world", sayHi)
	e.RegisterFunc("user", "create", createUser)

	e.Use(middleware.Recover)
	e.Use(middleware.Logger)

	return e
}

func sayHi(c *engine.GRpcContext) {
	name := c.StringParamDefault("name", "")
	c.SuccessResponse("hi " + name)
}

type User struct {
	Id int64
	Name string
	Age int
	Sex string
}

func createUser(c *engine.GRpcContext) {
	name := c.StringParam("name")
	sex := c.StringParam("sex")
	age := c.IntParam("age")

	user := User{
		Id: int64(uuid.New().ID()),
		Name: name,
		Age:  age,
		Sex:  sex,
	}
	// todo your business

	// return
	c.SuccessResponse(&user)
}