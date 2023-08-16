package queries

const (
	CreateComment = `INSERT INTO comments VALUES ($1, $2, $3, $4)`

	UpdateComment = `UPDATE comments SET issue_id=$2, user_id=$3, comment_text=$4 WHERE ID = $1`

	DeleteComment = `DELETE FROM comments WHERE ID = $1`

	GetComments = `SELECT (ID, issue_id, user_id, comment_text, created_at) FROM comments`

	GetCommentById = `SELECT * FROM comments WHERE ID = $1`
)
