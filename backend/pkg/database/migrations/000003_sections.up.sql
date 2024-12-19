CREATE TYPE SECTION_STATUS AS ENUM ('ACTIVE', 'PASSIVE', 'DRAFT');

CREATE TABLE IF NOT EXISTS sections
(
    id             SERIAL PRIMARY KEY,
    section_name_tk   VARCHAR(255) NOT NULL UNIQUE,
    section_name_ru   VARCHAR(255) NOT NULL UNIQUE,
    section_name_en   VARCHAR(255) NOT NULL UNIQUE,
    section_slug   VARCHAR(255) NOT NULL UNIQUE,
    section_status SECTION_STATUS DEFAULT 'DRAFT',
    section_icon   VARCHAR(255) NOT NULL,
    created_at     TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    deleted_at     TIMESTAMP      DEFAULT NULL
);