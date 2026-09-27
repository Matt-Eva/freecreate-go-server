SELECT (id, uuid, title, subtitle, description, topics, tags, writing_type, published, list_adds, likes, lib_adds, donations, flags, creators.id, creators.uuid, creators.name) 
FROM writings 
WHERE writings.user_id = @user_id AND writings.uuid = @uuid 
INNER JOIN ON creators
WHERE creators.id = writings.creator_id;