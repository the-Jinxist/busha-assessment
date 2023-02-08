CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "movie_id" varchar NOT NULL,
  "comment_ip_address" varchar NOT NULL,
  "comment" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "comments" ("movie_id");