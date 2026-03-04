CREATE TABLE IF NOT EXISTS records(
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    account_id uuid NOT NULL REFERENCES accounts(id),
    time real NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now()
);
