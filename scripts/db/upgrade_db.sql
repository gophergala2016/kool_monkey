ALTER USER kool_writer LOGIN;
ALTER USER kool_reader LOGIN;

CREATE TABLE result(
		id SERIAL PRIMARY KEY,
		agent_id integer NOT NULL REFERENCES agent(id),
        url text NOT NULL,
        response_time bigint NOT NULL,
		timestamp timestamp NOT NULL DEFAULT now()
);

GRANT SELECT,INSERT,DELETE,UPDATE ON result TO kool_writer;
GRANT USAGE,SELECT ON SEQUENCE result_id_seq TO kool_writer;
GRANT SELECT ON result TO kool_reader;

INSERT INTO agent(ip) VALUES ('127.0.0.1');
