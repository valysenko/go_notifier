-- +goose Up
-- +goose StatementBegin
RENAME TABLE device TO user_app;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
RENAME TABLE user_app TO device;
-- +goose StatementEnd
