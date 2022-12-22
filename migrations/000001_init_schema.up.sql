
CREATE TABLE "secrets" (
  "id" bigserial PRIMARY KEY,
  "content" varchar NOT NULL,
  "creator" varchar NOT NULL,
  "hashpassword" varchar NOT NULL,
  "isview" boolean NOT NULL DEFAULT FALSE,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);
