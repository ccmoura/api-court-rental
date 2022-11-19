database:
	docker run --name court_rental_db -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

create_db:
	docker exec -it court_rental_db createdb --username=root --owner=root court_rental_db

drop_db:
	docker exec -it court_rental_db dropdb court_rental

migrate_up:
	migrate -path api/database/migration -database "postgresql://root:secret@localhost:5432/court_rental_db?sslmode=disable" -verbose up

migrate_down:
	migrate -path api/database/migration -database "postgresql://root:secret@localhost:5432/court_rental_db?sslmode=disable" -verbose down

rebuild_compose:
	docker-compose up --build --remove-orphans --force-recreate

.PHONY: database create_db drop_db migrate_up migrate_down rebuild_compose
