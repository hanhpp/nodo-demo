-- Create the stock table
CREATE TABLE stock (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    current_price DECIMAL(10, 2),
    last_update TIMESTAMP
);

-- Insert sample data
INSERT INTO stock (name, current_price, last_update)
VALUES
    ('Stock 1', 100.50, NOW()),
    ('Stock 2', 75.25, NOW()),
    ('Stock 3', 50.75, NOW());
