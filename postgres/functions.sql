-- function that checks if an author is member of the targeted chat
create or replace function public.is_member_of_rooms(chat_id uuid, user_id uuid) returns boolean as
$$
select exists(select 1 from public.rooms where id = chat_id and (user1 = user_id or user2 = user_id));
$$ language sql;

alter table if exists public.messages
    add check (public.is_member_of_rooms(chat_id, author_id));


-- function that forbids the creation of 2 rooms with the same users in different order
create or replace function public.check_room_exist(u1 uuid, u2 uuid) returns boolean as
$$
select not (exists(select 1 from public.rooms where user1 = u1 and user2 = u2));
$$ language sql;

alter table if exists public.rooms
    add constraint check_room_with_inverse_users check (public.check_room_exist(user2, user1));


create or replace procedure public.block_user(initiator_user uuid, target_user uuid)
    language sql as
$$
begin;
begin transaction;

insert into relationships (initiator, target, type)
values (initiator_user, target_user, 'Block')
on conflict(initiator, target) do update set type='Block';

insert into relationships (initiator, target, type)
values (target_user, initiator_user, 'BlockedBy')
on conflict(target, initiator) do update set type='BlockedBy';

commit;
end;
$$;

-- Verify the uuid with the given UUID
create or replace procedure public.verify_user(user_id uuid)
    language plpgsql as
$$
begin
    update public.users set verified= true where id = user_id;
end
$$;


-- Triggers on user_oauth

create or replace function public.nullify_oauth_token() returns trigger
as
$nullify_oauth_token$
begin

    raise notice 'Begin of nullify_oauth_token trigger';
    raise notice 'access_token = %, refresh_token = %', new.access_token, new.refresh_token;


    if new.access_token = '' then
        raise notice 'access_token is empty, setting it to null';
        new.access_token = null;
    end if;

    if new.refresh_token = '' then
        raise notice 'refresh_token is empty, setting it to null';
        new.refresh_token = null;
    end if;

    raise notice 'end of  nullify_oauth_token trigger';
    return new;
end;
$nullify_oauth_token$
    language plpgsql;

create or replace function public.manage_value_constraints() returns trigger
as
$manage_value_constraints$

begin

    raise notice 'Begin of manage_value_constraints trigger';
    raise notice 'access_token = %, expiration = %, refresh_token = %, state = %', new.access_token, new.expiration, new.refresh_token, new.state;

    if new.access_token is not null then
        raise notice 'access token is not null, nullify state';
        new.state = null;
    end if;

    if new.refresh_token <> old.refresh_token and new.access_token is null then
        raise notice 'refresh token is updated and access token is null';
        raise exception E'refresh_token shouldn\'t be updated without updating access_token';
    end if;

    if old.access_token is not null and new.access_token is null and new.state is null then
        raise notice 'access token is nullified and no state is provided';
        raise exception E'value couldn\'t be updated. access_token can\'t be nullified if not replaced by a state';
    end if;

    if new.state is not null then
        raise notice 'state is not null, nullifying other values';
        new.access_token = null;
        new.refresh_token = null;
        new.expiration = null;
    end if;

    raise notice 'end of manage_value_constraints trigger';
    return new;

end ;
$manage_value_constraints$
    language plpgsql;

create trigger uot0_nullify_oauth_token
    before update of access_token, refresh_token
    on public.user_oauth
    for each row
    when (new.access_token = '' or new.refresh_token = '')
execute function public.nullify_oauth_token();

create trigger uot1_manage_value_constraints
    before update of state, refresh_token, access_token, expiration
    on public.user_oauth
    for each row
execute function public.manage_value_constraints();