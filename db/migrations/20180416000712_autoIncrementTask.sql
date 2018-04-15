
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
ALTER TABLE task MODIFY id INT UNSIGNED AUTO_INCREMENT NOT NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
ALTER TABLE task MODIFY id INT UNSIGNED NOT NULL;
