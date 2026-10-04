UPDATE writings 
SET 
    title = @title,
    subtitle = @subtitle,
    description = @description,
    writing_type = @writing_type,
    topics = @topics,
    tags = @tags,
    is_adult = @is_adult
WHERE
    uuid = @uuid
AND
    user_id = @user_id
RETURNING
    (uuid,
    title,
    subtitle,
    description,
    topics,
    tags,
    is_adult,
    writing_type,
    creator_id);
    