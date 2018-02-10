DROP TABLE IF EXISTS video;
CREATE TABLE video
(
    id            INT UNSIGNED UNIQUE NOT NULL AUTO_INCREMENT,
    video_key     VARCHAR(255) UNIQUE,
    title         VARCHAR(255)        NOT NULL,
    status        TINYINT                      DEFAULT 1,
    duration      INT UNSIGNED                 DEFAULT 0,
    url           VARCHAR(255)        NOT NULL,
    thumbnail_url VARCHAR(255)        NOT NULL DEFAULT '',
    PRIMARY KEY (id)
);

INSERT INTO
    video
SET
    video_key = 'd290f1ee-6c54-4b01-90e6-d701748f0851',
    title = 'Black Retrospective Woman',
    status = 3,
    duration = 127,
    url = '/content/d290f1ee-6c54-4b01-90e6-d701748f0851/index.mp4',
    thumbnail_url = '/content/d290f1ee-6c54-4b01-90e6-d701748f0851/screen.jpg';

INSERT INTO
    video
SET
    video_key = 'hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345',
    title = 'N Dance',
    status = 3,
    duration = 127,
    url = '/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/index.mp4',
    thumbnail_url = '/content/hjkhhjk3-23j4-j45k-erkj-kj3k4jl2k345/screen.jpg';

INSERT INTO
    video
SET
    video_key = 'sldjfl34-dfgj-523k-jk34-5jk3j45klj34',
    title = 'Cars',
    status = 3,
    duration = 127,
    url = '/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/index.mp4',
    thumbnail_url = '/content/sldjfl34-dfgj-523k-jk34-5jk3j45klj34/screen.jpg';
