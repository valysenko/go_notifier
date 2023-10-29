-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_campaign (
    id INT AUTO_INCREMENT PRIMARY KEY,
    campaign_id INT,
    user_id INT,
    FOREIGN KEY (campaign_id) REFERENCES campaign(id),
    FOREIGN KEY (user_id) REFERENCES user(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE user_campaign;
-- +goose StatementEnd
