# Define phony targets to prevent conflicts with files of the same name
.PHONY: all permission clean-build fast-build boot-jar clean run

# Default target when running just 'make'
all: permission clean-build

# Give execution permission to the Gradle wrapper (required on Linux/Ubuntu)
permission:
	@echo "Enabling execution permission for gradlew..."
	chmod +x gradlew

# Clean old build artifacts and rebuild the entire project (includes tests)
clean-build: permission
	./gradlew clean build

# Fast build by skipping all unit tests
fast-build: permission
	./gradlew build -x test

# Build executable JAR specifically for Spring Boot applications
boot-jar: permission
	./gradlew bootJar

# Remove the 'build' directory and clear cache
clean:
	./gradlew clean

# Run the compiled JAR file automatically from the build directory
run:
	@echo "Launching the application..."
	java -jar build/libs/*.jar
