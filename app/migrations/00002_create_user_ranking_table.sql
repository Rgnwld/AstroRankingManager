-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS AstroRankings.userRanking (
    id varchar(64)
    ,userId varchar(64)
    ,timeInSeconds integer
    ,mapId integer
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE AstroRankings.userRanking;
-- +goose StatementEnd
