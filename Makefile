pkg_migration := pkg/repository/db/migration

postgresUser := rhtPostgreUser
postgresPwd := rhtPostgrePassword

mongoUser := rhtMongoUser
mongoPwd := rhtMongoPwd

svcAuth_migrationDir := src/authservice/${pkg_migration}
svcAuth_dbName := auth

svcUser_migrationDir := src/userservice/${pkg_migration}
svcUser_dbName := user

createMigrateSvcUser:
	migrate create -ext sql -dir ${svcUser_migrationDir} -seq init_schema

createMigrateSvcAuth:
	migrate create -ext sql -dir ${svcAuth_migrationDir} -seq init_schema

migrateSvcUser:
	migrate -path ${svcUser_migrationDir} \
	-database "postgresql://${postgresUser}:${postgresPwd}@172.19.0.3:5432/${svcUser_dbName}?sslmode=disable" \
	-verbose up

migrateSvcAuth:
	migrate -path ${svcAuth_migrationDir} \
    -database "postgresql://${postgresUser}:${postgresPwd}@172.19.0.2:5432/${svcAuth_dbName}?sslmode=disable" \
    -verbose up

dbAuth:
	docker run --name rhtAuth-db -d \
	-e POSTGRES_USER=${postgresUser} \
	-e POSTGRES_PASSWORD=${postgresPwd} \
	-e POSTGRES_DB=${svcAuth_dbName} \
	-p 5432 \
	--network rht \
	postgres:13.4-alpine

dbUser:
	docker run --name rhtUser-db -d \
	-e POSTGRES_USER=${postgresUser} \
	-e POSTGRES_PASSWORD=${postgresPwd} \
	-e POSTGRES_DB=${svcUser_dbName} \
	-p 5432 \
	--network rht \
	postgres:13.4-alpine

dbTodo:
	docker run --name rhtTodo-db -d \
	-p 27017:27017 \
	-e MONGO_INITDB_ROOT_USERNAME=${mongoUser} \
	-e MONGO_INITDB_ROOT_PASSWORD=${mongoPwd} \
	--network rht \
	mongo:5.0-focal

runGateway:
	docker run --name gateway -d \
	-e SERVICE_AUTH_NAME=svc-auth \
	-e SERVICE_USER_NAME=svc-user \
	-p 8080:8080 \
	--network rht \
	arisygdc/rhttraining:demo1-gateway

runUser:
	docker run --name svc-user -d \
	-e DB_HOST=rhtUser-db \
    -e SERVICE_AUTH_NAME=svc-auth \
    -e SERVICE_USER_NAME=svc-user \
    --network rht \
    arisygdc/rhttraining:demo1-user

runAuth:
	docker run --name svc-auth -d \
	-e DB_HOST=rhtAuth-db \
    -e SERVICE_AUTH_NAME=svc-auth \
    --network rht \
    arisygdc/rhttraining:demo1-auth

runTodo:
	docker run --name svc-todo -d \
	-e DB_HOST=rhtTodo-db \
	-e SERVICE_TODO_NAME=svc-todo \
	-e DB_USER=${mongoUser} \
	-e DB_PASSWORD=${mongoPwd} \
	--network rht \
	arisygdc/rhttraining:demo1-todo

.PHONY: createMigrateSvcUser createMigrateSvcAuth migrateSvcUser createMigrateSvcAuth \
		dbAuth dbTodo dbUser dbTodo runGateway runUser runAuth