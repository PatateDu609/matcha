-- function that checks if an author is member of the targeted chat
create or replace function public.is_member_of_rooms(chat_id uuid, user_id uuid) returns boolean as
$$
select exists(select 1 from public.Rooms where id = chat_id and (user1 = user_id or user2 = user_id));
$$ language sql;

alter table if exists public.Messages
    add check (public.is_member_of_rooms(chat_id, author_id));


-- function that forbids the creation of 2 rooms with the same users in different order
create or replace function public.check_room_exist(u1 uuid, u2 uuid) returns boolean as
$$
select not (exists(select 1 from public.Rooms where user1 = u1 and user2 = u2));
$$ language sql;

alter table if exists public.Rooms
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