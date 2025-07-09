# Variables
APPNAME ?= example-app-1  # Default app name if not specified
TAG ?= latest
REGISTRY ?= morawskioz

# Build the specified application 
# Usage: make build APPNAME=example-app-1 TAG=latest REGISTRY=your-registry
.PHONY: build
build:
	@if [ ! -d "$(APPNAME)" ]; then \
		echo "Error: Application directory '$(APPNAME)' not found"; \
		exit 1; \
	fi
	@echo "Building $(APPNAME)..."
	docker build -t $(REGISTRY)/$(APPNAME):$(TAG) ./$(APPNAME)

# Push the specified application
# Usage: make push APPNAME=example-app-1 TAG=latest REGISTRY=your-registry
.PHONY: push
push:
	@echo "Pushing $(APPNAME)..."
	docker push $(REGISTRY)/$(APPNAME):$(TAG)

# Build and push in one command
# Usage: make all APPNAME=example-app-1 TAG=latest REGISTRY=your-registry
.PHONY: all
all: build push

# Clean up
# Usage: make clean APPNAME=example-app-1 TAG=latest REGISTRY=your-registry
.PHONY: clean
clean:
	@echo "Cleaning up Docker image for $(APPNAME)..."
	docker rmi $(REGISTRY)/$(APPNAME):$(TAG) || true
