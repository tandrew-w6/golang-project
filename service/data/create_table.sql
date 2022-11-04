CREATE TABLE IF NOT EXISTS m_restaurants (
    id SERIAL PRIMARY KEY NOT NULL,
    name varchar NOT NULL,
	address varchar NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    created_by varchar,
    updated_at timestamptz,
    updated_by varchar,
    deleted_at timestamptz,
    deleted_by varchar
);

CREATE TABLE IF NOT EXISTS m_dishes (
    id SERIAL PRIMARY KEY NOT NULL,
	restaurant_id integer NOT NULL,
    name varchar NOT NULL,
	price float NOT NULL,
    created_at timestamptz DEFAULT CURRENT_TIMESTAMP,
    created_by varchar,
    updated_at timestamptz,
    updated_by varchar,
    deleted_at timestamptz,
    deleted_by varchar
);

ALTER TABLE "m_dishes" ADD FOREIGN KEY ("restaurant_id") REFERENCES "m_restaurants" ("id");