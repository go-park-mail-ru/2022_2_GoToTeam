CREATE USER :user_name WITH PASSWORD ':user_password';
GRANT SELECT ON users, subscriptions, categories, articles, articles_likes, users_categories_subscriptions, comments, comments_likes, tags, tags_articles to :user_name;
GRANT INSERT ON users, subscriptions, articles, articles_likes, users_categories_subscriptions, comments, comments_likes, tags_articles to :user_name;
GRANT UPDATE ON users, subscriptions, articles, articles_likes, users_categories_subscriptions, comments, comments_likes, tags_articles to :user_name;
GRANT DELETE ON users, ubscriptions, articles, articles_likes, users_categories_subscriptions, comments, comments_likes, tags_articles to :user_name;
GRANT USAGE, SELECT ON ALL SEQUENCES IN SCHEMA public to :user_name;