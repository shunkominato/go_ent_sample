-- Modify "todos" table
ALTER TABLE "todos" DROP CONSTRAINT "todos_users_todos", ADD COLUMN "user_id" bigint NOT NULL, ADD CONSTRAINT "todos_users_todos" FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE NO ACTION;
