CREATE DATABASE IF NOT EXISTS `douyin` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;
USE `douyin`;
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`             bigint(20) unsigned  NOT NULL AUTO_INCREMENT      COMMENT '用户id',
    `name`           varchar(128)         NOT NULL DEFAULT ''          COMMENT '用户名',
    `password`       varchar(128)         NOT NULL DEFAULT ''          COMMENT '密码',
    `token`          varchar(128)         NOT NULL DEFAULT ''          COMMENT '用户鉴权',
    `followCount`    bigint(20)           NOT NULL DEFAULT 0           COMMENT '关注总数',
    `followerCount`  bigint(20)           NOT NULL DEFAULT 0           COMMENT '粉丝总数',
    `videoCount`     bigint(20)           NOT NULL DEFAULT 0           COMMENT '作品总数',
    `favouriteCount` bigint(20)           NOT NULL DEFAULT 0           COMMENT '喜欢总数',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='用户表';

INSERT INTO `user`
VALUES (1, 'glj123456', '123456', ' ', 0, 0, 0, 0),
       (2, 'lbz123456', '123456', ' ', 0, 0, 0, 0);


DROP TABLE IF EXISTS `video`;
CREATE TABLE `video`
(
    `id`                   bigint(20) unsigned  NOT NULL AUTO_INCREMENT            COMMENT '视频id',
    `userId`               bigint(20) unsigned  NOT NULL DEFAULT 0                 COMMENT '用户id',
    `playUrl`              varchar(128)         NOT NULL DEFAULT ''                COMMENT '视频播放地址',
    `coverUrl`             varchar(128)         NOT NULL DEFAULT ''                COMMENT '视频封面地址',
    `favouriteCount`       bigint(20)           NOT NULL DEFAULT 0                 COMMENT '视频的点赞总数',
    `commentCount`         bigint(20)           NOT NULL DEFAULT 0                 COMMENT '视频的评论总数',
    `title`                varchar(128)         NOT NULL DEFAULT ''                COMMENT '视频标题',
    `createTime`          timestamp            NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    FOREIGN KEY (userId) REFERENCES user(id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT ='视频表';

INSERT INTO `video`(id,userId, playUrl, coverUrl, favouriteCount, commentCount, title)
    values(1,1, '/videos/firstVideo.MP4/', '/photos/firstPhoto.png/', 0, 0, '标题1');
# INSERT INTO `video`(userId, playUrl, coverUrl, favouriteCount, commentCount, title)
#     values(2, 'http://yyy.com', 'http://xx.com', 0, 0, '标题2');