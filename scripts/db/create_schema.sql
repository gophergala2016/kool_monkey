DROP TABLE IF EXISTS result;
DROP TABLE IF EXISTS agents;

CREATE TABLE agent(
        id SERIAL PRIMARY KEY,
        ip cidr NOT NULL,
        last_alive timestamp NOT NULL DEFAULT now()
);

GRANT SELECT,INSERT,DELETE,UPDATE ON agent TO kool_writer;
GRANT SELECT ON agent TO kool_reader;

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
