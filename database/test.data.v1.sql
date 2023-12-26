INSERT INTO users (username, display_name) VALUES ('leila-codes', 'Leila Richardson-Noyes');

INSERT INTO projects (key, name, description, owner_id) VALUES ('TESTP', 'Test Project #1', 'This is a simple test project. \nYa know, for testing purposes.', 1);

INSERT INTO issue_types (project_key, issue_type_id, type)
VALUES ('TESTP', DEFAULT, 'Bug'),
       ('TESTP', DEFAULT, 'Enhancement'),
       ('TESTP', DEFAULT, 'User Story'),
       ('TESTP', DEFAULT, 'Feature'),
       ('TESTP', DEFAULT, 'Task')
;

INSERT INTO issue_statuses (project_key, issue_status_id, status)
VALUES ('TESTP', DEFAULT, 'Pending'),
       ('TESTP', DEFAULT, 'Approved'),
       ('TESTP', DEFAULT, 'In Progress'),
       ('TESTP', DEFAULT, 'Implementation Complete'),
       ('TESTP', DEFAULT, 'Review Passed'),
       ('TESTP', DEFAULT, 'Complete');

INSERT INTO issue_priorities (project_key, priority)
VALUES ('TESTP', 'Low'),
       ('TESTP', 'Medium'),
       ('TESTP', 'High'),
       ('TESTP', 'Highest');