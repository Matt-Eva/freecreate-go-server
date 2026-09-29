SELECT 
    (writings.id, 
    writings.uuid, 
    title, 
    subtitle, 
    description, 
    writings.topics, 
    writings.tags, 
    writing_type, 
    published, 
    list_adds, 
    likes, 
    lib_adds, 
    writings.donations, 
    writings.flags, 
    creators.id, 
    creators.uuid, 
    creators.name)
FROM writings 
INNER JOIN creators
ON 
    creators.id = writings.creator_id
WHERE 
    writings.user_id = @user_id 
AND 
    writings.uuid = @uuid;