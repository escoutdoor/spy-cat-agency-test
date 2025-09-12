-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS missions
(
    id uuid primary key default gen_random_uuid(),
    cat_id uuid NULL references cats(id) on delete set null,
    completed boolean not null default false,
    
    created_at timestamptz default now(),
    updated_at timestamptz default now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS missions;
-- +goose StatementEnd
