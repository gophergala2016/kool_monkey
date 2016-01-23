DROP TABLE IF EXISTS agents;

CREATE TABLE agent(
		id SERIAL PRIMARY KEY,
		ip cidr NOT NULL,
		last_alive timestamp NOT NULL DEFAULT now()
);

GRANT SELECT,INSERT,DELETE,UPDATE ON agent TO kool_writer;
GRANT SELECT ON agent TO kool_reader;
