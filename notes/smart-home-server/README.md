# hestia

Ultimate smart home server.

This is a work in progress and git commits will reveal each stage of implementation.

# References

- https://www.smarthomebeginner.com/docker-home-media-server-2018-basic/#OK_Great_but_why_build_a_Docker_Media_Server

# Preparation

Install docker, docker-compose in Ubuntu server 24.04.

## Edit env.sh

Specify POSTGRESQL user and password.
Specify value of PLEX_CLAIM ( visit plex website to obtain...)

and create the .env file:

```bash
./env.sh
```

# Initialize database

The first start of the postgres container initializes the database which can take some 
time. Initially one must bring up the various services individually until 
initialization is complete. Once done one can simply start and stop using docker-compose
normally.

Postgres must be started individually once to initialize internal database.
Adminer is then started individually and access to postgres tested.
Nextcloud is then started individually to allow initialization of a second
database in Postgres.

## Cleanup 

If you have attempted to start the system without doing the initialization then some 
cleanup is required.

### Stop the system

```bash
docker-compose down --rmi local -v --remove-orphans
```

### Cleanup nextcloud

```bash
docker images nextcloud

    REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
    nextcloud           15                  7f69ccc610f0        3 days ago          569MB
    nextcloud           latest              7f69ccc610f0        3 days ago          569MB

docker image rm 7f69ccc610f0

    Error response from daemon: conflict: unable to delete 7f69ccc610f0 (must be forced) - image is referenced in multiple repositories

docker image rm -f 7f69ccc610f0

    Untagged: nextcloud:15
    Untagged: nextcloud:latest
    Untagged: nextcloud@sha256:a2f2bd57fcfd92b3b6c23b6f34f65d59c9b25e7cc883b1ac67fff01229325692
    Deleted: sha256:7f69ccc610f0468f01106f40cb6b08f87b5e10a2627a765096dd5643ecef96d4
    .....
```

```bash
sudo rm -rf nextcloud/html/*
```

### Cleanup adminer

```bash
docker images adminer

    REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
    adminer             4.7                 6a66678d1bcc        2 weeks ago         81.7MB

docker image rm 6a66678d1bcc

    Untagged: adminer:4.7
    Untagged: adminer@sha256:26ed9238d44d2cc1aed8d760d2bdb7d03563f695a401ac58a929230301b18890
    Deleted: sha256:6a66678d1bcc8fc5c618283272126f578c02250d2dfe580d9f9af1753fcf1d6e
    .....
```

### Cleanup postgres

```bash
docker images postgres

    REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
    postgres            11.1-alpine         0c49db718c45        2 weeks ago         71.7MB

docker image rm 0c49db718c45

    Untagged: postgres:11.1-alpine
    Untagged: postgres@sha256:d965830b5148308278118f131e352d4c19366a7933c35a01f190b97dca1ad7d4
    Deleted: sha256:0c49db718c45f6df879711df7b88eda72e45aac8893c996665e08653959da45d
    ....
```

and remove any database files:


```bash
rm -rf postgres/data/*
```

## Initialization
 
### Initialization of postgres:

```bash
docker-compose up -d db

    Creating network "hestia_default" with the default driver
    Pulling db (postgres:11.1-alpine)...
    11.1-alpine: Pulling from library/postgres
    cd784148e348: Pull complete
    67451ade031c: Pull complete
    a9990a72e3ae: Pull complete
    cb46a6b8595e: Pull complete
    c4c5191c055e: Pull complete
    3eb34939ab82: Pull complete
    74bf1f0d5492: Pull complete
    d4fab745256a: Pull complete
    da4cf5e8ded8: Pull complete
    Creating hestia_db_1 ... done

docker-compose logs db

    Attaching to hestia_db_1
    db_1             | The files belonging to this database system will be owned by user "paul".
    db_1             | This user must also own the server process.
    db_1             | 
    db_1             | The database cluster will be initialized with locale "en_US.utf8".
    db_1             | The default database encoding has accordingly been set to "UTF8".
    db_1             | The default text search configuration will be set to "english".
    db_1             | 
    db_1             | Data page checksums are disabled.
    db_1             | 
    db_1             | fixing permissions on existing directory /var/lib/postgresql/data ... ok
    db_1             | creating subdirectories ... ok
    db_1             | selecting default max_connections ... 100
    db_1             | selecting default shared_buffers ... 128MB
    db_1             | selecting dynamic shared memory implementation ... posix
    db_1             | creating configuration files ... ok
    db_1             | running bootstrap script ... ok
    db_1             | sh: locale: not found
    db_1             | 2019-01-27 10:10:40.546 GMT [19] WARNING:  no usable system locales were found
    db_1             | performing post-bootstrap initialization ... ok
    db_1             | syncing data to disk ... ok
    db_1             | 
    db_1             | Success. You can now start the database server using:
    db_1             | 
    db_1             |     pg_ctl -D /var/lib/postgresql/data -l logfile start
    db_1             | 
    db_1             | 
    db_1             | WARNING: enabling "trust" authentication for local connections
    db_1             | You can change this by editing pg_hba.conf or using the option -A, or
    db_1             | --auth-local and --auth-host, the next time you run initdb.
    db_1             | waiting for server to start....2019-01-27 10:10:41.307 GMT [23] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
    db_1             | 2019-01-27 10:10:41.348 GMT [24] LOG:  database system was shut down at 2019-01-27 10:10:41 GMT
    db_1             | 2019-01-27 10:10:41.358 GMT [23] LOG:  database system is ready to accept connections
    db_1             |  done
    db_1             | server started
    db_1             | CREATE DATABASE
    db_1             | 
    db_1             | 
    db_1             | /usr/local/bin/docker-entrypoint.sh: ignoring /docker-entrypoint-initdb.d/*
    db_1             | 
    db_1             | waiting for server to shut down....2019-01-27 10:10:41.667 GMT [23] LOG:  received fast shutdown request
    db_1             | 2019-01-27 10:10:41.670 GMT [23] LOG:  aborting any active transactions
    db_1             | 2019-01-27 10:10:41.671 GMT [23] LOG:  background worker "logical replication launcher" (PID 30) exited with exit code 1
    db_1             | 2019-01-27 10:10:41.671 GMT [25] LOG:  shutting down
    db_1             | 2019-01-27 10:10:41.711 GMT [23] LOG:  database system is shut down
    db_1             |  done
    db_1             | server stopped
    db_1             | 
    db_1             | PostgreSQL init process complete; ready for start up.
    db_1             | 
    db_1             | 2019-01-27 10:10:41.785 GMT [1] LOG:  listening on IPv4 address "0.0.0.0", port 5432
    db_1             | 2019-01-27 10:10:41.785 GMT [1] LOG:  listening on IPv6 address "::", port 5432
    db_1             | 2019-01-27 10:10:41.801 GMT [1] LOG:  listening on Unix socket "/var/run/postgresql/.s.PGSQL.5432"
    db_1             | 2019-01-27 10:10:41.833 GMT [34] LOG:  database system was shut down at 2019-01-27 10:10:41 GMT
    db_1             | 2019-01-27 10:10:41.842 GMT [1] LOG:  database system is ready to accept connections
```

Check the log output for any errors.

### Initialization of adminer:

```bash
docker-compose up -d adminer

    Pulling adminer (adminer:4.7)...
    4.7: Pulling from library/adminer
    cd784148e348: Already exists
    d207535cd57f: Pull complete
    1167ab95319c: Pull complete
    bff34bff7f50: Pull complete
    152335cb99d2: Pull complete
    4f4bb6b6c24e: Pull complete
    16f2cd41ff90: Pull complete
    937eb7aa641a: Pull complete
    2308779f1613: Pull complete
    9321df607d85: Pull complete
    d40abead42f8: Pull complete
    40c5027dc3ef: Pull complete
    453bc81e4ff6: Pull complete
    49127472bf4e: Pull complete
    3cc4ea5dfc12: Pull complete
    f378ce480f70: Pull complete
    hestia_db_1 is up-to-date
    Creating hestia_adminer_1 ... done

docker-compose logs adminer

    Attaching to hestia_adminer_1
    adminer_1        | PHP 7.2.14 Development Server started at Sun Jan 27 10:13:18 2019
```

Point your browser to http://localhost:${PORT_ADMINER} and check that one can login and see the ${POSTGRES_DB} 
database.

- System: Postgresql
- Server: db
- Username: ${USER}
- Password: ${POSTGRES_PASSWORD}
- Database: ${POSTGRES_DB}


### Initialization of nextcloud:

```bash
docker-compose up -d nextcloud

    Pulling nextcloud (nextcloud:)...
    latest: Pulling from library/nextcloud
    5e6ec7f28fb7: Pull complete
    cf165947b5b7: Pull complete
    7bd37682846d: Pull complete
    99daf8e838e1: Pull complete
    ae320713efba: Pull complete
    ebcb99c48d8c: Pull complete
    9867e71b4ab6: Pull complete
    936eb418164a: Pull complete
    5d9617dfb66b: Pull complete
    8dd7afaae109: Pull complete
    8f207844da7e: Pull complete
    adb3ae5e4987: Pull complete
    44d7d07029db: Pull complete
    fb91064652b0: Pull complete
    50923e16d552: Pull complete
    a7cb9c70c5d2: Pull complete
    728e578e40fa: Pull complete
    4c3163d09df1: Pull complete
    842c4700643d: Pull complete
    cc1964f4bb3e: Pull complete
    125e01596295: Pull complete
    hestia_db_1 is up-to-date
    Creating hestia_nextcloud_1 ... done

docker-compose logs nextcloud

    Attaching to hestia_nextcloud_1
    nextcloud_1      | Initializing nextcloud 15.0.2.0 ...
    nextcloud_1      | Initializing finished
    nextcloud_1      | New nextcloud instance
    nextcloud_1      | Installing with PostgreSQL database
    nextcloud_1      | starting nextcloud installation
```

Point your browser to http://localhost:${PORT_ADMINER} and check that one can login and see the nextcloud
database.  

- System: Postgresql
- Server: db
- Username: ${USER}
- Password: ${POSTGRES_PASSWORD}
- Database: nextcloud

Point your browser to http://localhost:${PORT_NEXTCLOUD} and check that one can login to nextcloud.

- User: ${NEXTCLOUD_ADMIN_USER}
- Password: ${NEXTCLOUD_ADMIN_PASSWORD}

If you get an untrusted domain message (if the nextcloud service is running on a server and you are accessing from
a desktop) then one has to add to trusted_domains.

```bash
sudo vi nextcloud/html/config/config.php

    and add your server name to the trusted domains array:

        array (
          0 => 'localhost',
          1 => 'your server name',
        ),

docker-compose restart nextcloud.
```

Currently there appears to be a bug in nextcloud as the Web UI does not currently
let me login as the ADMIN user.

# Start and stop normally

```bash
docker-compose up -d 
docker-compose down --rmi local -v --remove-orphans
```

