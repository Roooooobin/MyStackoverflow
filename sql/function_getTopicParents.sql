# function to get parent topics in recursive way
CREATE FUNCTION `getTopicParents`(root_id int) RETURNS varchar(1000) CHARSET utf8mb4
BEGIN
    DECLARE res VARCHAR(1000) DEFAULT '$';
    DECLARE par int DEFAULT (SELECT parent_tid FROM TopicHierarchy WHERE tid = root_id);
    WHILE root_id != par DO
            SET res = CONCAT(res, ',', par);
            SET root_id = par;
            SET par = (SELECT parent_tid FROM TopicHierarchy WHERE tid = par);
        END WHILE;
    RETURN res;
END
