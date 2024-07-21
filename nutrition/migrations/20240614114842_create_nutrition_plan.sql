-- +goose Up
-- +goose StatementBegin
create table if not exists nutrition_plan (

    id uuid primary key default gen_random_uuid(),
    name text not null,
    day timestamp not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table nutrition_plan;
-- +goose StatementEnd
