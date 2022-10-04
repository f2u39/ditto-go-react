CREATE TABLE root (
  rid int,
  name varchar(20),
  PRIMARY KEY (rid)
);

CREATE TABLE main (
  rid int,
  mid int,
  name varchar(20),
  PRIMARY KEY (rid, mid)
);

CREATE TABLE sub (
  rid int,
  mid int,
  sid int,
  name varchar(20),
  PRIMARY KEY (rid, mid, sid)
);

CREATE TABLE time_rec (
  start varchar(20),
  stop varchar(20),
  duration int,
  type varchar(15),
  date varchar(8),
  created_at datetime
);

CREATE TABLE bal_rec (
  rid int,
  mid int,
  sid int,
  value int,
  description varchar(25),
  date varchar(8),
  created_at datetime
);