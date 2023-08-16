package queries

const (
	CreateProject = `INSERT INTO projects (name, resume, description, code, project_lead_id) VALUES ($1, $2, $3, $4, $5)`

	UpdateProject = `UPDATE projects SET name=$2, resume=$3, description=$4, code=$5, project_lead_id=$6 WHERE ID=$1`

	DeleteProject = `DELETE FROM projects WHERE ID = $1`

	GetProjects = `SELECT * FROM projects`

	GetProjectById = `SELECT * FROM projects WHERE ID = $1`
)
