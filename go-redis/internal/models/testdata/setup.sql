DROP TABLE IF EXISTS public.users;

CREATE TABLE public.users (
	id serial NOT NULL,
	"name" varchar NOT NULL,
	age int2 NULL,
	-- Define unique constraint for column "name"
	CONSTRAINT users_name_key UNIQUE (name),
	CONSTRAINT users_pk PRIMARY KEY (id)
);

