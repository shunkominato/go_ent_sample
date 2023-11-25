-- Modify "todos" table
ALTER TABLE "todos" ADD COLUMN "user_todos" bigint NOT NULL, ADD CONSTRAINT "todos_users_todos" FOREIGN KEY ("user_todos") REFERENCES "users" ("id") ON DELETE NO ACTION;
