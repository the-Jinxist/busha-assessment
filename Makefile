DB_URL=postgresql://root:secret@localhost:5432/comments_db?sslmode=disable
postgres:
	docker run --name assessment-image -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it assessment-image createdb --username=root --owner=root comments_db

drobdb:
	docker exec -it assessment-image dropdb comments_db

migrateup:
	migrate -path database/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path database/migration -database "$(DB_URL)" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

create_swapi_docker_service:
	docker run --name neo-swapi --network busha-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@assessment-image:5432/comments_db?sslmode=disable" -e REDIS_ADDRESS="redis:6379" busha-assessment:latest

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server