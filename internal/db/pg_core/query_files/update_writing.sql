UPDATE writings 
SET 
    title = @title,
    creator_id = @creator_id,
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
    writing_type,
    is_adult,
    creator_id);
    