run:
	@sh ./scripts/go-run.sh
swag-init:
	swag init --parseDependency --parseInternal --parseDepth 1 --overridesFile .swaggo
swag-fmt:
	swag fmt --exclude docs,scripts -g main.go