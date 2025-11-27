CREATE TABLE topups (
    id VARCHAR(255) PRIMARY KEY,
    order_id VARCHAR(255) UNIQUE NOT NULL,
    fee_id BIGINT NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    amount BIGINT NOT NULL,
    snap_token VARCHAR(255),
    snap_token_expiry TIMESTAMP,
    status VARCHAR(50) DEFAULT 'pending',
    payment_type VARCHAR(50),
    payment_code VARCHAR(255),
    transaction_time TIMESTAMP,
    settlement_time TIMESTAMP,
    raw_response TEXT,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    INDEX idx_fee_id (fee_id),
    INDEX idx_user_id (user_id)
);

ALTER TABLE `topups` ADD CONSTRAINT `topups_user_id_fkey`
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`)
    ON DELETE CASCADE ON UPDATE CASCADE;
