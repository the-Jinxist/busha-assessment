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

.PHONY: postgres createdb dropdb migrateup migratedown sqlc