CREATE TABLE IF NOT EXISTS comments (
    id SERIAL PRIMARY KEY, -- 用户 ID
    comment VARCHAR NOT NULL , -- 评论
    comment_date DATE DEFAULT CURRENT_DATE, -- 创建日期
    user_id BIGINT REFERENCES users(id) -- 关联用户 ID
);

INSERT INTO users(name) VALUES ('dev_test_user');
INSERT INTO comments (comment) VALUES ('first test comment');