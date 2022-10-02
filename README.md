# go-casbin-grpc-example
An example of how to use casbin as an RBAC GRPC service with MySQL

# Setup
## Install dependencies
```
$ go get github.com/casbin/casbin/v2
$ go get github.com/go-sql-driver/mysql
$ go get github.com/casbin/xorm-adapter/v2
$ go get google.golang.org/grpc
```

## Setup infrastructure
```
$ docker compose up &
```

## Recreating protobuf definition files
If necessary, you can recreate authz's protobuf definition files using below commands.
```
$ rm -f proto/authz/authz.pb.go
$ rm -f proto/authz/authz_grpc.pb.go
$ protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/authz/authz.proto
```

## Compiling
```
$ go build server/main.go
$ go build client/verify/main.go
$ go build client/get-roles-for-user/main.go
```

# Configuration
The RBAC policy is provided programmatically within the `newServer` function in `server/main.go`
```go
enforcer.AddPolicy("alice", "data1", "read")
enforcer.AddPolicy("data2_admin", "data2", "read")
enforcer.AddPolicy("data2_admin", "data2", "write")
enforcer.AddGroupingPolicy("alice", "data2_admin")
```

The definition of each line within the code quote defines that:
1. `alice` has `read` access to `data1`
2. `data2_admin` has `read` access to `data2`
3. `data2_admin` has `write` access to `data2`
4. `alice` has `data2_admin` role

# Running
## Server
To run the server, execute this following line:
```
$ go run server/main.go
```

## Verify client
To run the verify client, execute this following line:
```
$ go run client/verify/main.go <user> <resource> <action>
```

The following table provides the expected result of `go run` execution
| User | Resource | Action | Result |
| --- | --- | --- | --- |
| alice | data1 | read | Access allowed |
| alice | data1 | write | Access denied |
| alice | data2 | read | Access allowed |
| alice | data2 | write | Access allowed |
| bob | data1 | read | Access denied |

## Get-roles-for-user client
To run the get-roles-for-user client, execute this following line:
```
$ go run client/get-roles-for-user/main.go <user>
```

If you fill the user with `alice`, it prints `data2_admin`

# Cleaning-up
```
$ docker compose down
```
