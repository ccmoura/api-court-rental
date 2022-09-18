CREATE TABLE "owner" (
  "id" uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
  "name" varchar NOT NULL,
  "email" varchar NOT NULL UNIQUE,
  "phone" varchar UNIQUE,
  "area_code" varchar,
  "password" varchar NOT NULL,
  "cpf" varchar NOT NULL UNIQUE,
  "is_confirmed" boolean NOT NULL DEFAULT FALSE,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now()),
  "deleted_at" timestamptz
);
