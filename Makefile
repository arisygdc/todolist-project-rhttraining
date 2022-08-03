pkg_migration := pkg/repository/db/migration
svcAuth_migrationDir := src/authservice/${pkg_migration}


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

.PHONY: createMigrateSvcUser migrateSvcUser createMigrateSvcAuth