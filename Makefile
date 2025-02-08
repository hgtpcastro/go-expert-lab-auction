.PHONY: start
start:
		@docker compose up --build -d	

.PHONY: stop
stop:
	@docker compose stop

.PHONY: clean
clean: stop
	@docker compose down -v --remove-orphans	

.PHONY: test
test:	
	@go test -v ./...