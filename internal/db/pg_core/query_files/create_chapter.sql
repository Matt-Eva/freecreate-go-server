INSERT INTO chapters (user_id, creator_id, writing_id, title, chapter_number, content_db) 
    VALUES (@user_id, @creator_id, @writing_id, @title, @chapter_number, @content_db)
    RETURNING (title, chapter_number, published, uuid, content_db );