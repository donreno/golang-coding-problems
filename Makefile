
test: 
	@echo "Running tests..."
	@go test ./... -v

coverage: 
	@echo "Coverfile..."
	@go test ./...  --coverprofile coverfile_out >> /dev/null
	@go tool cover -func coverfile_out
	@go tool cover -html=coverfile_out -o coverfile_out.html

bench: 
	@echo "Running tests..."
	@go test ./... -bench=. -run=^a

.PHONY: test coverage bench
