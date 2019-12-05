package main

import(
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcimpl/proto"
	"net"
)

type server struct {

}

func main(){
	listener, err :=net.Listen("tcp",":4040")
	if err !=nil{
		panic(err)
	}
	serv := grpc.NewServer()
	proto.RegisterAddServiceServer(serv,&server{})
	reflection.Register(serv)
	if e:=  serv.Serve(listener); e!=nil{
		panic(e)
	}
}

func (s *server) Add(ctx context.Context, req *proto.Request) (*proto.Response, error ){
 a,b := req.GetA(),req.GetB()
 result := a+b
 return &proto.Response{C:result},nil
}

func (s *server) Sub(ctx context.Context, req *proto.Request) (*proto.Response, error )  {
	a,b := req.GetA(),req.GetB()
	result := a-b
	return &proto.Response{C:result},nil
}
