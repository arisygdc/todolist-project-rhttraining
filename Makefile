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
	-database "postgresql://${svcUser_dbUser}:${svcUser_dbPassword}@localhost:5432/${svcUser_dbName}?sslmode=disable" \
	-verbose up

migrateSvcAuth:
	migrate -path ${svcAuth_migrationDir} \
    -database "postgresql://${svcAuth_dbUser}:${svcAuth_dbPassword}@127.0.0.1:5432/${svcAuth_dbName}?sslmode=disable" \
    -verbose up

dbTestInit:
	docker run --name rhtAuth-db -d \
	-e POSTGRES_USER=${svcAuth_dbUser} \
	-e POSTGRES_PASSWORD=${svcAuth_dbPassword} \
	-e POSTGRES_DB=${svcAuth_dbName} \
	-p 5432:5432 \
	postgres:13.4-alpine

.PHONY: createMigrateSvcUser createMigrateSvcAuth migrateSvcUser createMigrateSvcAuth