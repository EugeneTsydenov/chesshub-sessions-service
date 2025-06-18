ALTER TABLE sessions ADD COLUMN lifetime_new INTERVAL NOT NULL DEFAULT INTERVAL '30 days';

ALTER TABLE sessions DROP COLUMN lifetime;

ALTER TABLE sessions RENAME COLUMN lifetime_new TO lifetime;