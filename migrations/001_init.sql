create table users (
    id UUID primary key ,
    email text unique not null,
    password_hash text not null,
    created_at TIMESTAMPTZ NOT NULL
);
CREATE TABLE notifications (
    id UUID PRIMARY KEY ,
    user_id UUID NOT NULL REFERENCES users(id),
    channel TEXT NOT NULL CHECK (channel IN ('email', 'sms', 'push')),
    recipient TEXT NOT NULL,
    message TEXT NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'processing', 'sent', 'failed')),
    attempts INT NOT NULL DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL 
);
CREATE TABLE notification_logs (
    id UUID PRIMARY KEY ,
    notification_id UUID NOT NULL REFERENCES notifications(id),
    result TEXT NOT NULL CHECK (result IN ('success', 'failure')),
    error_message TEXT,
    created_at TIMESTAMPTZ NOT NULL
);

CREATE TABLE refresh_tokens (
    id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(id),
    token_hash TEXT NOT NULL,
    expires_at TIMESTAMPTZ NOT NULL,
    revoked BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMPTZ NOT NULL
);