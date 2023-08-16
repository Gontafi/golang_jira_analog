package queries

const (
	CreateAttachment = `INSERT INTO attachments VALUES ($1, $2, $3, $4, $5)`

	UpdateAttachment = `UPDATE attachments SET issue_id=$2, file_name=$3, file_size=$4,
                      uploaded_by_user_id=$5, uploaded_date=$6 WHERE ID = $1`

	DeleteAttachment = `DELETE FROM attachments WHERE ID = $1`

	GetAttachments = `SELECT (ID, issue_id, file_name, file_size, uploaded_by_user_id, uploaded_date) FROM attachments`

	GetAttachmentById = `SELECT * FROM attachments WHERE ID = $1`
)
