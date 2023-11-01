-- +goose Up
-- +goose StatementBegin
CREATE DATABASE IF NOT EXISTS AstroRankings;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP DATABASE AstroRankings;
-- +goose StatementEnd
