CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE houses (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    price NUMERIC(10, 2) NOT NULL,
    owner_id INTEGER REFERENCES users(id)
);
-- Insert 5 users
INSERT INTO users (name, email, password) VALUES ('John Doe', 'john@example.com', 'password1');
INSERT INTO users (name, email, password) VALUES ('Jane Smith', 'jane@example.com', 'password2');
INSERT INTO users (name, email, password) VALUES ('Bob Johnson', 'bob@example.com', 'password3');
INSERT INTO users (name, email, password) VALUES ('Alice Lee', 'alice@example.com', 'password4');
INSERT INTO users (name, email, password) VALUES ('Charlie Brown', 'charlie@example.com', 'password5');

-- Insert 10 houses for each user
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 1', 'Spacious 3-bedroom house', 500000, 1);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 2', 'Cozy 2-bedroom bungalow', 350000, 1);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 3', 'Modern 4-bedroom townhouse', 650000, 1);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 4', 'Luxury 5-bedroom villa', 1000000, 1);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 5', 'Charming 1-bedroom cottage', 250000, 1);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 6', 'Beautiful 2-bedroom apartment', 400000, 2);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 7', 'Stylish 3-bedroom condo', 550000, 2);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 8', 'Spacious 4-bedroom house', 700000, 2);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 9', 'Elegant 5-bedroom mansion', 1200000, 2);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 10', 'Cozy studio apartment', 200000, 2);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 11', 'Modern 1-bedroom loft', 300000, 3);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 12', 'Spacious 2-bedroom townhouse', 450000, 3);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 13', 'Luxury 3-bedroom penthouse', 800000, 3);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 14', 'Charming 4-bedroom cottage', 550000, 3);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 15', 'Beautiful 5-bedroom villa', 900000, 3);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 16', 'Stylish 1-bedroom apartment', 250000, 4);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 17', 'Spacious 2-bedroom condo', 400000, 4);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 18', 'Luxury 3-bedroom house', 750000, 4);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 19', 'Charming 4-bedroom bungalow', 600000, 4);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 20', 'Beautiful 5-bedroom mansion', 1100000, 4);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 21', 'Cozy studio loft', 220000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 22', 'Modern 1-bedroom apartment', 280000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 23', 'Spacious 2-bedroom townhouse', 420000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 24', 'Luxury 3-bedroom penthouse', 780000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 25', 'Charming 4-bedroom cottage', 580000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 26', 'Beautiful 5-bedroom villa', 950000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 27', 'Stylish 1-bedroom condo', 320000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 28', 'Spacious 2-bedroom house', 480000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 29', 'Luxury 3-bedroom townhouse', 720000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 30', 'Charming 4-bedroom bungalow', 620000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 31', 'Beautiful 5-bedroom mansion', 1050000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 32', 'Cozy studio apartment', 210000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 33', 'Modern 1-bedroom loft', 270000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 34', 'Spacious 2-bedroom condo', 430000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 35', 'Luxury 3-bedroom penthouse', 750000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 36', 'Charming 4-bedroom cottage', 550000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 37', 'Beautiful 5-bedroom villa', 900000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 38', 'Stylish 1-bedroom apartment', 250000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 39', 'Spacious 2-bedroom townhouse', 450000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 40', 'Luxury 3-bedroom house', 800000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 41', 'Charming 4-bedroom bungalow', 600000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 42', 'Beautiful 5-bedroom mansion', 1100000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 43', 'Cozy studio loft', 220000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 44', 'Modern 1-bedroom apartment', 280000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 45', 'Spacious 2-bedroom condo', 420000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 46', 'Luxury 3-bedroom penthouse', 780000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 47', 'Charming 4-bedroom cottage', 580000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 48', 'Beautiful 5-bedroom villa', 950000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 49', 'Stylish 1-bedroom condo', 320000, 5);
INSERT INTO houses (title, description, price, owner_id) VALUES ('House 50', 'Spacious 2-bedroom house', 480000, 5);
