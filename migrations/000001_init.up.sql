CREATE TABLE roles(
    id serial not null unique,
    name varchar(255) unique,
    description text
);

CREATE TABLE users (
    id serial not null unique,
    username varchar(255) not null unique,
    password_hash varchar(255) not null,
    full_name varchar(255) not null,
    email varchar not null unique,
    role_id int references roles(id) on delete cascade,
    created_at timestamp not null,
    updated_at timestamp not null
);

CREATE TABLE tags (
    id serial not null unique,
    name varchar(255)
);

CREATE TABLE issue_types (
    id serial not null unique,
    name varchar(255),
    description text
);

CREATE TABLE stages (
    id serial not null unique,
    name varchar(255),
    description text
);

CREATE TABLE statuses (
    id serial not null unique,
    name varchar(255),
    description text
);


CREATE TABLE projects (
    id serial not null unique,
    name varchar(255),
    resume varchar(255),
    description text,
    code varchar(255),
    project_lead_id int references users(id) on delete cascade not null,
    project_start_date timestamp,
    project_end_date timestamp
);

CREATE TABLE issues (
    id serial not null unique,
    project_id int references projects(id) on delete cascade not null,
    issue_type_id int references issue_types(id) on delete cascade not null,
    issue_summary text,
    issue_description text,
    reporter_id int references users(id) on delete cascade not null,
    assignee_id int references users(id) on delete cascade not null,
    stage_id int references stages(id) on delete cascade not null,
    status_id int references statuses(id) on delete cascade not null,
    created_at timestamp not null,
    updated_at timestamp not null,
    resolved_at timestamp not null
);
CREATE TABLE attachments (
    id serial not null unique,
    issue_id int references issues(id) on delete cascade not null not null,
    file_name varchar(255) not null,
    file_size int not null,
    uploaded_by_user_id int references users(id) on delete cascade not null,
    uploaded_date timestamp not null
);

CREATE TABLE comments (
    id serial not null unique,
    issue_id int references issues(id) on delete cascade not null,
    user_id int references users(id) on delete cascade not null,
    comment_text text,
    created_at timestamp
);

CREATE TABLE users_projects (
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    project_id int references projects(id) on delete cascade not null
);

-- Indexes for 'users' table
CREATE INDEX idx_users_username ON users (username);
CREATE INDEX idx_users_email ON users (email);

-- Indexes for 'projects' table
CREATE INDEX idx_projects_name ON projects (name);

-- Indexes for 'stages' table
CREATE INDEX idx_stages_name ON stages (name);

-- Indexes for 'issues' table
CREATE INDEX idx_issues_project_id ON issues (project_id);
CREATE INDEX idx_issues_reporter_id ON issues (reporter_id);
CREATE INDEX idx_issues_assignee_id ON issues (assignee_id);
CREATE INDEX idx_issues_status_id ON issues (status_id);

-- Indexes for 'attachments' table
CREATE INDEX idx_attachments_issue_id ON attachments (issue_id);

-- Indexes for 'comments' table
CREATE INDEX idx_comments_issue_id ON comments (issue_id);
CREATE INDEX idx_comments_user_id ON comments (user_id);

