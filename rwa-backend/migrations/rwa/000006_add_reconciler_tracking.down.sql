ALTER TABLE orders DROP COLUMN IF EXISTS settlement_attempts;
ALTER TABLE orders DROP COLUMN IF EXISTS last_settlement_error;
ALTER TABLE orders DROP COLUMN IF EXISTS last_settlement_attempt_at;
ALTER TABLE orders DROP COLUMN IF EXISTS needs_manual_review;