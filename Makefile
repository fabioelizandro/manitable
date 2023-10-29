default: help

.gobin/air:
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b .gobin

.PHONY: help
help:
	@echo "Usage: make <TARGET>\n"
	@grep -E "^[\. a-zA-Z_-]+:.*?## .*$$" $(firstword $(MAKEFILE_LIST)) | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

.PHONY: dev
dev: .gobin/air ## Run the development server
	@echo "Starting development server..."
	@.gobin/air