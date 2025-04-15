DROP TABLE items;
CREATE TABLE items (
    id INT PRIMARY KEY AUTO_INCREMENT,
    checklist_id INT NOT NULL,
    item_name VARCHAR(100) NOT NULL,
    description TEXT,
    status ENUM('pending', 'completed') DEFAULT 'pending',
    CONSTRAINT fk_items_checklists FOREIGN KEY (checklist_id) REFERENCES checklists(id) ON DELETE CASCADE ON UPDATE CASCADE
);