#!/bin/bash
# first argument: postgres user name for auth microservice
# second argument: database name
# third argument: postgres user password for main microservice
touch profile.toml
echo "serverAddress = \":8083\"
prometheusServerAddress = \":8085\"
logLevel = \"debug\"
logFilePath = \"logs/userProfileService/logs.log\"

databaseUser = \"$1\"
databaseName = \"$2\"
databasePassword = \"$3\"
databaseHost = \"127.0.0.1\"
databasePort = \"5432\"
databaseMaxOpenConnections = \"10\"

authSessionServiceAddress = \":8081\"" > profile.toml
