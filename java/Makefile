.PHONY: clean build server client proto test

# Generate Protobuf code only (useful when you update .proto files and want IDE suggestions)
proto:
	@echo "Generating Protobuf files..."
	chmod +x gradlew
	./gradlew generateProto

# Grant execution permission and build the project (skip tests for faster build)
build:
	@echo "Building the project..."
	chmod +x gradlew
	./gradlew clean build -x test

# Clean up old build files
clean:
	@echo "Cleaning old files..."
	./gradlew clean

# Start gRPC Server using Java (Ready to receive requests from Java or Go Client)
server:
	@echo "Starting Java Server..."
	./gradlew runServer

# Run gRPC Client using Java (Call to Java or Go Server)
client:
	@echo "Running Java Client..."
	./gradlew runClient

# Run unit tests
test:
	@echo "Running unit tests..."
	chmod +x gradlew
	./gradlew test