GENPORTOFF?=0
genport=$(shell expr ${GENPORTOFF} + \( $(shell id -u) - \( $(shell id -u) / 100 \) \* 100 \) \* 200 + 20000 + 0${1})

TOPDIR?=$(realpath .)

PROJECT=$(TOPDIR)
RPM_DIR=$(TOPDIR)/rpm
RUN_DIR=$(TOPDIR)/dev-env
CONF_DIR=$(RUN_DIR)/conf
BIN_DIR=$(RUN_DIR)/bin
LOG_DIR=$(RUN_DIR)/log

SCRIPTS_DIR=$(RUN_DIR)/scripts

USR_BIN=/usr/bin
USR_SBIN=/usr/sbin
RM=rm -fr
MKDIR=mkdir -p

DATABASE=monkey

PGSQL_VERSION=$(shell psql -V | awk -F' ' '{ print $$3 }' | awk -F'.' '{ if ($$2 != null) print $$1"."$$2 }')

ifdef PROD_BUILD
	DB_SCRIPTS_DIR=/opt/kool_monkey/share

	PGSQL_HOST=127.0.0.1
	PGSQL_PORT=5432
	PGSQL_USER=postgres
	PGSQL_PASSWD=
else
	DB_SCRIPTS_DIR=$(SCRIPTS_DIR)

	PGSQL_HOST=$(RUN_DIR)/data
	PGSQL_PORT=$(call genport,10)
	PGSQL_USER=$(shell id -un)
	PGSQL_PASSWD=
endif

PGSQL_DATA=$(PGSQL_HOST)
PGSQL_DIR=$(RUN_DIR)/pgsql
PGSQL_LOGDIR=$(LOG_DIR)
PGSQL_LOG=$(PGSQL_LOGDIR)/pgsql.log
PGSQL_SCHEMA=$(DB_SCRIPTS_DIR)/create_db.sql

ifeq (UNAME),Darwin)
	PGSQL_BIN?=/usr/lib/postgresql/$(PGSQL_VERSION)/bin
else ifeq (,$(wildcard /etc/redhat-release))
	PGSQL_BIN?=/usr/lib/postgresql/$(PGSQL_VERSION)/bin
else
	PGSQL_BIN?=/usr/pgsql-$(PGSQL_VERSION)/bin
endif

export
