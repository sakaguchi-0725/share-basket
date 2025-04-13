up:
	docker-compose up
build:
	docker-compose build
down:
	docker-compose down
exec-bff:
	docker-compose exec bff sh
exec-auth:
	docker-compose exec auth sh
cp-proto:
	buf export proto -o services/bff/proto