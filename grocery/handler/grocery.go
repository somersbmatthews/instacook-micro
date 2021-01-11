package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	grocery "grocery/proto"
)

type Grocery struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Grocery) Call(ctx context.Context, req *grocery.Request, rsp *grocery.Response) error {
	log.Info("Received Grocery.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Grocery) Stream(ctx context.Context, req *grocery.StreamingRequest, stream grocery.Grocery_StreamStream) error {
	log.Infof("Received Grocery.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&grocery.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Grocery) PingPong(ctx context.Context, stream grocery.Grocery_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&grocery.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
