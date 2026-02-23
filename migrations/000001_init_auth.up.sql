create extension if not exists citext;
create extension if not exists pgcrypto;

CREATE TABLE users (
  id uuid primary key default gen_random_uuid(),
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),

  session_version integer not null default 0,
  disabled_at timestamptz
);

CREATE index users_disabled on users(disabled_at);

create table user_identities (
  id uuid primary key default gen_random_uuid(),
  user_id uuid not null references users(id) on delete cascade,

  provider text not null,
  identifier text not null,

  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),

  unique (provider, identifier)
);

create index user_identities_lookup_idx on user_identities(provider, identifier);

create table user_passwords (
  user_id uuid primary key references users(id) on delete cascade,
  password_hash text not null,
  updated_at timestamptz not null default now()
);

create table user_profiles (
  user_id uuid primary key references users(id) on delete cascade,

  display_name text,
  first_name text,
  last_name text,
  avatar_url text,

  updated_at timestamptz not null default now()
);

create table sessions (
  id uuid primary key default gen_random_uuid(),
  user_id uuid not null references users(id) on delete cascade,

  session_version integer not null,

  token_hash bytea not null,
  created_at timestamptz not null default now(),
  last_seen_at timestamptz not null default now(),
  expires_at timestamptz not null,
  revoked_at timestamptz,

  ip inet,
  user_agent text,
  device_id text,

  replaced_by uuid references sessions(id)
);

create unique index sessions_token_hash_uq on sessions(token_hash);
create index sessions_user_id_idx on sessions(user_id);
create index sessions_expires_at_idx on sessions(expires_at);
create index sessions_active_user_idx on sessions(user_id) where revoked_at is null;