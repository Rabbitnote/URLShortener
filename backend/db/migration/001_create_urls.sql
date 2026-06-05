CREATE TABLE urls (
  id serial PRIMARY KEY,
  original_url text NOT NULL,
  short_code varchar(10) NOT NULL UNIQUE,
  created_at timestamp DEFAULT now(),
  click_count integer DEFAULT 0,
  expires_at timestamp
);