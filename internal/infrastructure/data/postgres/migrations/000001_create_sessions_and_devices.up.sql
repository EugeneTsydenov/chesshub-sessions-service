CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS pg_cron;

SELECT cron.schedule('*/5 * * * *', $$
    UPDATE sessions
    SET is_active = FALSE
    WHERE (last_active_at + lifetime) < NOW()
      AND is_active = TRUE
$$);

SELECT cron.schedule('daily_cleanup_sessions', '0 0 * * *', $$
    DELETE FROM sessions
    WHERE is_active = FALSE
      AND last_active_at < NOW() - INTERVAL '5 days';
$$);

CREATE TABLE sessions
(
    id UUID NOT NULL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    device_type SMALLINT NOT NULL,
    device_name VARCHAR(50),
    app_type SMALLINT NOT NULL,
    app_version VARCHAR(20) NOT NULL,
    os VARCHAR(20) NOT NULL,
    os_version VARCHAR(50),
    device_model VARCHAR(50),
    ip_address VARCHAR(100) NOT NULL,
    city VARCHAR(100),
    country VARCHAR(70),
    is_active BOOLEAN DEFAULT TRUE NOT NULL,
    lifetime INTERVAL NOT NULL,
    last_active_at TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    updated_at TIMESTAMP DEFAULT now() NOT NULL
);

CREATE INDEX idx_sessions_user_id ON sessions(user_id);
CREATE INDEX idx_sessions_user_active ON sessions(user_id, is_active);
CREATE INDEX idx_sessions_expiry_expr ON sessions((last_active_at + lifetime));
CREATE INDEX idx_sessions_cleanup ON sessions(is_active, updated_at);