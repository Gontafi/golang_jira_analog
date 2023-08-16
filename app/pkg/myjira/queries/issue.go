package queries

const (
	CreateIssue = `INSERT INTO issues VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	UpdateIssue = `UPDATE issues SET project_id=$2, issue_type_id=$3, 
                issue_summary=$4, issue_description=$5, reporter_id=$6, assignee_id=$7,
                stage_id=$8, status_id=$9, updated_at=$10, updated_at=$11, resolved_at=$12 WHERE id=$1`

	DeleteIssue = `DELETE FROM issues WHERE ID = $1`

	GetIssues = `SELECT * FROM issues`

	GetIssueById = `SELECT * FROM issues WHERE ID = $1`

	// Issue Type

	CreateIssueType = `INSERT INTO issue_types VALUES ($1, $2)`

	UpdateIssueType = `UPDATE issue_types SET name=$2, description=$3 WHERE ID = $1`

	DeleteIssueType = `DELETE FROM issue_types WHERE ID = $1`

	GetIssueTypes = `SELECT (ID, Name, Description) FROM issue_types`

	GetIssueTypeById = `SELECT * FROM issue_types WHERE ID = $1`

	// Tags

	CreateTag = `INSERT INTO tags VALUES ($1)`

	UpdateTag = `UPDATE tags SET name=$2 WHERE ID = $1`

	DeleteTag = `DELETE FROM tags WHERE ID = $1`

	GetTags = `SELECT (ID, Name) FROM tags`

	GetTagById = `SELECT * FROM tags WHERE ID = $1`
)
