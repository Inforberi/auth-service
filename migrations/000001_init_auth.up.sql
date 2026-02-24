create extension if not exists pgcrypto;

create table users (
  id uuid primary key default gen_random_uuid(),
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),

  session_version integer not null default 0,
  disabled_at timestamptz
);

create index users_disabled_idx on users(disabled_at);

create table auth_providers (
  code text primary key,
  kind text not null check (kind in ('email', 'phone', 'oauth')),
  enabled boolean not null default false,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now()
);

insert into auth_providers (code, kind, enabled)
values
  ('email', 'email', true),
  ('phone', 'phone', false),
  ('oauth_google', 'oauth', false),
  ('oauth_apple', 'oauth', false)
on conflict (code) do nothing;

create table user_identities (
  id uuid primary key default gen_random_uuid(),
  user_id uuid not null references users(id) on delete cascade,

  provider_code text not null references auth_providers(code),
  identifier text not null,
  identifier_normalized text not null,
  is_verified boolean not null default false,
  verified_at timestamptz,

  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),

  unique (provider_code, identifier_normalized),
  check (char_length(identifier_normalized) > 0),
  check (provider_code <> 'email' or identifier_normalized = lower(identifier_normalized)),
  check (verified_at is null or is_verified = true)
);

create index user_identities_user_id_idx on user_identities(user_id);
create index user_identities_lookup_idx on user_identities(provider_code, identifier_normalized);

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
