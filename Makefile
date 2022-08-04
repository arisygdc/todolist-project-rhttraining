pkg_migration := pkg/repository/db/migration

svcAuth_migrationDir := src/authservice/${pkg_migration}
svcAuth_dbUser := rhtPostgreUser
svcAuth_dbPassword := rhtPostgrePassword
svcAuth_dbName := auth

svcUser_migrationDir := src/userservice/${pkg_migration}
svcUser_dbUser := rhtPostgreUser
svcUser_dbPassword := rhtPostgrePassword
svcUser_dbName := user

createMigrateSvcUser:
	migrate create -ext sql -dir ${svcUser_migrationDir} -seq init_schema

createMigrateSvcAuth:
	migrate create -ext sql -dir ${svcAuth_migrationDir} -seq init_schema

migrateSvcUser:
	migrate -path ${svcUser_migrationDir} \
	-database "postgresql://${svcUser_dbUser}:${svcUser_dbPassword}@172.19.0.3:5432/${svcUser_dbName}?sslmode=disable" \
	-verbose up

migrateSvcAuth:
	migrate -path ${svcAuth_migrationDir} \
    -database "postgresql://${svcAuth_dbUser}:${svcAuth_dbPassword}@172.19.0.2:5432/${svcAuth_dbName}?sslmode=disable" \
    -verbose up

dbAuth:
	docker run --name rhtAuth-db -d \
	-e POSTGRES_USER=${svcAuth_dbUser} \
	-e POSTGRES_PASSWORD=${svcAuth_dbPassword} \
	-e POSTGRES_DB=${svcAuth_dbName} \
	-p 5432 \
	--network rht \
	postgres:13.4-alpine

dbUser:
	docker run --name rhtUser-db -d \
	-e POSTGRES_USER=${svcUser_dbUser} \
	-e POSTGRES_PASSWORD=${svcUser_dbPassword} \
	-e POSTGRES_DB=${svcUser_dbName} \
	-p 5432 \
	--network rht \
	postgres:13.4-alpine

runGateway:
	docker run --name gateway -d \
	-e SERVICE_AUTH_NAME=svc-auth \
	-e SERVICE_USER_NAME=svc-user \
	-p 8080:8080 \
	--network rht \
	arisygdc/rhttraininggateway:v0.1

runUser:
	docker run --name svc-user -d \
	-e DB_HOST=rhtUser-db \
    -e SERVICE_AUTH_NAME=svc-auth \
    -e SERVICE_USER_NAME=svc-user \
    --network rht \
    arisygdc/rhttraininguser:v0.1

runAuth:
	docker run --name svc-auth -d \
	-e DB_HOST=rhtAuth-db \
    -e SERVICE_AUTH_NAME=svc-auth \
    --network rht \
    arisygdc/rhttrainingauth:v0.1

.PHONY: createMigrateSvcUser createMigrateSvcAuth migrateSvcUser createMigrateSvcAuth dbAuth dbUser runGateway runUser runAuth