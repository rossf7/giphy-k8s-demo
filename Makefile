default: build

# Build Docker image
build: docker_build output

# Settings can be overidden with an env var.
DOCKER_IMAGE ?= rossf7/giphy-k8s-demo
DOCKER_TAG ?= latest
BINARY ?= giphydemo

# Get the latest commit.
GIT_COMMIT = $(strip $(shell git rev-parse --short HEAD))

SOURCES := $(shell find . -name '*.go')

clean: 
	rm $(BINARY)

$(BINARY): $(SOURCES)
	# Compile for Linux
	GOOS=linux go build -o $(BINARY)	

docker_build: $(BINARY)
	# Build Docker image
	docker build \
 	--build-arg BUILD_DATE=`date -u +"%Y-%m-%dT%H:%M:%SZ"` \
  --build-arg VCS_URL=`git config --get remote.origin.url` \
  --build-arg VCS_REF=$(GIT_COMMIT) \
	-t $(DOCKER_IMAGE):$(DOCKER_TAG) .

output:
	@echo Docker Image: $(DOCKER_IMAGE):$(DOCKER_TAG)
