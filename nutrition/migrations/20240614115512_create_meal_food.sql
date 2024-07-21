-- +goose Up
-- +goose StatementBegin
CREATE TABLE if NOT EXISTS meal_food (
    id uuid primary key default gen_random_uuid(),
    meal_id uuid not null references meal(id) on delete cascade,
    food_id uuid not null references food(id) on delete cascade,
    weight smallint null,
    created_at timestamp not null default now(),
    updated_at timestamp
)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE meal_food;
-- +goose StatementEnd
