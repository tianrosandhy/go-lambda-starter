# Go Lambda Starter 

This repository contains the starter-code that help you develop AWS Lambda function with golang. 

AWS Lambda support many trigger from basic HTTP API gateway request, event-based from RabbitMQ, or just a direct event entry. This package will help you more focused on business logic instead of thinking about parsing the different request type. You can easily debug from local with the example payload too, and deploy after that to prevent too many error occured in deployed lambda.

### Installation
1. Clone or download this repository : `git clone https://github.com/tianrosandhy/go-lambda-starter`
2. Install dependencies `go get .`
3. `cp .env.example .env`
4. Make sure the environment value of `RUN_AS_LOCAL` is true. After that you can try to run : `go run .`

### Directory Structure
No many changes required in main.go since that file just contain the entry point that can be directed to lambda.Start, or just a direct access to handler. 

If you want to register some external service like Database, Redis, RabbitMQ, Kafka, etc you can define them in `./bootstrap/application.go`. that `Application` struct will be available in your business logic handler, and you can access the external connection there.

Struct is stored in `./src/types`. By default we use `EventPayload` as entry point of expected struct of lambda request. In your handler you will access this struct instead of completed one from AWS. 

Business logic handler is stored in `./src/handler : StartEventProcess()`. In this file you can handle what the function need to do with the request payload. 

Mock request for local only is stored in `./mock-event/request.json`. This file is only used if `RUN_AS_LOCAL=true`. So you can just run `go run .` and the function will be running locally (No need to deploy to test)

### Build and Deployment 
By default, AWS Lambda runtime that we will use is : `provided.al2023` (Amazon Linux 2023). So AWS will expect us to deploy the file as .zip file, with the content of binary file with filename `bootstrap`. 
We've prepared makefile to automate that : 
- `make build` : to build the binary file to ./build/bootstrap
- `make zip` : compress the file as zip to be sent to AWS to ./build/
- `make all` : build and compress the file (same as 2 command above)
- `make deploy` : deploy with serverless (note: you need to setup the serverless manually)

You can upload the zip file to AWS Lambda, and test them in AWS environment.
