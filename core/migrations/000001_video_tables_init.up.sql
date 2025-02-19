CREATE TABLE tbl_videos
(
    `id`            INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name`          VARCHAR(255) NOT NULL,
    `description`   TEXT,
    `video_url`     VARCHAR(255) NOT NULL,
    `config`        TEXT         NOT NULL,
    `path_resource` TEXT         NOT NULL,
    `level_system`  VARCHAR(255) NOT NULL,
    `status`        VARCHAR(255) NOT NULL,
    `note`          TEXT,
    `assign`        VARCHAR(255) NOT NULL,
    `Author`        VARCHAR(255) NOT NULL,
    `created_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`    TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`    TIMESTAMP    NULL     DEFAULT NULL
) ENGINE = InnoDB;

CREATE TABLE tbl_reactions
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `video_id`    INT UNSIGNED NOT NULL,
    `description` TEXT         NOT NULL,
    `name`        VARCHAR(50)  NOT NULL,
    `number`      INT UNSIGNED NOT NULL,
    `time_point`  INT UNSIGNED NOT NULL,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  TIMESTAMP    NULL     DEFAULT NULL,
    FOREIGN KEY (video_id) REFERENCES tbl_videos (id)
) ENGINE = InnoDB;

CREATE TABLE tbl_comments
(
    `id`          INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `video_id`    INT UNSIGNED NOT NULL,
    `description` TEXT,
    `comment`     TEXT         NOT NULL,
    `user_name`   VARCHAR(100) NOT NULL,
    `avatar`      VARCHAR(255) NOT NULL DEFAULT '',
    `time_point`  INT UNSIGNED NOT NULL,
    `created_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`  TIMESTAMP    NULL     DEFAULT NULL,
    FOREIGN KEY (video_id) REFERENCES tbl_videos (id)
) ENGINE = InnoDB;

CREATE TABLE tbl_viewers
(
    `id`         INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `video_id`   INT UNSIGNED NOT NULL,
    `number`     INT UNSIGNED NOT NULL,
    `time_point` INT UNSIGNED NOT NULL,
    `created_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` TIMESTAMP    NULL     DEFAULT NULL,
    FOREIGN KEY (video_id) REFERENCES tbl_videos (id)
) ENGINE = InnoDB;

CREATE TABLE tbl_objects
(
    `id`           INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `video_id`     INT UNSIGNED NOT NULL,
    `description`  TEXT,
    `coordinate_x` INT UNSIGNED NOT NULL,
    `coordinate_y` INT UNSIGNED NOT NULL,
    `length`       INT UNSIGNED NOT NULL,
    `width`        INT UNSIGNED NOT NULL,
    `order`        INT UNSIGNED NOT NULL,
    `time_point`   INT UNSIGNED NOT NULL,
    `created_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`   TIMESTAMP    NULL     DEFAULT NULL,
    FOREIGN KEY (video_id) REFERENCES tbl_videos (id)
) ENGINE = InnoDB;

CREATE TABLE tbl_activity_histories
(
    `id`           INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `video_id`     INT UNSIGNED NOT NULL,
    `note`         TEXT,
    `level_system` VARCHAR(255) NOT NULL,
    `actions`      VARCHAR(255) NOT NULL,
    `user`         VARCHAR(255) NOT NULL,
    `created_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at`   TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at`   TIMESTAMP    NULL     DEFAULT NULL,
    FOREIGN KEY (video_id) REFERENCES tbl_videos (id)
) ENGINE = InnoDB;