# Домашнее задание к занятию "6.2. SQL"

## Введение

Перед выполнением задания вы можете ознакомиться с 
[дополнительными материалами](https://github.com/netology-code/virt-homeworks/tree/master/additional/README.md).

## Задача 1
<details>
<summary>Задание</summary>
Используя docker поднимите инстанс PostgreSQL (версию 12) c 2 volume, 
в который будут складываться данные БД и бэкапы.

Приведите получившуюся команду или docker-compose манифест.
</details>


<details>
<summary>Решение</summary>  

Подготовим манифест в файле `docker-compose.yml`
```shell
root@deb10-test50:~# cat ./docker-compose.yml
```

```yaml
version: '3.6'

volumes:
  data: {}
  backup: {}

services:

  postgres:
    image: postgres:12
    container_name: pg-sql01
    ports:
      - "0.0.0.0:5432:5432"
    volumes:
      - data:/var/lib/postgresql/data
      - backup:/media/postgresql/backup
    environment:
      POSTGRES_USER: "test-admin-user"
      POSTGRES_PASSWORD: "testpwd777"
      POSTGRES_DB: "test_db"
    restart: always
```

Запустим docker-compose  

```shell
root@deb10-test50:~# docker-compose up -d
```

```shell
Creating network "root_default" with the default driver
Creating volume "root_data" with default driver
Creating volume "root_backup" with default driver
Pulling postgres (postgres:12)...
12: Pulling from library/postgres
b85a868b505f: Pull complete
b53bada42f30: Pull complete
303bde9620f5: Pull complete
5c32c0c0a1b9: Pull complete
302630a57c06: Pull complete
ddfead4dfb39: Pull complete
03d9917b9309: Pull complete
4bb0d8ea11e0: Pull complete
18b4b6185066: Pull complete
9e8d0d57b0f9: Pull complete
68987ba225b7: Pull complete
297860b39beb: Pull complete
ee6b41e09bf5: Pull complete
Digest: sha256:e6ffad42c91a4d5a29257a27ac4e160c3ae7196696b37bf2e80410024ed95951
Status: Downloaded newer image for postgres:12
Creating pg-sql01 ... done
```

</details>

## Задача 2

<details>
   <summary>Задание</summary>

В БД из задачи 1: 
- создайте пользователя test-admin-user и БД test_db
- в БД test_db создайте таблицу orders и clients (спeцификация таблиц ниже)
- предоставьте привилегии на все операции пользователю test-admin-user на таблицы БД test_db
- создайте пользователя test-simple-user  
- предоставьте пользователю test-simple-user права на SELECT/INSERT/UPDATE/DELETE данных таблиц БД test_db

Таблица orders:
- id (serial primary key)
- наименование (string)
- цена (integer)

Таблица clients:
- id (serial primary key)
- фамилия (string)
- страна проживания (string, index)
- заказ (foreign key orders)

Приведите:
- итоговый список БД после выполнения пунктов выше,
- описание таблиц (describe)
- SQL-запрос для выдачи списка пользователей с правами над таблицами test_db
- список пользователей с правами над таблицами test_db
</details>


<details>
<summary>Решение</summary> 

Подключимся к `bash` внутри контейнера  

```shell
root@deb10-test50:~# docker exec -it pg-sql01 bash
root@3547d3cbfc1d:/#
```

Создадим пользователя `test-admin-user` и БД `test_db`  

```shell
root@3547d3cbfc1d:/# export PGPASSWORD=testpwd777 && psql -h localhost -U test-admin-user test_db
psql (12.11 (Debian 12.11-1.pgdg110+1))
Type "help" for help.

test_db=#
```
Создадим указанные в задании таблицы и пользователей и назначим им права:  

```jql
CREATE TABLE orders (
    id SERIAL,
    наименование VARCHAR, 
    цена INTEGER,
    PRIMARY KEY (id)
);

CREATE TABLE clients (
    id SERIAL,
    фамилия VARCHAR,
    "страна проживания" VARCHAR, 
    заказ INTEGER,
    PRIMARY KEY (id),
    CONSTRAINT fk_заказ  
    FOREIGN KEY(заказ)  
    REFERENCES orders(id)  
);

CREATE INDEX ON clients("страна проживания");

GRANT ALL ON TABLE orders, clients TO "test-admin-user";

CREATE USER "test-simple-user" WITH PASSWORD 'netology';

GRANT CONNECT ON DATABASE test_db TO "test-simple-user";
GRANT USAGE ON SCHEMA public TO "test-simple-user";
GRANT SELECT, INSERT, UPDATE, DELETE ON orders, clients TO "test-simple-user";
```

Результат выполнения

```shell
CREATE TABLE
CREATE TABLE
CREATE INDEX
GRANT
CREATE ROLE
GRANT
GRANT
GRANT
```

итоговый список БД после выполнения пунктов выше,

```commandline
test_db=# \l+
                                                                               List of databases
   Name    |      Owner      | Encoding |  Collate   |   Ctype    |            Access privileges            |  Size   | Tablespace |
                Description
-----------+-----------------+----------+------------+------------+-----------------------------------------+---------+------------+
--------------------------------------------
 postgres  | test-admin-user | UTF8     | en_US.utf8 | en_US.utf8 |                                         | 7969 kB | pg_default |
 default administrative connection database
 template0 | test-admin-user | UTF8     | en_US.utf8 | en_US.utf8 | =c/"test-admin-user"                   +| 7825 kB | pg_default |
 unmodifiable empty database
           |                 |          |            |            | "test-admin-user"=CTc/"test-admin-user" |         |            |

 template1 | test-admin-user | UTF8     | en_US.utf8 | en_US.utf8 | =c/"test-admin-user"                   +| 7825 kB | pg_default |
 default template for new databases
           |                 |          |            |            | "test-admin-user"=CTc/"test-admin-user" |         |            |

 test_db   | test-admin-user | UTF8     | en_US.utf8 | en_US.utf8 | =Tc/"test-admin-user"                  +| 8121 kB | pg_default |

           |                 |          |            |            | "test-admin-user"=CTc/"test-admin-user"+|         |            |

           |                 |          |            |            | "test-simple-user"=c/"test-admin-user"  |         |            |

(4 rows)
```

описание таблиц (describe)  

```commandline
test_db=# \d+ clients
                                                           Table "public.clients"
      Column       |       Type        | Collation | Nullable |               Default               | Storage  | Stats target | Desc
ription
-------------------+-------------------+-----------+----------+-------------------------------------+----------+--------------+-----
--------
 id                | integer           |           | not null | nextval('clients_id_seq'::regclass) | plain    |              |
 фамилия           | character varying |           |          |                                     | extended |              |
 страна проживания | character varying |           |          |                                     | extended |              |
 заказ             | integer           |           |          |                                     | plain    |              |
Indexes:
    "clients_pkey" PRIMARY KEY, btree (id)
    "clients_страна проживания_idx" btree ("страна проживания")
Foreign-key constraints:
    "fk_заказ" FOREIGN KEY ("заказ") REFERENCES orders(id)
Access method: heap

test_db=# test_db=# \d+ orders
                                                        Table "public.orders"
    Column    |       Type        | Collation | Nullable |              Default               | Storage  | Stats target | Descriptio
n
--------------+-------------------+-----------+----------+------------------------------------+----------+--------------+-----------
--
 id           | integer           |           | not null | nextval('orders_id_seq'::regclass) | plain    |              |
 наименование | character varying |           |          |                                    | extended |              |
 цена         | integer           |           |          |                                    | plain    |              |
Indexes:
    "orders_pkey" PRIMARY KEY, btree (id)
Referenced by:
    TABLE "clients" CONSTRAINT "fk_заказ" FOREIGN KEY ("заказ") REFERENCES orders(id)
Access method: heap
```

SQL-запрос для выдачи списка пользователей с правами над таблицами test_db  

```commandline
SELECT grantee, table_name, privilege_type 
  FROM information_schema.table_privileges 
  WHERE grantee in ('test-admin-user','test-simple-user')
        and table_name in ('clients','orders')
  order by 1,2,3
; 
```

список пользователей с правами над таблицами test_db  

```commandline
     grantee      | table_name | privilege_type
------------------+------------+----------------
 test-admin-user  | clients    | DELETE
 test-admin-user  | clients    | INSERT
 test-admin-user  | clients    | REFERENCES
 test-admin-user  | clients    | SELECT
 test-admin-user  | clients    | TRIGGER
 test-admin-user  | clients    | TRUNCATE
 test-admin-user  | clients    | UPDATE
 test-admin-user  | orders     | DELETE
 test-admin-user  | orders     | INSERT
 test-admin-user  | orders     | REFERENCES
 test-admin-user  | orders     | SELECT
 test-admin-user  | orders     | TRIGGER
 test-admin-user  | orders     | TRUNCATE
 test-admin-user  | orders     | UPDATE
 test-simple-user | clients    | DELETE
 test-simple-user | clients    | INSERT
 test-simple-user | clients    | SELECT
 test-simple-user | clients    | UPDATE
 test-simple-user | orders     | DELETE
 test-simple-user | orders     | INSERT
 test-simple-user | orders     | SELECT
 test-simple-user | orders     | UPDATE
(22 rows)
```

</details>  

## Задача 3

<details>
   <summary>Задание</summary>

Используя SQL синтаксис - наполните таблицы следующими тестовыми данными:

Таблица orders

|Наименование|цена|
|------------|----|
|Шоколад| 10 |
|Принтер| 3000 |
|Книга| 500 |
|Монитор| 7000|
|Гитара| 4000|

Таблица clients

|ФИО|Страна проживания|
|------------|----|
|Иванов Иван Иванович| USA |
|Петров Петр Петрович| Canada |
|Иоганн Себастьян Бах| Japan |
|Ронни Джеймс Дио| Russia|
|Ritchie Blackmore| Russia|

Используя SQL синтаксис:
- вычислите количество записей для каждой таблицы 
- приведите в ответе:
    - запросы 
    - результаты их выполнения.
</details>

<details>
<summary>Решение</summary>  

Используя SQL синтаксис - наполните таблицы ... тестовыми данными:

Запрос к `orders`
```commandline
INSERT INTO orders 
   VALUES 
     (1, 'Шоколад', 10), 
	 (2, 'Принтер', 3000), 
	 (3, 'Книга', 500), 
	 (4, 'Монитор', 7000), 
	 (5, 'Гитара', 4000)
;
```

Результат выполнения  

```commandline
test_db=# INSERT INTO orders
   VALUES
     (1, 'Шоколад', 10),
         (2, 'Принтер', 3000),
         (3, 'Книга', 500),
         (4, 'Монитор', 7000),
         (5, 'Гитара', 4000)
;
INSERT 0 5
```

Запрос к `clients`

```commandline
INSERT INTO clients 
   VALUES 
    (1, 'Иванов Иван Иванович', 'USA'), 
	(2, 'Петров Петр Петрович', 'Canada'), 
	(3, 'Иоганн Себастьян Бах', 'Japan'), 
	(4, 'Ронни Джеймс Дио', 'Russia'), 
	(5, 'Ritchie Blackmore', 'Russia')
;
```

Результат выполнения  

```commandline
test_db=# INSERT INTO clients
   VALUES
    (1, 'Иванов Иван Иванович', 'USA'),
        (2, 'Петров Петр Петрович', 'Canada'),
        (3, 'Иоганн Себастьян Бах', 'Japan'),
        (4, 'Ронни Джеймс Дио', 'Russia'),
        (5, 'Ritchie Blackmore', 'Russia')
;
INSERT 0 5
```
Количество записей для каждой таблицы

```commandline
test_db=# SELECT * FROM orders;
 id | наименование | цена
----+--------------+------
  1 | Шоколад      |   10
  2 | Принтер      | 3000
  3 | Книга        |  500
  4 | Монитор      | 7000
  5 | Гитара       | 4000
(5 rows)

test_db=# SELECT count(1) FROM orders;
 count
-------
     5
(1 row)
```

```commandline
test_db=# SELECT * FROM clients;
 id |       фамилия        | страна проживания | заказ
----+----------------------+-------------------+-------
  1 | Иванов Иван Иванович | USA               |
  2 | Петров Петр Петрович | Canada            |
  3 | Иоганн Себастьян Бах | Japan             |
  4 | Ронни Джеймс Дио     | Russia            |
  5 | Ritchie Blackmore    | Russia            |
(5 rows)

test_db=# SELECT count(1) FROM clients;
 count
-------
     5
(1 row)
```

</details>

## Задача 4

<details>
   <summary>Задание</summary>

Часть пользователей из таблицы clients решили оформить заказы из таблицы orders.

Используя foreign keys свяжите записи из таблиц, согласно таблице:

|ФИО|Заказ|
|------------|----|
|Иванов Иван Иванович| Книга |
|Петров Петр Петрович| Монитор |
|Иоганн Себастьян Бах| Гитара |

Приведите SQL-запросы для выполнения данных операций.

Приведите SQL-запрос для выдачи всех пользователей, которые совершили заказ, а также вывод данного запроса.
 
Подсказк - используйте директиву `UPDATE`.
</details>


<details>
<summary>Решение</summary>  

Запрос

```commandline
UPDATE clients 
    SET "заказ" = (SELECT id FROM orders WHERE "наименование"='Книга') 
    WHERE "фамилия"='Иванов Иван Иванович';
UPDATE clients 
    SET "заказ" = (SELECT id FROM orders WHERE "наименование"='Монитор')
    WHERE "фамилия"='Петров Петр Петрович';
UPDATE clients 
    SET "заказ" = (SELECT id FROM orders WHERE "наименование"='Гитара')
    WHERE "фамилия"='Иоганн Себастьян Бах';
```

Результат выполнения  

```commandline
test_db=# UPDATE clients
    SET "заказ" = (SELECT id FROM orders WHERE "наименование"='Книга')
    WHERE "фамилия"='Иванов Иван Иванович';
UPDATE clients
    SET "заказ" = (SELECT id FROM orders WHERE "наименование"='Монитор')
    WHERE "фамилия"='Петров Петр Петрович';
UPDATE clients
    SET "заказ" = (SELECT id FROM orders WHERE "наименование"='Гитара')
    WHERE "фамилия"='Иоганн Себастьян Бах';
UPDATE 1
UPDATE 1
UPDATE 1
```

Запрос для выдачи всех пользователей, которые совершили заказ

```commandline
SELECT c.* FROM clients c JOIN orders o ON c.заказ = o.id;
```

Результат

```commandline
test_db=# SELECT c.* FROM clients c JOIN orders o ON c.заказ = o.id;
 id |       фамилия        | страна проживания | заказ
----+----------------------+-------------------+-------
  1 | Иванов Иван Иванович | USA               |     3
  2 | Петров Петр Петрович | Canada            |     4
  3 | Иоганн Себастьян Бах | Japan             |     5
(3 rows)

```

</details> 


## Задача 5

<details>
   <summary>Задание</summary>

Получите полную информацию по выполнению запроса выдачи всех пользователей из задачи 4 
(используя директиву EXPLAIN).

Приведите получившийся результат и объясните что значат полученные значения.
</details>


<details>
<summary>Решение</summary>  

Запрос  

```commandline
EXPLAIN SELECT c.* FROM clients c JOIN orders o ON c.заказ = o.id;
```
Результат выполнения  

```commandline
test_db=# EXPLAIN SELECT c.* FROM clients c JOIN orders o ON c.заказ = o.id;
                               QUERY PLAN
------------------------------------------------------------------------
 Hash Join  (cost=37.00..57.24 rows=810 width=72)
   Hash Cond: (c."заказ" = o.id)
   ->  Seq Scan on clients c  (cost=0.00..18.10 rows=810 width=72)
   ->  Hash  (cost=22.00..22.00 rows=1200 width=4)
         ->  Seq Scan on orders o  (cost=0.00..22.00 rows=1200 width=4)
(5 rows)
```
Интерпретация:
* `Seq Scan on orders o` - Последовательно прочитана `orders`  
* `Hash` - Произведено хэширование  
* `Seq Scan on clients c` - Последовательно прочитана `clients`  
* `Hash Cond: (c."заказ" = o.id)` - Произведено хэширование строк с истинным значением `c."заказ" = o.id`  
* `Hash Join...` - объединение строк в результирующую таблицу.  

</details> 


## Задача 6

<details>
   <summary>Задание</summary>

Создайте бэкап БД test_db и поместите его в volume, предназначенный для бэкапов (см. Задачу 1).

Остановите контейнер с PostgreSQL (но не удаляйте volumes).

Поднимите новый пустой контейнер с PostgreSQL.

Восстановите БД test_db в новом контейнере.

Приведите список операций, который вы применяли для бэкапа данных и восстановления. 
</details>


<details>
<summary>Решение</summary>  

Создайте бэкап БД test_db и поместите его в volume, предназначенный для бэкапов  

```commandline
root@3547d3cbfc1d:/# export PGPASSWORD=testpwd777 && pg_dumpall -h localhost -U test-admin-user > /media/postgresql/backup/test_db.sql                                                                                                       root@3547d3cbfc1d:/#
root@3547d3cbfc1d:/# ls /media/postgresql/backup/
test_db.sql
root@3547d3cbfc1d:/#
root@3547d3cbfc1d:/# exit
exit
```

Остановите контейнер с PostgreSQL (но не удаляйте volumes).  

```commandline
root@deb10-test50:~# docker-compose stop
Stopping pg-sql01 ... done
root@deb10-test50:~#
root@deb10-test50:~# docker ps -a
CONTAINER ID   IMAGE         COMMAND                  CREATED             STATUS                     PORTS     NAMES
3547d3cbfc1d   postgres:12   "docker-entrypoint.s…"   About an hour ago   Exited (0) 7 seconds ago             pg-sql01
dbed7f5c74fb   hello-world   "/hello"                 5 hours ago         Exited (0) 5 hours ago               priceless_bassi
root@deb10-test50:~#
```

Поднимите новый пустой контейнер с PostgreSQL (`pg-sql02`).

```commandline
root@deb10-test50:~# docker run --rm -d -e POSTGRES_USER=test-admin-user \
>     -e POSTGRES_PASSWORD=testpwd777 \
> -e POSTGRES_DB=test_db \
> -v root_backup:/media/postgresql/backup --name pg-sql02 postgres:12
8ab1180a848630e23475a4c82f58a84ae1654b83880ead81d9c9e94d90e6a98f
root@deb10-test50:~#
root@deb10-test50:~# docker ps -a
CONTAINER ID   IMAGE         COMMAND                  CREATED             STATUS                      PORTS      NAMES
8ab1180a8486   postgres:12   "docker-entrypoint.s…"   8 seconds ago       Up 7 seconds                5432/tcp   pg-sql02
3547d3cbfc1d   postgres:12   "docker-entrypoint.s…"   About an hour ago   Exited (0) 29 seconds ago              pg-sql01
dbed7f5c74fb   hello-world   "/hello"                 5 hours ago         Exited (0) 5 hours ago                 priceless_bassi
root@deb10-test50:~#
```

Восстановите БД test_db в новом контейнере.  

```commandline
root@deb10-test50:~# docker exec -it pg-sql02 bash
root@8ab1180a8486:/#
root@8ab1180a8486:/# ls /media/postgresql/backup/
test_db.sql
root@8ab1180a8486:/# export PGPASSWORD=testpwd777 && psql -h localhost -U test-admin-user -f /media/postgresql/backup/test_db.sql test_db
SET
SET
SET
psql:/media/postgresql/backup/test_db.sql:14: ERROR:  role "test-admin-user" already exists
ALTER ROLE
CREATE ROLE
ALTER ROLE
You are now connected to database "template1" as user "test-admin-user".
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
You are now connected to database "postgres" as user "test-admin-user".
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
psql:/media/postgresql/backup/test_db.sql:110: ERROR:  database "test_db" already exists
ALTER DATABASE
You are now connected to database "test_db" as user "test-admin-user".
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
CREATE TABLE
ALTER TABLE
CREATE SEQUENCE
ALTER TABLE
ALTER SEQUENCE
ALTER TABLE
ALTER TABLE
COPY 5
COPY 5
 setval
--------
      1
(1 row)

 setval
--------
      1
(1 row)

ALTER TABLE
ALTER TABLE
CREATE INDEX
ALTER TABLE
GRANT
GRANT
GRANT
GRANT
root@8ab1180a8486:/#
```

Пользователи и таблицы в базе в новом контейнере  

```commandline
root@8ab1180a8486:/# psql -h localhost -U test-admin-user test_db
psql (12.11 (Debian 12.11-1.pgdg110+1))
Type "help" for help.

test_db=# \l+
                                                                               List of databases
   Name    |      Owner      | Encoding |  Collate   |   Ctype    |            Access privileges            |  Size   | Tablespace |                Description
-----------+-----------------+----------+------------+------------+-----------------------------------------+---------+------------+--------------------------------------------
 postgres  | test-admin-user | UTF8     | en_US.utf8 | en_US.utf8 |                                         | 7969 kB | pg_default | default administrative connection database
 template0 | test-admin-user | UTF8     | en_US.utf8 | en_US.utf8 | =c/"test-admin-user"                   +| 7825 kB | pg_default | unmodifiable empty database
           |                 |          |            |            | "test-admin-user"=CTc/"test-admin-user" |         |            |
 template1 | test-admin-user | UTF8     | en_US.utf8 | en_US.utf8 | =c/"test-admin-user"                   +| 7969 kB | pg_default | default template for new databases
           |                 |          |            |            | "test-admin-user"=CTc/"test-admin-user" |         |            |
 test_db   | test-admin-user | UTF8     | en_US.utf8 | en_US.utf8 | =Tc/"test-admin-user"                  +| 8161 kB | pg_default |
           |                 |          |            |            | "test-admin-user"=CTc/"test-admin-user"+|         |            |
           |                 |          |            |            | "test-simple-user"=c/"test-admin-user"  |         |            |
(4 rows)

test_db=#
test_db=# \d clients
                                       Table "public.clients"
      Column       |       Type        | Collation | Nullable |               Default
-------------------+-------------------+-----------+----------+-------------------------------------
 id                | integer           |           | not null | nextval('clients_id_seq'::regclass)
 фамилия           | character varying |           |          |
 страна проживания | character varying |           |          |
 заказ             | integer           |           |          |
Indexes:
    "clients_pkey" PRIMARY KEY, btree (id)
    "clients_страна проживания_idx" btree ("страна проживания")
Foreign-key constraints:
    "fk_заказ" FOREIGN KEY ("заказ") REFERENCES orders(id)

test_db=#
test_db=# \d orders
                                    Table "public.orders"
    Column    |       Type        | Collation | Nullable |              Default
--------------+-------------------+-----------+----------+------------------------------------
 id           | integer           |           | not null | nextval('orders_id_seq'::regclass)
 наименование | character varying |           |          |
 цена         | integer           |           |          |
Indexes:
    "orders_pkey" PRIMARY KEY, btree (id)
Referenced by:
    TABLE "clients" CONSTRAINT "fk_заказ" FOREIGN KEY ("заказ") REFERENCES orders(id)

```

</details>


---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
