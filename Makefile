include .env
export

DATABASE_DSN ?= $(DATABASE_USERNAME):$(DATABASE_PASSWORD)@tcp($(DATABASE_HOST):$(DATABASE_PORT))/$(DATABASE_DATABASE)

migrate-setup:
	@if [ -z "$$(which migrate)" ]; then echo "Installing migrate command..."; go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate; fi


migrate-up: migrate-setup
	@ migrate -database 'mysql://$(DATABASE_DSN)?multiStatements=true' -path database/migrations up $(N)


migrate-down: migrate-setup
	@ migrate -database 'mysql://$(DATABASE_DSN)?multiStatements=true' -path database/migrations down $(N)


migrate-to-version: migrate-setup
	@ migrate -database 'mysql://$(DATABASE_DSN)?multiStatements=true' -path database/migrations goto $(V)


drop-db: migrate-setup
	@ migrate -database 'mysql://$(DATABASE_DSN)?multiStatements=true' -path database/migrations drop


force-version: migrate-setup
	@ migrate -database 'mysql://$(DATABASE_DSN)?multiStatements=true' -path database/migrations force $(V)


migration-version: migrate-setup
	@ migrate -database 'mysql://$(DATABASE_DSN)?multiStatements=true' -path database/migrations version


build:
	@ go build main.go

run: build
	$(env) go run .