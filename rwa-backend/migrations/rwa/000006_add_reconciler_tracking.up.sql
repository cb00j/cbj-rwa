-- Add the tracking fields required for "on-chain settlement retry" to the orders table.
-- Used to determine: the database status has been updated, but the corresponding on-chain transaction hash is still empty (indicating that the on-chain call failed or has not been initiated yet).
ALTER TABLE orders ADD COLUMN IF NOT EXISTS settlement_attempts INT NOT NULL DEFAULT 0;
ALTER TABLE orders ADD COLUMN IF NOT EXISTS last_settlement_error TEXT;
ALTER TABLE orders ADD COLUMN IF NOT EXISTS last_settlement_attempt_at TIMESTAMPTZ;
ALTER TABLE orders ADD COLUMN IF NOT EXISTS needs_manual_review BOOLEAN NOT NULL DEFAULT false;
