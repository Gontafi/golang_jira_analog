package queries

const (
	CreateUserProject   = `INSERT INTO users_projects(user_id, project_id)  VALUES ($1, $2) RETURNING id`
	GetUsersFromProject = `SELECT u.id, u.username, u.full_name, u.email, u.role_id 
							FROM users u JOIN users_projects p ON u.id = p.user_id WHERE project_id = $1`
	GetProjectFromUsers = `SELECT p.id, p.name, p.resume, p.description, p.code, p.project_lead_id
							FROM projects p JOIN users_projects pu ON p.id = pu.id WHERE user_id = $1`
	RemoveUserFromProject = `DELETE FROM users_projects WHERE user_id = $1 and project_id = $2`
)
