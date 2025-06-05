DROP EXTENSION IF EXISTS "uuid-ossp";

DROP TYPE IF EXISTS device_type;
DROP TYPE IF EXISTS app_name;

DROP TABLE IF EXISTS sessions;

DROP INDEX IF EXISTS idx_sessions_user_id;
DROP INDEX IF EXISTS idx_sessions_user_active;
DROP INDEX IF EXISTS idx_sessions_expiry_expr;
DROP INDEX IF EXISTS idx_sessions_cleanup;