
# Домашнее задание к занятию "5.2. Применение принципов IaaC в работе с виртуальными машинами"


## Задача 1

- Опишите своими словами основные преимущества применения на практике IaaC паттернов.
- Какой из принципов IaaC является основополагающим?
  
### Решение:
  
***Опишите своими словами основные преимущества применения на практике IaaC паттернов.***

Основные преимущества IaaC:  

1. **Ускорение производства и вывода продукта на рынок.**  
  
Стандартизация компонентов среды со стороны заказчика упрощает работу с ней и ускоряет создание требуемой инфраструктуры  

2. **Стабильность среды, устранение дрейфа конфигураций.**  
  
Уменьшение человеческого фактора, стандартизация компонентов среды для заказчика, повторяемость решений на разном железе.  

3. **Более быстрая и эффективная разработка.**  

  Ускорение достигается за счёт более формализированного описания и создания сред (разработки, тестирования и продуктовых)
и возможности оперативно менять возможностьи среды в зависимости от потребности в ресурсах.  
  
___
## Задача 2

- Чем Ansible выгодно отличается от других систем управление конфигурациями?
- Какой, на ваш взгляд, метод работы систем конфигурации более надёжный push или pull?
  
### Решение:
  
***Чем Ansible выгодно отличается от других систем управление конфигурациями?***  
- Работа без агентов. На управляемые узлы не нужно устанавливать никакого дополнительного ПО, всё работает через SSH;  
- Идемпотентность (повторяемость результатов при повторном применении)  
- Относительная простота освоения Python и YAML; относительно низкий порог вхождения: обучиться работе с Ansible можно за очень короткое время;  
- Хорошая документация и поддержка большим количеством производителей;  
- Возможность работы в push и pull режимах;  
- Поддержка последовательного обновления состояния узлов (rolling update).  
  
***Какой, на ваш взгляд, метод работы систем конфигурации более надёжный push или pull?***
  
По материалам лекции метод Push кажется более надёжным, но сложнее масштабируюшимся.
___
## Задача 3

### Установить на личный компьютер:

- VirtualBox
- Vagrant
- Ansible

#### Установка Oracle VirtualBox
```shell
sudo apt install virtualbox -y
```
#### Установка Vagrant из apt repository

Установим зависимости для работы с дополнительными репозиториями:
```shell
sudo apt update
sudo apt -y install apt-transport-https ca-certificates curl software-properties-common
```
  
Добавим официальный репозиторий Vagrant в apt:
```shell
curl -fs https://apt.releases.hashicorp.com/gpg | sudo apt-key add -
sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"
```
  
Обновим `APT` и установим `Vagrant`:
```shell
sudo apt update
sudo apt install vagrant
```
  
#### Установка "Ansible"
```shell
sudo apt install -y ansible
```
  
### Приложить вывод команд установленных версий каждой из программ, оформленный в `markdown`.

```shell
root@ubhost50:~# vboxmanage --version
6.1.26_Ubuntur145957
root@ubhost50:~# vagrant version
Installed Version: 2.2.19
Latest Version: 2.2.19

You''re running an up-to-date version of Vagrant!
root@ubhost50:~# ansible --version
ansible 2.9.6
  config file = /etc/ansible/ansible.cfg
  configured module search path = ['/root/.ansible/plugins/modules', '/usr/share/ansible/plugins/modules']
  ansible python module location = /usr/lib/python3/dist-packages/ansible
  executable location = /usr/bin/ansible
  python version = 3.8.10 (default, Nov 26 2021, 20:14:08) [GCC 9.3.0]
root@ubhost50:~#

```
___
## Задача 4 (*)

Воспроизвести практическую часть лекции самостоятельно.

- Создать виртуальную машину.
- Зайти внутрь ВМ, убедиться, что Docker установлен с помощью команды
```
docker ps
```
  
## Решение
  
### Создать виртуальную машину  
  
Скопируем файлы конфигурации в каталог, где будет развёрнута ВМ и создадим ВМ запустив  команду  `vagrant up`  
  
```shell
sysadmin@ubhost50:~/GIT/virt-homeworks/05-virt-02-iaac$ cd /home/sysadmin/VMs/hw5.2/vagrant
sysadmin@ubhost50:~/VMs/hw5.2/vagrant$ nano Vagrantfile 
sysadmin@ubhost50:~/VMs/hw5.2/vagrant$ vagrant up
Bringing machine 'server1.netology' up with 'virtualbox' provider...
==> server1.netology: Box 'bento/ubuntu-20.04' could not be found. Attempting to find and install...
    server1.netology: Box Provider: virtualbox
    server1.netology: Box Version: >= 0
==> server1.netology: Loading metadata for box 'bento/ubuntu-20.04'
    server1.netology: URL: https://vagrantcloud.com/bento/ubuntu-20.04
==> server1.netology: Adding box 'bento/ubuntu-20.04' (v202112.19.0) for provider: virtualbox
    server1.netology: Downloading: https://vagrantcloud.com/bento/boxes/ubuntu-20.04/versions/202112.19.0/providers/virtualbox.box
==> server1.netology: Successfully added box 'bento/ubuntu-20.04' (v202112.19.0) for 'virtualbox'!
==> server1.netology: Importing base box 'bento/ubuntu-20.04'...
==> server1.netology: Matching MAC address for NAT networking...
==> server1.netology: Checking if box 'bento/ubuntu-20.04' version '202112.19.0' is up to date...
==> server1.netology: Setting the name of the VM: server1.netology
Vagrant is currently configured to create VirtualBox synced folders with
the `SharedFoldersEnableSymlinksCreate` option enabled. If the Vagrant
guest is not trusted, you may want to disable this option. For more
information on this option, please refer to the VirtualBox manual:

  https://www.virtualbox.org/manual/ch04.html#sharedfolders

This option can be disabled globally with an environment variable:

  VAGRANT_DISABLE_VBOXSYMLINKCREATE=1

or on a per folder basis within the Vagrantfile:

  config.vm.synced_folder '/host/path', '/guest/path', SharedFoldersEnableSymlinksCreate: false
==> server1.netology: Clearing any previously set network interfaces...
==> server1.netology: Preparing network interfaces based on configuration...
    server1.netology: Adapter 1: nat
    server1.netology: Adapter 2: hostonly
==> server1.netology: Forwarding ports...
    server1.netology: 22 (guest) => 20011 (host) (adapter 1)
    server1.netology: 22 (guest) => 2222 (host) (adapter 1)
==> server1.netology: Running 'pre-boot' VM customizations...
==> server1.netology: Booting VM...
==> server1.netology: Waiting for machine to boot. This may take a few minutes...
    server1.netology: SSH address: 127.0.0.1:2222
    server1.netology: SSH username: vagrant
    server1.netology: SSH auth method: private key
    server1.netology: 
    server1.netology: Vagrant insecure key detected. Vagrant will automatically replace
    server1.netology: this with a newly generated keypair for better security.
    server1.netology: 
    server1.netology: Inserting generated public key within guest...
    server1.netology: Removing insecure key from the guest if it''s present...
    server1.netology: Key inserted! Disconnecting and reconnecting using new SSH key...
==> server1.netology: Machine booted and ready!
==> server1.netology: Checking for guest additions in VM...
==> server1.netology: Setting hostname...
==> server1.netology: Configuring and enabling network interfaces...
==> server1.netology: Mounting shared folders...
    server1.netology: /vagrant => /home/sysadmin/VMs/hw5.2/vagrant
==> server1.netology: Running provisioner: ansible...
    server1.netology: Running ansible-playbook...

PLAY [nodes] *******************************************************************

TASK [Gathering Facts] *********************************************************
ok: [server1.netology]

TASK [Create directory for ssh-keys] *******************************************
ok: [server1.netology]

TASK [Adding rsa-key in /root/.ssh/authorized_keys] ****************************
changed: [server1.netology]

TASK [Checking DNS] ************************************************************
changed: [server1.netology]

TASK [Installing tools] ********************************************************
ok: [server1.netology] => (item=['git', 'curl'])

TASK [Installing docker] *******************************************************
changed: [server1.netology]

TASK [Add the current user to docker group] ************************************
changed: [server1.netology]

PLAY RECAP *********************************************************************
server1.netology           : ok=7    changed=4    unreachable=0    failed=0    skipped=0    rescued=0    ignored=0   
```
  
### Зайти внутрь ВМ, убедиться, что Docker установлен с помощью команды `docker ps`

```shell
sysadmin@ubhost50:~/VMs/hw5.2/vagrant$ vagrant ssh
Welcome to Ubuntu 20.04.3 LTS (GNU/Linux 5.4.0-91-generic x86_64)

 * Documentation:  https://help.ubuntu.com
 * Management:     https://landscape.canonical.com
 * Support:        https://ubuntu.com/advantage

 System information disabled due to load higher than 1.0


This system is built by the Bento project by Chef Software
More information can be found at https://github.com/chef/bento
Last login: Thu Jan 27 18:01:25 2022 from 10.0.2.2
```

  
```shell
vagrant@server1:~$ docker ps
CONTAINER ID   IMAGE     COMMAND   CREATED   STATUS    PORTS     NAMES
```
  
### Дополнительные команды из лабораторной работы, показанной на лекции:
  
```shell
vagrant@server1:~$ docker run hello-world
Unable to find image 'hello-world:latest' locally
latest: Pulling from library/hello-world
2db29710123e: Pull complete 
Digest: sha256:507ecde44b8eb741278274653120c2bf793b174c06ff4eaa672b713b3263477b
Status: Downloaded newer image for hello-world:latest

Hello from Docker!
This message shows that your installation appears to be working correctly.

To generate this message, Docker took the following steps:
 1. The Docker client contacted the Docker daemon.
 2. The Docker daemon pulled the "hello-world" image from the Docker Hub.
    (amd64)
 3. The Docker daemon created a new container from that image which runs the
    executable that produces the output you are currently reading.
 4. The Docker daemon streamed that output to the Docker client, which sent it
    to your terminal.

To try something more ambitious, you can run an Ubuntu container with:
 $ docker run -it ubuntu bash

Share images, automate workflows, and more with a free Docker ID:
 https://hub.docker.com/

For more examples and ideas, visit:
 https://docs.docker.com/get-started/
```
  
```shell
vagrant@server1:~$ cat /etc/*release
DISTRIB_ID=Ubuntu
DISTRIB_RELEASE=20.04
DISTRIB_CODENAME=focal
DISTRIB_DESCRIPTION="Ubuntu 20.04.3 LTS"
NAME="Ubuntu"
VERSION="20.04.3 LTS (Focal Fossa)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 20.04.3 LTS"
VERSION_ID="20.04"
HOME_URL="https://www.ubuntu.com/"
SUPPORT_URL="https://help.ubuntu.com/"
BUG_REPORT_URL="https://bugs.launchpad.net/ubuntu/"
PRIVACY_POLICY_URL="https://www.ubuntu.com/legal/terms-and-policies/privacy-policy"
VERSION_CODENAME=focal
UBUNTU_CODENAME=focal
vagrant@server1:~$ 

```