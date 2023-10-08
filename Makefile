.PHONY: run

test:
	go test ./...

run:
	@bash ./scripts/run.sh
