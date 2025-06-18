ALTER TABLE sessions
    ALTER COLUMN device_type TYPE NUMERIC USING device_type::NUMERIC;

ALTER TABLE sessions
    ALTER COLUMN app_type TYPE NUMERIC USING app_type::NUMERIC;