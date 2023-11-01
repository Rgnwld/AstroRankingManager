-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS AstroRankings.userTable (
    id varchar(64)
    ,username varchar(32)
    ,hashedpassword varchar(64)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE AstroRankings.userTable;
-- +goose StatementEnd
