#installs DB. Requres installed and running postgresql
DB_NAME=ve_ru
DB_USER=ve_ru_user
DB_USER_PASS=ve_ru_password
DB_AUTH_USER=ve_ru_auth_user
DB_PROFILE_USER=ve_ru_profile_user
su postgres <<EOF
createdb  $DB_NAME;
psql -c "CREATE USER $DB_USER WITH PASSWORD '$DB_USER_PASS';"
psql -U postgres -d $DB_NAME -f PostgreSQLdbGenScript.sql
psql -U postgres -d  $DB_NAME -c "GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public to $DB_USER;"
psql -U postgres -d  $DB_NAME -c "GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public to $DB_USER;"
echo "Postgres User '$DB_USER' and database '$DB_NAME' created."
psql -c "CREATE USER $DB_AUTH_USER WITH PASSWORD '$DB_USER_PASS';"
psql -U postgres -d  $DB_NAME -c "GRANT SELECT ON users to $DB_AUTH_USER;"
echo "Postgres User '$DB_AUTH_USER' created."
psql -c "CREATE USER $DB_PROFILE_USER WITH PASSWORD '$DB_USER_PASS';"
psql -U postgres -d  $DB_NAME -c "GRANT SELECT, UPDATE ON users to $DB_PROFILE_USER;"
echo "Postgres User '$DB_PROFILE_USER' created."
EOF
