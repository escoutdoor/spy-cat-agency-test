-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS cats
(
    id uuid primary key default gen_random_uuid(),
    name text not null,
    years_of_experience integer not null,
    breed text not null,
    salary decimal(12,2) not null,
    
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS cats;
-- +goose StatementEnd
