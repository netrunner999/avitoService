create table if not exists users (
    id UUID primary key default gen_random_uuid(),
    login text not null unique,
    password varchar(64) not null, -- bcrypt-hash
    email text,
    phone text,
    is_active boolean not null default true,
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create table if not exists sessions (
    session_id text primary key,
    user_id UUID not null references users(id) on delete cascade,
    created_at timestamptz not null default now(),
    expires_at timestamptz not null
);

create table if not exists adverts (
    id UUID primary key default gen_random_uuid(),
    user_id UUID not null references users(id) on delete cascade,
    title varchar(200) not null,
    description text,
    price decimal(10,2) check (price >= 0),
    status varchar(20) not null default 'active'
        check (status in ('active', 'sold', 'cancelled')),
    created_at timestamptz not null default now(),
    completed_at timestamptz,
    updated_at timestamptz not null default now()
);
