
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

![Отображение стартовой сраницы Nginx в Firefox](imgs/5.3-Img01.png)

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
  
___
## Задача 3

- Запустите первый контейнер из образа ***centos*** c любым тэгом в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Запустите второй контейнер из образа ***debian*** в фоновом режиме, подключив папку ```/data``` из текущей рабочей директории на хостовой машине в ```/data``` контейнера;
- Подключитесь к первому контейнеру с помощью ```docker exec``` и создайте текстовый файл любого содержания в ```/data```;
- Добавьте еще один файл в папку ```/data``` на хостовой машине;
- Подключитесь во второй контейнер и отобразите листинг и содержание файлов в ```/data``` контейнера.
  
### Решение:
  
___
## Задача 4 (*)

Воспроизвести практическую часть лекции самостоятельно.

Соберите Docker образ с Ansible, загрузите на Docker Hub и пришлите ссылку вместе с остальными ответами к задачам.

  
### Решение:

---