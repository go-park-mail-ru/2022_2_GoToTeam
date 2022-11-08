#installs DB. Requres installed and running postgresql
DB_NAME=ve_ru
DB_USER=ve_ru_user
DB_USER_PASS=ve_ru_password
su postgres <<EOF
psql -U postgres -d $DB_NAME -f PostgreSQLInsertValuesScript.sql
EOF
