package queries

const (
	GetUserByUsernameAndPassword = `SELECT id, username, created_at, 
      							 updated_at, full_name, email, role_id 
								FROM users WHERE username=$1 AND password_hash=$2`

	UpdateUserPassword = `UPDATE users SET password_hash = $2 WHERE username = $1`
)
