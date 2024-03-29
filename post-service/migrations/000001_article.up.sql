CREATE TABLE post (
    id UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    picture VARCHAR(100),
    title VARCHAR(100),
    article VARCHAR(64),
    owner_id UUID NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updeted_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
)