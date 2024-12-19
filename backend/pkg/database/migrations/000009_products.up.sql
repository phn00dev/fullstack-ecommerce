CREATE TYPE PRODUCT_STATUS AS ENUM ('ACTIVE', 'PASSIVE', 'DRAFT');

CREATE TABLE IF NOT EXISTS products
(
    id                            SERIAL PRIMARY KEY,
    product_name_tk               VARCHAR(255) NOT NULL UNIQUE,
    product_name_ru               VARCHAR(255) NOT NULL UNIQUE,
    product_name_en               VARCHAR(255) NOT NULL UNIQUE,
    product_short_description_tk  TEXT         NOT NULL,
    product_short_description_ru  TEXT         NOT NULL,
    product_short_description_en  TEXT         NOT NULL,
    product_long_description_tk   TEXT         NOT NULL,
    product_long_description_ru   TEXT         NOT NULL,
    product_long_description_en   TEXT         NOT NULL,
    product_all_specifications_tk TEXT         NOT NULL,
    product_all_specifications_ru TEXT         NOT NULL,
    product_all_specifications_en TEXT         NOT NULL,
    product_price                 FLOAT        NOT NULL,
    product_count                 INT          NOT NULL,
    product_remaining_number      INT          NOT NULL,
    product_main_image_one        VARCHAR(255) NOT NULL,
    product_main_image_two        VARCHAR(255) NOT NULL,
    product_status                PRODUCT_STATUS DEFAULT 'DRAFT',
    section_id                    INT REFERENCES sections (id) ON DELETE CASCADE,
    category_id                   INT REFERENCES categories (id) ON DELETE CASCADE,
    brand_id                      INT REFERENCES brands (id) ON DELETE CASCADE,
    created_at                    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    updated_at                    TIMESTAMP      DEFAULT CURRENT_TIMESTAMP,
    deleted_at                    TIMESTAMP      DEFAULT NULL
);