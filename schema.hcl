table "cars" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "model" {
    null = false
    type = character_varying
  }
  column "registered_at" {
    null = false
    type = timestamptz
  }
  column "user_cars" {
    null = true
    type = bigint
  }
  primary_key {
    columns = [column.id]
  }
  foreign_key "cars_users_cars" {
    columns     = [column.user_cars]
    ref_columns = [table.users.column.id]
    on_update   = NO_ACTION
    on_delete   = SET_NULL
  }
}
table "groups" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "name" {
    null = false
    type = character_varying
  }
  primary_key {
    columns = [column.id]
  }
}
table "users" {
  schema = schema.public
  column "id" {
    null = false
    type = bigint
    identity {
      generated = BY_DEFAULT
    }
  }
  column "age" {
    null = false
    type = bigint
  }
  column "name" {
    null    = false
    type    = character_varying
    default = "-"
  }
  primary_key {
    columns = [column.id]
  }
}
schema "public" {
  comment = "standard public schema"
}
