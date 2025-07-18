package grpcserver

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	host   string
	port   string
	server *grpc.Server
	//service *service.Services
	logger *logrus.Logger
}

//func New(host string, port string, serv)
