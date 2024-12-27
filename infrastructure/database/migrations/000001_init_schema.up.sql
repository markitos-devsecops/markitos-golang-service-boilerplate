CREATE TABLE boilers (
    id TEXT PRIMARY KEY UNIQUE CHECK (char_length(id) = 36),
    message TEXT NOT NULL CHECK (char_length(message) <= 200),
    createdAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updatedAt TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT message_length CHECK (char_length(message) <= 200)
);

CREATE INDEX idx_boilers_createdAt ON boilers (createdAt);
CREATE INDEX idx_boilers_message ON boilers (message);

CREATE OR REPLACE FUNCTION update_updatedAt_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updatedAt = CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_updatedAt
BEFORE UPDATE ON boilers
FOR EACH ROW
EXECUTE FUNCTION update_updatedAt_column();
