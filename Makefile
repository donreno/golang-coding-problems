
test: 
	@echo "Running tests..."
	@go test ./... -v

coverage: 
	@echo "Coverfile..."
	@go test ./...  --covermode=count --coverprofile coverage.out >> /dev/null
	@go tool cover -func coverage.out
	@go tool cover -html=coverage.out -o coverfile_out.html

bench: 
	@echo "Running tests..."
	@go test ./... -bench=. -run=^a

badge:
	@echo "Creating Badge..."
	@./badge.sh > /dev/null
	@echo "Badge created!"

.PHONY: test coverage bench
