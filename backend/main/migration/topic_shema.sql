CREATE TABLE IF NOT EXISTS topic(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    title VARCHAR(225) NOT NULL UNIQUE,
    description MEDIUMTEXT NOT NULL,
    imag_url MEDIUMTEXT NULL,
    first_category VARCHAR(225) NOT NULL,
    second_category VARCHAR(225) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    -- Foreign key constraint
    FOREIGN KEY fk_partner_user_id (user_id) 
        REFERENCES users(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE,
        
    FOREIGN KEY fk_partner_partner_id (partner_id) 
        REFERENCES users(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE,

    -- Indexes for performance
    INDEX idx_user_id (user_id)
    INDEX idx_partner_id (upartner_id),
    INDEX idx_created_at (created_at),
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;