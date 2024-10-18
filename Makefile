export MYSQL_URL = 'mysql://root:root@tcp(localhost:3306)/simple-forum'

migrate-create migrate_create:
	@ migrate create -ext sql -dir scripts/migrations -seq $(name)

migrate-up migrate_up:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations up 1

migrate-down migrate_down:
	@ migrate -database ${MYSQL_URL} -path scripts/migrations down 1

d-restart d_restart:
	@ docker-compose down && docker-compose build && docker-compose up -d

d-mysql d_mysql:
	@ docker-compose down && docker-compose up local-mysql-simple-forum-db -d

d-down d_down:
	@ docker-compose down