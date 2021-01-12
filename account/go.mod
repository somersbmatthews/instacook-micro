module github.com/somersbmatthews/account

go 1.15

require (
	github.com/golang/protobuf v1.4.3 // indirect
	github.com/jackc/pgproto3/v2 v2.0.7 // indirect
	github.com/micro/micro/v3 v3.0.4
	github.com/soheilhy/cmux v0.1.4 // indirect
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	golang.org/x/text v0.3.5 // indirect
	golang.org/x/tools v0.0.0-20210108195828-e2f9c7f1fc8e // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gorm.io/driver/postgres v1.0.6
	gorm.io/gorm v1.20.11
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0
