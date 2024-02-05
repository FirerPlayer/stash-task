-- migrate:up
-- TABELA DE USUÁRIOS
create table public.users
(
    id         uuid      default gen_random_uuid() not null
        primary key,
    email      text                                not null
        unique,
    avatar     text,
    username   text                                not null,
    password   text                                not null,
    bio        text,
    created_at timestamp default now()             not null,
    updated_at timestamp
);

alter table public.users
    owner to micael;

-- TABELA DE TAREFAS

create table public.tasks
(
    id           uuid      default gen_random_uuid() not null,
    "user"       uuid                                not null
        references public.users,
    title        text                                not null,
    description  text,
    completed_at timestamp,
    created_at   timestamp default now()             not null,
    updated_at   timestamp,
    priority     integer   default 3
);

comment on column public.tasks.priority is 'prioridade da tarefa, quanto menor o valor mais prioridade terá, 1 é a prioridade mais alta.';

alter table public.tasks
    owner to micael;
-- migrate:down
drop table users;
drop table tasks;
