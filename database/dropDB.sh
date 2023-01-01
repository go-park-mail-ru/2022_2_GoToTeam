#Drop ve_ru database
DB_NAME=ve_ru
DB_USER=ve_ru_user
DB_AUTH_USER=ve_ru_auth_user
DB_PROFILE_USER=ve_ru_profile_user
su postgres <<EOF
dropdb $DB_NAME
psql -c "DROP ROLE $DB_USER;"
psql -c "DROP ROLE $DB_AUTH_USER;"
psql -c "DROP ROLE $DB_PROFILE_USER;"
EOF