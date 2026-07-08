CREATE TABLE IF NOT EXISTS notes(
    id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    user_id BIGINT UNSIGNED NOT NULL,
    partner_id BIGINT UNSIGNED NULL, 
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    star BOOLEAN DEFAULT FALSE,
    is_partner_note BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL,

    -- Foreign key constraint
    FOREIGN KEY fk_notes_user_id (user_id) 
        REFERENCES users(id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE,
        
    FOREIGN KEY fk_notes_partner_id (partner_id) 
        REFERENCES users(id) 
        ON DELETE SET NULL 
        ON UPDATE CASCADE,

    -- Indexes for performance
    INDEX idx_user_id (user_id),
    INDEX idx_partner_id (partner_id),
    INDEX idx_star (star),
    INDEX idx_created_at (created_at),
    INDEX idx_title (title(255))
    INDEX idx_deleted_at (deleted_at)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;