-- +goose Up
-- +goose StatementBegin
ALTER TABLE campaign
MODIFY COLUMN days_of_week SET('0', '1', '2', '3', '4', '5', '6') NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE campaign
MODIFY COLUMN days_of_week SET('1', '2', '3', '4', '5', '6', '7') NULL;
-- +goose StatementEnd
