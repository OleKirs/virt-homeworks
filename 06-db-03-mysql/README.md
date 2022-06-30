# Домашнее задание к занятию "6.3. MySQL"

## Введение

Перед выполнением задания вы можете ознакомиться с 
[дополнительными материалами](https://github.com/netology-code/virt-homeworks/tree/master/additional/README.md).

## Задача 1

<details>
   <summary>Задание</summary>

Используя docker поднимите инстанс MySQL (версию 8). Данные БД сохраните в volume.

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-03-mysql/test_data) и 
восстановитесь из него.

Перейдите в управляющую консоль `mysql` внутри контейнера.

Используя команду `\h` получите список управляющих команд.

Найдите команду для выдачи статуса БД и **приведите в ответе** из ее вывода версию сервера БД.

Подключитесь к восстановленной БД и получите список таблиц из этой БД.

**Приведите в ответе** количество записей с `price` > 300.

В следующих заданиях мы будем продолжать работу с данным контейнером.

</details>

<details>
<summary>Решение</summary>  

Используя docker поднимите инстанс MySQL (версию 8). Данные БД сохраните в volume.

```shell
root@deb10-test50:~# docker run \
>   --rm -d \
>   --name mysql_01 \
>   -p 3306:3306 \
>   -e MYSQL_DATABASE=test_db \
>   -e MYSQL_ROOT_PASSWORD=netology \
>   -v $PWD/mysql/data_vol:/var/lib/mysql \
>   -v $PWD/mysql/config_vol:/etc/mysql/conf.d \
>   -v $PWD/mysql/backup_vol:/media/mysql/backup \
>   mysql:8.0
01ae15ec1ff336df5aa9cbfd49a9515eb1fa0dfd89cb3ed94e759d0e68ee115f
root@deb10-test50:~#
```

Изучите [бэкап БД](https://github.com/netology-code/virt-homeworks/tree/master/06-db-03-mysql/test_data) и 
восстановитесь из него.

```shell
root@deb10-test50:~# wget -q  https://raw.githubusercontent.com/netology-code/virt-homeworks/master/06-db-03-mysql/test_data/test_dump.sql                                                                                                   
root@deb10-test50:~# docker cp test_dump.sql mysql_01:/media/mysql/backup/test_dump.sql
root@deb10-test50:~# docker exec -it mysql_01 bash
```
```shell
root@01ae15ec1ff3:/# mysql -u root -p test_db < /media/mysql/backup/test_dump.sql
Enter password:
root@01ae15ec1ff3:/#
```

Перейдите в управляющую консоль `mysql` внутри контейнера.

```shell
root@01ae15ec1ff3:/# mysql -u root -p
Enter password:
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 9
Server version: 8.0.29 MySQL Community Server - GPL

Copyright (c) 2000, 2022, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql>
```

Используя команду `\h` получите список управляющих команд.

```shell
mysql> \h

For information about MySQL products and services, visit:
   http://www.mysql.com/
For developer information, including the MySQL Reference Manual, visit:
   http://dev.mysql.com/
To buy MySQL Enterprise support, training, or other products, visit:
   https://shop.mysql.com/

List of all MySQL commands:
Note that all text commands must be first on line and end with ';'
?         (\?) Synonym for `help'.
clear     (\c) Clear the current input statement.
connect   (\r) Reconnect to the server. Optional arguments are db and host.
delimiter (\d) Set statement delimiter.
edit      (\e) Edit command with $EDITOR.
ego       (\G) Send command to mysql server, display result vertically.
exit      (\q) Exit mysql. Same as quit.
go        (\g) Send command to mysql server.
help      (\h) Display this help.
nopager   (\n) Disable pager, print to stdout.
notee     (\t) Don't write into outfile.
pager     (\P) Set PAGER [to_pager]. Print the query results via PAGER.
print     (\p) Print current command.
prompt    (\R) Change your mysql prompt.
quit      (\q) Quit mysql.
rehash    (\#) Rebuild completion hash.
source    (\.) Execute an SQL script file. Takes a file name as an argument.
status    (\s) Get status information from the server.
system    (\!) Execute a system shell command.
tee       (\T) Set outfile [to_outfile]. Append everything into given outfile.
use       (\u) Use another database. Takes database name as argument.
charset   (\C) Switch to another charset. Might be needed for processing binlog with multi-byte charsets.
warnings  (\W) Show warnings after every statement.
nowarning (\w) Don't show warnings after every statement.
resetconnection(\x) Clean session context.
query_attributes Sets string parameters (name1 value1 name2 value2 ...) for the next query to pick up.
ssl_session_data_print Serializes the current SSL session data to stdout or file

For server side help, type 'help contents'
```

Найдите команду для выдачи статуса БД и **приведите в ответе** из ее вывода версию сервера БД.

```shell
mysql> \s
--------------
mysql  Ver 8.0.29 for Linux on x86_64 (MySQL Community Server - GPL)

Connection id:          11
Current database:       test_db
Current user:           root@localhost
SSL:                    Not in use
Current pager:          stdout
Using outfile:          ''
Using delimiter:        ;
Server version:         8.0.29 MySQL Community Server - GPL
Protocol version:       10
Connection:             Localhost via UNIX socket
Server characterset:    utf8mb4
Db     characterset:    utf8mb4
Client characterset:    latin1
Conn.  characterset:    latin1
UNIX socket:            /var/run/mysqld/mysqld.sock
Binary data as:         Hexadecimal
Uptime:                 10 min 4 sec

Threads: 2  Questions: 54  Slow queries: 0  Opens: 168  Flush tables: 3  Open tables: 86  Queries per second avg: 0.089
--------------
```

Подключитесь к восстановленной БД и получите список таблиц из этой БД.

```shell
mysql> show tables;
+-------------------+
| Tables_in_test_db |
+-------------------+
| orders            |
+-------------------+
1 row in set (0.00 sec)

```

**Приведите в ответе** количество записей с `price` > 300.

```shell
mysql> select count(*) from orders where price >300;
+----------+
| count(*) |
+----------+
|        1 |
+----------+
1 row in set (0.01 sec)
```

</details>

## Задача 2

<details>
   <summary>Задание</summary>

Создайте пользователя test в БД c паролем test-pass, используя:
- плагин авторизации mysql_native_password
- срок истечения пароля - 180 дней 
- количество попыток авторизации - 3 
- максимальное количество запросов в час - 100
- аттрибуты пользователя:
    - Фамилия "Pretty"
    - Имя "James"

Предоставьте привелегии пользователю `test` на операции SELECT базы `test_db`.
    
Используя таблицу INFORMATION_SCHEMA.USER_ATTRIBUTES получите данные по пользователю `test` и 
**приведите в ответе к задаче**.

</details>

<details>
<summary>Решение</summary>  

Создайте пользователя `test` в БД с заданными параметрами

```shell
mysql> CREATE USER 'test'@'localhost'
    ->   IDENTIFIED WITH mysql_native_password BY 'test-pass'
    ->   WITH MAX_CONNECTIONS_PER_HOUR 100
    ->   PASSWORD EXPIRE INTERVAL 180 DAY
    ->   FAILED_LOGIN_ATTEMPTS 3 PASSWORD_LOCK_TIME 3
    ->   ATTRIBUTE '{"first_name":"James", "last_name":"Pretty"}';
Query OK, 0 rows affected (0.00 sec)
```

Предоставьте привелегии пользователю `test` на операции SELECT базы `test_db`.

```shell
mysql> GRANT SELECT ON test_db.* TO test@localhost;
Query OK, 0 rows affected (0.00 sec)
```

Используя таблицу INFORMATION_SCHEMA.USER_ATTRIBUTES получите данные по пользователю `test` и 

```shell
mysql> SELECT * FROM INFORMATION_SCHEMA.USER_ATTRIBUTES WHERE USER = 'test';
+------+-----------+------------------------------------------------+
| USER | HOST      | ATTRIBUTE                                      |
+------+-----------+------------------------------------------------+
| test | localhost | {"last_name": "Pretty", "first_name": "James"} |
+------+-----------+------------------------------------------------+
1 row in set (0.00 sec)
```

</details>

## Задача 3

<details>
   <summary>Задание</summary>

Установите профилирование `SET profiling = 1`.
Изучите вывод профилирования команд `SHOW PROFILES;`.

Исследуйте, какой `engine` используется в таблице БД `test_db` и **приведите в ответе**.

Измените `engine` и **приведите время выполнения и запрос на изменения из профайлера в ответе**:
- на `MyISAM`
- на `InnoDB`

</details>

<details>
<summary>Решение</summary>  

Установите профилирование `SET profiling = 1`.

```shell
mysql> SET profiling = 1;
Query OK, 0 rows affected, 1 warning (0.00 sec)
```

Изучите вывод профилирования команд `SHOW PROFILES;`.

```shell
mysql> SHOW PROFILES;
Empty set, 1 warning (0.00 sec)
```

Исследуйте, какой `engine` используется в таблице БД `test_db` и **приведите в ответе**.

```shell
mysql> SELECT TABLE_SCHEMA,TABLE_NAME,ENGINE FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'test_db';
+--------------+------------+--------+
| TABLE_SCHEMA | TABLE_NAME | ENGINE |
+--------------+------------+--------+
| test_db      | orders     | InnoDB |
1 row in set (0.01 sec)
```

Измените `engine` и **приведите время выполнения и запрос на изменения из профайлера в ответе**:
- на `MyISAM`

```shell
mysql> ALTER TABLE orders ENGINE = MyISAM;
Query OK, 5 rows affected (0.03 sec)
Records: 5  Duplicates: 0  Warnings: 0
```

- на `InnoDB`

```shell
mysql> ALTER TABLE orders ENGINE = InnoDB;
Query OK, 5 rows affected (0.04 sec)
Records: 5  Duplicates: 0  Warnings: 0
```

Длительность выполнения с большей точностью отражена в выводе запроса `SHOW PROFILES;`

```shell
mysql> SHOW PROFILES;
+----------+------------+------------------------------------------------------------------------------------------------------+
| Query_ID | Duration   | Query                                                                                                |
+----------+------------+------------------------------------------------------------------------------------------------------+
|        1 | 0.00143800 | SELECT TABLE_SCHEMA,TABLE_NAME,ENGINE FROM information_schema.TABLES WHERE TABLE_SCHEMA = 'test_db'  |
|        2 | 0.03352875 | ALTER TABLE orders ENGINE = MyISAM                                                                   |
|        3 | 0.03602775 | ALTER TABLE orders ENGINE = InnoDB                                                                   |
+----------+------------+------------------------------------------------------------------------------------------------------+
3 rows in set, 1 warning (0.00 sec)
```

</details>

## Задача 4 

<details>
   <summary>Задание</summary>

Изучите файл `my.cnf` в директории /etc/mysql.

Измените его согласно ТЗ (движок InnoDB):
- Скорость IO важнее сохранности данных
- Нужна компрессия таблиц для экономии места на диске
- Размер буффера с незакомиченными транзакциями 1 Мб
- Буффер кеширования 30% от ОЗУ
- Размер файла логов операций 100 Мб

Приведите в ответе измененный файл `my.cnf`.

</details>

<details>
<summary>Решение</summary>  

Определим объём ОЗУ доступного из контейнера 

```shell
root@01ae15ec1ff3:/# cat /proc/meminfo | grep MemTotal
MemTotal:        1926820 kB
```

Файл `/etc/mysql/my.cnf` с изменениями

```shell
root@01ae15ec1ff3:/# cat /etc/mysql/my.cnf
# Copyright (c) 2017, Oracle and/or its affiliates. All rights reserved.
#
# This program is free software; you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation; version 2 of the License.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program; if not, write to the Free Software
# Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301 USA

#
# The MySQL  Server configuration file.
#
# For explanations see
# http://dev.mysql.com/doc/mysql/en/server-system-variables.html

[mysqld]
pid-file        = /var/run/mysqld/mysqld.pid
socket          = /var/run/mysqld/mysqld.sock
datadir         = /var/lib/mysql
secure-file-priv= NULL

# Custom config should go here
!includedir /etc/mysql/conf.d/

#Max IO Speed (Скорость IO важнее сохранности данных)
innodb_flush_log_at_trx_commit = 0 

#Set compression (Нужна компрессия таблиц для экономии места на диске)
# Использовать для хранения формат файла с сжатием (Barracuda)
innodb_file_format=Barracuda
# (Дополнительно) Хранить каждую таблицу для InnoDB в отдельном файле
innodb_file_per_table=1;

#Set buffer (Размер буфера с незакомиченными транзакциями 1 Мб)
innodb_log_buffer_size	= 1M

#Set buffer_pool size (Буффер кеширования 30% от ОЗУ (30% от 2ГБ = 682MБ))
innodb_buffer_pool_size = 682M

#Set log size (Размер файла логов операций 100 Мб)
max_binlog_size	= 100M

```

</details>

---

### Как оформить ДЗ?

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
