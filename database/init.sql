-- Create the employees table
CREATE TABLE IF NOT EXISTS employees (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    position VARCHAR(100) NOT NULL,
    salary INT NOT NULL
);

-- Insert sample data
INSERT INTO employees (name, position, salary) VALUES
('Alice Johnson', 'Software Engineer', 95000),
('Bob Smith', 'Project Manager', 110000),
('Charlie Brown', 'Product Designer', 87000),
('Diana Prince', 'QA Engineer', 78000);
