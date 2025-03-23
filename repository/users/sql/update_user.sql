UPDATE users
SET first_name = $1, last_name = $2, sex = $3, updated_at = now()
WHERE id = $4;