-- +goose Up
-- +goose StatementBegin
CREATE TABLE if NOT EXISTS food (
    id uuid primary key default gen_random_uuid(),
    name text not null,
    calorie smallint null,
    proteins smallint  null,
    fats smallint null,
    carbs smallint null,
    created_at timestamp null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE food;
-- +goose StatementEnd
