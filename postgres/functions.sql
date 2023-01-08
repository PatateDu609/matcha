-- function that checks if an author is member of the targeted chat
create or replace function public.is_member_of_chat(chat_id uuid, user_id uuid) returns boolean as
$$
select exists(select 1 from public.Chats where id = chat_id and (user1 = user_id or user2 = user_id));
$$ language sql;

alter table if exists public.Messages
    add check (public.is_member_of_chat(chat_id, author_id));



create or replace procedure public.block_user(initiator uuid, target uuid)
    language sql AS
$$
begin;
begin transaction;

insert into relationships (initiator, target, type)
values (initiator, target, 'Block')
on conflict(initiator, target) do update set type='Block';

insert into relationships (initiator, target, type)
values (target, initiator, 'BlockedBy')
on conflict(target, initiator) do update set type='BlockedBy';

commit;
end;
$$;
