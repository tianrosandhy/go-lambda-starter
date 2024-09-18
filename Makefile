# Targets
.PHONY: all build clean

# Default target
all: build zip

# Build target
build: 
	@mv .env .env.bak
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o build/bootstrap .
	@mv .env.bak .env

# Clean target
clean:
	rm -rf build/

zip:
	zip -j build/golambdastarter.zip build/bootstrap
	
deploy:
	serverless deploy