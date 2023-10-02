-- Create the stock table
CREATE TABLE stock (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    current_price DECIMAL(10, 2),
    last_update TIMESTAMP
);

-- Insert sample data
INSERT INTO stock (id,name, current_price, last_update)
VALUES
    (1,'Apple', 100.50, NOW()),
    (2,'Microsoft', 75.25, NOW()),
    (3,'Samsung', 50.75, NOW());

