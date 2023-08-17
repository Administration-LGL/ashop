
.PHONY: startdb
startdb:
	cd deploy/mysql && docker-compose up -d

.PHONY: stopdb
stopdb:
	cd deploy/mysql && docker-compose down