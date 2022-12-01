#Drop ve_ru database
DB_NAME=ve_ru
su postgres <<EOF
dropdb $DB_NAME
EOF