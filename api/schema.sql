-- Usuário
DROP TABLE IF EXISTS "users";
CREATE TABLE IF NOT EXISTS users (
  id          UUID PRIMARY KEY,
  email       varchar(100) not null,
  password    varchar(100) not null,
  is_active   boolean not null default true,
  user_type   varchar(50) not null,
  created_at  timestamp not null,
  updated_at  timestamp not null
);
create unique index email_idx on users (email);

-- Proprietário
CREATE TABLE owners (
  ID 							UUID PRIMARY KEY,
  people_type			VARCHAR(1) NOT NULL,
  is_active 			BOOLEAN DEFAULT TRUE,
  bucket_id 			UUID NOT NULL,
  created_at  		TIMESTAMP NOT NULL,
	updated_at 		 	TIMESTAMP NOT NULL
);

ALTER TABLE
   "owners"
ADD
   FOREIGN KEY ("bucket_id") REFERENCES "owners" ("id");
