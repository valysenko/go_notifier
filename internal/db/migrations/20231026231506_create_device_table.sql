-- +goose Up
-- +goose StatementBegin
CREATE TABLE device (
    id INT AUTO_INCREMENT PRIMARY KEY,
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INT,
    FOREIGN KEY (user_id) REFERENCES user(id),
    UNIQUE (token)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE device;
-- +goose StatementEnd
