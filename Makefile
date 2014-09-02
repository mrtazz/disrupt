#
# simple makefile to run and build things
#
PROJECT=github.com/mrtazz/disrupt

.phony: test benchmark format

deps:
	@echo "Running 'go get ./...' to get dependencies..."
	@go get ./...

test: deps
	@echo "Running tests..."
	@go test ${PROJECT}
	@go test ${PROJECT}/pushover
	@go test ${PROJECT}/twilio

benchmark: deps
	@echo "Running tests..."
	@go test -bench=. ${PROJECT}
	@go test -bench=. ${PROJECT}/pushover
	@go test -bench=. ${PROJECT}/twilio

format:
	@go fmt ${PROJECT}
	@go fmt ${PROJECT}/pushover
	@go fmt ${PROJECT}/twilio
