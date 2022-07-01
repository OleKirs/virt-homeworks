# Домашнее задание к занятию "6.4. PostgreSQL"

## Задача 1

<details>
   <summary>Задание</summary>

Используя docker поднимите инстанс PostgreSQL (версию 13). Данные БД сохраните в volume.

Подключитесь к БД PostgreSQL используя `psql`.

Воспользуйтесь командой `\?` для вывода подсказки по имеющимся в `psql` управляющим командам.

**Найдите и приведите** управляющие команды для:
- вывода списка БД
- подключения к БД
- вывода списка таблиц
- вывода описания содержимого таблиц
- выхода из psql

</details>

<details>
<summary>Решение</summary>  
  
Используя docker поднимите инстанс PostgreSQL (версию 13). Данные БД сохраните в volume.
  
```shell
root@deb10-test50:~# docker pull postgres:13
13: Pulling from library/postgres
Digest: sha256:55e5270ce9644c62bbac55ac98cf6921e4c59ee3b21a762ea5ce9f01b6f743f6
Status: Image is up to date for postgres:13
docker.io/library/postgres:13
root@deb10-test50:~#
```
```shell
root@deb10-test50:~# docker run --rm -d \
>   --name pgsql-01 \
>   -e POSTGRES_PASSWORD=netology \
>   -v pgsql_data:/var/lib/postgresql/data \
>   -v pgsql_backup:/media/backup \
>   -p 5432:5432 \
>   postgres:13
5bd0c94161d4094b5462d881f31f93cf00afe6641fbaeb0eaffbaf8151496e0e
root@deb10-test50:~#
```
```shell
root@deb10-test50:~# docker ps -a
CONTAINER ID   IMAGE         COMMAND                  CREATED          STATUS                    PORTS                                       NAMES
5bd0c94161d4   postgres:13   "docker-entrypoint.s…"   10 seconds ago   Up 9 seconds              0.0.0.0:5432->5432/tcp, :::5432->5432/tcp   pgsql-01
dbed7f5c74fb   hello-world   "/hello"                 23 hours ago     Exited (0) 23 hours ago                                               priceless_bassi
```

Подключитесь к БД PostgreSQL используя `psql`.

```shell
root@deb10-test50:~# docker exec -it pgsql-01 bash
root@5bd0c94161d4:/# psql -U postgres
psql (13.7 (Debian 13.7-1.pgdg110+1))
Type "help" for help.

postgres=#
```

Воспользуйтесь командой `\?` для вывода подсказки по имеющимся в `psql` управляющим командам.

```shell
postgres=# \?
```
<details>
 <summary>Результат команды</summary>

```shell  
General
  \copyright             show PostgreSQL usage and distribution terms
  \crosstabview [COLUMNS] execute query and display results in crosstab
  \errverbose            show most recent error message at maximum verbosity
  \g [(OPTIONS)] [FILE]  execute query (and send results to file or |pipe);
                         \g with no arguments is equivalent to a semicolon
  \gdesc                 describe result of query, without executing it
  \gexec                 execute query, then execute each value in its result
  \gset [PREFIX]         execute query and store results in psql variables
  \gx [(OPTIONS)] [FILE] as \g, but forces expanded output mode
  \q                     quit psql
  \watch [SEC]           execute query every SEC seconds

Help
  \? [commands]          show help on backslash commands
  \? options             show help on psql command-line options
  \? variables           show help on special variables
  \h [NAME]              help on syntax of SQL commands, * for all commands

Query Buffer
  \e [FILE] [LINE]       edit the query buffer (or file) with external editor
  \ef [FUNCNAME [LINE]]  edit function definition with external editor
  \ev [VIEWNAME [LINE]]  edit view definition with external editor
  \p                     show the contents of the query buffer
  \r                     reset (clear) the query buffer
  \s [FILE]              display history or save it to file
  \w FILE                write query buffer to file

Input/Output
  \copy ...              perform SQL COPY with data stream to the client host
  \echo [-n] [STRING]    write string to standard output (-n for no newline)
  \i FILE                execute commands from file
  \ir FILE               as \i, but relative to location of current script
  \o [FILE]              send all query results to file or |pipe
  \qecho [-n] [STRING]   write string to \o output stream (-n for no newline)
  \warn [-n] [STRING]    write string to standard error (-n for no newline)

Conditional
  \if EXPR               begin conditional block
  \elif EXPR             alternative within current conditional block
  \else                  final alternative within current conditional block
  \endif                 end conditional block

Informational
  (options: S = show system objects, + = additional detail)
  \d[S+]                 list tables, views, and sequences
  \d[S+]  NAME           describe table, view, sequence, or index
  \da[S]  [PATTERN]      list aggregates
  \dA[+]  [PATTERN]      list access methods
  \dAc[+] [AMPTRN [TYPEPTRN]]  list operator classes
  \dAf[+] [AMPTRN [TYPEPTRN]]  list operator families
  \dAo[+] [AMPTRN [OPFPTRN]]   list operators of operator families
  \dAp[+] [AMPTRN [OPFPTRN]]   list support functions of operator families
  \db[+]  [PATTERN]      list tablespaces
  \dc[S+] [PATTERN]      list conversions
  \dC[+]  [PATTERN]      list casts
  \dd[S]  [PATTERN]      show object descriptions not displayed elsewhere
  \dD[S+] [PATTERN]      list domains
  \ddp    [PATTERN]      list default privileges
  \dE[S+] [PATTERN]      list foreign tables
  \det[+] [PATTERN]      list foreign tables
  \des[+] [PATTERN]      list foreign servers
  \deu[+] [PATTERN]      list user mappings
  \dew[+] [PATTERN]      list foreign-data wrappers
  \df[anptw][S+] [PATRN] list [only agg/normal/procedures/trigger/window] functions
  \dF[+]  [PATTERN]      list text search configurations
  \dFd[+] [PATTERN]      list text search dictionaries
  \dFp[+] [PATTERN]      list text search parsers
  \dFt[+] [PATTERN]      list text search templates
  \dg[S+] [PATTERN]      list roles
  \di[S+] [PATTERN]      list indexes
  \dl                    list large objects, same as \lo_list
  \dL[S+] [PATTERN]      list procedural languages
  \dm[S+] [PATTERN]      list materialized views
  \dn[S+] [PATTERN]      list schemas
  \do[S+] [PATTERN]      list operators
  \dO[S+] [PATTERN]      list collations
  \dp     [PATTERN]      list table, view, and sequence access privileges
  \dP[itn+] [PATTERN]    list [only index/table] partitioned relations [n=nested]
  \drds [PATRN1 [PATRN2]] list per-database role settings
  \dRp[+] [PATTERN]      list replication publications
  \dRs[+] [PATTERN]      list replication subscriptions
  \ds[S+] [PATTERN]      list sequences
  \dt[S+] [PATTERN]      list tables
  \dT[S+] [PATTERN]      list data types
  \du[S+] [PATTERN]      list roles
  \dv[S+] [PATTERN]      list views
  \dx[+]  [PATTERN]      list extensions
  \dy[+]  [PATTERN]      list event triggers
  \l[+]   [PATTERN]      list databases
  \sf[+]  FUNCNAME       show a function's definition
  \sv[+]  VIEWNAME       show a view's definition
  \z      [PATTERN]      same as \dp

Formatting
  \a                     toggle between unaligned and aligned output mode
  \C [STRING]            set table title, or unset if none
  \f [STRING]            show or set field separator for unaligned query output
  \H                     toggle HTML output mode (currently off)
  \pset [NAME [VALUE]]   set table output option
                         (border|columns|csv_fieldsep|expanded|fieldsep|
                         fieldsep_zero|footer|format|linestyle|null|
                         numericlocale|pager|pager_min_lines|recordsep|
                         recordsep_zero|tableattr|title|tuples_only|
                         unicode_border_linestyle|unicode_column_linestyle|
                         unicode_header_linestyle)
  \t [on|off]            show only rows (currently off)
  \T [STRING]            set HTML <table> tag attributes, or unset if none
  \x [on|off|auto]       toggle expanded output (currently off)

Connection
  \c[onnect] {[DBNAME|- USER|- HOST|- PORT|-] | conninfo}
                         connect to new database (currently "postgres")
  \conninfo              display information about current connection
  \encoding [ENCODING]   show or set client encoding
  \password [USERNAME]   securely change the password for a user

Operating System
  \cd [DIR]              change the current working directory
  \setenv NAME [VALUE]   set or unset environment variable
  \timing [on|off]       toggle timing of commands (currently off)
  \! [COMMAND]           execute command in shell or start interactive shell

Variables
  \prompt [TEXT] NAME    prompt user to set internal variable
  \set [NAME [VALUE]]    set internal variable, or list all if no parameters
  \unset NAME            unset (delete) internal variable

Large Objects
  \lo_export LOBOID FILE
  \lo_import FILE [COMMENT]
  \lo_list
  \lo_unlink LOBOID      large object operations
```
</details>

**Найдите и приведите** управляющие команды для:
- вывода списка БД
```shell
postgres=# \l+
                                                                   List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges   |  Size   | Tablespace |                Descripti
on
-----------+----------+----------+------------+------------+-----------------------+---------+------------+-------------------------
-------------------
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |                       | 7901 kB | pg_default | default administrative c
onnection database
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +| 7753 kB | pg_default | unmodifiable empty datab
ase
           |          |          |            |            | postgres=CTc/postgres |         |            |
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 | =c/postgres          +| 7753 kB | pg_default | default template for new
 databases
           |          |          |            |            | postgres=CTc/postgres |         |            |
(3 rows)
```
- подключения к БД
```shell
postgres=# \conninfo
You are connected to database "postgres" as user "postgres" via socket in "/var/run/postgresql" at port "5432".
```
- вывода списка таблиц  

```shell
postgres=# \dtS
```

<details>
 <summary>Результат команды</summary>

```shell
                    List of relations
   Schema   |          Name           | Type  |  Owner
------------+-------------------------+-------+----------
 pg_catalog | pg_aggregate            | table | postgres
 pg_catalog | pg_am                   | table | postgres
 pg_catalog | pg_amop                 | table | postgres
 pg_catalog | pg_amproc               | table | postgres
 pg_catalog | pg_attrdef              | table | postgres
 pg_catalog | pg_attribute            | table | postgres
 pg_catalog | pg_auth_members         | table | postgres
 pg_catalog | pg_authid               | table | postgres
 pg_catalog | pg_cast                 | table | postgres
 pg_catalog | pg_class                | table | postgres
 pg_catalog | pg_collation            | table | postgres
 pg_catalog | pg_constraint           | table | postgres
 pg_catalog | pg_conversion           | table | postgres
 pg_catalog | pg_database             | table | postgres
 pg_catalog | pg_db_role_setting      | table | postgres
 pg_catalog | pg_default_acl          | table | postgres
 pg_catalog | pg_depend               | table | postgres
 pg_catalog | pg_description          | table | postgres
 pg_catalog | pg_enum                 | table | postgres
 pg_catalog | pg_event_trigger        | table | postgres
 pg_catalog | pg_extension            | table | postgres
 pg_catalog | pg_foreign_data_wrapper | table | postgres
 pg_catalog | pg_foreign_server       | table | postgres
 pg_catalog | pg_foreign_table        | table | postgres
 pg_catalog | pg_index                | table | postgres
 pg_catalog | pg_inherits             | table | postgres
 pg_catalog | pg_init_privs           | table | postgres
 pg_catalog | pg_language             | table | postgres
 pg_catalog | pg_largeobject          | table | postgres
 pg_catalog | pg_largeobject_metadata | table | postgres
 pg_catalog | pg_namespace            | table | postgres
 pg_catalog | pg_opclass              | table | postgres
 pg_catalog | pg_operator             | table | postgres
 pg_catalog | pg_opfamily             | table | postgres
 pg_catalog | pg_partitioned_table    | table | postgres
 pg_catalog | pg_policy               | table | postgres
 pg_catalog | pg_proc                 | table | postgres
 pg_catalog | pg_publication          | table | postgres
 pg_catalog | pg_publication_rel      | table | postgres
 pg_catalog | pg_range                | table | postgres
 pg_catalog | pg_replication_origin   | table | postgres
 pg_catalog | pg_rewrite              | table | postgres
 pg_catalog | pg_seclabel             | table | postgres
 pg_catalog | pg_sequence             | table | postgres
 pg_catalog | pg_shdepend             | table | postgres
 pg_catalog | pg_shdescription        | table | postgres
 pg_catalog | pg_shseclabel           | table | postgres
 pg_catalog | pg_statistic            | table | postgres
 pg_catalog | pg_statistic_ext        | table | postgres
 pg_catalog | pg_statistic_ext_data   | table | postgres
 pg_catalog | pg_subscription         | table | postgres
 pg_catalog | pg_subscription_rel     | table | postgres
 pg_catalog | pg_tablespace           | table | postgres
 pg_catalog | pg_transform            | table | postgres
 pg_catalog | pg_trigger              | table | postgres
 pg_catalog | pg_ts_config            | table | postgres
 pg_catalog | pg_ts_config_map        | table | postgres
 pg_catalog | pg_ts_dict              | table | postgres
 pg_catalog | pg_ts_parser            | table | postgres
 pg_catalog | pg_ts_template          | table | postgres
 pg_catalog | pg_type                 | table | postgres
 pg_catalog | pg_user_mapping         | table | postgres
(62 rows)

```

</details>

- вывода описания содержимого таблиц

```shell
postgres=# \dS+
```

<details>
 <summary>Результат команды</summary>

```shell
                                            List of relations
   Schema   |              Name               | Type  |  Owner   | Persistence |    Size    | Description
------------+---------------------------------+-------+----------+-------------+------------+-------------
 pg_catalog | pg_aggregate                    | table | postgres | permanent   | 56 kB      |
 pg_catalog | pg_am                           | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_amop                         | table | postgres | permanent   | 80 kB      |
 pg_catalog | pg_amproc                       | table | postgres | permanent   | 64 kB      |
 pg_catalog | pg_attrdef                      | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_attribute                    | table | postgres | permanent   | 456 kB     |
 pg_catalog | pg_auth_members                 | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_authid                       | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_available_extension_versions | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_available_extensions         | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_cast                         | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_class                        | table | postgres | permanent   | 136 kB     |
 pg_catalog | pg_collation                    | table | postgres | permanent   | 240 kB     |
 pg_catalog | pg_config                       | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_constraint                   | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_conversion                   | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_cursors                      | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_database                     | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_db_role_setting              | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_default_acl                  | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_depend                       | table | postgres | permanent   | 488 kB     |
 pg_catalog | pg_description                  | table | postgres | permanent   | 376 kB     |
 pg_catalog | pg_enum                         | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_event_trigger                | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_extension                    | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_file_settings                | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_foreign_data_wrapper         | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_foreign_server               | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_foreign_table                | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_group                        | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_hba_file_rules               | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_index                        | table | postgres | permanent   | 64 kB      |
 pg_catalog | pg_indexes                      | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_inherits                     | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_init_privs                   | table | postgres | permanent   | 56 kB      |
 pg_catalog | pg_language                     | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_largeobject                  | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_largeobject_metadata         | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_locks                        | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_matviews                     | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_namespace                    | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_opclass                      | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_operator                     | table | postgres | permanent   | 144 kB     |
 pg_catalog | pg_opfamily                     | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_partitioned_table            | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_policies                     | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_policy                       | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_prepared_statements          | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_prepared_xacts               | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_proc                         | table | postgres | permanent   | 688 kB     |
 pg_catalog | pg_publication                  | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_publication_rel              | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_publication_tables           | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_range                        | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_replication_origin           | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_replication_origin_status    | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_replication_slots            | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_rewrite                      | table | postgres | permanent   | 656 kB     |
 pg_catalog | pg_roles                        | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_rules                        | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_seclabel                     | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_seclabels                    | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_sequence                     | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_sequences                    | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_settings                     | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_shadow                       | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_shdepend                     | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_shdescription                | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_shmem_allocations            | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_shseclabel                   | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_stat_activity                | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_all_indexes             | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_all_tables              | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_archiver                | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_bgwriter                | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_database                | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_database_conflicts      | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_gssapi                  | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_progress_analyze        | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_progress_basebackup     | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_progress_cluster        | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_progress_create_index   | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_progress_vacuum         | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_replication             | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_slru                    | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_ssl                     | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_subscription            | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_sys_indexes             | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_sys_tables              | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_user_functions          | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_user_indexes            | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_user_tables             | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_wal_receiver            | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_xact_all_tables         | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_xact_sys_tables         | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_xact_user_functions     | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stat_xact_user_tables        | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statio_all_indexes           | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statio_all_sequences         | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statio_all_tables            | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statio_sys_indexes           | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statio_sys_sequences         | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statio_sys_tables            | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statio_user_indexes          | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statio_user_sequences        | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statio_user_tables           | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_statistic                    | table | postgres | permanent   | 248 kB     |
 pg_catalog | pg_statistic_ext                | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_statistic_ext_data           | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_stats                        | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_stats_ext                    | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_subscription                 | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_subscription_rel             | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_tables                       | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_tablespace                   | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_timezone_abbrevs             | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_timezone_names               | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_transform                    | table | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_trigger                      | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_ts_config                    | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_ts_config_map                | table | postgres | permanent   | 56 kB      |
 pg_catalog | pg_ts_dict                      | table | postgres | permanent   | 48 kB      |
 pg_catalog | pg_ts_parser                    | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_ts_template                  | table | postgres | permanent   | 40 kB      |
 pg_catalog | pg_type                         | table | postgres | permanent   | 120 kB     |
 pg_catalog | pg_user                         | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_user_mapping                 | table | postgres | permanent   | 8192 bytes |
 pg_catalog | pg_user_mappings                | view  | postgres | permanent   | 0 bytes    |
 pg_catalog | pg_views                        | view  | postgres | permanent   | 0 bytes    |
(129 rows)
```

</details>

- выхода из psql

```shell
postgres=# \q
```

</details>



## Задача 2

<details>
   <summary>Задание</summary>

Используя `psql` создайте БД `test_database`.

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-04-postgresql/test_data).

Восстановите бэкап БД в `test_database`.

Перейдите в управляющую консоль `psql` внутри контейнера.

Подключитесь к восстановленной БД и проведите операцию ANALYZE для сбора статистики по таблице.

Используя таблицу [pg_stats](https://postgrespro.ru/docs/postgresql/12/view-pg-stats), найдите столбец таблицы `orders` 
с наибольшим средним значением размера элементов в байтах.

**Приведите в ответе** команду, которую вы использовали для вычисления и полученный результат.

</details>

<details>
<summary>Решение</summary>  

Используя `psql` создайте БД `test_database`.

```shell
root@3d2228715ed3:/# psql -U postgres
psql (13.7 (Debian 13.7-1.pgdg110+1))
Type "help" for help.

postgres=# CREATE DATABASE test_database;
CREATE DATABASE
postgres=# \q
root@3d2228715ed3:/#
```

Восстановите бэкап БД в `test_database`

```shell
root@deb10-test50:~# wget -q https://raw.githubusercontent.com/netology-code/virt-homeworks/master/06-db-04-postgresql/test_data/test_dump.sql
root@deb10-test50:~#
root@deb10-test50:~# docker cp test_dump.sql pgsql-01:/media/backup/test_dump.sql
root@deb10-test50:~#
root@deb10-test50:~# docker exec -it pgsql-01 bash
root@3d2228715ed3:/# psql -U postgres -f /media/backup/test_dump.sql  test_database
```

<details>
 <summary>Результат команды</summary>

```shell
SET
SET
SET
SET
SET
 set_config
------------

(1 row)

SET
SET
SET
SET
SET
SET
CREATE TABLE
ALTER TABLE
CREATE SEQUENCE
ALTER TABLE
ALTER SEQUENCE
ALTER TABLE
COPY 8
 setval
--------
      8
(1 row)

ALTER TABLE
root@3d2228715ed3:/#

```
</details>

Подключитесь к восстановленной БД и проведите операцию ANALYZE для сбора статистики по таблице.

```shell
root@3d2228715ed3:/# psql -U postgres
psql (13.7 (Debian 13.7-1.pgdg110+1))
Type "help" for help.

postgres=# \c test_database
You are now connected to database "test_database" as user "postgres".
test_database=#
test_database=# ANALYZE VERBOSE public.orders;
INFO:  analyzing "public.orders"
INFO:  "orders": scanned 1 of 1 pages, containing 8 live rows and 0 dead rows; 8 rows in sample, 8 estimated total rows
ANALYZE
```

Используя таблицу `pg_stats`, найдите столбец таблицы `orders` с наибольшим средним значением размера элементов в байтах.

```shell
test_database=# SELECT avg_width FROM pg_stats WHERE tablename='orders';
 avg_width
-----------
         4
        16
         4
(3 rows)
```

</details>

## Задача 3

<details>
   <summary>Задание</summary>

Архитектор и администратор БД выяснили, что ваша таблица orders разрослась до невиданных размеров и
поиск по ней занимает долгое время. Вам, как успешному выпускнику курсов DevOps в нетологии предложили
провести разбиение таблицы на 2 (шардировать на orders_1 - price>499 и orders_2 - price<=499).

Предложите SQL-транзакцию для проведения данной операции.

Можно ли было изначально исключить "ручное" разбиение при проектировании таблицы orders?

</details>

<details>
<summary>Решение</summary>  

Предложите SQL-транзакцию для проведения данной операции.

```shell
CREATE TABLE orders_1 (CHECK (price > 499)) INHERITS (orders);
INSERT INTO orders_1 SELECT * FROM orders WHERE price > 499;
CREATE TABLE orders_2 (CHECK (price <= 499)) INHERITS (orders);
INSERT INTO orders_2 SELECT * FROM orders WHERE price <= 499;
DELETE FROM ONLY orders;
```

<details>
 <summary>Результат транзакции</summary>

```shell
test_database=# CREATE TABLE orders_1 (CHECK (price > 499)) INHERITS (orders);
INSERT INTO orders_1 SELECT * FROM orders WHERE price > 499;
CREATE TABLE orders_2 (CHECK (price <= 499)) INHERITS (orders);
INSERT INTO orders_2 SELECT * FROM orders WHERE price <= 499;
DELETE FROM ONLY orders;
CREATE TABLE
INSERT 0 3
CREATE TABLE
INSERT 0 5
DELETE 8
```

Таблицы

```shell
test_database=# \dt
          List of relations
 Schema |   Name   | Type  |  Owner
--------+----------+-------+----------
 public | orders   | table | postgres
 public | orders_1 | table | postgres
 public | orders_2 | table | postgres
(3 rows)
```

</details>

Можно ли было изначально исключить "ручное" разбиение при проектировании таблицы orders?

Да, можно.
Есть вариант с партиционированием таблицы `orders` по столбцу `price` или добавлением правил для операции `Insert` 

```shell
CREATE RULE orders_ins_gt AS ON INSERT TO orders WHERE ( price > 499 ) DO INSTEAD INSERT INTO orders_1 VALUES (NEW.*);
CREATE RULE orders_ins_leq AS ON INSERT TO orders WHERE ( price <= 499 ) DO INSTEAD INSERT INTO orders_2 VALUES (NEW.*);
```


</details>

## Задача 4

<details>
   <summary>Задание</summary>

Используя утилиту `pg_dump` создайте бекап БД `test_database`.

Как бы вы доработали бэкап-файл, чтобы добавить уникальность значения столбца `title` для таблиц `test_database`?

</details>

<details>
<summary>Решение</summary>  

Используя утилиту `pg_dump` создайте бекап БД `test_database`.

```shell
test_database=# \q
root@3d2228715ed3:/# export PGPASSWORD=netology && pg_dump -h localhost -U postgres test_database > /media/backup/test_database_backup_V2.sql
root@3d2228715ed3:/# ls -l /media/backup/
total 8
-rw-r--r-- 1 root root 3503 Jul  1 06:25 test_database_backup_V2.sql
-rw-r--r-- 1 root root 2082 Jul  1 06:03 test_dump.sql
```

Как бы вы доработали бэкап-файл, чтобы добавить уникальность значения столбца `title` для таблиц `test_database`?

Например, можно изменить в секции `CREATE TABLE public.orders` строку :

```shell
    title character varying(80) NOT NULL,
```

добавив в неё требование к уникальности значения (`UNIQUE`)

```shell
    title character varying(80) UNIQUE NOT NULL,
```

</details>

---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
