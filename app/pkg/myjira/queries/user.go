package queries

const (
	CreateUser = `INSERT INTO users (username, password_hash, full_name, email, 
                  role_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING Id`

	UpdateUser = `UPDATE users SET username=$2, updated_at=$3, password_hash=$4,
                full_name=$5, email=$6, role_id=$7 WHERE ID=$1`

	DeleteUser = `DELETE FROM users WHERE ID = $1`

	GetUsers = `SELECT (ID, Username, password_hash, full_name, Email, 
                  role_id, created_at, updated_at) FROM users`

	GetUserById = `SELECT * FROM users WHERE ID = $1`

	GetUserByEmail = `SELECT id, username, full_name, email, 
                      role_id, created_at, updated_at FROM users WHERE email = $1`

	GetUserByUsername = `SELECT id, username, full_name, email, 
                        role_id, created_at, updated_at FROM users WHERE username = $1`
)
