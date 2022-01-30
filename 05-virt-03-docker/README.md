
# Домашнее задание к занятию "5.3. Введение. Экосистема. Архитектура. Жизненный цикл Docker контейнера"

---

## Задача 1

Сценарий выполения задачи:

- создайте свой репозиторий на https://hub.docker.com;
- выберете любой образ, который содержит веб-сервер Nginx;
- создайте свой fork образа;
- реализуйте функциональность:
запуск веб-сервера в фоне с индекс-страницей, содержащей HTML-код ниже:
```
<html>
<head>
Hey, Netology
</head>
<body>
<h1>I’m DevOps Engineer!</h1>
</body>
</html>
```
Опубликуйте созданный форк в своем репозитории и предоставьте ответ в виде ссылки на https://hub.docker.com/username_repo.

### Решение:
  
репозиторий - [olekirs/netology_nginx:1.21.6](https://hub.docker.com/repository/docker/olekirs/netology_nginx)

Запуск на Docker:
```shell
vagrant@server1:~/netology_nginx$ docker run -d -p 8080:80 olekirs/netology_nginx:1.21.6
Unable to find image 'olekirs/netology_nginx:1.21.6' locally
1.21.6: Pulling from olekirs/netology_nginx
5eb5b503b376: Pull complete 
1ae07ab881bd: Pull complete 
78091884b7be: Pull complete 
091c283c6a66: Pull complete 
55de5851019b: Pull complete 
b559bad762be: Pull complete 
cdbaff6c9faa: Pull complete 
7b91ead5a882: Pull complete 
Digest: sha256:bd208fa0653be6954ed044eab540d26005de60193c8331f943695a299661c9c7
Status: Downloaded newer image for olekirs/netology_nginx:1.21.6
7a4f47427f97d0d07967d3fdd188aa4ec4a92ce41715d580adaf423755d57454
```
  
Проверим, что контейнер запущен и получим его ID:
  
```shell
vagrant@server1:~/netology_nginx$ docker ps 
CONTAINER ID   IMAGE                           COMMAND                  CREATED          STATUS         PORTS                                   NAMES
7a4f47427f97   olekirs/netology_nginx:1.21.6   "/docker-entrypoint.…"   11 seconds ago   Up 9 seconds   0.0.0.0:8080->80/tcp, :::8080->80/tcp   thirsty_lewin
```
  
Получим сетевой адрес запущённого контейнера:
  
```shell
vagrant@server1:~/netology_nginx$ docker container inspect 7a4f47427f97 | grep \"IPAddress
            "IPAddress": "172.17.0.2",
                    "IPAddress": "172.17.0.2",
```
  
Получим содержимое страницы http://172.17.0.2
  
```shell

vagrant@server1:~/netology_nginx$ curl http://172.17.0.2
<html>
  <head>
    Hey, Netology
  </head>
  <body>
    <h1>I`m DevOps Engineer!</h1>
  </body>
</html>

```

Для проверки правильности работы Nginx получим страницу в браузере Firefox на хосте (порт хоста 8080 проброшен на порт 80 в контейнер)

![Отображение стартовой сраницы Nginx в Firefox](https://github.com/OleKirs/virt-homeworks/blob/05-virt-03-docker/05-virt-03-docker/imgs/5.3-img01.jpg)

___
## Задача 2

Посмотрите на сценарий ниже и ответьте на вопрос:
"Подходит ли в этом сценарии использование Docker контейнеров или лучше подойдет виртуальная машина, физическая машина? Может быть возможны разные варианты?"

Детально опишите и обоснуйте свой выбор.

--

Сценарий:

- Высоконагруженное монолитное java веб-приложение;
- Nodejs веб-приложение;
- Мобильное приложение c версиями для Android и iOS;
- Шина данных на базе Apache Kafka;
- Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana;
- Мониторинг-стек на базе Prometheus и Grafana;
- MongoDB, как основное хранилище данных для java-приложения;
- Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.
  
### Решение:

•	**Высоконагруженное монолитное java веб-приложение:**  
В общем случае - предпочтительна реализация на физической машине, чтобы не тратить ресурсы на виртуализацию.  

•	**Nodejs веб-приложение:**   
Скорее всего Docker подойдёт хорошо, т.к. позволит реализовать преимущества в скорости и стабильности развёртывания приложения.  

•	**Мобильное приложение c версиями для Android и iOS:**  
Создание среды разработки мобильного приложения - это задача для полной аппаратной виртуализации (для создания сред с разными версиями ОC, эмулирующих разные мобильные устройства).  
Выдача приложения потребителям в проде – это задача магазинов приложений или корпоративных web-серверов, которые могут использовать Docker в работе.  

•	**Шина данных на базе Apache Kafka**  
По описанию технологии, подразумевается активное масштабирование приложения с одинаковой архитектурой, а это представляется вполне применимым сценарием для Docker.  

•	**Elasticsearch кластер для реализации логирования продуктивного веб-приложения - три ноды elasticsearch, два logstash и две ноды kibana:**  
Docker, т.к. все приложения могут работать с одним ядром ОС и при этом использование контейнеров позволит оптимально утилизировать ресурсы.  

•	**Мониторинг-стек на базе Prometheus и Grafana**  
Docker представляется вполне приемлемым решением, особенно при потребности в масштабировании вслед за продуктивными серверами.  

•	**MongoDB, как основное хранилище данных для java-приложения**  
При необходимости сегментирования и быстрого масштабирования – Docker кажется предпочтительным вариантом. При относительно стабильных объёмах возможно стоит рассмотреть реализацию на отдельной ВМ.   

•	**Gitlab сервер для реализации CI/CD процессов и приватный (закрытый) Docker Registry.**    
Предпочтительна реализация на ВМ, т.к. не вижу необходимости в активном масштабировании ресурсов (кроме места для хранения образов) и при этом шире представлены инструменты работы с артефактами.  

___
## Задача 3

- Запустите первый контейнер из образа ***centos*** c любым тэгом в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Запустите второй контейнер из образа ***debian*** в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Подключитесь к первому контейнеру с помощью ```docker exec``` и создайте текстовый файл любого содержания в ```/data```;
- Добавьте еще один файл в папку ```/data``` на хостовой машине;
- Подключитесь во второй контейнер и отобразите листинг и содержание файлов в ```/data``` контейнера.
  
### Решение:
#### Запустите первый контейнер из образа ***centos*** c любым тэгом в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;

```shell
vagrant@server1:~/netology_nginx$ docker run -it --rm -d --name centos -v $(pwd)/data:/data centos:latest
Unable to find image 'centos:latest' locally
latest: Pulling from library/centos
a1d0c7532777: Pull complete 
Digest: sha256:a27fd8080b517143cbbbab9dfb7c8571c40d67d534bbdee55bd6c473f432b177
Status: Downloaded newer image for centos:latest
f28215c8ae8791b6514472424b1dfbb470eea6bd2fb94e440304c18059648156

vagrant@server1:~/netology_nginx$ docker ps
CONTAINER ID   IMAGE                           COMMAND                  CREATED          STATUS          PORTS                                   NAMES
f28215c8ae87   centos:latest                   "/bin/bash"              14 seconds ago   Up 12 seconds                                           centos
7a4f47427f97   olekirs/netology_nginx:1.21.6   "/docker-entrypoint.…"   3 hours ago      Up 3 hours      0.0.0.0:8080->80/tcp, :::8080->80/tcp   thirsty_lewin
```
#### Запустите второй контейнер из образа ***debian*** в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;

```shell
vagrant@server1:~/netology_nginx$ docker pull debian:stretch
stretch: Pulling from library/debian
a834d7c95167: Pull complete 
Digest: sha256:4bb600434787c903886fe33526d19ff33114a33b433a4a4cdbdf9b8543f1ab5d
Status: Downloaded newer image for debian:stretch
docker.io/library/debian:stretch

vagrant@server1:~/netology_nginx$ docker run -d -it --name debian -v $(pwd)/data:/data debian:stretch
98ebde5c90e63e86fc33102bc35d2c5c8622e626b8696e1431d2d7eb185d5f32

vagrant@server1:~/netology_nginx$ docker ps
CONTAINER ID   IMAGE                           COMMAND                  CREATED          STATUS         PORTS                                   NAMES
98ebde5c90e6   debian:stretch                  "bash"                   10 seconds ago   Up 9 seconds                                           debian
f28215c8ae87   centos:latest                   "/bin/bash"              7 minutes ago    Up 7 minutes                                           centos
7a4f47427f97   olekirs/netology_nginx:1.21.6   "/docker-entrypoint.…"   3 hours ago      Up 3 hours     0.0.0.0:8080->80/tcp, :::8080->80/tcp   thirsty_lewin
```

#### Подключитесь к первому контейнеру с помощью ```docker exec``` и создайте текстовый файл любого содержания в ```/data```;
```shell
vagrant@server1:~/netology_nginx$ docker exec -it centos /bin/bash

[root@f28215c8ae87 /]# echo 'Test message' > /data/testfile

[root@f28215c8ae87 /]# cat /data/testfile
Test message

[root@f28215c8ae87 /]# exit
exit
```

#### Добавьте еще один файл в папку ```/data``` на хостовой машине;
```shell
vagrant@server1:~/netology_nginx$ sudo su -

root@server1:~# echo 'Second test message' > /home/vagrant/netology_nginx/data/testfile_02

root@server1:~# exit
logout

```
#### Подключитесь во второй контейнер и отобразите листинг и содержание файлов в ```/data``` контейнера.
```shell
vagrant@server1:~/netology_nginx$ docker exec -it debian /bin/bash

root@98ebde5c90e6:/# ls -la /data/
total 16
drwxr-xr-x 2 root root 4096 Jan 30 17:34 .
drwxr-xr-x 1 root root 4096 Jan 30 17:23 ..
-rw-r--r-- 1 root root   13 Jan 30 17:29 testfile
-rw-r--r-- 1 root root   20 Jan 30 17:34 testfile_02

root@98ebde5c90e6:/# cat /data/testfile
Test message

root@98ebde5c90e6:/# cat /data/testfile_02 
Second test message

root@98ebde5c90e6:/# exit
exit

vagrant@server1:~/netology_nginx$
```
 

___
## Задача 4 (*)

Воспроизвести практическую часть лекции самостоятельно.

Соберите Docker образ с Ansible, загрузите на Docker Hub и пришлите ссылку вместе с остальными ответами к задачам.

  
### Решение:

[https://hub.docker.com/repository/docker/olekirs/ansible](https://hub.docker.com/repository/docker/olekirs/ansible)

---
