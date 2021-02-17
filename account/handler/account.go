package handler

import (
	// "google.golang.org/grpc/status"
	"context"

	"github.com/google/uuid"
	log "github.com/micro/micro/v3/service/logger"

	account "account/proto"
	"account/redis"
)

type Account struct {
	email    string
	password string
}

var redisClient = redis.Client{}

func newAccount() *Account {
	return &Account{}
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Account) Call(ctx context.Context, req *account.Request, rsp *account.Response) error {
	log.Info("Received Account.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Account) Stream(ctx context.Context, req *account.StreamingRequest, stream account.Account_StreamStream) error {
	log.Infof("Received Account.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Infof("Responding: %d", i)
		if err := stream.Send(&account.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Account) PingPong(ctx context.Context, stream account.Account_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&account.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

func (e *Account) Register(ctx context.Context, req *account.AccountInfo, res *account.AccountResponse) error {
	schema := newSchema()
	e.email = req.Email
	e.password = req.Password
	AccountSchema.Create(schema, ctx, e)

	uuid, err := uuid.NewRandom()
	if err != nil {
		log.Fatal(err)
	}

	JWT, err := NewJWT(userID)

	// TODO: finish

	return nil
}

func (e *Account) Authenticate(ctx context.Context, req *account.AccountInfo, res *account.AccountResponse) error {
	schema := newSchema()
	// TODO: finish
	AccountSchema.Find(schema, ctx, e)
	return nil
}

func (e *Account) UpdateAccount(ctx context.Context, new *account.NewAccountInfo, res *account.AccountResponse) error {
	// schema := newSchema()
	// oldAccount := newAccount()
	// oldAccount := &Account{
	// 	email:    old.Email,
	// 	password: old.Password,
	// }

	// newAccount := &Account{
	// 	email:    new.Email,
	// 	password: new.Password,
	// }
	// err := accountSchema.Update(schema, ctx, oldAccount, newAccount)
	// if err != nil {
	// 	return err

	// }
	return nil
}
