
CREATE TABLE usuario(
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL
) ENGINE=InnoDB;

START TRANSACTION;
INSERT INTO usuario(name, email) VALUES
    ('Lautaro', 'lau@email.com'),
    ('Mauri', 'mau@email.com'),
    ('Aizen', 'aizen_perri@email.com'),
    ('Alma', 'almitus@email.com');
COMMIT;

SELECT * FROM usuario;