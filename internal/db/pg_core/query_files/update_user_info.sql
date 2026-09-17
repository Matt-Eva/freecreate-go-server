UPDATE users 
SET username = @username, user_handle = @user_handle, reading_history = @reading_history, is_adult = @is_adult 
WHERE id = @user_id 
RETURNING username, user_handle, reading_history, is_adult;