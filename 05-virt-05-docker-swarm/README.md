# Домашнее задание к занятию "5.5. Оркестрация кластером Docker контейнеров на примере Docker Swarm"

___

## Задача 1

Дайте письменые ответы на следующие вопросы:

- В чём отличие режимов работы сервисов в Docker Swarm кластере: replication и global?
- Какой алгоритм выбора лидера используется в Docker Swarm кластере?
- Что такое Overlay Network?
  
### Решение
  
___

## Задача 2

Создать ваш первый Docker Swarm кластер в Яндекс.Облаке

Для получения зачета, вам необходимо предоставить скриншот из терминала (консоли), с выводом команды:
```
docker node ls
```
  
### Решение
  
#### Подготовим образ для развёртываемых ВМ:
Инициализируем (повторно) профиль в Yandex облаке с помощью утилиты `yc`:
```shell
root@D10:~/virt-homeworks/05-virt-04-docker-compose/src/packer# yc init
Welcome! This command will take you through the configuration process.
Pick desired action:
 [1] Re-initialize this profile 'default' with new settings
 [2] Create a new profile
Please enter your numeric choice: 1

Please go to https://oauth.yandex.ru/authorize?response_type=token&client_id=1a6990*****ec2fb in order to obtain OAuth token.

Please enter OAuth token: AQAAAAABFdXXAAT*******************ZjMI

You have one cloud available: 'cloud-kirs-corp' (id = 'b1g3*************2q'). It is going to be used by default.
Please choose folder to use:
 [1] netology (id = b1g***************44)
 [2] Create a new folder
Please enter your numeric choice: 1

Your current folder has been set to 'netology' (id = b1g***************44).
Do you want to configure a default Compute zone? [Y/n] y

Which zone do you want to use as a profile default?
 [1] ru-central1-a
 [2] ru-central1-b
 [3] ru-central1-c
 [4] Don\'t set default zone
Please enter your numeric choice: 1
Your profile default Compute zone has been set to 'ru-central1-a'.
```
Создадим в облаке сеть с именем `net`
```shell
root@D10:~/virt-homeworks/05-virt-04-docker-compose/src/packer# yc vpc network create  --name net
id: enp1s*************pr0g
folder_id: b1g***************44
created_at: "2022-02-12T21:24:03Z"
name: net
```
Создадим в сети `net` подсеть с именем `subnet`
```shell
root@D10:~/virt-homeworks/05-virt-04-docker-compose/src/packer# yc vpc subnet create  --name my-subnet-a --zone ru-central1-a --range 10.1.2.0/24 --network-name net --description 'my first subnet via yc'
id: e9b9***********cgir2
folder_id: b1g***************44
created_at: "2022-02-12T21:24:24Z"
name: my-subnet-a
description: my first subnet via yc
network_id: enp1s*************pr0g
zone_id: ru-central1-a
v4_cidr_blocks:
- 10.1.2.0/24
```
Изменим значения в файле конфигурации `.../src/packer/centos-7-base.json`
```shell
root@D10:~/virt-homeworks/05-virt-04-docker-compose/src/packer# nano centos-7-base.json
root@D10:~/virt-homeworks/05-virt-04-docker-compose/src/packer# cat centos-7-base.json
{
  "builders": [
    {
      "disk_type": "network-nvme",
      "folder_id": "b1g***************44",
      "image_description": "by packer",
      "image_family": "centos",
      "image_name": "centos-7-base",
      "source_image_family": "centos-7",
      "ssh_username": "centos",
      "subnet_id": "e9b9a3*******cgir2",
      "token": "AQAAAAABFdXXAAT*******************ZjMI",
      "type": "yandex",
      "use_ipv4_nat": true,
      "zone": "ru-central1-a"
    }
  ],
  "provisioners": [
    {
      "inline": [
        "sudo yum -y update",
        "sudo yum -y install bridge-utils bind-utils iptables curl net-tools tcpdump rsync telnet openssh-server"
      ],
      "type": "shell"
    }
  ]
}
```
Проверим правильность конфигурации `packer`
```shell
root@D10:~/virt-homeworks/05-virt-04-docker-compose/src/packer# packer validate centos-7-base.json
The configuration is valid.
```
Запустим на выполнение `packer` с конфигурацией `centos-7-base.json`
```shell
root@D10:~/virt-homeworks/05-virt-04-docker-compose/src/packer# packer build centos-7-base.json
yandex: output will be in this color.

==> yandex: Creating temporary RSA SSH key for instance...
==> yandex: Using as source image: fd8gdnd09d0iqdu7ll2a (name: "centos-7-v20220207", family: "centos-7")
==> yandex: Use provided subnet id e9b9a3qig1ek7qicgir2
==> yandex: Creating disk...
==> yandex: Creating instance...
==> yandex: Waiting for instance with id fhmsen9j3gboom8boh3s to become active...
    yandex: Detected instance IP: 62.84.113.249
==> yandex: Using SSH communicator to connect: 62.84.113.249
==> yandex: Waiting for SSH to become available...
==> yandex: Connected to SSH!
==> yandex: Provisioning with shell script: /tmp/packer-shell125394147
...
    yandex: Complete!
==> yandex: Stopping instance...
==> yandex: Deleting instance...
    yandex: Instance has been deleted!
==> yandex: Creating image: centos-7-base
==> yandex: Waiting for image to complete...
==> yandex: Success image create...
==> yandex: Destroying boot disk...
    yandex: Disk has been deleted!
Build 'yandex' finished after 2 minutes 2 seconds.

==> Wait completed after 2 minutes 2 seconds

==> Builds finished. The artifacts of successful builds are:
--> yandex: A disk image was created: centos-7-base (id: fd8795*******2jamnl) with family name centos
```
**Образ подготовлен**

___

## Задача 3

Создать ваш первый, готовый к боевой эксплуатации кластер мониторинга, состоящий из стека микросервисов.

Для получения зачета, вам необходимо предоставить скриншот из терминала (консоли), с выводом команды:
```
docker service ls
```
  
### Решение
  
___

## Задача 4 (*)

Выполнить на лидере Docker Swarm кластера команду (указанную ниже) и дать письменное описание её функционала, что она делает и зачем она нужна:
```
# см.документацию: https://docs.docker.com/engine/swarm/swarm_manager_locking/
docker swarm update --autolock=true
```
