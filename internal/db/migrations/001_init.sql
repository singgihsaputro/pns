create extension if not exists "uuid-ossp";
create extension if not exists pgcrypto;

create table if not exists users (
    id uuid primary key default gen_random_uuid(),
    email text not null unique,
    password_hash text not null,
    created_at timestamptz not null default now()
);

create table if not exists device_tokens (
    id bigserial primary key,
    device_id text not null,
    fcm_device_token text not null,
    status text not null check (status in ('active', 'inactive')),
    created_at timestamptz not null default now(),
    unique (device_id, fcm_device_token)
);

create index if not exists idx_device_tokens_device_id on device_tokens(device_id);
create index if not exists idx_device_tokens_active on device_tokens(device_id, status);

create table if not exists promotions (
    id uuid primary key default gen_random_uuid(),
    title text not null,
    image_url text,
    description text,
    start_date date not null,
    end_date date not null,
    status text not null check (status in ('created', 'active', 'inactive')),
    created_at timestamptz not null default now()
);

