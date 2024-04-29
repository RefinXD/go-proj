BEGIN;
CREATE TABLE employees (
  id           BIGSERIAL PRIMARY KEY,
  name         text      NOT NULL,
  dob          date      NOT NULL,
  department   text      NOT NULL,
  job_title    text      NOT NULL,
  address      text      NOT NULL,
  join_date    date      NOT NULL,
  created_date date      NOT NULL,
  updated_date date      DEFAULT NULL,
  deleted_date date      DEFAULT NULL,
  UNIQUE(name)
);
END;