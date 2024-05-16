CREATE TABLE patiens (
    identity_number BIGINT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(100) NOT NULL,
    birth_date VARCHAR(100) NOT NULL,
    gender VARCHAR(100) NOT NULL,
    identity_card_scan_img VARCHAR(100),
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);
