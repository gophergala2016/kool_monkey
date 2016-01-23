GENPORTOFF?=0
genport=$(shell expr ${GENPORTOFF} + \( $(shell id -u) - \( $(shell id -u) / 100 \) \* 100 \) \* 200 + 20000 + 0${1})

TOPDIR?=$(realpath .)
PROJECT=$(TOPDIR)
RPM_DIR=$(TOPDIR)/rpm
RUN_DIR=$(TOPDIR)/dev-env
CONF_DIR=$(RUN_DIR)/conf
LOG_DIR=$(RUN_DIR)/log

SCRIPTS_DIR=$(TOPDIR)/scripts
DB_SCRIPTS_DIR=$(SCRIPTS_DIR)/db

USR_BIN=/usr/bin
USR_SBIN=/usr/sbin
RM=rm -fr
MKDIR=mkdir -p

DATABASE=monkey

PGSQL_VERSION=$(shell psql -V | awk -F' ' '{ print $$3 }' | awk -F'.' '{ if ($$2 != null) print $$1"."$$2 }')

PGSQL_HOST=$(RUN_DIR)/data
PGSQL_PORT=$(call genport,10)
PGSQL_USER=$(shell id -un)
PGSQL_PASSWD=

PGSQL_DATA=$(PGSQL_HOST)
PGSQL_DIR=$(RUN_DIR)/pgsql
PGSQL_BIN?=/usr/lib/postgresql/$(PGSQL_VERSION)/bin
PGSQL_LOGDIR=$(LOG_DIR)
PGSQL_LOG=$(PGSQL_LOGDIR)/pgsql.log
PGSQL_SCHEMA=$(DB_SCRIPTS_DIR)/create_db.sql

export
