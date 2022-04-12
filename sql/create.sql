CREATE TABLE `AnswerLikes` (
                               `uid` int NOT NULL,
                               `aid` int NOT NULL,
                               `time` datetime DEFAULT CURRENT_TIMESTAMP,
                               PRIMARY KEY (`uid`,`aid`),
                               KEY `f_al_aid_idx` (`aid`),
                               CONSTRAINT `f_al_aid` FOREIGN KEY (`aid`) REFERENCES `Answers` (`aid`),
                               CONSTRAINT `f_al_uid` FOREIGN KEY (`uid`) REFERENCES `Users` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `Answers` (
                           `aid` int NOT NULL AUTO_INCREMENT,
                           `qid` int NOT NULL,
                           `uid` int NOT NULL,
                           `body` text,
                           `time` datetime DEFAULT CURRENT_TIMESTAMP,
                           `is_best` tinyint DEFAULT '0',
                           `likes` int DEFAULT '0',
                           PRIMARY KEY (`aid`),
                           KEY `idx_qid` (`qid`),
                           KEY `idx_user` (`uid`),
                           FULLTEXT KEY `fullidx_answer_body` (`body`),
                           CONSTRAINT `f_a_qid` FOREIGN KEY (`qid`) REFERENCES `Questions` (`qid`),
                           CONSTRAINT `f_a_uid` FOREIGN KEY (`uid`) REFERENCES `Users` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `AnswerTopic` (
                               `aid` int NOT NULL,
                               `tid` int NOT NULL,
                               PRIMARY KEY (`aid`,`tid`),
                               KEY `f_at_tid_idx` (`tid`),
                               CONSTRAINT `f_at_aid` FOREIGN KEY (`aid`) REFERENCES `Answers` (`aid`),
                               CONSTRAINT `f_at_tid` FOREIGN KEY (`tid`) REFERENCES `Topics` (`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `QuestionLikes` (
                                 `uid` int NOT NULL,
                                 `qid` int NOT NULL,
                                 `time` datetime DEFAULT CURRENT_TIMESTAMP,
                                 PRIMARY KEY (`uid`,`qid`),
                                 KEY `f_ql_qid_idx` (`qid`),
                                 CONSTRAINT `f_ql_qid` FOREIGN KEY (`qid`) REFERENCES `Questions` (`qid`),
                                 CONSTRAINT `f_ql_uid` FOREIGN KEY (`uid`) REFERENCES `Users` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `Questions` (
                             `qid` int NOT NULL AUTO_INCREMENT,
                             `uid` int NOT NULL,
                             `title` text NOT NULL,
                             `body` text,
                             `time` datetime DEFAULT CURRENT_TIMESTAMP,
                             `is_resolved` tinyint DEFAULT '0',
                             `best_aid` int DEFAULT NULL,
                             PRIMARY KEY (`qid`),
                             KEY `f_user_idx` (`uid`),
                             KEY `f_q_aid_idx` (`best_aid`),
                             FULLTEXT KEY `fullidx_question_body` (`body`),
                             FULLTEXT KEY `fullidx_title` (`title`),
                             CONSTRAINT `f_q_aid` FOREIGN KEY (`best_aid`) REFERENCES `Answers` (`aid`),
                             CONSTRAINT `f_q_uid` FOREIGN KEY (`uid`) REFERENCES `Users` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `QuestionTopic` (
                                 `qid` int NOT NULL,
                                 `tid` int NOT NULL,
                                 PRIMARY KEY (`qid`,`tid`),
                                 KEY `f_qt_tid_idx` (`tid`),
                                 CONSTRAINT `f_qt_qid` FOREIGN KEY (`qid`) REFERENCES `Questions` (`qid`),
                                 CONSTRAINT `f_qt_tid` FOREIGN KEY (`tid`) REFERENCES `Topics` (`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `TopicHierarchy` (
                                 `tid` int NOT NULL,
                                 `parent_tid` int NOT NULL,
                                 PRIMARY KEY (`tid`,`parent_tid`),
                                 KEY `f_th_ptid_idx` (`parent_tid`),
                                 CONSTRAINT `f_th_ptid` FOREIGN KEY (`parent_tid`) REFERENCES `Topics` (`tid`),
                                 CONSTRAINT `f_th_tid` FOREIGN KEY (`tid`) REFERENCES `Topics` (`tid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `Topics` (
                          `tid` int NOT NULL AUTO_INCREMENT,
                          `topic_name` varchar(45) NOT NULL,
                          PRIMARY KEY (`tid`),
                          KEY `idx_tname` (`topic_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
CREATE TABLE `Users` (
                         `uid` int NOT NULL AUTO_INCREMENT,
                         `username` varchar(45) NOT NULL,
                         `status` varchar(20) NOT NULL DEFAULT 'Basic',
                         `email` varchar(100) NOT NULL,
                         `password` varchar(255) NOT NULL,
                         `city` varchar(45) NOT NULL,
                         `state` varchar(45) NOT NULL,
                         `country` varchar(45) NOT NULL,
                         `profile` varchar(200) NOT NULL DEFAULT '',
                         PRIMARY KEY (`uid`),
                         UNIQUE KEY `index2` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
