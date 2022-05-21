--
-- PostgreSQL database dump
--
-- Dumped from database version 14.2 (Debian 14.2-1.pgdg110+1)
-- Dumped by pg_dump version 14.3
SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;
SET default_tablespace = '';
SET default_table_access_method = heap;
--
-- Name: goose_db_version; Type: TABLE; Schema: public; Owner: -
--
CREATE TABLE public.goose_db_version (
    id integer NOT NULL,
    version_id bigint NOT NULL,
    is_applied boolean NOT NULL,
    tstamp timestamp without time zone DEFAULT now()
);
--
-- Name: goose_db_version_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--
CREATE SEQUENCE public.goose_db_version_id_seq AS integer START WITH 1 INCREMENT BY 1 NO MINVALUE NO MAXVALUE CACHE 1;
--
-- Name: goose_db_version_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--
ALTER SEQUENCE public.goose_db_version_id_seq OWNED BY public.goose_db_version.id;
--
-- Name: orders; Type: TABLE; Schema: public; Owner: -
--
CREATE TABLE public.orders (
    id uuid NOT NULL,
    from_addr text NOT NULL,
    to_addr text NOT NULL,
    sender_name text NOT NULL,
    sender_passport_num text NOT NULL,
    recipient_name text NOT NULL,
    weight smallint NOT NULL
);
--
-- Name: goose_db_version id; Type: DEFAULT; Schema: public; Owner: -
--
ALTER TABLE ONLY public.goose_db_version
ALTER COLUMN id
SET DEFAULT nextval('public.goose_db_version_id_seq'::regclass);
--
-- Name: goose_db_version goose_db_version_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--
ALTER TABLE ONLY public.goose_db_version
ADD CONSTRAINT goose_db_version_pkey PRIMARY KEY (id);
--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--
ALTER TABLE ONLY public.orders
ADD CONSTRAINT orders_pkey PRIMARY KEY (id);
--
-- Name: orders_sender_passport_num_inx; Type: INDEX; Schema: public; Owner: -
--
CREATE INDEX orders_sender_passport_num_inx ON public.orders USING btree (sender_passport_num);
--
-- PostgreSQL database dump complete
--