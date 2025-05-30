CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    quantity INT NOT NULL
);

INSERT INTO books (title, author, quantity) VALUES
('The Go Programming Language', 'Alan A. A. Donovan', 3),
('Introducing Go', 'Caleb Doxsey', 5);
