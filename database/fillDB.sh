#installs DB. Requres installed and running postgresql
DB_NAME=ve_ru
su postgres <<EOF
psql -U postgres -d $DB_NAME -f ./sql/PostgreSQLInsertValuesScript.sql
EOF
