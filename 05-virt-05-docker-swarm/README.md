# Домашнее задание к занятию "5.5. Оркестрация кластером Docker контейнеров на примере Docker Swarm"

___

## Задача 1

Дайте письменые ответы на следующие вопросы:

- В чём отличие режимов работы сервисов в Docker Swarm кластере: replication и global?
- Какой алгоритм выбора лидера используется в Docker Swarm кластере?
- Что такое Overlay Network?
  
### Решение

#### В чём отличие режимов работы сервисов в Docker Swarm кластере: replication и global?  
***global*** - в этом режиме сервис запускается сразу на всех узлах, удовлетворяющих условиям ограничений для сервиса.  
***replication*** - запускается в заданном (ограниченном) количестве экземпляров и количество одновременно работающих сервисов может меняться в зависимости от потребности в масштабировании сервиса.   
  
#### Какой алгоритм выбора лидера используется в Docker Swarm кластере?    
Для выбора лидера среди master node используется алгоритм поддержания распределенного консенсуса — ***Raft***  

#### Что такое Overlay Network?
***Overlay Network*** - это распределенная сеть между несколькими узлами демона Docker.  
Эта сеть находится поверх (перекрывает) сети, специфичные для хоста, позволяя контейнерам, подключенным к ней (включая контейнеры службы swarm), безопасно обмениваться данными при включенном шифровании.   
Docker, используя Overlay Network, прозрачно обрабатывает маршрутизацию каждого входящего и исходящего пакета, направляя его к нужным Docker-хостам и контейнерам, даже если изначально пакет пришел на другой хост.  
___

## Задача 2

Создать ваш первый Docker Swarm кластер в Яндекс.Облаке

Для получения зачета, вам необходимо предоставить скриншот из терминала (консоли), с выводом команды:
```
docker node ls
```
  
### Решение
  
#### 2.1. Подготовим образ для развёртываемых ВМ:  
<details>
  
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
  
Удалим оставшиеся от создания образа подсети и сети  
  
```shell
root@deb10-gw-k11:~/virt-homeworks/05-virt-05-docker-swarm/src/terraform# yc vpc subnet list
+----------------------+-------------+----------------------+----------------+---------------+---------------+
|          ID          |    NAME     |      NETWORK ID      | ROUTE TABLE ID |     ZONE      |     RANGE     |
+----------------------+-------------+----------------------+----------------+---------------+---------------+
| e9b9a**********cgir2 | my-subnet-a | enp1***********pr0g  |                | ru-central1-a | [10.1.2.0/24] |
+----------------------+-------------+----------------------+----------------+---------------+---------------+

root@deb10-gw-k11:~/virt-homeworks/05-virt-05-docker-swarm/src/terraform# yc vpc subnet delete e9b9a********cgir2
done (2s)

root@deb10-gw-k11:~/virt-homeworks/05-virt-05-docker-swarm/src/terraform# yc vpc network list
+----------------------+------+
|          ID          | NAME |
+----------------------+------+
| enp1****pr0g | net  |
+----------------------+------+

root@deb10-gw-k11:~/virt-homeworks/05-virt-05-docker-swarm/src/terraform# yc vpc network delete enp1****pr0g
```
  
</details>   
  
#### 2.2. Создадим ВМ:  
  
<details>
  
Перейдем в каталог с конфигурацией terraform и проверим правильность значений в файле `variables.tf`  
  
```shell

root@D10:~/virt-homeworks/05-virt-05-docker-swarm/src/packer# cd ../terraform/

root@D10:~/virt-homeworks/05-virt-05-docker-swarm/src/terraform# cat variables.tf
# https://console.cloud.yandex.ru/cloud?section=overview
variable "yandex_cloud_id" {
  default = "b1g3a****nf2q"
}

# https://console.cloud.yandex.ru/cloud?section=overview
variable "yandex_folder_id" {
  default = "b1g8******0ih44"
}

# Image ID (from console YC or from command `yc compute image list`)
variable "centos-7-base" {
  default = "fd879*****jamnl"
}
```
  
Для автоматического развёртывания используем ранее сгенерированный ключ `key.json` (см. ["Получение IAM-токена для сервисного аккаунта"](https://cloud.yandex.ru/docs/iam/operations/iam-token/create-for-sa) )
  
Выполним инициализацию terraform:  
  
```shell
root@D10:~/virt-homeworks/05-virt-05-docker-swarm/src/terraform#  terraform init

Initializing the backend...

Initializing provider plugins...
- Reusing previous version of yandex-cloud/yandex from the dependency lock file
- Reusing previous version of hashicorp/null from the dependency lock file
- Reusing previous version of hashicorp/local from the dependency lock file
- Using previously-installed yandex-cloud/yandex v0.71.0
- Using previously-installed hashicorp/null v3.1.0
- Using previously-installed hashicorp/local v2.1.0

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```
  
Выполним предварительную оценку развертывания командой `terraform plan` (сокращённый вывод):  
  
```shell
root@D10:~/virt-homeworks/05-virt-05-docker-swarm/src/terraform#  terraform plan

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

...

Plan: 13 to add, 0 to change, 0 to destroy.

Changes to Outputs:
  + external_ip_address_node01 = (known after apply)
  + external_ip_address_node02 = (known after apply)
  + external_ip_address_node03 = (known after apply)
  + external_ip_address_node04 = (known after apply)
  + external_ip_address_node05 = (known after apply)
  + external_ip_address_node06 = (known after apply)
  + internal_ip_address_node01 = "192.168.101.11"
  + internal_ip_address_node02 = "192.168.101.12"
  + internal_ip_address_node03 = "192.168.101.13"
  + internal_ip_address_node04 = "192.168.101.14"
  + internal_ip_address_node05 = "192.168.101.15"
  + internal_ip_address_node06 = "192.168.101.16"


Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these actions if you run "terraform apply" now.

```
  
И запустим развёртывание командой `terraform apply` (сокращённый вывод):  
  
```shell
root@D10:~/virt-homeworks/05-virt-05-docker-swarm/src/terraform#  terraform apply -auto-approve

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create

Terraform will perform the following actions:

...

Apply complete! Resources: 13 added, 0 changed, 0 destroyed.

Outputs:

external_ip_address_node01 = "51.250.0.8"
external_ip_address_node02 = "51.250.1.222"
external_ip_address_node03 = "51.250.4.181"
external_ip_address_node04 = "51.250.1.61"
external_ip_address_node05 = "51.250.2.21"
external_ip_address_node06 = "51.250.11.245"
internal_ip_address_node01 = "192.168.101.11"
internal_ip_address_node02 = "192.168.101.12"
internal_ip_address_node03 = "192.168.101.13"
internal_ip_address_node04 = "192.168.101.14"
internal_ip_address_node05 = "192.168.101.15"
internal_ip_address_node06 = "192.168.101.16"

```
  
</details>  
  
#### 2.3. Подключимся к ВМ node01, повысим привилегии и выполним команду `docker node ls`  
  
![Результат выполнения команды `docker node ls`](Imgs/HW5.5-img1.png "HW5.5-img1")
  
___

## Задача 3

Создать ваш первый, готовый к боевой эксплуатации кластер мониторинга, состоящий из стека микросервисов.

Для получения зачета, вам необходимо предоставить скриншот из терминала (консоли), с выводом команды:
```
docker service ls
```
  
### Решение  
  
В консоли ВМ node01 выполним команду `docker service ls`:  
  
![Результат выполнения команды `docker service ls`](Imgs/HW5.5-img2.png "HW5.5-img2") 
  
___

## Задача 4 (*)

Выполнить на лидере Docker Swarm кластера команду (указанную ниже) и дать письменное описание её функционала, что она делает и зачем она нужна:
```
# см.документацию: https://docs.docker.com/engine/swarm/swarm_manager_locking/
docker swarm update --autolock=true
```
