-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_app
CHANGE token identifier VARCHAR(255) NOT NULL,
ADD type ENUM('firebase', 'sms') NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_app
    CHANGE identifier token VARCHAR(255) NULL,
    DROP COLUMN type;
-- +goose StatementEnd
