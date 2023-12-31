-- Create "groups" table
CREATE TABLE "groups" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "name" character varying NOT NULL, PRIMARY KEY ("id"));
-- Create "users" table
CREATE TABLE "users" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "age" bigint NOT NULL, "name" character varying NOT NULL DEFAULT '-', PRIMARY KEY ("id"));
-- Create "cars" table
CREATE TABLE "cars" ("id" bigint NOT NULL GENERATED BY DEFAULT AS IDENTITY, "model" character varying NOT NULL, "registered_at" timestamptz NOT NULL, "user_cars" bigint NULL, PRIMARY KEY ("id"), CONSTRAINT "cars_users_cars" FOREIGN KEY ("user_cars") REFERENCES "users" ("id") ON DELETE SET NULL);
