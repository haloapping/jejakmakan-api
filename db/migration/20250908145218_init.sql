-- +goose Up
-- +goose StatementBegin
CREATE TABLE users(
    id TEXT NOT NULL PRIMARY KEY,
    
    profile_picture TEXT DEFAULT NULL,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE,
    fullname TEXT NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE owners(
    id TEXT NOT NULL PRIMARY KEY,

    images TEXT DEFAULT NULL,
    name TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE owner_images(
    id TEXT NOT NULL PRIMARY KEY,
    owner_id TEXT NOT NULL,

    images TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

     FOREIGN KEY(owner_id) REFERENCES owners(id)
);

CREATE TABLE locations(
    id TEXT NOT NULL PRIMARY KEY,

    district TEXT NOT NULL,
    city TEXT NOT NULL,
    province TEXT NOT NULL,
    postal_code TEXT NOT NULL,
    details TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE foods(
    id TEXT NOT NULL PRIMARY KEY,
    user_id TEXT NOT NULL,
    owner_id TEXT NOT NULL,
    location_id TEXT NOT NULL,

    images TEXT DEFAULT NULL,
    name TEXT NOT NULL,
    description TEXT,
    price BIGINT NOT NULL,
    review TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

    FOREIGN KEY(user_id) REFERENCES users(id),
    FOREIGN KEY(owner_id) REFERENCES owners(id),
    FOREIGN KEY(location_id) REFERENCES locations(id)
);

CREATE TABLE food_images(
    id TEXT NOT NULL PRIMARY KEY,
    food_id TEXT NOT NULL,

    images TEXT DEFAULT NULL,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

    FOREIGN KEY(food_id) REFERENCES foods(id) ON DELETE CASCADE
);

CREATE TABLE food_stats(
    id TEXT NOT NULL PRIMARY KEY,
    food_id TEXT NOT NULL UNIQUE,

    order_count BIGINT DEFAULT 0,
    most_order TEXT DEFAULT NULL,
    total_spend_order BIGINT DEFAULT 0,
    cheapest_order_name TEXT DEFAULT NULL,
    cheapest_order_price BIGINT DEFAULT 0, 
    most_expensive_order_name TEXT DEFAULT NULL,
    most_expensive_order_price BIGINT DEFAULT 0,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NULL,

    FOREIGN KEY(food_id) REFERENCES foods(id) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users CASCADE;
DROP TABLE IF EXISTS owners CASCADE;
DROP TABLE IF EXISTS owner_images CASCADE;
DROP TABLE IF EXISTS locations CASCADE;
DROP TABLE IF EXISTS foods CASCADE;
DROP TABLE IF EXISTS food_images CASCADE;
DROP TABLE IF EXISTS food_stats CASCADE;
-- +goose StatementEnd
