#!/bin/bash
# first argument: Front Address for CORS
# second argument: path to folder to save static at. Should end by '/'. If starts from '/', then search from root of the file system. If value is "", then absolute dir is backend project root dir
# third argument: path to save profile photos. Should not start from '/'. Should not end by '/'.
# fourth argument: path to SSL certificate
# fifth argument: path to SSL private key
# sixth argument: postgres user name for main microservice
# seventh argument: database name
# eighth argument: postgres user password for main microservice
touch main.toml
echo "serverAddress = \":8080\"
originsAddressesCORS = [\"$1\"]
logLevel = \"debug\"
logFilePath = \"logs/serverRestApi/logs.log\"

staticDirAbsolutePath = \"$2\"
profilePhotosDirPath = \"$3\"

enableEchoCsrfToken = true
enableEchoSecurity = true

enableHttpsWithTLS = true
TLSCertificateFilePath = \"$4\"
TLSCertificateKeyFilePath = \"$5\"

databaseUser = \"$6\"
databaseName = \"$7\"
databasePassword = \"$8\"
databaseHost = \"127.0.0.1\"
databasePort = \"5432\"
databaseMaxOpenConnections = \"10\"

authSessionServiceAddress = \":8081\"
userProfileServiceAddress = \":8083\"" > main.toml
