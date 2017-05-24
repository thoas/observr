app.conf ?= config.yml

run-server:
	@(echo "-> Run web server ...")
	@(OBSERVR_CONF=`pwd`/$(app.conf) gin -i --appPort 4001 -p 4000 run run)
