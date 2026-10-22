UPDATE chapters
SET chapter_number = @chapter_number
WHERE user_id = @user_id AND uuid = @uuid
RETURNING chapter_number;