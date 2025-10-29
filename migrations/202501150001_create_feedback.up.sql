CREATE TABLE IF NOT EXISTS feedback (
    id SERIAL PRIMARY KEY,
    title VARCHAR(200) NOT NULL,
    description TEXT NULL,
    type INT NOT NULL,
    tags TEXT[] NULL,
    sentiment INT NOT NULL DEFAULT 0,
    sentiment_score DECIMAL(3,2) NULL,
    votes INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


