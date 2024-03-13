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

-- Pessoa
CREATE TABLE IF NOT EXISTS people (
  ID 							UUID PRIMARY KEY,
  first_name 			VARCHAR(100) NOT NULL,
  last_name 			VARCHAR(100) NOT NULL,
  email           VARCHAR(255) NOT NULL,
  phone						VARCHAR(20) DEFAULT '',
  cell_phone      VARCHAR(20) NOT NULL,
  personable_id 	UUID NOT NULL,
  personable_type VARCHAR(255) NOT NULL,
  created_at      TIMESTAMP NOT NULL,
	updated_at      TIMESTAMP NOT NULL
);

-- Endereço
CREATE TABLE IF NOT EXISTS addresses (
  ID 								UUID PRIMARY KEY,
  public_place 			VARCHAR(255) DEFAULT '',
  complement 				VARCHAR(255) DEFAULT '',
  neighborhood 			VARCHAR(255) DEFAULT '',
  city 							VARCHAR(255) DEFAULT '',
  state 						VARCHAR(255) DEFAULT '',
  zip_code 					VARCHAR(255) DEFAULT '',
  addressable_id 		UUID NOT NULL,
  addressable_type 	VARCHAR(255) NOT NULL,
  created_at       	TIMESTAMP NOT NULL,
	updated_at       	TIMESTAMP NOT NULL
);
