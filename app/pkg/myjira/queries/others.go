package queries

const (
	//Role

	CreateRole = `INSERT INTO roles VALUES ($1, $2)`

	UpdateRole = `UPDATE roles SET name=$2, description=$3 WHERE ID = $1`

	DeleteRole = `DELETE FROM roles WHERE ID = $1`

	GetRoles = `SELECT (id, name, description) FROM roles`

	GetRoleById = `SELECT * FROM roles WHERE ID = $1`

	// Status

	CreateStatus = `INSERT INTO statuses VALUES ($1, $2)`

	UpdateStatus = `UPDATE statuses SET name=$2, description=$3 WHERE ID = $1`

	DeleteStatus = `DELETE FROM statuses WHERE ID = $1`

	GetStatuses = `SELECT (ID, Name, Description) FROM statuses`

	GetStatusById = `SELECT * FROM statuses WHERE ID = $1`

	// Priority

	CreatePriority = `INSERT INTO stages VALUES ($1, $2)`

	UpdatePriority = `UPDATE stages SET name=$2, description=$3 WHERE ID = $1`

	DeletePriority = `DELETE FROM stages WHERE ID = $1`

	GetPriorities = `SELECT (ID, Name, Description) FROM stages`

	GetPriorityById = `SELECT * FROM stages WHERE ID = $1`
)
