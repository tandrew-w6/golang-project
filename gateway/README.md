# Requirement
1. Go 1.16 or higher

# How to run this gateway
1. Run terminal and run command below in the same folder main.go file
    make run
2. Gateway running when you see message like this "Listening and serving HTTP on localhost:8080"
3. Congrats your gateway server is running

# Architecture
domain/new/handler/http_handler 
1. handler.go -> request and return data from grpc service

proto
1. request.proto -> same with grpc service
2. return.proto -> same with grpc service
3. service.proto -> same with grpc service

route
1. new.go -> define API and connect to handler