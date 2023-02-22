createdb:
	createdb movie_ratings

dropdb:
	dropdb movie_ratings

migrateup:
	migrate -path db/migration -database "postgresql:///movie_ratings?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql:///movie_ratings?sslmode=disable" -verbose down

sqlc:
	sqlc generate
	
.PHONY: craetedb dropdb migrateup migratedown sqlc