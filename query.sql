-- name: GetUser :one
SELECT * FROM users
WHERE username = $1;

-- name: GetUserID :one
SELECT * FROM users
WHERE id = $1;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id;

-- name: CreateUser :one
INSERT INTO users (username, bio, avatar, email, password)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: UpdateUser :one
UPDATE users
SET bio = $2, avatar = $3
WHERE id = $1
RETURNING *;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1
RETURNING *;


-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id;

-- name: ListPostsUser :many
SELECT * FROM posts
WHERE userId= $1;

-- name: ListCommentsPost :many
SELECT * FROM comments
WHERE postId= $1;

-- name: LikesPost :many
SELECT * FROM postLikes
WHERE postId= $1;

-- name: CreatePost :one
INSERT INTO posts (url, caption, userId)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdatePost :one
UPDATE posts
SET url = $2, caption = $3
WHERE id = $1
RETURNING *;

-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1
RETURNING *;

-- name: GetComment :one
SELECT * FROM comments
WHERE id = $1;

-- name: ListComments :many
SELECT * FROM comments
ORDER BY id;

-- name: CreateComment :one
INSERT INTO comments (contents, userId, postId)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateComment :one
UPDATE comments
SET contents = $2
WHERE id = $1
RETURNING *;

-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1
RETURNING *;

-- name: ListPostLikes :many
SELECT * FROM postLikes
ORDER BY id;

-- name: GetPostLike :one
SELECT * FROM postLikes
WHERE id = $1;

-- name: CreatePostLike :one
INSERT INTO postLikes (userId, postId)
VALUES ($1, $2)
RETURNING *;

-- name: DeletePostLike :one
DELETE FROM postLikes
WHERE id = $1
RETURNING *;

-- name: ListCommentLikes :many
SELECT * FROM commentLikes
ORDER BY id;

-- name: LikesComment :many
SELECT * FROM commentLikes
WHERE commentId = $1;

-- name: GetCommentLike :one
SELECT * FROM commentLikes
WHERE id = $1;

-- name: CreateCommentLike :one
INSERT INTO commentLikes (userId, commentId)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteCommentLike :one
DELETE FROM commentLikes
WHERE id = $1
RETURNING *;

-- name: GetHashtag :one
SELECT * FROM hashtags
WHERE title = $1;

-- name: ListHashtags :many
SELECT * FROM hashtags
ORDER BY id;

-- name: CreateHashtag :one
INSERT INTO hashtags (title)
VALUES ($1)
RETURNING *;

-- name: UpdateHashtag :one
UPDATE hashtags
SET title = $2
WHERE id = $1
RETURNING *;

-- name: DeleteHashtag :one
DELETE FROM hashtags
WHERE id = $1
RETURNING *;

-- name: GetFollower :one
SELECT * FROM followers
WHERE id = $1;

-- name: ListFollowers :many
SELECT * FROM followers
ORDER BY id;

-- name: CreateFollower :one
INSERT INTO followers (leaderId, followerId)
VALUES ($1, $2)
RETURNING *;

-- name: DeleteFollower :one
DELETE FROM followers
WHERE id = $1
RETURNING *;