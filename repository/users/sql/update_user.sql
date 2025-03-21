UPDATE users
SET title = $1, description = $2, status = $3, updated_at = now()
WHERE id = $4;