CREATE TABLE checklists(
    id INT PRIMARY KEY AUTO_INCREMENT,
    user_username VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    CONSTRAINT fk_checklists_users FOREIGN KEY (user_username) REFERENCES users(username)
);