#!/bin/bash
# first argument: postgres user name for auth microservice
# second argument: database name
# third argument: postgres user password for main microservice
touch auth.toml
echo "serverAddress = \":8081\"
prometheusServerAddress = \":8082\"
logLevel = \"debug\"
logFilePath = \"../backend/logs/authSessionService/logs.log\"

databaseUser = \"$1\"
databaseName = \"$2\"
databasePassword = \"$3\"
databaseHost = \"127.0.0.1\"
databasePort = \"5432\"
databaseMaxOpenConnections = \"30\"" > auth.toml
