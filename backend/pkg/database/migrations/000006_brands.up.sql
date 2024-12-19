CREATE TYPE BRAND_STATUS AS ENUM ('ACTIVE', 'PASSIVE', 'DRAFT');

CREATE TABLE IF NOT EXISTS brands
(
    id           SERIAL PRIMARY KEY,
    brand_name_tk   VARCHAR(255) NOT NULL UNIQUE,
    brand_name_ru   VARCHAR(255) NOT NULL UNIQUE,
    brand_name_en   VARCHAR(255) NOT NULL UNIQUE,
    brand_slug   VARCHAR(255) NOT NULL UNIQUE,
    brand_status BRAND_STATUS DEFAULT 'DRAFT',
    brand_icon   VARCHAR(255) NOT NULL,
    created_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    updated_at   TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    deleted_at   TIMESTAMP    DEFAULT NULL

);