-- name: GetUser :one 
SELECT * FROM users 
WHERE id=$1 LIMIT 1;

-- name: ListUsers :many 
SELECT id, username, email, profile
FROM users 
ORDER BY username; 

-- name: CreateUser :one 
INSERT INTO users (username, password_hash, email, profile)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: UpdateUser :exec
UPDATE users 
    set 
        username = $2,
        email = $3,
        profile = $4
WHERE id = $1; 

-- name: DeleteUser :exec 
DELETE FROM users 
WHERE id = $1;

-- name: SavePasswordResetToken :exec
INSERT INTO password_reset (user_email, user_token)
VALUES($1, $2); 

-- name: UpdateUserPassword :exec
UPDATE users 
    set 
        password_hash = $2
WHERE email = $1;  

-- name: CreateWorkout :one
INSERT INTO workouts (user_id, name, description)
VALUES($1, $2, $3)
RETURNING id, user_id, name, description, date, created_at, updated_at;

-- name: GetWorkoutByUserID :many
SELECT id, user_id, name, description, date, created_at, updated_at
FROM workouts
WHERE user_id = $1;

-- name: GetWorkoutByID :one
SELECT id, user_id, name, description, date, created_at, updated_at
FROM workouts
WHERE id = $1;
