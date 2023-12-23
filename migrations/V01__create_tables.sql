CREATE TABLE IF NOT EXISTS soures (
  id            BIGSERIAL     NOT NULL  PRIMARY KEY,
  url           TEXT          NOT NULL  UNIQUE,
  title         TEXT,
  link          TEXT          NOT NULL,
  description   TEXT          NOT NULL,
  is_active     BOOLEAN       NOT NULL  DEFAULT true,
  created_at    TIMESTAMPTZ   NOT NULL  DEFAULT NOW(),
  updated_at    TIMESTAMPTZ,
  deleted_at    TIMESTAMPTZ
);

CREATE INDEX "sources_is_active" ON "soures" ("is_active");
CREATE INDEX "sources_created_at" ON "soures" ("created_at");
CREATE INDEX "sources_updated_at" ON "soures" ("updated_at");
CREATE INDEX "sources_deleted_at" ON "soures" ("deleted_at");

CREATE TABLE IF NOT EXISTS posts (
  id                BIGSERIAL     NOT NULL  PRIMARY KEY,
  source_id         BIGINT        NOT NULL,
  title             TEXT          NOT NULL,
  body              TEXT          NOT NULL,
  link              TEXT          NOT NULL,
  author            TEXT          NOT NULL,
  has_readed        BOOLEAN       NOT NULL  DEFAULT false,
  additional_info   JSONB,
  posted_at         TIMESTAMPTZ   NOT NULL,
  created_at        TIMESTAMPTZ   NOT NULL  DEFAULT NOW()
);

CREATE INDEX "posts_has_readed" ON "posts" ("has_readed");
CREATE INDEX "posts_posted_at" ON "posts" ("posted_at");
CREATE INDEX "posts_created_at" ON "posts" ("created_at");

CREATE TABLE IF NOT EXISTS channels (
  channel_tg_id       BIGINT        NOT NULL PRIMARY KEY,
  added_by_tg_id      BIGINT        NOT NULL,
  sources             BIGINT[]      NOT NULL,
  title               TEXT          NOT NULL,
  is_active           BOOLEAN       NOT NULL DEFAULT false,
  created_at          TIMESTAMPTZ   NOT NUlL DEFAULT NOW(),
  updated_at          TIMESTAMPTZ,
  deleted_at          TIMESTAMPTZ
);

CREATE INDEX "channels_created_at" ON "channels" ("created_at");
CREATE INDEX "channels_updated_at" ON "channels" ("updated_at");
CREATE INDEX "channels_deleted_at" ON "channels" ("deleted_at");