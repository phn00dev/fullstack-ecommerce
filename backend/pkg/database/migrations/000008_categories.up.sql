CREATE TYPE CATEGORY_STATUS AS ENUM ('ACTIVE', 'PASSIVE', 'DRAFT');

CREATE TABLE IF NOT EXISTS categories
(
    id                SERIAL PRIMARY KEY,
    category_name_tk  VARCHAR(255) NOT NULL UNIQUE,
    category_name_ru  VARCHAR(255) NOT NULL UNIQUE,
    category_name_en  VARCHAR(255) NOT NULL UNIQUE,
    category_slug     VARCHAR(255) NOT NULL UNIQUE,
    category_status   CATEGORY_STATUS DEFAULT 'DRAFT',
    category_icon     VARCHAR(255) NOT NULL,
    created_at        TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    updated_at        TIMESTAMP    DEFAULT CURRENT_TIMESTAMP,
    deleted_at        TIMESTAMP    DEFAULT NULL,
    section_id        INT REFERENCES sections(id) ON DELETE CASCADE
);