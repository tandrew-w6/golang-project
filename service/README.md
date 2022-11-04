# Requirement
1. Go 1.16 or higher
2. Postgresql (any version)

# Where is the data?
There are 2 sql files in data folder and this data can be used for this service

# How to run this service
1. Open config.toml
2. Change db connection with your connection on section [PostgresDB]
3. Run terminal and run command below in the same folder main.go file
    make run
4. Service running when you see message like this "NEW-SERVICE running on port 10001"
5. Congrats your service server is running

# Architecture
domain/new
1. Handler -> receive request grpc from gateway and the function must be match with rpc in file proto/service.proto
2. Usecase -> logic code
3. Repo -> query to db
4. Model -> define structure for repo and usecase

proto
1. request.proto -> grpc request param
2. return.proto -> grpc return data
3. service.proto -> grpc function

pb
1. request.pb.go -> generate from proto
2. return.pb.go -> generate from proto
3. service_grpc.pb.go -> generate from proto
4. service.pb.go -> generate from proto

route
1. route.go -> register grpc

config.toml -> config for db