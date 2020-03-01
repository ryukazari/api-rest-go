--
-- PostgreSQL database dump
--

-- Dumped from database version 12.2
-- Dumped by pg_dump version 12.2

-- Started on 2020-03-01 18:59:03

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

--
-- TOC entry 6 (class 2615 OID 16394)
-- Name: api_rest_v1; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA api_rest_v1;


ALTER SCHEMA api_rest_v1 OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 206 (class 1259 OID 16408)
-- Name: categoria; Type: TABLE; Schema: api_rest_v1; Owner: postgres
--

CREATE TABLE api_rest_v1.categoria (
    id_categoria integer NOT NULL,
    nombre character varying NOT NULL,
    descripcion character varying
);


ALTER TABLE api_rest_v1.categoria OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 16406)
-- Name: categoria_id_categoria_seq; Type: SEQUENCE; Schema: api_rest_v1; Owner: postgres
--

CREATE SEQUENCE api_rest_v1.categoria_id_categoria_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE api_rest_v1.categoria_id_categoria_seq OWNER TO postgres;

--
-- TOC entry 2837 (class 0 OID 0)
-- Dependencies: 205
-- Name: categoria_id_categoria_seq; Type: SEQUENCE OWNED BY; Schema: api_rest_v1; Owner: postgres
--

ALTER SEQUENCE api_rest_v1.categoria_id_categoria_seq OWNED BY api_rest_v1.categoria.id_categoria;


--
-- TOC entry 204 (class 1259 OID 16397)
-- Name: post; Type: TABLE; Schema: api_rest_v1; Owner: postgres
--

CREATE TABLE api_rest_v1.post (
    id_post integer NOT NULL,
    categoria integer NOT NULL,
    nombre character varying NOT NULL,
    descripcion character varying,
    enlace character varying NOT NULL,
    imagen character varying
);


ALTER TABLE api_rest_v1.post OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 16395)
-- Name: post_id_post_seq; Type: SEQUENCE; Schema: api_rest_v1; Owner: postgres
--

CREATE SEQUENCE api_rest_v1.post_id_post_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE api_rest_v1.post_id_post_seq OWNER TO postgres;

--
-- TOC entry 2838 (class 0 OID 0)
-- Dependencies: 203
-- Name: post_id_post_seq; Type: SEQUENCE OWNED BY; Schema: api_rest_v1; Owner: postgres
--

ALTER SEQUENCE api_rest_v1.post_id_post_seq OWNED BY api_rest_v1.post.id_post;


--
-- TOC entry 2697 (class 2604 OID 16411)
-- Name: categoria id_categoria; Type: DEFAULT; Schema: api_rest_v1; Owner: postgres
--

ALTER TABLE ONLY api_rest_v1.categoria ALTER COLUMN id_categoria SET DEFAULT nextval('api_rest_v1.categoria_id_categoria_seq'::regclass);


--
-- TOC entry 2696 (class 2604 OID 16400)
-- Name: post id_post; Type: DEFAULT; Schema: api_rest_v1; Owner: postgres
--

ALTER TABLE ONLY api_rest_v1.post ALTER COLUMN id_post SET DEFAULT nextval('api_rest_v1.post_id_post_seq'::regclass);


--
-- TOC entry 2831 (class 0 OID 16408)
-- Dependencies: 206
-- Data for Name: categoria; Type: TABLE DATA; Schema: api_rest_v1; Owner: postgres
--

COPY api_rest_v1.categoria (id_categoria, nombre, descripcion) FROM stdin;
\.


--
-- TOC entry 2829 (class 0 OID 16397)
-- Dependencies: 204
-- Data for Name: post; Type: TABLE DATA; Schema: api_rest_v1; Owner: postgres
--

COPY api_rest_v1.post (id_post, categoria, nombre, descripcion, enlace, imagen) FROM stdin;
3	1	PUT method	Probando el metodo PUT con GO	http://localhost:3000/post/3	
1	1	GET method	Probando el metodo GET con GO	http://localhost:3000/post/1	
\.


--
-- TOC entry 2839 (class 0 OID 0)
-- Dependencies: 205
-- Name: categoria_id_categoria_seq; Type: SEQUENCE SET; Schema: api_rest_v1; Owner: postgres
--

SELECT pg_catalog.setval('api_rest_v1.categoria_id_categoria_seq', 1, false);


--
-- TOC entry 2840 (class 0 OID 0)
-- Dependencies: 203
-- Name: post_id_post_seq; Type: SEQUENCE SET; Schema: api_rest_v1; Owner: postgres
--

SELECT pg_catalog.setval('api_rest_v1.post_id_post_seq', 3, true);


--
-- TOC entry 2701 (class 2606 OID 16416)
-- Name: categoria categoria_pk; Type: CONSTRAINT; Schema: api_rest_v1; Owner: postgres
--

ALTER TABLE ONLY api_rest_v1.categoria
    ADD CONSTRAINT categoria_pk PRIMARY KEY (id_categoria);


--
-- TOC entry 2699 (class 2606 OID 16405)
-- Name: post post_pk; Type: CONSTRAINT; Schema: api_rest_v1; Owner: postgres
--

ALTER TABLE ONLY api_rest_v1.post
    ADD CONSTRAINT post_pk PRIMARY KEY (id_post);


-- Completed on 2020-03-01 18:59:04

--
-- PostgreSQL database dump complete
--

