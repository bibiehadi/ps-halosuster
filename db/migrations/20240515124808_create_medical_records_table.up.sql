CREATE TABLE medical_records (
    id SERIAL PRIMARY KEY,
    patien_id BIGINT NOT NULL,
    sympthoms VARCHAR(100) NOT NULL,
    medications VARCHAR(100) NOT NULL,
    created_by INT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (patien_id) REFERENCES patiens(identity_number) ON DELETE CASCADE,
    FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);
