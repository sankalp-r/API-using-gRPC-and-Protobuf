package main

import (
	"google.golang.org/grpc"
	"grpcimpl/proto"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func main()  {
	conn,err := grpc.Dial("localhost:4040",grpc.WithInsecure())
	if err!=nil{
		panic(err)
	}
	cli := proto.NewAddServiceClient(conn)
	g := gin.Default()
	g.GET("/add/:a/:b", func(context *gin.Context) {
			a,err:=strconv.ParseUint(context.Param("a"),10,64)
			if err!=nil{
				context.JSON(http.StatusBadRequest,gin.H{"error":"Invalid a"})
				return
			}
			b,err:=strconv.ParseUint(context.Param("b"),10,64)
			if err!=nil{
				context.JSON(http.StatusBadRequest,gin.H{"error":"Invalid a"})
				return
			}
			req:= &proto.Request{A:int64(a),B:int64(b)}
			if res,err:=cli.Add(context,req); err==nil{
				context.JSON(http.StatusOK,gin.H{"result":res.C})
			} else{
				context.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
			}

	})
	g.GET("/sub/:a/:b", func(context *gin.Context) {
		a,err:=strconv.ParseUint(context.Param("a"),10,64)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{"error":"Invalid a"})
			return
		}
		b,err:=strconv.ParseUint(context.Param("b"),10,64)
		if err!=nil{
			context.JSON(http.StatusBadRequest,gin.H{"error":"Invalid a"})
			return
		}
		req:= &proto.Request{A:int64(a),B:int64(b)}
		if res,err:=cli.Sub(context,req); err==nil{
			context.JSON(http.StatusOK,gin.H{"result":res.C})
		} else {
			context.JSON(http.StatusInternalServerError,gin.H{"error":err.Error()})
		}
	})

	if err:=g.Run(":8080"); err!=nil{
		log.Fatalf("Server failed: %v",err)
	}
}
