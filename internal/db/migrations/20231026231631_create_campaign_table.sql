-- +goose Up
-- +goose StatementBegin
CREATE TABLE campaign (
    id INT AUTO_INCREMENT PRIMARY KEY,
    uuid CHAR(36) NOT NULL,
    name VARCHAR(255) NOT NULL,
    message VARCHAR(255) NOT NULL,
    is_active BOOLEAN DEFAULT true,
    time TIME,
    days_of_week SET("1", "2", "3", "4", "5", "6", "7")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE campaign;
-- +goose StatementEnd
