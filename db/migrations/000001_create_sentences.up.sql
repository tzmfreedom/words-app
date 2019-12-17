CREATE TABLE sentences (
  id SERIAL not null,
  value varchar(1000) not null,
  created_at timestamp default CURRENT_TIMESTAMP not null,
  updated_at timestamp default CURRENT_TIMESTAMP not null,
  PRIMARY KEY (id)
);

