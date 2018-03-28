-- Table: public.users

-- DROP TABLE public.users;

CREATE TABLE public.users
(
    id integer NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    email character varying(50) COLLATE pg_catalog."default",
    phone character varying(11) COLLATE pg_catalog."default",
    password character varying(500) COLLATE pg_catalog."default" NOT NULL,
    state integer NOT NULL DEFAULT 0,
    createtime timestamp without time zone NOT NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.users
    OWNER to postgres;



-- Table: public.channels

-- DROP TABLE public.channels;

CREATE TABLE public.channels
(
    id integer NOT NULL DEFAULT nextval('"channels_ID_seq"'::regclass),
    name character varying(30) COLLATE pg_catalog."default" NOT NULL,
    remark text COLLATE pg_catalog."default" DEFAULT ''::text,
    createtime timestamp without time zone NOT NULL DEFAULT now(),
    sort integer DEFAULT 0,
    CONSTRAINT channels_pkey PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE public.channels
    OWNER to postgres;