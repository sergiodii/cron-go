#!/bin/bash

set -e
set -u
# postgres -c logging_collector=on -c log_destination=stderr -c log_directory=/logs
function create_user_and_database() {
	local database=$1
	echo "  Creating user and database '$database'"
	psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
	    CREATE USER $database;
	    CREATE DATABASE $database;
	    GRANT ALL PRIVILEGES ON DATABASE $database TO $database;
EOSQL
}

if [ -n "$POSTGRES_MULTIPLE_DATABASES" ]; then
    apt-get install -y iputils-ping net-tools netcat
    echo "INSTALADOOOOOOOO"
    find /tmp/ -name .s.PGSQL.5432
    CONTADOR=0
    while [  $CONTADOR -lt 1000 ]; do
        # echo "$CONTADOR" ps -ef | grep postgres;
        # ps -ef | grep postgres
        if ($( nc -z postgrescrongo 5432 || exit 1 )) ; then
            echo "ping response succsess!!!"
        fi
        let CONTADOR=CONTADOR+1;
    done
	echo "Multiple database creation requested: $POSTGRES_MULTIPLE_DATABASES"
	for db in $(echo $POSTGRES_MULTIPLE_DATABASES | tr ',' ' '); do
		create_user_and_database $db
	done
	echo "Multiple databases created"
fi