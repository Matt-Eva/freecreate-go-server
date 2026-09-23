INSERT INTO writings (user_id, creator_id, title, writing_type) 
VALUES (@user_id, @creator_id, @title, @writing_type)
RETURNING user_id, creator_id, title, subtitle, writing_type, topics, tags, description;