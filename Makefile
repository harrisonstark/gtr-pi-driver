# Build the Docker container
build:
	docker build -t gtr-pi-driver .

# Run the Docker container and enter a bash shell
run:
	docker run -it --rm -p 7070:7070 --env-file .env gtr-pi-driver

# Shortcut to build, run, install, and start the app in development mode
start: build run