create extension if not exists "uuid-ossp";

create type public.gender as enum ('Male', 'Female', 'Other');
create type public.orientation as enum ('Heterosexuality', 'Homosexuality', 'Bisexuality');
create type public.relation as enum ('Like', 'Dislike', 'Block', 'BlockedBy', 'Report', 'Connected');
create type public.event as enum ('Liked', 'LikedBack', 'Messaged', 'Unliked', 'ProfileChecked');

create table if not exists public.Tags
(
    id    uuid not null default uuid_generate_v4() primary key,
    value text not null unique
);

comment on column public.Tags.value is 'This is the actual tag value';

create table if not exists public.Users
(
    id          uuid               not null default uuid_generate_v4() primary key,
    first_name  varchar(256)       not null,
    last_name   varchar(256)       not null,
    full_name   varchar(512)       not null generated always as (first_name || ' ' || last_name) stored,
    gender      public.gender      not null default 'Other'::public.gender,
    birth_date  date               not null,
    orientation public.orientation not null default 'Bisexuality'::public.orientation,
    email       varchar(512)       not null unique,
    verified    boolean            not null default false,
    password    varchar(256)       not null check (password <> ''),
    position    point              null     default null,
    fame_rating bigint             not null default 0,
    biography   text               null     default null
);

create table if not exists public.UserTags
(
    user_id uuid not null references public.Users (id),
    tag_id  uuid not null references public.Tags (id),
    primary key (user_id, tag_id)
);

create table if not exists public.Chats
(
    id    uuid not null default uuid_generate_v4() primary key,
    user1 uuid not null references public.Users (id),
    user2 uuid not null references public.Users (id) check (user1 <> user2),

    unique (user1, user2)
);

create table if not exists public.Messages
(
    id        uuid                     not null default uuid_generate_v4() primary key,
    chat_id   uuid                     not null references public.Chats (id),
    author_id uuid                     not null references public.Users (id),
    content   text                     not null,
    date      timestamp with time zone not null default now()
);

create table if not exists public.Relationships
(
    initiator uuid            not null references public.Users (id),
    target    uuid            not null references public.Users (id) check ( target <> initiator ),
    type      public.relation not null,
    primary key (initiator, target)
);

create table if not exists public.Notifications
(
    id          uuid                     not null default uuid_generate_v4() primary key,
    origin_user uuid                     not null references public.Users (id),
    target_user uuid                     not null references public.Users (id),
    type        public.event             not null,
    seen        boolean                  not null default false,
    content     text                     null     default null
        check ((type = 'Messaged'::public.event and content is not null) or
               (type <> 'Messaged'::public.event and content is null)),
    date        timestamp with time zone not null default now()
);

comment on column public.Notifications.content is 'this is used only if the notification type is Messaged, otherwise it must be null';


create index on public.UserTags (user_id);
create index on public.UserTags (tag_id);

create index on public.Users (gender);
create index on public.Users (orientation);
create index on public.Users (fame_rating integer_ops desc);
create index on public.Users (birth_date date_ops);