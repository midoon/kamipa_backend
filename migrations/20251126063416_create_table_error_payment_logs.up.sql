CREATE TABLE error_payment_logs (
    id VARCHAR(255) PRIMARY KEY,
    order_id VARCHAR(255) NOT NULL,
    fee_id BIGINT NOT NULL,
    raw TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);