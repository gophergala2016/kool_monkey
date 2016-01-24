GRANT SELECT ON result_id_seq TO kool_reader;

CREATE TABLE test(
		id SERIAL PRIMARY KEY,
		targetUrl VARCHAR(512) NOT NULL,
		frequency integer NOT NULL default 30
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
