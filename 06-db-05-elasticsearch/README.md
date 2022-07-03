# Домашнее задание к занятию "6.5. Elasticsearch"

## Задача 1

<details>
   <summary>Задание</summary>


В этом задании вы потренируетесь в:
- установке elasticsearch
- первоначальном конфигурировании elastcisearch
- запуске elasticsearch в docker

Используя докер образ [centos:7](https://hub.docker.com/_/centos) как базовый и 
[документацию по установке и запуску Elastcisearch](https://www.elastic.co/guide/en/elasticsearch/reference/current/targz.html):

- составьте Dockerfile-манифест для elasticsearch
- соберите docker-образ и сделайте `push` в ваш docker.io репозиторий
- запустите контейнер из получившегося образа и выполните запрос пути `/` c хост-машины

Требования к `elasticsearch.yml`:
- данные `path` должны сохраняться в `/var/lib`
- имя ноды должно быть `netology_test`

В ответе приведите:
- текст Dockerfile манифеста
- ссылку на образ в репозитории dockerhub
- ответ `elasticsearch` на запрос пути `/` в json виде

Подсказки:
- возможно вам понадобится установка пакета perl-Digest-SHA для корректной работы пакета shasum
- при сетевых проблемах внимательно изучите кластерные и сетевые настройки в elasticsearch.yml
- при некоторых проблемах вам поможет docker директива ulimit
- elasticsearch в логах обычно описывает проблему и пути ее решения

Далее мы будем работать с данным экземпляром elasticsearch.

</details>

<details>
<summary>Решение</summary>  

Текст Dockerfile манифеста  

```shell
root@deb10-test50:~# cat ~/elasticsearch-docker/Dockerfile
FROM centos:7

EXPOSE 9200 9300

USER 0

RUN export ES_HOME="/var/lib/elasticsearch" && \
    yum -y install wget && \
    wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.17.0-linux-x86_64.tar.gz && \
    wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-7.17.0-linux-x86_64.tar.gz.sha512 && \
    sha512sum -c elasticsearch-7.17.0-linux-x86_64.tar.gz.sha512 && \
    tar -xzf elasticsearch-7.17.0-linux-x86_64.tar.gz && \
    rm -f elasticsearch-7.17.0-linux-x86_64.tar.gz* && \
    mv elasticsearch-7.17.0 ${ES_HOME} && \
    useradd -m -u 1000 elasticsearch && \
    chown elasticsearch:elasticsearch -R ${ES_HOME} && \
    yum -y remove wget && \
    yum clean all

COPY --chown=elasticsearch:elasticsearch config/* /var/lib/elasticsearch/config/

RUN mkdir -p /var/lib/elasticsearch/logs &&\
    mkdir -p /var/lib/elasticsearch/data &&\
    chown -R elasticsearch:elasticsearch /var/lib/elasticsearch/

USER 1000

ENV ES_HOME="/var/lib/elasticsearch" \
    ES_PATH_CONF="/var/lib/elasticsearch/config"
WORKDIR ${ES_HOME}

CMD ["sh", "-c", "${ES_HOME}/bin/elasticsearch"]

```

Ссылку на образ в репозитории dockerhub:  
[olekirs/centos7-elasticsearch:7.17](https://hub.docker.com/repository/docker/olekirs/centos7-elasticsearch)
```shell
docker push olekirs/centos7-elasticsearch:7.17
```

Ответ `elasticsearch` на запрос пути `/` в json виде  

```shell
docker run --rm -d \
  --name centos-elastic \
  -p 9200:9200 \
  -p 9300:9300 \
  olekirs/centos7-elasticsearch:7.17
```

```shell
root@deb10-test50:~/elasticsearch-docker# docker run --rm -d \
>   --name centos-elastic \
>   -p 9200:9200 \
>   -p 9300:9300 \
>   olekirs/centos7-elasticsearch:7.17
0457ceb7b08eaa434f526be3108bda1e072952bcbbc7bca5f9c376bb0d103c0a
```

```shell
root@deb10-test50:~/elasticsearch-docker# docker ps
CONTAINER ID   IMAGE                                COMMAND                  CREATED              STATUS              PORTS                                                                                  NAMES
0457ceb7b08e   olekirs/centos7-elasticsearch:7.17   "sh -c ${ES_HOME}/bi…"   About a minute ago   Up About a minute   0.0.0.0:9200->9200/tcp, :::9200->9200/tcp, 0.0.0.0:9300->9300/tcp, :::9300->9300/tcp   centos-elastic
```

```shell
root@deb10-test50:~/elasticsearch-docker# curl -X GET 'localhost:9200/'
{
  "name" : "netology_test",
  "cluster_name" : "netology",
  "cluster_uuid" : "_3-BX5VhSB2nVBI_Pmo5wA",
  "version" : {
    "number" : "7.17.0",
    "build_flavor" : "default",
    "build_type" : "tar",
    "build_hash" : "bee86328705acaa9a6daede7140defd4d9ec56bd",
    "build_date" : "2022-01-28T08:36:04.875279988Z",
    "build_snapshot" : false,
    "lucene_version" : "8.11.1",
    "minimum_wire_compatibility_version" : "6.8.0",
    "minimum_index_compatibility_version" : "6.0.0-beta1"
  },
  "tagline" : "You Know, for Search"
}
```

Файл `elasticsearch.yml`
```shell
root@deb10-test50:~/elasticsearch-docker# grep "^[^#*/;]" config/elasticsearch.yml
cluster.name: netology
discovery.type: single-node
node.name: netology_test
path.data: /var/lib/elasticsearch/data
path.logs: /var/lib/elasticsearch/logs
network.host: 0.0.0.0
discovery.seed_hosts: ["127.0.0.1", "[::1]"]

```

</details>


## Задача 2

<details>
   <summary>Задание</summary>


В этом задании вы научитесь:
- создавать и удалять индексы
- изучать состояние кластера
- обосновывать причину деградации доступности данных

Ознакомтесь с [документацией](https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html) 
и добавьте в `elasticsearch` 3 индекса, в соответствии со таблицей:

| Имя | Количество реплик | Количество шард |
|-----|-------------------|-----------------|
| ind-1| 0 | 1 |
| ind-2 | 1 | 2 |
| ind-3 | 2 | 4 |

Получите список индексов и их статусов, используя API и **приведите в ответе** на задание.

Получите состояние кластера `elasticsearch`, используя API.

Как вы думаете, почему часть индексов и кластер находится в состоянии yellow?

Удалите все индексы.

**Важно**

При проектировании кластера elasticsearch нужно корректно рассчитывать количество реплик и шард,
иначе возможна потеря данных индексов, вплоть до полной, при деградации системы.

</details>

<details>
<summary>Решение</summary>  

Добавьте в `elasticsearch` 3 индекса, в соответствии со таблицей:

```shell
curl -X PUT "localhost:9200/ind-1?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "number_of_replicas": 0,
    "number_of_shards": 1
  }
}
'
```
```shell
curl -X PUT "localhost:9200/ind-2?pretty" -H 'Content-Type: application/json' -d'
  "settings": {
    "number_of_replicas": 1,
    "number_of_shards": 2
  }
}
'
```
```shell
root@deb10-test50:~/elasticsearch-docker# curl -X PUT "localhost:9200/ind-3?pretty" -H 'Content-Type: application/json' -d'
  "settings": {
    "number_of_replicas": 2,
    "number_of_shards": 4
  }
}
'
```

Получите список индексов и их статусов

```shell
root@deb10-test50:~/elasticsearch-docker# curl 'localhost:9200/_cat/indices?v'
health status index            uuid                   pri rep docs.count docs.deleted store.size pri.store.size
green  open   .geoip_databases OhaZFzrBQaqi31ZHoYiFdQ   1   0         40            0       38mb           38mb
green  open   ind-1            jB5KilTLQU2xYALqN7Ji5A   1   0          0            0       226b           226b
yellow open   ind-3            -VmL5j33ScO8_J-2prOT-w   4   2          0            0       904b           904b
yellow open   ind-2            FxN6lZteQhiARtB26U-TkQ   2   1          0            0       452b           452b
```

Получите состояние кластера `elasticsearch`

```shell
root@deb10-test50:~/elasticsearch-docker# curl -X GET "localhost:9200/_cluster/health?pretty"
{
  "cluster_name" : "netology",
  "status" : "yellow",
  "timed_out" : false,
  "number_of_nodes" : 1,
  "number_of_data_nodes" : 1,
  "active_primary_shards" : 10,
  "active_shards" : 10,
  "relocating_shards" : 0,
  "initializing_shards" : 0,
  "unassigned_shards" : 10,
  "delayed_unassigned_shards" : 0,
  "number_of_pending_tasks" : 0,
  "number_of_in_flight_fetch" : 0,
  "task_max_waiting_in_queue_millis" : 0,
  "active_shards_percent_as_number" : 50.0
}
```

Почему часть индексов и кластер находится в состоянии yellow?

Кластер состоит из одной ноды, поэтому реплики индексов `ind-2` и `ind-3` не могут быть запущены и индексы помечены как `yellow`, т.е. в режиме ограниченной функциональности.
Состояние кластера отражает наихудшее состояние любого из индексов в кластере.

Удалите все индексы.

```shell
root@deb10-test50:~/elasticsearch-docker# curl -X DELETE 'http://localhost:9200/*'
{"acknowledged":true} 
root@deb10-test50:~/elasticsearch-docker#
root@deb10-test50:~/elasticsearch-docker# curl 'localhost:9200/_cat/indices?pretty'
green open .geoip_databases OhaZFzrBQaqi31ZHoYiFdQ 1 0 40 0 38mb 38mb

```

</details>


## Задача 3

<details>
   <summary>Задание</summary>


В данном задании вы научитесь:
- создавать бэкапы данных
- восстанавливать индексы из бэкапов

Создайте директорию `{путь до корневой директории с elasticsearch в образе}/snapshots`.

Используя API [зарегистрируйте](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-register-repository.html#snapshots-register-repository) 
данную директорию как `snapshot repository` c именем `netology_backup`.

**Приведите в ответе** запрос API и результат вызова API для создания репозитория.

Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.

[Создайте `snapshot`](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-take-snapshot.html) 
состояния кластера `elasticsearch`.

**Приведите в ответе** список файлов в директории со `snapshot`ами.

Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

[Восстановите](https://www.elastic.co/guide/en/elasticsearch/reference/current/snapshots-restore-snapshot.html) состояние
кластера `elasticsearch` из `snapshot`, созданного ранее. 

**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

Подсказки:
- возможно вам понадобится доработать `elasticsearch.yml` в части директивы `path.repo` и перезапустить `elasticsearch`

</details>

<details>
<summary>Решение</summary>  

Создайте директорию `{путь до корневой директории с elasticsearch в образе}/snapshots`

```shell
root@deb10-test50:~# docker exec -it centos-elastic bash
[elasticsearch@0457ceb7b08e elasticsearch]$ mkdir -p  "$ES_HOME/config/elasticsearch.yml"

```
Используя API зарегистрируйте данную директорию как `snapshot repository` c именем `netology_backup`.

```shell
[elasticsearch@0457ceb7b08e elasticsearch]$ echo path.repo: [ "/var/lib/elasticsearch/snapshots" ] >> "$ES_HOME/config/elasticsearch.yml"

[elasticsearch@0457ceb7b08e elasticsearch]$ exit
exit

root@deb10-test50:~# docker restart centos-elastic
centos-elastic
```

Запрос  

```shell
curl -X PUT "localhost:9200/_snapshot/netology_backup?pretty" -H 'Content-Type: application/json' -d'
{
  "type": "fs",
  "settings": {
    "location": "/var/lib/elasticsearch/snapshots",
    "compress": true
  }
}
'
```

Результат

```shell
root@deb10-test50:~# curl -X PUT "localhost:9200/_snapshot/netology_backup?pretty" -H 'Content-Type: application/json' -d'
{
  "type": "fs",
  "settings": {
    "location": "/var/lib/elasticsearch/snapshots",
    "compress": true
  }
}
'
{
  "acknowledged" : true
}
```

Создайте индекс `test` с 0 реплик и 1 шардом и **приведите в ответе** список индексов.

Запрос  

```shell
curl -X PUT "localhost:9200/test?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "number_of_replicas": 0,
	"number_of_shards": 1
  }
}
'
```

Результат  

```shell
root@deb10-test50:~# curl -X PUT "localhost:9200/test?pretty" -H 'Content-Type: application/json' -d'
> {
>   "settings": {
>     "number_of_replicas": 0,
> "number_of_shards": 1
>   }
> }
> '
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "test"
}
```

Список индексов  

```shell
root@deb10-test50:~# curl 'localhost:9200/_cat/indices?pretty'
green open .geoip_databases 4Bpkuh3wT6ON8bPGi2T6ag 1 0 40 0 38mb 38mb
green open test             33TFHzAwTT2jjaXVCsIBbA 1 0  0 0 226b 226b
```

Создайте `snapshot` состояния кластера `elasticsearch`.  

```shell
root@deb10-test50:~# curl -X PUT "localhost:9200/_snapshot/netology_backup/snapshot_01?wait_for_completion=true&pretty"
{
  "snapshot" : {
    "snapshot" : "snapshot_01",
    "uuid" : "r4quGLDJTsKBIhtUmBmoTQ",
    "repository" : "netology_backup",
    "version_id" : 7170099,
    "version" : "7.17.0",
    "indices" : [
      ".ds-.logs-deprecation.elasticsearch-default-2022.07.03-000001",
      ".geoip_databases",
      ".ds-ilm-history-5-2022.07.03-000001",
      "test"
    ],
    "data_streams" : [
      "ilm-history-5",
      ".logs-deprecation.elasticsearch-default"
    ],
    "include_global_state" : true,
    "state" : "SUCCESS",
    "start_time" : "2022-07-03T20:43:02.961Z",
    "start_time_in_millis" : 1656880982961,
    "end_time" : "2022-07-03T20:43:04.362Z",
    "end_time_in_millis" : 1656880984362,
    "duration_in_millis" : 1401,
    "failures" : [ ],
    "shards" : {
      "total" : 4,
      "failed" : 0,
      "successful" : 4
    },
    "feature_states" : [
      {
        "feature_name" : "geoip",
        "indices" : [
          ".geoip_databases"
        ]
      }
    ]
  }
}
```

**Приведите в ответе** список файлов в директории со `snapshot`ами.  

```shell
root@deb10-test50:~# docker exec -it centos-elastic ls -l /var/lib/elasticsearch/snapshots/
total 28
-rw-r--r-- 1 elasticsearch elasticsearch 1423 Jul  3 20:43 index-0
-rw-r--r-- 1 elasticsearch elasticsearch    8 Jul  3 20:43 index.latest
drwxr-xr-x 6 elasticsearch elasticsearch 4096 Jul  3 20:43 indices
-rw-r--r-- 1 elasticsearch elasticsearch 9758 Jul  3 20:43 meta-r4quGLDJTsKBIhtUmBmoTQ.dat
-rw-r--r-- 1 elasticsearch elasticsearch  458 Jul  3 20:43 snap-r4quGLDJTsKBIhtUmBmoTQ.dat
```

Удалите индекс `test` и создайте индекс `test-2`. **Приведите в ответе** список индексов.

```shell
root@deb10-test50:~# curl -X DELETE "localhost:9200/test?pretty"
{
  "acknowledged" : true
}
root@deb10-test50:~#
root@deb10-test50:~# curl -X PUT "localhost:9200/test-2?pretty" -H 'Content-Type: application/json' -d'
> {
>   "settings": {
>     "number_of_shards": 1,
>     "number_of_replicas": 0
>   }
> }
> '
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "index" : "test-2"
}
```

Список индексов  

```shell
root@deb10-test50:~# curl 'localhost:9200/_cat/indices?pretty'
green open test-2           j0mAqtFmQg-JWANvyGNmdA 1 0  0 0 226b 226b
green open .geoip_databases OhaZFzrBQaqi31ZHoYiFdQ 1 0 40 0 38mb 38mb
root@deb10-test50:~#
```

Восстановите состояние кластера `elasticsearch` из `snapshot`, созданного ранее. 
**Приведите в ответе** запрос к API восстановления и итоговый список индексов.

- Восстановление индекса `test` можно провести одним запросом.  
- Восстановление состояния кластера требует вывода на обслуживание, т.к. нужно закрывать системые индексы, иначе при попытке восстановления выдаётся ошибка *"нельзя восстановить открытые индексы"*
  
Получилось сделать таким образом:    
1. Остановить `GeoIP database downloader`  
```shell
root@deb10-test50:~# curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
> {
>   "persistent": {
>     "ingest.geoip.downloader.enabled": false
>   }
> }
> '
{
  "acknowledged" : true,
  "persistent" : {
    "ingest" : {
      "geoip" : {
        "downloader" : {
          "enabled" : "false"
        }
      }
    }
  },
  "transient" : { }
}
```
2. Закрыть открытые индексы (в т.ч. скрытые)  
```shell
root@deb10-test50:~# curl -X POST "localhost:9200/*/_close?pretty"
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "indices" : {
    "test-2" : {
      "closed" : true
    }
  }
}
root@deb10-test50:~# curl -X POST "localhost:9200/.*/_close?pretty"
{
  "acknowledged" : true,
  "shards_acknowledged" : true,
  "indices" : {
    ".ds-.logs-deprecation.elasticsearch-default-2022.07.03-000001" : {
      "closed" : true
    },
    ".ds-ilm-history-5-2022.07.03-000001" : {
      "closed" : true
    }
  }
}
```
3. Восстановить индексы из снапшота, включая `global_state`  
```shell
root@deb10-test50:~# curl -X POST "localhost:9200/_snapshot/netology_backup/snapshot_01/_restore?pretty" -H 'Content-Type: application/json' -d'
> {
>   "indices": "*",
>   "include_global_state": true
> }
> '
{
  "accepted" : true
}
```
4. Открыть ранее закрытые индексы  
```shell
root@deb10-test50:~# curl -X POST "localhost:9200/*/_open?pretty"
{
  "acknowledged" : true,
  "shards_acknowledged" : true
}
root@deb10-test50:~# curl -X POST "localhost:9200/.*/_open?pretty"
{
  "acknowledged" : true,
  "shards_acknowledged" : true
}
```

5. Запустить `GeoIP database downloader`  
```shell
root@deb10-test50:~# curl -X PUT "localhost:9200/_cluster/settings?pretty" -H 'Content-Type: application/json' -d'
> {
>   "persistent": {
>     "ingest.geoip.downloader.enabled": true
>   }
> }
> '
{
  "acknowledged" : true,
  "persistent" : {
    "ingest" : {
      "geoip" : {
        "downloader" : {
          "enabled" : "true"
        }
      }
    }
  },
  "transient" : { }
}
```

Результат (список индексов):

```shell
root@deb10-test50:~# curl 'localhost:9200/_cat/indices?pretty'
green open test-2           j0mAqtFmQg-JWANvyGNmdA 1 0  0 0 248b 248b
green open .geoip_databases 4Bpkuh3wT6ON8bPGi2T6ag 1 0 40 0 38mb 38mb
green open test             33TFHzAwTT2jjaXVCsIBbA 1 0  0 0 226b 226b
```

</details>


---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
