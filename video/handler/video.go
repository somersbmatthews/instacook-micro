package handler

import (
	"context"

	log "github.com/micro/micro/v3/service/logger"

	video "video/proto"
)

type Video struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Video) Call(ctx context.Context, req *video.Request, rsp *video.Response) error {
	log.Info("Received Video.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Video) Stream(ctx context.Context, req *video.StreamingRequest, stream video.Video_StreamStream) error {
	log.Infof("Received Video.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&video.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Video) PingPong(ctx context.Context, stream video.Video_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&video.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
