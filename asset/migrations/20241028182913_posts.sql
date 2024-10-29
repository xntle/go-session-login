-- +goose Up
-- +goose StatementBegin
CREATE TABLE posts (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    createdAt DATETIME NOT NULL
);

INSERT INTO posts (title, content, createdAt) VALUES 
('Introduction to SQL', 'This post introduces SQL basics for beginners.', '2024-10-01 10:00:00'),
('Understanding Primary Keys', 'Learn about primary keys and their role in databases.', '2024-10-02 12:30:00'),
('Advanced SQL Queries', 'A deep dive into complex SQL queries.', '2024-10-03 14:45:00'),
('Database Indexing Techniques', 'An overview of indexing and its performance benefits.', '2024-10-04 09:00:00'),
('Foreign Keys in SQL', 'This post explains foreign keys and relational database concepts.', '2024-10-05 11:15:00'),
('SQL Joins Explained', 'Understand different types of SQL joins and their uses.', '2024-10-06 16:30:00'),
('Data Types in SQL', 'A guide to data types in SQL and how to use them effectively.', '2024-10-07 08:45:00'),
('Stored Procedures in SQL', 'Learn about stored procedures and how they can simplify tasks.', '2024-10-08 13:00:00'),
('Working with NULL Values', 'Handling NULL values in SQL can be tricky; here’s how to manage them.', '2024-10-09 15:00:00'),
('Optimizing SQL Queries', 'Tips and tricks to improve the performance of your SQL queries.', '2024-10-10 10:30:00');

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE posts; 
-- +goose StatementEnd
