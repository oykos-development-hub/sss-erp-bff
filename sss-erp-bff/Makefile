BINARY_NAME=bffMsApp

build:
	@echo "Building BFF..."
	@go build -o tmp/${BINARY_NAME} .
	@echo "BFF built!"

run: build
	@echo "Starting BFF..."
	@./tmp/${BINARY_NAME} 
	@echo "BFF started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm tmp/${BINARY_NAME}
	@echo "Cleaned!"

test:
	@echo "Testing..."
	@go test ./...
	@echo "Done!"

start: run

stop:
	@echo "Stopping BFF..."
	@-pkill -SIGTERM -f "./tmp/${BINARY_NAME}"
	@echo "Stopped BFF!"

restart: stop start