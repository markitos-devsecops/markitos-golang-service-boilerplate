DROP TRIGGER IF EXISTS set_updatedAt ON boilers;
DROP FUNCTION IF EXISTS update_updatedAt_column;

DROP INDEX IF EXISTS idx_boilers_createdAt;
DROP INDEX IF EXISTS idx_boilers_message;

DROP TABLE IF EXISTS boilers;
