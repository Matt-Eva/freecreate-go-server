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
-- Name: creators; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.creators (
    id bigint NOT NULL,
    uuid uuid DEFAULT uuidv7() NOT NULL,
    user_id bigint NOT NULL,
    name character varying(100) NOT NULL
);


--
-- Name: creators_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.creators ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.creators_id_seq
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
-- Name: users; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    uuid uuid DEFAULT uuidv7() NOT NULL,
    email character varying(255) NOT NULL,
    username character varying(100),
    reading_history boolean DEFAULT false,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL
);


--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: writings; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.writings (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    creator_id bigint NOT NULL,
    uuid uuid DEFAULT uuidv7(),
    title character varying(50) NOT NULL,
    subtitle character varying(50),
    description character varying(300),
    writing_type text NOT NULL,
    topics text[] DEFAULT ARRAY[]::text[] NOT NULL,
    tags text[] DEFAULT ARRAY[]::text[] NOT NULL,
    rank bigint DEFAULT 0 NOT NULL,
    rel_rank bigint DEFAULT 0 NOT NULL,
    views bigint DEFAULT 0 NOT NULL,
    list_adds bigint DEFAULT 0 NOT NULL,
    likes bigint DEFAULT 0 NOT NULL,
    lib_adds bigint DEFAULT 0 NOT NULL,
    donations bigint DEFAULT 0 NOT NULL,
    flags bigint DEFAULT 0 NOT NULL,
    rank_tracker bigint DEFAULT 0 NOT NULL,
    created_at timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    published boolean DEFAULT false NOT NULL,
    published_before boolean DEFAULT false NOT NULL,
    last_published timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    CONSTRAINT writings_tags_check CHECK ((cardinality(tags) <= 20)),
    CONSTRAINT writings_topics_check CHECK ((cardinality(topics) <= 3))
);


--
-- Name: writings_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

ALTER TABLE public.writings ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.writings_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: creators creators_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.creators
    ADD CONSTRAINT creators_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: users users_username_key; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_username_key UNIQUE (username);


--
-- Name: writings writings_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.writings
    ADD CONSTRAINT writings_pkey PRIMARY KEY (id);


--
-- Name: idx_creators_name_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX idx_creators_name_user_id ON public.creators USING btree (user_id, name);


--
-- Name: idx_creators_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_creators_user_id ON public.creators USING btree (user_id);


--
-- Name: idx_creators_uuid; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_creators_uuid ON public.creators USING btree (uuid);


--
-- Name: idx_users_uuid; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_users_uuid ON public.users USING btree (uuid);


--
-- Name: idx_writings_creator_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_writings_creator_id ON public.writings USING btree (creator_id);


--
-- Name: idx_writings_last_published; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_writings_last_published ON public.writings USING btree (last_published);


--
-- Name: idx_writings_rank; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_writings_rank ON public.writings USING btree (rank);


--
-- Name: idx_writings_rel_rank; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_writings_rel_rank ON public.writings USING btree (rel_rank);


--
-- Name: idx_writings_tags; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_writings_tags ON public.writings USING gin (tags);


--
-- Name: idx_writings_topics; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_writings_topics ON public.writings USING gin (topics);


--
-- Name: idx_writings_user_id; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_writings_user_id ON public.writings USING btree (user_id);


--
-- Name: idx_writings_uuid; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX idx_writings_uuid ON public.writings USING btree (uuid);


--
-- Name: creators creators_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.creators
    ADD CONSTRAINT creators_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- Name: writings writings_creator_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.writings
    ADD CONSTRAINT writings_creator_id_fkey FOREIGN KEY (creator_id) REFERENCES public.creators(id) ON DELETE CASCADE;


--
-- Name: writings writings_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.writings
    ADD CONSTRAINT writings_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

\unrestrict dbmate


--
-- Dbmate schema migrations
--

INSERT INTO public.schema_migrations (version) VALUES
    ('20260623031619'),
    ('20260624164610'),
    ('20260813232323');
