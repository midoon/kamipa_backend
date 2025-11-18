CREATE TABLE topup_logs (
    id VARCHAR(255) PRIMARY KEY,
    order_id VARCHAR(255) NOT NULL,
    event VARCHAR(100) NOT NULL,
    status VARCHAR(50),
    raw JSON NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);