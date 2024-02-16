CREATE SCHEMA IF NOT EXISTS feedgram;

CREATE TABLE IF NOT EXISTS feedgram.soures (
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

CREATE INDEX 'sources__is_active' ON 'feedgram.soures' ('is_active');
CREATE INDEX 'sources__created_at' ON 'feedgram.soures' ('created_at');
CREATE INDEX 'sources__updated_at' ON 'feedgram.soures' ('updated_at');
CREATE INDEX 'sources__deleted_at' ON 'feedgram.soures' ('deleted_at');

CREATE TABLE IF NOT EXISTS feedgram.posts (
  id                BIGSERIAL     NOT NULL  PRIMARY KEY,
  source_id         BIGINT        NOT NULL,
  title             TEXT          NOT NULL,
  body              TEXT          NOT NULL,
  link              TEXT          NOT NULL,
  author            TEXT          NOT NULL,
  has_read          BOOLEAN       NOT NULL  DEFAULT false,
  additional_info   JSONB,
  created_at        TIMESTAMPTZ   NOT NULL  DEFAULT NOW(),
  posted_at         TIMESTAMPTZ   NOT NULL,
  updated_at        TIMESTAMPTZ
);

CREATE UNIQUE INDEX 'posts__link' ON 'feedgram.posts' ('link');
CREATE INDEX 'posts__has_read' ON 'feedgram.posts' ('has_read');
CREATE INDEX 'posts__created_at' ON 'feedgram.posts' ('created_at');
CREATE INDEX 'posts__posted_at' ON 'feedgram.posts' ('posted_at');

