-- +goose Up
-- +goose StatementBegin
CREATE TABLE if NOT EXISTS meal (
    id uuid primary key default gen_random_uuid(),
    nutrition_plan_id uuid not null references nutrition_plan(id) on delete cascade,
    name text not null,
    time timestamp not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE meal;
-- +goose StatementEnd
