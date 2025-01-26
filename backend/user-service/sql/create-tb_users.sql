-- Table: public.tb_users

-- DROP TABLE IF EXISTS public.tb_users;

CREATE TABLE IF NOT EXISTS public.tb_users
(
    user_id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    email text COLLATE pg_catalog."default" NOT NULL,
    password text COLLATE pg_catalog."default" NOT NULL,
    first_name text COLLATE pg_catalog."default" NOT NULL,
    last_name text COLLATE pg_catalog."default" NOT NULL,
    address text COLLATE pg_catalog."default" NOT NULL,
    is_seller boolean NOT NULL DEFAULT false,
    CONSTRAINT tb_users_pkey PRIMARY KEY (user_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tb_users
    OWNER to natutitato;

COMMENT ON COLUMN public.tb_users.is_seller
    IS 'true = seller 
false = not seller';