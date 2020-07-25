package mygrpc

import (
	"context"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"reflect"
	"syscall"
	"time"
)

//Server wrapper of grpc server
type App struct {
	s         *grpc.Server
	sigChan   chan os.Signal
	registers map[reflect.Value]interface{}
	sopts     ServerOption
}

func NewApp(opts ...ServerOptions) *App {
	var servWrapper App
	for _, opt := range opts {
		opt(&servWrapper.sopts)
	}
	servWrapper.sigChan = make(chan os.Signal, 1)

	servWrapper.registers = make(map[reflect.Value]interface{}, 0)
	servWrapper.s = grpc.NewServer(servWrapper.sopts.grpcOpts...)
	return &servWrapper
}

func (sw *App) register() {
	for register, receiver := range sw.registers {
		rv := register
		rt := rv.Type()
		if rt.Kind() != reflect.Func {
			panic("register must be func")
		}

		cv := reflect.ValueOf(receiver)
		ct := cv.Type()
		if ct.Kind() == reflect.Ptr {
			ct = ct.Elem()
		}
		if ct.Kind() != reflect.Struct {
			panic("receiver must be struct")
		}

		params := make([]reflect.Value, 2)
		params[0] = reflect.ValueOf(sw.s)
		params[1] = cv

		rv.Call(params)
	}
}

// Register register
func (s *App) Register(register interface{}, receiver interface{}) {
	s.registers[reflect.ValueOf(register)] = receiver
}

func (sw *App) DumpGrpcName() {
	for fm, info := range sw.s.GetServiceInfo() {
		for _, method := range info.Methods {
			log.Info("GRPC Method: ", fm, method.Name)
		}
	}
	log.Info("GRPC Listen On: ", sw.sopts.address)
}

//Start start running server
func (sw *App) Start() error {
	lis, err := net.Listen("tcp", sw.sopts.address)
	if err != nil {
		return err
	}

	sw.register()

	//registry
	if sw.sopts.registry != nil {
		err := sw.sopts.registry.Register(
			context.TODO(),
			sw.sopts.serverName,
			sw.sopts.address,
		)
		if err != nil {
			return err
		}
	} else {
		log.Info("registry is nil")
	}
	sw.hookSignals()
	sw.DumpGrpcName()
	// Register reflection service on gRPC server.
	reflection.Register(sw.s)
	if err := sw.s.Serve(lis); err != nil {
		return err
	}
	return nil
}

//Stop stop tht server
func (sw *App) gracefulStop() {
	if sw.sopts.registry != nil {
		_ = sw.sopts.registry.Unregister(
			context.TODO(),
			sw.sopts.serverName,
			sw.sopts.address,
		)
		sw.sopts.registry.Close()
	}
	log.Warn("Receive Signal gracefulStop")

	sw.s.GracefulStop()
}

func (sw *App) stop() {
	if sw.sopts.registry != nil {
		_ = sw.sopts.registry.Unregister(
			context.TODO(),
			sw.sopts.serverName,
			sw.sopts.address,
		)
		sw.sopts.registry.Close()
	}
	log.Warn("Receive Signal stop")

	sw.s.Stop()
}

func (app *App) hookSignals() {
	signal.Notify(
		app.sigChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGSTOP,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
		syscall.SIGKILL,
	)

	go func() {
		var sig os.Signal
		for {
			sig = <-app.sigChan
			log.Warnf("Receive Signal %v", sig)
			time.Sleep(time.Second)
			switch sig {
			case syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGSTOP, syscall.SIGUSR1:
				app.gracefulStop() // graceful stop
			case syscall.SIGINT, syscall.SIGKILL, syscall.SIGUSR2, syscall.SIGTERM:
				app.stop() // terminalte now
			}
			time.Sleep(time.Second)
		}
	}()
}
