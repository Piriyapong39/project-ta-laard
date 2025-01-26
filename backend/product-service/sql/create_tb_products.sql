-- Table: public.tb_products

-- DROP TABLE IF EXISTS public.tb_products;

CREATE TABLE IF NOT EXISTS public.tb_products
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    product_id text COLLATE pg_catalog."default" NOT NULL,
    product_name text COLLATE pg_catalog."default" NOT NULL,
    product_desc text COLLATE pg_catalog."default",
    product_price numeric(18,2) NOT NULL,
    product_stock bigint NOT NULL,
    main_cate bigint NOT NULL,
    sub_cate bigint NOT NULL,
    product_image text[] COLLATE pg_catalog."default" NOT NULL,
    user_id bigint NOT NULL,
    is_active boolean NOT NULL DEFAULT true,
    CONSTRAINT tb_products_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tb_products
    OWNER to natutitato;