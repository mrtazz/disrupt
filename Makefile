#
# simple makefile to run and build things
#
PROJECT=github.com/mrtazz/disrupt

.phony: test benchmark format

test:
	@go test ${PROJECT}
	@go test ${PROJECT}/pushover
	@go test ${PROJECT}/twilio

benchmark:
	@echo "Running tests..."
	@go test -bench=. ${PROJECT}
	@go test -bench=. ${PROJECT}/pushover
	@go test -bench=. ${PROJECT}/twilio

format:
	@go fmt ${PROJECT}
	@go fmt ${PROJECT}/pushover
	@go fmt ${PROJECT}/twilio
