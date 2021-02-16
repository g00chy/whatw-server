
-- +migrate Up
create table displays (
  id integer primary key,
  maker text,
  model text,
  size integer,
  hi integer,
  low integer,
  created_at text,
  updated_at text
);
-- +migrate Down
drop table displays;