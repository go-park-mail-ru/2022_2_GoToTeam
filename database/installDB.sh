#installs DB. Requires installed and running postgresql
# first argument: database name
# second argument: database main server user name
# third argument: database auth server user name
# fourth argument: database profile server user name
# fifth argument: database users password
DB_NAME=$1
DB_USER=$2
DB_AUTH_USER=$3
DB_PROFILE_USER=$4
DB_USER_PASS=$5
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
