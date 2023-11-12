\c luna_track;
-- User reference table for FK's
CREATE TABLE users
(
    user_id  BIGSERIAL NOT NULL,
    username TEXT NOT NULL,
    display_name text NULL,
    PRIMARY KEY (user_id)
);

-- Projects Table
CREATE TABLE projects
(
    key         varchar(6)   NOT NULL,
    name        varchar(255) not null,
    description text null,
    owner_id    int          not null,
    CONSTRAINT fk_owner FOREIGN KEY (owner_id) REFERENCES users(user_id),
    primary key (key)
);

-- Project Members
CREATE TABLE project_members
(
    project_key varchar(6) NOT NULL,
    user_id     BIGINT     NOT NULL,
    CONSTRAINT fk_project_key FOREIGN KEY (project_key) REFERENCES projects(key),
    CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- Issue Types
CREATE TABLE issue_types
(
    project_key   varchar(6)   NOT NULL,
    issue_type_id SERIAL       NOT NULL,
    type          varchar(255) NOT NULL,
    CONSTRAINT fk_project_key FOREIGN KEY (project_key) REFERENCES projects(key),
    PRIMARY KEY (issue_type_id)
);

-- Issue Statuses
CREATE TABLE issue_statuses
(
    project_key     varchar(6)   NOT NULL,
    issue_status_id SERIAL       NOT NULL,
    status          varchar(255) NOT NULL,
    CONSTRAINT fk_project_key FOREIGN KEY (project_key) REFERENCES projects(key),
    PRIMARY KEY (issue_status_id)
);

-- Issue Priorities
CREATE TABLE issue_priorities
(
    project_key varchar(6)   NOT NULL,
    priority_id SERIAL       NOT NULL,
    priority    varchar(255) NOT NULL,
    CONSTRAINT fk_project_key FOREIGN KEY (project_key) REFERENCES projects(key),
    PRIMARY KEY (priority_id)
);

-- Issues
CREATE TABLE issues
(
    project_key  varchar(6)   NOT NULL,
    issue_id     BIGSERIAL    NOT NULL,
    issue_type   INT          NOT NULL,
    issue_status INT          NOT NULL,
    priority     INT          NOT NULL,
    summary      varchar(255) NOT NULL,
    description  text         not null,
    created_at   timestamp    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at   timestamp NULL,
    deleted_at   timestamp NULL,
    created_by   INT          NOT NULL,
    CONSTRAINT fk_project_key FOREIGN KEY (project_key) REFERENCES projects(key),
    CONSTRAINT fk_issue_type FOREIGN KEY (issue_type) REFERENCES issue_types(issue_type_id),
    CONSTRAINT fk_issue_status FOREIGN KEY (issue_status) REFERENCES issue_statuses(issue_status_id),
    CONSTRAINT fk_issue_priority FOREIGN KEY (priority) REFERENCES issue_priorities(priority_id),
    CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(user_id),
    PRIMARY KEY (issue_id)
);

-- Link relationships
-- Issue Priorities
CREATE TABLE issue_link_relationships
(
    project_key     varchar(6)   NOT NULL,
    relationship_id SERIAL       NOT NULL,
    relationship    varchar(255) NOT NULL,
    CONSTRAINT fk_project_key FOREIGN KEY (project_key) REFERENCES projects(key),
    PRIMARY KEY (relationship_id)
);

-- Issue Links
CREATE TABLE issue_links
(
    issue_link_id   BIGSERIAL NOT NULL,
    parent_issue_id INT       NOT NULL,
    child_issue_id  INT       NOT NULL,
    relationship_id INT       NOT NULL,
    CONSTRAINT fk_issue_parent FOREIGN KEY (parent_issue_id) REFERENCES issues(issue_id),
    CONSTRAINT fk_issue_child FOREIGN KEY (child_issue_id) REFERENCES issues(issue_id),
    CONSTRAINT fk_relationship FOREIGN KEY (relationship_id) REFERENCES issue_link_relationships(relationship_id),
    PRIMARY KEY (issue_link_id)
);

-- Comments
CREATE TABLE issue_comments
(
    comment_id BIGSERIAL NOT NULL,
    issue_id   BIGINT    NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL,
    deleted_at timestamp NULL,
    user_id    BIGINT    NOT NULL,
    content    TEXT      NOT NULL,
    CONSTRAINT fk_issue FOREIGN KEY (issue_id) REFERENCES issues(issue_id),
    CONSTRAINT fk_author FOREIGN KEY (user_id) REFERENCES users(user_id),
    PRIMARY KEY (comment_id)
)