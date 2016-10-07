
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
create table settings(
  id serial primary key,
  key varchar(255) not null,
  val bytea not null,
  flag boolean not null default false,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_settings_key on settings(key);

create table locales(
  id serial primary key,
  code varchar(255) not null,
  lang varchar(8) not null default 'en-US',
  message text not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);
create unique index idx_locales_code_lang on locales(code, lang);
create index idx_locales_code on locales(code);
create index idx_locales_lang on locales(lang);

create table notices(
  id serial primary key,
  content text not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);

create table leave_words(
  id serial primary key,
  content text not null,
  created_at timestamp without time zone not null default now(),
  updated_at timestamp without time zone not null
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

drop table leave_words;
drop table notices;
drop table locales;
drop table settings;
