CREATE TABLE IF NOT EXISTS votes (
    id SERIAL PRIMARY KEY,
    feedback_id INT NOT NULL REFERENCES feedback(id) ON DELETE CASCADE,
    user_id INT NULL,
    ip_address VARCHAR(45) NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CHECK (user_id IS NOT NULL OR ip_address IS NOT NULL)
);

-- Prevent duplicate votes from authenticated users
CREATE UNIQUE INDEX IF NOT EXISTS ux_votes_feedback_user
    ON votes (feedback_id, user_id)
    WHERE user_id IS NOT NULL;

-- Prevent duplicate votes from same IP when anonymous
CREATE UNIQUE INDEX IF NOT EXISTS ux_votes_feedback_ip
    ON votes (feedback_id, ip_address)
    WHERE ip_address IS NOT NULL;

-- Performance indexes
CREATE INDEX IF NOT EXISTS ix_votes_feedback ON votes (feedback_id);
CREATE INDEX IF NOT EXISTS ix_votes_created_at ON votes (created_at);


