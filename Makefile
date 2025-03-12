GOCMD = go

install:
	${GOCMD} mod download && ${GOCMD} mod vendor


dev:
	${GOCMD} run main.go serve-http


env:
	cp .env.example .env

# create new sql db migration file in database/migration
# example usage: make new-db-migration name="create_user_table"
new-db-migration:
	${GOCMD} run main.go db:migrate create $(name) sql

run-db-migration:
	${GOCMD} run main.go db:migrate up