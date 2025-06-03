CREATE TABLE IF NOT EXISTS books (
    id INT AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    author VARCHAR(255) NOT NULL,
    quantity INT NOT NULL
);

INSERT INTO books (title, author, quantity) VALUES
('The Go Programming Language', 'Alan A. A. Donovan', 3),
('Go in Action', 'William Kennedy', 2),
('Go Web Programming', 'Sau Sheong Chang', 4),
('Go Programming Blueprints', 'Mat Ryer', 6),
('Learning Go', 'Jon Bodner', 1),
('Go Design Patterns', 'Mario Castro Contreras', 2),
('Go Systems Programming', ' Mihalis Tsoukalos', 3),
('Concurrency in Go', 'Katherine Cox-Buday', 4),
('Go in Practice', 'Matt Butcher', 2),
('Go Cookbook', 'Aaron Torres', 5),
('Mastering Go', 'Mihalis Tsoukalos', 3),
('Go Programming by Example', 'Agus Kurniawan', 4),
('Go for DevOps', 'Josh Armitage', 2),
('Go in Action, Second Edition', 'William Kennedy', 1),
('Introducing Go', 'Caleb Doxsey', 5);

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    email VARCHAR(255) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL
);
