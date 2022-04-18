CREATE FUNCTION `calculateUserStatus`(id int) RETURNS varchar(20) CHARSET utf8mb4
BEGIN
    DECLARE i, cnt_answers int default 0;
    DECLARE points int default 0;
    DECLARE total_likes, total_best_answer int default 0;
    SET total_likes = (
        SELECT sum(likes)
        FROM Answers
        WHERE uid = id
    );
    # double points for likes for the answer posted within a month
    SET total_likes = total_likes + (
        SELECT sum(likes)
        FROM Answers
        WHERE uid = id AND DATE_SUB(now(), INTERVAL 1 MONTH) <= time
    );
    # add likes for question, also double points if posted within a month
    SET total_likes = total_likes + (
        SELECT sum(likes)
        FROM Questions
        WHERE uid = id
    ) + (
                          SELECT sum(likes)
                          FROM Questions
                          WHERE uid = id AND DATE_SUB(now(), INTERVAL 1 MONTH) <= time
                      );
    SET total_best_answer = (
        SELECT count(is_best)
        FROM Answers
        WHERE uid = id
    );
    SET points = total_likes + 100 * total_best_answer;
    IF (points <= 200) then
        RETURN 'basic';
    ELSEIF (points <= 1000) then
        RETURN 'intermediate';
    ELSEIF (points <= 3000) then
        RETURN 'advanced';
    ELSEIF (points <= 10000) then
        RETURN 'expert';
    ELSE
        RETURN 'master';
    END IF;
END