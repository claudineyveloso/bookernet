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

-- Usuário do proprietário
CREATE TABLE IF NOT EXISTS owners_users (
	owner_id   UUID NOT NULL,
	user_id 	 UUID NOT NULL,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL
);

ALTER TABLE
   "owners_users"
ADD
	FOREIGN KEY (owner_id) REFERENCES owners(id);

ALTER TABLE
   "owners_users"
ADD
	FOREIGN KEY (user_id) REFERENCES users(id);

-- Cliente
CREATE TABLE IF NOT EXISTS customers (
  ID UUID PRIMARY KEY,
  birthday DATE,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL
);

-- Chamadas atendidas
CREATE TABLE IF NOT EXISTS call_terminals (
	ID 							UUID PRIMARY KEY,
	call_type 			VARCHAR(30) NOT NULL,
	origin 					VARCHAR(50) NOT NULL,
	destination 		VARCHAR(50) NOT NULL,
	name_file 			VARCHAR(255) NOT NULL,
	call_flow 			VARCHAR(1) NOT NULL,
	call_status 		VARCHAR(10) NOT NULL,
	size 						INTEGER NOT NULL,
	duration 				VARCHAR(255) NOT NULL,
	file_name 			VARCHAR(255) NOT NULL,
	etag 						VARCHAR(50) NOT NULL,
	input_date 			TIMESTAMP NOT NULL,
	beginning_call 	TIMESTAMP NOT NULL,
	closing_call 		TIMESTAMP NOT NULL,
	owner_id 				UUID NOT NULL,
	bucket_id 			UUID NOT NULL,
	created_at  		TIMESTAMP NOT NULL,
	updated_at 		 	TIMESTAMP NOT NULL
);

ALTER TABLE
   "call_terminals"
ADD
	FOREIGN KEY (owner_id) REFERENCES owners(id);

ALTER TABLE
   "call_terminals"
ADD
	FOREIGN KEY (bucket_id) REFERENCES buckets(id);

-- Intervalo de atendimento
CREATE TABLE IF NOT EXISTS intervals (
	ID 								UUID PRIMARY KEY,
	owner_id 					UUID NOT NULL,
	interval_minutes 	INTEGER NOT NULL,
	created_at  			TIMESTAMP NOT NULL,
	updated_at 		 		TIMESTAMP NOT NULL
);

ALTER TABLE
   "intervals"
ADD
  FOREIGN KEY (owner_id) REFERENCES owners(id);

-- Tipo de serviço
CREATE TABLE IF NOT EXISTS type_services (
	ID 					UUID PRIMARY KEY,
	name				VARCHAR(100) NOT NULL,
	duration		INTEGER NOT NULL,
	created_at  TIMESTAMP NOT NULL,
	updated_at 	TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX type_service_name_idx ON type_services (name);

-- Atendimento/Agenda
-- Status - Aberta - Efetivada - Cancelada - Finalizada

CREATE TABLE IF NOT EXISTS attendances (
	ID 							UUID PRIMARY KEY,
	date_service		TIMESTAMP NOT NULL,
	start_service		TIME NOT NULL,
	end_service			TIME NOT NULL,
	status					VARCHAR(1) NOT NULL,
	reminder				INTEGER NOT NULL,
	owner_id 				UUID NOT NULL,
	type_service_id   UUID NOT NULL,
	created_at  		TIMESTAMP NOT NULL,
	updated_at 		 	TIMESTAMP NOT NULL
);
ALTER TABLE
   "attendances"
ADD
  FOREIGN KEY (owner_id) REFERENCES owners(id);

ALTER TABLE
   "attendances"
ADD
  FOREIGN KEY (type_service_id) REFERENCES type_services(id);

-- Plano de saúde
-- Diurno - Vespertino - Integral
CREATE TABLE IF NOT EXISTS insurances (
	ID 					UUID PRIMARY KEY,
	name				VARCHAR(100) NOT NULL,
	period			VARCHAR(1) NOT NULL,
	created_at  TIMESTAMP NOT NULL,
	updated_at 	TIMESTAMP NOT NULL
);
CREATE UNIQUE INDEX insurances_name_idx ON insurances (name);

-- Regras
CREATE TABLE IF NOT EXISTS roles (
  ID 					UUID PRIMARY KEY,
  name 				VARCHAR(100),
  model 			VARCHAR(100),
  action 			VARCHAR(100),
  created_at  TIMESTAMP NOT NULL,
	updated_at  TIMESTAMP NOT NULL
);

-- Regras por usuário
CREATE TABLE IF NOT EXISTS roles_users (
 	ID          UUID PRIMARY KEY,
 	role_id     UUID NOT NULL,
 	user_id     UUID NOT NULL,
  created_at  TIMESTAMP NOT NULL,
	updated_at  TIMESTAMP NOT NULL
);
ALTER TABLE
   "roles_users"
ADD
	FOREIGN KEY (role_id) REFERENCES roles(id);

ALTER TABLE
   "roles_users"
ADD
	FOREIGN KEY (user_id) REFERENCES users(id);
