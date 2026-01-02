-- +goose Up
-- +goose StatementBegin
CREATE TABLE links(
    id int AUTO_INCREMENT PRIMARY KEY,
    original_url text not null,
    short_url VARCHAR(255) not null UNIQUE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS links;
-- +goose StatementEnd
