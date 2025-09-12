-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS targets
(
    id uuid primary key default gen_random_uuid(),
    mission_id uuid references missions(id) on delete cascade,
    name text not null,
    country text not null,
    notes text default '',
    completed boolean not null default false,
    
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS targets;
-- +goose StatementEnd
