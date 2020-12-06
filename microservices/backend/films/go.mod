module github.com/sebach1/coding-challenge/microservices/backend/people

go 1.14

require (
	github.com/athomecomar/storeql v1.5.4
	github.com/athomecomar/xerrors v1.2.1
	github.com/golang-migrate/migrate/v4 v4.11.0
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/jmoiron/sqlx v1.2.1-0.20191203222853-2ba0fc60eb4a
	github.com/lib/pq v1.7.0
	github.com/pkg/errors v0.9.1
	golang.org/x/sys v0.0.0-20200610111108-226ff32320da // indirect
	google.golang.org/genproto v0.0.0-20200612171551-7676ae05be11 // indirect
	google.golang.org/grpc v1.29.1
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15 // indirect
)

replace google.golang.org/grpc => github.com/grpc/grpc-go v1.29.1