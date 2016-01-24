TOPDIR?=$(realpath ..)
include $(TOPDIR)/Config.mk

DIRS  = 
SRCS  = postgresql.conf
SRCS += pg_hba.conf
SRCS += kool-server.conf

install:
	@echo "\\033[1;35m+++ Installing conf\\033[39;0m"
	@mkdir -p $(CONF_DIR)
	@for i in $(DIRS) ; do $(MAKE) -C $$i install ; done
	@for i in $(SRCS); do \
		cp $$i $(CONF_DIR) ; \
		perl -i -pne 's/%(\w+?)%/$$ENV{$$1}/ge' $(CONF_DIR)/$$i ; \
	done
	@echo "\\033[1;35m+++ Conf installed\\033[39;0m"
