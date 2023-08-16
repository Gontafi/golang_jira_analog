package queries

const (
	CreateIssue = `INSERT INTO issues VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`

	UpdateIssue = `UPDATE issues SET project_id=$1, issue_type_id=$2, 
                issue_summary=$3, issue_description=$4, reporter_id=$5, assignee_id=$6,
                stage_id=$7, status_id=$8, updated_at=$9, updated_at=$10, resolved_at=$11`

	DeleteIssue = `DELETE FROM issues WHERE ID = $1`

	GetIssues = `SELECT * FROM issues`

	GetIssueById = `SELECT * FROM issues WHERE ID = $1`

	// Issue Type

	CreateIssueType = `INSERT INTO issue_types VALUES ($1, $2)`

	UpdateIssueType = `UPDATE issue_types SET Name=$1, Description=$2`

	DeleteIssueType = `DELETE FROM issue_types WHERE ID = $1`

	GetIssueTypes = `SELECT (ID, Name, Description) FROM issue_types`

	GetIssueTypeById = `SELECT * FROM issue_types WHERE ID = $1`

	// Tags

	CreateTag = `INSERT INTO tags VALUES ($1)`

	UpdateTag = `UPDATE tags SET name=$1`

	DeleteTag = `DELETE FROM tags WHERE ID = $1`

	GetTags = `SELECT (ID, Name) FROM tags`

	GetTagById = `SELECT * FROM tags WHERE ID = $1`
)
