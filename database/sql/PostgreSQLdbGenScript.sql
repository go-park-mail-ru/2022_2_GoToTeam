-- Users accounts
CREATE TABLE users
(
    user_id             SERIAL,
    email               VARCHAR(255) UNIQUE NOT NULL,
    login               VARCHAR(255) UNIQUE NOT NULL,
    password            VARCHAR(255)        NOT NULL,
    username            VARCHAR(255),
    sex                 CHAR(1) CHECK (sex IN ('M', 'F')),
    date_of_birth       DATE,
    avatar_img_path     VARCHAR(1024),
    registration_date   TIMESTAMP           NOT NULL DEFAULT CURRENT_TIMESTAMP,
    subscribers_count   INT                 NOT NULL DEFAULT 0 CHECK (subscribers_count >= 0),
    subscriptions_count INT                 NOT NULL DEFAULT 0 CHECK (subscribers_count >= 0),

    PRIMARY KEY (user_id)
);

CREATE TABLE subscriptions
(
    subscribe_date   TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id          INT       NOT NULL,
    subscribed_to_id INT       NOT NULL CHECK (subscribed_to_id != subscriptions.user_id),

    PRIMARY KEY (user_id, subscribed_to_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (subscribed_to_id) REFERENCES users (user_id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Trigger that increments subscribers_count and subscriptions_count for the users table after inserting in subscriptions.
CREATE FUNCTION USF_TRIGGER_subscriptions_INSERT() RETURNS TRIGGER AS
$$
BEGIN
    UPDATE users SET subscribers_count = subscribers_count + 1 WHERE users.user_id = NEW.subscribed_to_id;
    UPDATE users SET subscriptions_count = subscriptions_count + 1 WHERE users.user_id = NEW.user_id;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER subscriptions_INSERT
    AFTER INSERT
    ON subscriptions
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_subscriptions_INSERT();

-- Trigger that decrements subscribers_count and subscriptions_count for the users table after deletion in subscriptions.
CREATE FUNCTION USF_TRIGGER_subscriptions_DELETE() RETURNS TRIGGER AS
$$
BEGIN
    UPDATE users SET subscribers_count = subscribers_count - 1 WHERE users.user_id = OLD.subscribed_to_id;
    UPDATE users SET subscriptions_count = subscriptions_count - 1 WHERE users.user_id = OLD.user_id;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER subscriptions_DELETE
    AFTER DELETE
    ON subscriptions
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_subscriptions_DELETE();

CREATE TABLE categories
(
    category_id       SERIAL,
    category_name              VARCHAR(255) UNIQUE NOT NULL,
    description       TEXT UNIQUE         NOT NULL,
    subscribers_count INT                 NOT NULL DEFAULT 0 CHECK (subscribers_count >= 0),

    PRIMARY KEY (category_id)
);

CREATE TABLE articles
(
    article_id     SERIAL,
    title          VARCHAR(1024) NOT NULL,
    description    TEXT,
    rating         INT           NOT NULL DEFAULT 0,
    comments_count INT           NOT NULL DEFAULT 0 CHECK (comments_count >= 0),
    content        TEXT          NOT NULL,
    cover_img_path VARCHAR(1024),
    co_author_id   INT,
    publisher_id   INT           NOT NULL,
    category_id    INT,

    PRIMARY KEY (article_id),
    FOREIGN KEY (co_author_id) REFERENCES users (user_id) ON DELETE SET NULL ON UPDATE CASCADE,
    FOREIGN KEY (publisher_id) REFERENCES users (user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (category_id) ON DELETE SET NULL ON UPDATE CASCADE
);

CREATE TABLE articles_likes
(
    is_like      BOOLEAN   NOT NULL,
    publish_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id      INT       NOT NULL,
    article_id   INT       NOT NULL,

    PRIMARY KEY (user_id, article_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (article_id) REFERENCES articles (article_id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Trigger that increments rating for the articles table after inserting in articles_likes.
CREATE FUNCTION USF_TRIGGER_articles_likes_INSERT() RETURNS TRIGGER AS
$$
BEGIN
    IF NEW.is_like THEN
        UPDATE articles SET rating = rating + 1 WHERE article_id = NEW.article_id;
    ELSE
        UPDATE articles SET rating = rating - 1 WHERE article_id = NEW.article_id;
    END IF;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER articles_likes_INSERT
    AFTER INSERT
    ON articles_likes
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_articles_likes_INSERT();

-- Trigger that decrements rating for the articles table after deleting in articles_likes.
CREATE FUNCTION USF_TRIGGER_articles_likes_DELETE() RETURNS TRIGGER AS
$$
BEGIN
    IF OLD.is_like THEN
        UPDATE articles SET rating = rating - 1 WHERE article_id = OLD.article_id;
    ELSE
        UPDATE articles SET rating = rating + 1 WHERE article_id = OLD.article_id;
    END IF;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER articles_likes_DELETE
    AFTER DELETE
    ON articles_likes
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_articles_likes_DELETE();

CREATE TABLE users_categories_subscriptions
(
    subscribe_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id        INT       NOT NULL,
    category_id    INT       NOT NULL,

    PRIMARY KEY (user_id, category_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (category_id) REFERENCES categories (category_id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Trigger that increments subscribers_count for the categories table after inserting in users_categories_subscriptions.
CREATE FUNCTION USF_TRIGGER_users_categories_subscriptions_INSERT() RETURNS TRIGGER AS
$$
BEGIN
    UPDATE categories SET subscribers_count = subscribers_count + 1 WHERE categories.category_id = NEW.category_id;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER users_categories_subscriptions_INSERT
    AFTER INSERT
    ON users_categories_subscriptions
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_users_categories_subscriptions_INSERT();

-- Trigger that decrements subscribers_count for the categories table after deleting in users_categories_subscriptions.
CREATE FUNCTION USF_TRIGGER_users_categories_subscriptions_DELETE() RETURNS TRIGGER AS
$$
BEGIN
    UPDATE categories SET subscribers_count = subscribers_count - 1 WHERE categories.category_id = OLD.category_id;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER users_categories_subscriptions_DELETE
    AFTER DELETE
    ON users_categories_subscriptions
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_users_categories_subscriptions_DELETE();

CREATE TABLE comments
(
    comment_id             SERIAL,
    content                TEXT      NOT NULL,
    publish_date           TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    rating                 INT       NOT NULL DEFAULT 0,
    publisher_id           INT       NOT NULL,
    article_id             INT       NOT NULL,
    comment_for_comment_id INT,

    PRIMARY KEY (comment_id),
    FOREIGN KEY (publisher_id) REFERENCES users (user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (comment_for_comment_id) REFERENCES comments (comment_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (article_id) REFERENCES articles (article_id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Trigger that increments comments_count for the articles table after inserting in comments.
CREATE FUNCTION USF_TRIGGER_comments_INSERT() RETURNS TRIGGER AS
$$
BEGIN
    UPDATE articles SET comments_count = comments_count + 1 WHERE articles.article_id = NEW.article_id;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER comments_INSERT
    AFTER INSERT
    ON comments
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_comments_INSERT();

-- Trigger that decrements comments_count for the articles table after deleting in comments.
CREATE FUNCTION USF_TRIGGER_comments_DELETE() RETURNS TRIGGER AS
$$
BEGIN
    UPDATE articles SET comments_count = comments_count - 1 WHERE articles.article_id = OLD.article_id;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER comments_DELETE
    AFTER DELETE
    ON comments
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_comments_DELETE();

CREATE TABLE comments_likes
(
    is_like      BOOLEAN   NOT NULL,
    publish_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    user_id      INT       NOT NULL,
    comment_id   INT       NOT NULL,

    PRIMARY KEY (user_id, comment_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (comment_id) REFERENCES comments (comment_id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Trigger that increments rating for the comments table after inserting in comments_likes.
CREATE FUNCTION USF_TRIGGER_comments_likes_INSERT() RETURNS TRIGGER AS
$$
BEGIN
    IF NEW.is_like THEN
        UPDATE comments SET rating = rating + 1 WHERE comment_id = NEW.comment_id;
    ELSE
        UPDATE comments SET rating = rating - 1 WHERE comment_id = NEW.comment_id;
    END IF;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER comments_likes_INSERT
    AFTER INSERT
    ON comments_likes
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_comments_likes_INSERT();

-- Trigger that decrements rating for the comments table after deleting in comments_likes.
CREATE FUNCTION USF_TRIGGER_comments_likes_DELETE() RETURNS TRIGGER AS
$$
BEGIN
    IF OLD.is_like THEN
        UPDATE comments SET rating = rating - 1 WHERE comment_id = OLD.comment_id;
    ELSE
        UPDATE comments SET rating = rating + 1 WHERE comment_id = OLD.comment_id;
    END IF;
    RETURN NULL;
END
$$ LANGUAGE plpgSQL;

CREATE TRIGGER comments_likes_DELETE
    AFTER DELETE
    ON comments_likes
    FOR EACH ROW
EXECUTE PROCEDURE USF_TRIGGER_comments_likes_DELETE();

CREATE TABLE tags
(
    tag_id   SERIAL,
    tag_name VARCHAR(255) UNIQUE NOT NULL,

    PRIMARY KEY (tag_id)
);

CREATE TABLE tags_articles
(
    article_id INT NOT NULL,
    tag_id     INT NOT NULL,

    PRIMARY KEY (article_id, tag_id),
    FOREIGN KEY (article_id) REFERENCES articles (article_id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (tag_id) REFERENCES tags (tag_id) ON DELETE CASCADE ON UPDATE CASCADE
);
