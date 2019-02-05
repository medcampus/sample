package server

import (
	"github.com/medcampus/backend/clap/config"
	"github.com/medcampus/backend/clap/db"
	"github.com/medcampus/backend/clap/libraries"
	"github.com/medcampus/backend/clap/proto"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"

	context "golang.org/x/net/context"
)

func Run() {

	initializeLibraries()

	listener, err := net.Listen("tcp", viper.GetString("app.port"))
	if err != nil {
		log.Errorf("Unable to create listener: %v", err)
	}

	srv := grpc.NewServer()
	pb.RegisterClapServer(srv, &ClapServer{})

	log.Info("Listening server at port 8080")
	srv.Serve(listener)
}

func initializeLibraries() {
	config.InitConfig()

	db.CreateDBSession()

	libraries.InitLog()
}

type ClapServer struct {}

func (s *ClapServer) Add(ctx context.Context,in *pb.Request) (*pb.Response, error) {

	session := db.GetMongoSession()
	defer session.Close()

	handler := db.NewHandler(session)
	claps, err := handler.AddClap(in.ServiceId)
	return &pb.Response{
		Claps: claps,
	}, err
}

func (s *ClapServer) Get(ctx context.Context, in *pb.Request) (*pb.Response, error) {

	session := db.GetMongoSession()
	defer session.Close()

	handler := db.NewHandler(session)
	claps, err := handler.GetClaps(in.ServiceId)
	return &pb.Response{
		Claps: claps,
	}, err
}