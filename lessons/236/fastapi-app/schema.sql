CREATE TABLE IF NOT EXISTS python_device (
    id SERIAL PRIMARY KEY,
    uuid UUID DEFAULT NULL,
    mac VARCHAR(255) DEFAULT NULL,
    firmware VARCHAR(255) DEFAULT NULL,
    created_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP,
        updated_at TIMESTAMP
    WITH
        TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_device_uuid ON python_device (uuid);