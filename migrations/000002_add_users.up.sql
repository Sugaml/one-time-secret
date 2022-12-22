
CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "hashpassword" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "isactive" boolean NOT NULL DEFAULT FALSE,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
