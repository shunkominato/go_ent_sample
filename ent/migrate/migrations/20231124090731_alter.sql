-- Modify "cars" table
ALTER TABLE "cars" DROP CONSTRAINT "cars_users_cars", ADD COLUMN "user_id" bigint NOT NULL, ADD CONSTRAINT "cars_users_cars" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION;
