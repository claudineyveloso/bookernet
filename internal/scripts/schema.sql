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
  people_type			varchar(1) not null,
  is_active       boolean not null default true,
  bucket_id 			UUID not null,
  created_at      timestamp not null,
  updated_at      timestamp not null
);

ALTER TABLE
   "owners"
ADD
   FOREIGN KEY ("bucket_id") REFERENCES "buckets" ("id");

-- Pessoa
CREATE TABLE IF NOT EXISTS people (
  ID 							UUID PRIMARY KEY,
  first_name 			varchar(100) not null,
  last_name 			varchar(100) not null,
  email           varchar(255) not null,
  phone						varchar(20) DEFAULT '',
  cell_phone      varchar(20) not null,
  personable_id 	UUID not null,
  personable_type varchar(255) not null,
  created_at      timestamp not null,
  updated_at      timestamp not null
);
create unique index first_name_idx on people (first_name);
create unique index last_name_idx on people (last_name);
create unique index email_people_idx on people (email);

-- Endereço
CREATE TABLE IF NOT EXISTS addresses (
  ID 								UUID PRIMARY KEY,
  public_place 			varchar(255) default '',
  complement 				varchar(255) default '',
  neighborhood 			varchar(255) default '',
  city 							varchar(255) default '',
  state 						varchar(255) default '',
  zip_code 					varchar(255) default '',
  addressable_id 		UUID NOT NULL,
  addressable_type 	varchar(255) not null,
  created_at        timestamp not null,
  updated_at        timestamp not null
);

-- Bucket s3
CREATE TABLE IF NOT EXISTS buckets (
	ID 										UUID PRIMARY KEY,
	description 					varchar(100) NOT NULL,
	name 									varchar(100) NOT NULL,
	aws_access_key_id 		varchar(150) NOT NULL,
	aws_secret_access_key varchar(100) NOT NULL,
	aws_region 						varchar(50) NOT NULL,
  created_at            timestamp not null,
  updated_at            timestamp not null
);
create unique index name_bucket_idx on buckets (name);
