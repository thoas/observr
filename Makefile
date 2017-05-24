ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
app.conf ?= config.yml
app.test.conf ?= $(ROOT_DIR)/config.test.yml

run-server:
	@(echo "-> Run web server ...")
	@(OBSERVR_CONF=`pwd`/$(app.conf) gin -i --appPort 4001 -p 4000 run run)

bootstrap-test:
	@(psql -d postgres -c "drop database if exists observr_test;")
	@(psql -d postgres -c "create database observr_test with owner observr;")
	@(alembic -c `pwd`/alembic.test.ini upgrade bdd74fcd8a5c)

test: bootstrap-test
	OBSERVR_CONF=$(app.test.conf) go test -v $(shell go list ./... | grep -v /vendor/)
