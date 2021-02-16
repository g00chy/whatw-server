
-- +migrate Up
create table displays (
  id integer primary key,
  maker text,
  model text,
  size integer,
  hi integer,
  low integer
);
-- +migrate Down
drop table displays;