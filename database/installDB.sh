#installs DB. Requires installed and running postgresql
DB_NAME=ve_ru
DB_USER=ve_ru_user
DB_USER_PASS=ve_ru_password
DB_AUTH_USER=ve_ru_auth_user
DB_PROFILE_USER=ve_ru_profile_user
su postgres <<EOF
createdb  $DB_NAME;
psql -U postgres -d $DB_NAME -f ./sql/PostgreSQLdbGenScript.sql
echo "database '$DB_NAME' created."
psql -U postgres -d  $DB_NAME -f ./sql/PostreSQLCreateMainUser.sql -v user_name='$DB_USER' -v user_password='$DB_USER_PASS'
echo "Postgres User '$DB_USER' created."
psql -U postgres -d  $DB_NAME -f ./sql/PostreSQLCreateAuthUser.sql -v user_name='$DB_AUTH_USER' -v user_password='$DB_USER_PASS'
echo "Postgres User '$DB_AUTH_USER' created."
psql -U postgres -d  $DB_NAME -f ./sql/PostreSQLCreateProfileUser.sql -v user_name='$DB_PROFILE_USER' -v user_password='$DB_USER_PASS'
echo "Postgres User '$DB_PROFILE_USER' created."
EOF
