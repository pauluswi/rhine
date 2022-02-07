-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied

--CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS trx_history (
    "id" INT PRIMARY KEY,
    "trx_id" VARCHAR NOT NULL,
    "customer_id" VARCHAR NOT NULL,
    "cd" VARCHAR NOT NULL,
    "status" VARCHAR NOT NULL,
    "amount" INT NOT NULL,
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now(),
    "updated_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT now()
);

