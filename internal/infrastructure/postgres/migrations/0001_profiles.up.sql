create table if not exists profiles (
                                        user_id uuid primary key,
                                        username text not null,
                                        email text not null unique,
                                        phone text not null unique,
                                        language text not null default 'ru',
                                        has_Gamification boolean not null default false,
                                        created_at timestamptz not null default now(),
                                        updated_at timestamptz null default now()
    );