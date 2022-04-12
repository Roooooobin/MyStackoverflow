# function to get parent topics in recursive way
CREATE DEFINER=`root`@`localhost` FUNCTION `f_getTopicParents`(root_id int) RETURNS varchar(1000) CHARSET utf8mb4
BEGIN
    DECLARE res VARCHAR(1000) DEFAULT '$';
    DECLARE par int DEFAULT 1;
    WHILE root_id != 1 DO
            SET par = (SELECT parent_tid FROM TopicHierarchy WHERE tid = root_id);
            IF par > 0 THEN
                SET res = CONCAT(res, ',', par);
            END IF;
            SET root_id = par;
        END WHILE;
    RETURN res;
END