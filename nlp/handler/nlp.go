package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	nlp "nlp/proto"
)

type Nlp struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Nlp) Call(ctx context.Context, req *nlp.Request, rsp *nlp.Response) error {
	log.Info("Received Nlp.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Nlp) Stream(ctx context.Context, req *nlp.StreamingRequest, stream nlp.Nlp_StreamStream) error {
	log.Infof("Received Nlp.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&nlp.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Nlp) PingPong(ctx context.Context, stream nlp.Nlp_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&nlp.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
