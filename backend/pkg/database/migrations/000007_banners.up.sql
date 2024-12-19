CREATE TYPE BANNER_STATUS AS ENUM ('ACTIVE', 'PASSIVE', 'DRAFT');

CREATE TABLE IF NOT EXISTS banners
(
    id            SERIAL PRIMARY KEY,
    banner_status BANNER_STATUS DEFAULT 'DRAFT',
    banner_image  VARCHAR(255) NOT NULL,
    created_at    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    updated_at    TIMESTAMP     DEFAULT CURRENT_TIMESTAMP,
    deleted_at    TIMESTAMP     DEFAULT NULL
);