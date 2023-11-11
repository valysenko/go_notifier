-- +goose Up
-- +goose StatementBegin
CREATE TABLE user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    email VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    timezone VARCHAR(255),
    is_active BOOLEAN DEFAULT true,
    UNIQUE (uuid),
    UNIQUE (email)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user;
-- +goose StatementEnd
