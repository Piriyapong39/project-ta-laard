-- Table: public.tb_products_category

-- DROP TABLE IF EXISTS public.tb_products_category;

CREATE TABLE IF NOT EXISTS public.tb_products_category
(
    main_cate bigint NOT NULL,
    sub_cate bigint NOT NULL,
    name_cate text COLLATE pg_catalog."default" NOT NULL
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.tb_products_category
    OWNER to natutitato;