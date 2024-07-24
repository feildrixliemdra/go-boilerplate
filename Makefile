GOCMD = go

install:
	${GOCMD} mod download && ${GOCMD} mod vendor


dev:
	${GOCMD} run main.go serve-http


env:
	cp config/config.yaml.example config/config.yaml

# create new sql db migration file in database/migration
# example usage: make new-db-migration name="create_user_table"
new-db-migration:
	${GOCMD} run main.go db:migrate create $(name) sql

run-db-migration:
	${GOCMD} run main.go db:migrate up