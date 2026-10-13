\restrict dbmate

-- Dumped from database version 18.6 (Homebrew)
-- Dumped by pg_dump version 18.6 (Homebrew)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
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
-- Name: chapter_content; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.chapter_content (
    id bigint NOT NULL,
    chapter_uuid uuid NOT NULL,
    json_content jsonb DEFAULT '{}'::jsonb NOT NULL,
    html_content text DEFAULT ''::text NOT NULL
);


--
-- Name: chapter_content_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.chapter_content ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.chapter_content_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.schema_migrations (
    version character varying NOT NULL
);


--
-- Name: chapter_content chapter_content_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.chapter_content
    ADD CONSTRAINT chapter_content_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: idx_chapter_content_chapter_uuid; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_chapter_content_chapter_uuid ON public.chapter_content USING btree (chapter_uuid);


--
-- PostgreSQL database dump complete
--

\unrestrict dbmate


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20261013211521');
