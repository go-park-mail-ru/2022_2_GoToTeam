#shows all tables with everything in it (only for tests, might has too big output)
DB_NAME=ve_ru
su postgres <<EOF
psql -U postgres  -d $DB_NAME -c "SELECT * FROM users;"
psql -U postgres  -d $DB_NAME -c "SELECT * FROM subscriptions;"
psql -U postgres  -d $DB_NAME -c "SELECT * FROM categories;"
psql -U postgres  -d $DB_NAME -c "SELECT * FROM articles;"
psql -U postgres  -d $DB_NAME -c "SELECT * FROM articles_likes;"
psql -U postgres  -d $DB_NAME -c "SELECT * FROM users_categories_subscriptions;"
psql -U postgres  -d $DB_NAME -c "SELECT * FROM comments;"
psql -U postgres  -d $DB_NAME -c "SELECT * FROM comments_likes;"
psql -U postgres  -d $DB_NAME -c "SELECT * FROM tags;"
psql -U postgres  -d $DB_NAME -c "SELECT * FROM tags_articles;"
echo "Database '$DB_NAME' is above."
EOF
