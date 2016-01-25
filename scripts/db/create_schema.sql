DROP TABLE IF EXISTS result;
DROP TABLE IF EXISTS testAgent;
DROP TABLE IF EXISTS test;
DROP TABLE IF EXISTS agents;

CREATE TABLE agent(
		id SERIAL PRIMARY KEY,
		ip cidr NOT NULL,
		last_alive timestamp NOT NULL DEFAULT now()
);

GRANT SELECT,INSERT,DELETE,UPDATE ON agent TO kool_writer;
GRANT SELECT ON agent TO kool_reader;
GRANT SELECT,INSERT,DELETE,UPDATE ON agent_id_seq TO kool_writer;
GRANT SELECT ON agent_id_seq TO kool_reader;

CREATE TABLE test(
		id SERIAL PRIMARY KEY,
		targetUrl VARCHAR(512) NOT NULL,
		frequency INTEGER NOT NULL default 30
);

GRANT SELECT,INSERT,DELETE,UPDATE ON test TO kool_writer;
GRANT USAGE,SELECT ON SEQUENCE test_id_seq TO kool_writer;
GRANT SELECT ON test TO kool_reader;
GRANT SELECT ON test_id_seq TO kool_reader;

CREATE TABLE testAgent(
		idAgent INTEGER NOT NULL REFERENCES agent(id),
		idTest INTEGER NOT NULL REFERENCES test(id),
		PRIMARY KEY (idAgent, idTest)
);

GRANT SELECT,INSERT,DELETE,UPDATE ON testAgent TO kool_writer;
GRANT SELECT ON testAgent TO kool_reader;

CREATE TABLE result(
		id SERIAL PRIMARY KEY,
		agent_id INTEGER NOT NULL REFERENCES agent(id),
		test_id INTEGER NOT NULL REFERENCES test(id),
		url text NOT NULL,
		response_time BIGINT NOT NULL,
		timestamp TIMESTAMP NOT NULL DEFAULT now()
);

GRANT SELECT,INSERT,DELETE,UPDATE ON result TO kool_writer;
GRANT USAGE,SELECT ON SEQUENCE result_id_seq TO kool_writer;
GRANT SELECT ON result TO kool_reader;
GRANT SELECT ON result_id_seq TO kool_reader;

INSERT INTO agent(ip) VALUES ('127.0.0.1');
