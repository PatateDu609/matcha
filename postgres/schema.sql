create extension if not exists "uuid-ossp";

create type public.gender as enum ('Male', 'Female', 'Other');
create type public.orientation as enum ('Heterosexuality', 'Homosexuality', 'Bisexuality');
create type public.relation as enum ('Like', 'Dislike', 'Block', 'BlockedBy', 'Report', 'Connected');
create type public.event as enum ('Liked', 'LikedBack', 'Messaged', 'Unliked', 'ProfileChecked');
create type public.oauth_provider as enum ('google', 'discord', 'github', '42');
create domain public.rating as int8 check (value between 1 and 5);

create table if not exists public.tags
(
    id    uuid not null default uuid_generate_v4() primary key,
    value text not null unique
);

comment on column public.Tags.value is 'This is the actual tag value';

create table if not exists public.users
(
    id          uuid               not null default uuid_generate_v4() primary key,

    -- given during the registration process
    username    varchar(256)       not null unique,
    first_name  varchar(256)       not null,
    last_name   varchar(256)       not null,
    full_name   varchar(512)       not null generated always as (first_name || ' ' || last_name) stored,
    email       varchar(512)       not null unique,
    verified    boolean            not null default false,
    password    varchar(256)       not null check (password <> ''),

    -- given during the onboarding process (after registration...)
    gender      public.gender      not null default 'Other'::public.gender,
    birth_date  date               null,
    orientation public.orientation not null default 'Bisexuality'::public.orientation,
    fame_rating bigint             not null default 0,
    position    point              null     default null,
    biography   text               null     default null
);

create table if not exists public.user_oauth
(
    id            uuid                     not null default uuid_generate_v4() primary key,

    user_id       uuid                     null     default null references public.users (id),
    provider      oauth_provider           not null,
    state         uuid                     null     default uuid_generate_v4(),
    access_token  varchar(1024)            null     default null,
    refresh_token varchar(1024)            null     default null,
    expiration    timestamp with time zone null     default null,

    unique (user_id, provider),
    unique (state),
    check (
            (state is not null and (access_token is null and refresh_token is null and expiration is null)) or
            (state is null and (access_token is not null and expiration is not null))
        )
);

create table if not exists public.user_tags
(
    user_id uuid not null references public.users (id),
    tag_id  uuid not null references public.tags (id),
    primary key (user_id, tag_id)
);

create table if not exists public.rooms
(
    id    uuid not null default uuid_generate_v4() primary key,
    user1 uuid not null references public.users (id),
    user2 uuid not null references public.users (id)
        constraint check_room_same_user check (user1 <> user2),

    unique (user1, user2)
);

create table if not exists public.messages
(
    id        uuid                     not null default uuid_generate_v4() primary key,
    chat_id   uuid                     not null references public.rooms (id),
    author_id uuid                     not null references public.users (id),
    content   text                     not null,
    date      timestamp with time zone not null default now()
);

create table if not exists public.relationships
(
    initiator uuid            not null references public.users (id),
    target    uuid            not null references public.users (id) check ( target <> initiator ),
    type      public.relation not null,
    primary key (initiator, target)
);

create table if not exists public.Images
(
    owner     uuid            not null references public.Users (id),
    path      varchar(256)    not null,
    number    int             not null
);

create table if not exists public.grades
(
    initiator uuid          not null references public.users (id),
    target    uuid          not null references public.users (id) check ( target <> initiator ),
    grade     public.rating not null,
    primary key (initiator, target)
);

create table if not exists public.notifications
(
    id          uuid                     not null default uuid_generate_v4() primary key,
    origin_user uuid                     not null references public.users (id),
    target_user uuid                     not null references public.users (id),
    type        public.event             not null,
    seen        boolean                  not null default false,
    content     text                     null     default null
        check ((type = 'Messaged'::public.event and content is not null) or
               (type <> 'Messaged'::public.event and content is null)),
    date        timestamp with time zone not null default now()
);

comment on column public.notifications.content is 'this is used only if the notification type is Messaged, otherwise it must be null';


create index on public.user_tags (user_id);
create index on public.user_tags (tag_id);

create index on public.users (gender);
create index on public.users (orientation);
create index on public.users (fame_rating desc);
create index on public.users (birth_date date_ops);