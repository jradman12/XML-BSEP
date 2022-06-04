package startup

import (
	"common/module/interceptor"
	postsProto "common/module/proto/posts_service"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"log"
	"net"
	"post/module/application"
	"post/module/domain/repositories"
	"post/module/infrastructure/handlers"
	"post/module/infrastructure/persistence"
	"post/module/startup/config"
)

type Server struct {
	config *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{config: config}
}

func (server *Server) Start() {
	mongoClient := server.InitMongoClient()
	postRepo := server.InitPostsRepo(mongoClient)
	postService := server.InitPostService(postRepo)
	postHandler := server.InitPostHandler(postService)
	server.StartGrpcServer(postHandler)
}

func (server *Server) InitMongoClient() *mongo.Client {
	client, err := persistence.GetClient(server.config.PostDBHost, server.config.PostDBPort)
	if err != nil {
		log.Fatalln(err)
	}
	return client
}

func (server *Server) InitPostsRepo(client *mongo.Client) repositories.PostRepository {
	return persistence.NewPostRepositoryImpl(client)
}

func (server *Server) InitPostService(repo repositories.PostRepository) *application.PostService {
	return application.NewPostService(repo)
}

func (server *Server) InitPostHandler(service *application.PostService) *handlers.PostHandler {
	return handlers.NewPostHandler(service)
}

func (server *Server) StartGrpcServer(handler *handlers.PostHandler) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(server.config.PublicKey))
	interceptor := interceptor.NewAuthInterceptor(config.AccessibleRoles(), publicKey)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(interceptor.Unary()))
	postsProto.RegisterPostServiceServer(grpcServer, handler)
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
