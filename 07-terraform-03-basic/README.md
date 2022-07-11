# Домашнее задание к занятию "7.3. Основы и принцип работы Терраформ"


> Выполнял с использованием `Yandex Cloud`  
> ***Рабочий репозиторий:*** [https://github.com/OleKirs/07-terraform.git](https://github.com/OleKirs/07-terraform.git)



## Задача 1. Создадим бэкэнд в S3 (необязательно, но крайне желательно).

Если в рамках предыдущего задания у вас уже есть аккаунт AWS, то давайте продолжим знакомство со взаимодействием
терраформа и aws.
2. Создайте s3 бакет, iam роль и пользователя от которого будет работать терраформ. Можно создать отдельного пользователя,
а можно использовать созданного в рамках предыдущего задания, просто добавьте ему необходимы права, как описано 
[здесь](https://www.terraform.io/docs/backends/types/s3.html).
3. Зарегистрируйте бэкэнд в терраформ проекте как описано по ссылке выше. 

<details> 
  <summary>Часть кода из `main.tf`</summary>

```shell
terraform {
  # Moved to ./versions.tf
  #  required_providers {
  #    yandex = {
  #      source = "yandex-cloud/yandex"
  #    }
  #  }

  backend "s3" {
    endpoint = "storage.yandexcloud.net"
    bucket   = "olekirs-netology"
    region   = "ru-central1"
    key      = "[terraform.workspace]/terraform.tfstate"

    # Used \"AWS_ACCESS_KEY_ID\" and \"AWS_SECRET_ACCESS_KEY\" in user ENV
    #    access_key = "<идентификатор статического ключа>"
    #    secret_key = "<секретный ключ>"

    skip_region_validation      = true
    skip_credentials_validation = true
  }
}

```

</details>
  
## Задача 2. Инициализируем проект и создаем воркспейсы. 

1. Выполните `terraform init`:
    * если был создан бэкэнд в S3, то терраформ создат файл стейтов в S3 и запись в таблице 
dynamodb.
    * иначе будет создан локальный файл со стейтами. 

```shell
terraform init
```

<details> 
  <summary>Вывод команды</summary>

```shell
root@deb11-test50:~/olekirs/07-terraform# terraform init
Initializing modules...

Initializing the backend...

Successfully configured the backend "s3"! Terraform will automatically
use this backend unless the backend configuration changes.

Initializing provider plugins...
- Reusing previous version of yandex-cloud/yandex from the dependency lock file
- Using previously-installed yandex-cloud/yandex v0.61.0

Terraform has been successfully initialized!

You may now begin working with Terraform. Try running "terraform plan" to see
any changes that are required for your infrastructure. All Terraform commands
should now work.

If you ever set or change modules or backend configuration for Terraform,
rerun this command to reinitialize your working directory. If you forget, other
commands will detect it and remind you to do so if necessary.
```

</details>

1. Создайте два воркспейса `stage` и `prod`.

```shell
terraform workspace new prod
```

```shell
terraform workspace new stage
```

<details> 
  <summary>Вывод команд</summary>

```shell
root@deb11-test50:~/olekirs/07-terraform# terraform workspace list
* default
```

```shell

root@deb11-test50:~/olekirs/07-terraform# terraform workspace new prod
Created and switched to workspace "prod"!

You're now on a new, empty workspace. Workspaces isolate their state,
so if you run "terraform plan" Terraform will not see any existing state
for this configuration.

```

```shell

root@deb11-test50:~/olekirs/07-terraform# terraform workspace new stage
Created and switched to workspace "stage"!

You're now on a new, empty workspace. Workspaces isolate their state,
so if you run "terraform plan" Terraform will not see any existing state
for this configuration.
```

```shell

root@deb11-test50:~/olekirs/07-terraform# terraform workspace select prod
Switched to workspace "prod".
```

</details>

2. В уже созданный `aws_instance` добавьте зависимость типа инстанса от вокспейса, что бы в разных ворскспейсах 
использовались разные `instance_type`.
3. Добавим `count`. Для `stage` должен создаться один экземпляр `ec2`, а для `prod` два. 
4. Создайте рядом еще один `aws_instance`, но теперь определите их количество при помощи `for_each`, а не `count`.
5. Что бы при изменении типа инстанса не возникло ситуации, когда не будет ни одного инстанса добавьте параметр
жизненного цикла `create_before_destroy = true` в один из рессурсов `aws_instance`.
6. При желании поэкспериментируйте с другими параметрами и ресурсами.

В виде результата работы пришлите:
* Вывод команды `terraform workspace list`.

```shell
root@deb11-test50:~/olekirs/07-terraform# terraform workspace list
  default
* prod
  stage

```

* Вывод команды `terraform plan` для воркспейса `prod`.  

<details> 
  <summary>Вывод команды</summary>  

```shell
root@deb11-test50:~/olekirs/07-terraform# terraform plan
module.vpc.data.yandex_compute_image.nat_instance: Reading...
module.vpc.data.yandex_compute_image.nat_instance: Read complete after 2s [id=fd89681vdciaeqsurfhv]

Terraform used the selected providers to generate the following execution plan. Resource actions are indicated with the following symbols:
  + create
 <= read (data resources)

Terraform will perform the following actions:

  # module.instances.data.yandex_compute_image.image will be read during apply
  # (depends on a resource or a module with changes pending)
 <= data "yandex_compute_image" "image" {
      + created_at    = (known after apply)
      + description   = (known after apply)
      + family        = "ubuntu-2004-lts"
      + folder_id     = (known after apply)
      + id            = (known after apply)
      + image_id      = (known after apply)
      + labels        = (known after apply)
      + min_disk_size = (known after apply)
      + name          = (known after apply)
      + os_type       = (known after apply)
      + product_ids   = (known after apply)
      + size          = (known after apply)
      + status        = (known after apply)
    }

  # module.instances.yandex_compute_instance.instance[0] will be created
  + resource "yandex_compute_instance" "instance" {
      + created_at                = (known after apply)
      + description               = "Demo count"
      + folder_id                 = (known after apply)
      + fqdn                      = (known after apply)
      + hostname                  = "vm-c-1"
      + id                        = (known after apply)
      + metadata                  = {
          + "user-data" = <<-EOT
                sysadmin:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDFaeqemGCwv4VgiWEr4Ljlc+s9BozJSUIlHkmnYFBRo5uNk8qhibML9/q5LYqAhXUmZXw+PjOWOLP9GGyTi6a93qc5/GTuYipgqPAqf/Pc/kw87jm7ePJg4KIrWZ+hbSOguqkYEI9lwEbmaQKItvhp1ormF9df7FIfYGUc7cNJXF2crcrXLdrqFC/AWAfVNzzJUG4AanqGe64aCRnwC4evFOQTckDG2BmSJjCsPCdeL37XNiPT6Q5pTwiF3ani0vz+iM6As880xHOyqiCFDXE3U8PBfqTACzjPdmuYm3jhzLvZiJ0IQQnCT+OY+IW4nVWzucmKmx8KuEv7f3Zgwln1lvE5CwfoovUGf0B2qpOWMd8SZfSxDqcQRUKxNglnxKO2nRbuVTAQGG4HqpFWbqjBfFfeJk8c5vHAmgnTrFU4mF4eUu9bzx/frpZ2S62VxjJH71lBzPP6wGkRZkHm3D/3lQXoD+t5rMn3R6CT7Z72wo98MxVoLVdHWZmxHB60h4c= root@deb11-test50
            EOT
        }
      + name                      = "vm-c1"
      + network_acceleration_type = "standard"
      + platform_id               = "standard-v2"
      + service_account_id        = (known after apply)
      + status                    = (known after apply)
      + zone                      = "ru-central1-a"

      + boot_disk {
          + auto_delete = true
          + device_name = (known after apply)
          + disk_id     = (known after apply)
          + mode        = (known after apply)

          + initialize_params {
              + description = (known after apply)
              + image_id    = (known after apply)
              + name        = (known after apply)
              + size        = 20
              + snapshot_id = (known after apply)
              + type        = "network-ssd"
            }
        }

      + network_interface {
          + index              = (known after apply)
          + ip_address         = (known after apply)
          + ipv4               = true
          + ipv6               = false
          + ipv6_address       = (known after apply)
          + mac_address        = (known after apply)
          + nat                = true
          + nat_ip_address     = (known after apply)
          + nat_ip_version     = (known after apply)
          + security_group_ids = (known after apply)
          + subnet_id          = (known after apply)
        }

      + placement_policy {
          + placement_group_id = (known after apply)
        }

      + resources {
          + core_fraction = 100
          + cores         = 2
          + memory        = 2
        }

      + scheduling_policy {
          + preemptible = (known after apply)
        }
    }

  # module.instances.yandex_compute_instance.instance[1] will be created
  + resource "yandex_compute_instance" "instance" {
      + created_at                = (known after apply)
      + description               = "Demo count"
      + folder_id                 = (known after apply)
      + fqdn                      = (known after apply)
      + hostname                  = "vm-c-2"
      + id                        = (known after apply)
      + metadata                  = {
          + "user-data" = <<-EOT
                sysadmin:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDFaeqemGCwv4VgiWEr4Ljlc+s9BozJSUIlHkmnYFBRo5uNk8qhibML9/q5LYqAhXUmZXw+PjOWOLP9GGyTi6a93qc5/GTuYipgqPAqf/Pc/kw87jm7ePJg4KIrWZ+hbSOguqkYEI9lwEbmaQKItvhp1ormF9df7FIfYGUc7cNJXF2crcrXLdrqFC/AWAfVNzzJUG4AanqGe64aCRnwC4evFOQTckDG2BmSJjCsPCdeL37XNiPT6Q5pTwiF3ani0vz+iM6As880xHOyqiCFDXE3U8PBfqTACzjPdmuYm3jhzLvZiJ0IQQnCT+OY+IW4nVWzucmKmx8KuEv7f3Zgwln1lvE5CwfoovUGf0B2qpOWMd8SZfSxDqcQRUKxNglnxKO2nRbuVTAQGG4HqpFWbqjBfFfeJk8c5vHAmgnTrFU4mF4eUu9bzx/frpZ2S62VxjJH71lBzPP6wGkRZkHm3D/3lQXoD+t5rMn3R6CT7Z72wo98MxVoLVdHWZmxHB60h4c= root@deb11-test50
            EOT
        }
      + name                      = "vm-c2"
      + network_acceleration_type = "standard"
      + platform_id               = "standard-v2"
      + service_account_id        = (known after apply)
      + status                    = (known after apply)
      + zone                      = "ru-central1-a"

      + boot_disk {
          + auto_delete = true
          + device_name = (known after apply)
          + disk_id     = (known after apply)
          + mode        = (known after apply)

          + initialize_params {
              + description = (known after apply)
              + image_id    = (known after apply)
              + name        = (known after apply)
              + size        = 20
              + snapshot_id = (known after apply)
              + type        = "network-ssd"
            }
        }

      + network_interface {
          + index              = (known after apply)
          + ip_address         = (known after apply)
          + ipv4               = true
          + ipv6               = false
          + ipv6_address       = (known after apply)
          + mac_address        = (known after apply)
          + nat                = true
          + nat_ip_address     = (known after apply)
          + nat_ip_version     = (known after apply)
          + security_group_ids = (known after apply)
          + subnet_id          = (known after apply)
        }

      + placement_policy {
          + placement_group_id = (known after apply)
        }

      + resources {
          + core_fraction = 100
          + cores         = 2
          + memory        = 2
        }

      + scheduling_policy {
          + preemptible = (known after apply)
        }
    }

  # module.instances_fe.data.yandex_compute_image.image_fe will be read during apply
  # (depends on a resource or a module with changes pending)
 <= data "yandex_compute_image" "image_fe" {
      + created_at    = (known after apply)
      + description   = (known after apply)
      + family        = "ubuntu-2004-lts"
      + folder_id     = (known after apply)
      + id            = (known after apply)
      + image_id      = (known after apply)
      + labels        = (known after apply)
      + min_disk_size = (known after apply)
      + name          = (known after apply)
      + os_type       = (known after apply)
      + product_ids   = (known after apply)
      + size          = (known after apply)
      + status        = (known after apply)
    }

  # module.instances_fe.yandex_compute_instance.instance_fe["vm-fe-1"] will be created
  + resource "yandex_compute_instance" "instance_fe" {
      + created_at                = (known after apply)
      + description               = "Demo for_each"
      + folder_id                 = (known after apply)
      + fqdn                      = (known after apply)
      + hostname                  = "vm-fe-1"
      + id                        = (known after apply)
      + metadata                  = {
          + "user-data" = <<-EOT
                sysadmin:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDFaeqemGCwv4VgiWEr4Ljlc+s9BozJSUIlHkmnYFBRo5uNk8qhibML9/q5LYqAhXUmZXw+PjOWOLP9GGyTi6a93qc5/GTuYipgqPAqf/Pc/kw87jm7ePJg4KIrWZ+hbSOguqkYEI9lwEbmaQKItvhp1ormF9df7FIfYGUc7cNJXF2crcrXLdrqFC/AWAfVNzzJUG4AanqGe64aCRnwC4evFOQTckDG2BmSJjCsPCdeL37XNiPT6Q5pTwiF3ani0vz+iM6As880xHOyqiCFDXE3U8PBfqTACzjPdmuYm3jhzLvZiJ0IQQnCT+OY+IW4nVWzucmKmx8KuEv7f3Zgwln1lvE5CwfoovUGf0B2qpOWMd8SZfSxDqcQRUKxNglnxKO2nRbuVTAQGG4HqpFWbqjBfFfeJk8c5vHAmgnTrFU4mF4eUu9bzx/frpZ2S62VxjJH71lBzPP6wGkRZkHm3D/3lQXoD+t5rMn3R6CT7Z72wo98MxVoLVdHWZmxHB60h4c= root@deb11-test50
            EOT
        }
      + name                      = "vm-fe-1"
      + network_acceleration_type = "standard"
      + platform_id               = "standard-v2"
      + service_account_id        = (known after apply)
      + status                    = (known after apply)
      + zone                      = "ru-central1-a"

      + boot_disk {
          + auto_delete = true
          + device_name = (known after apply)
          + disk_id     = (known after apply)
          + mode        = (known after apply)

          + initialize_params {
              + description = (known after apply)
              + image_id    = (known after apply)
              + name        = (known after apply)
              + size        = 20
              + snapshot_id = (known after apply)
              + type        = "network-ssd"
            }
        }

      + network_interface {
          + index              = (known after apply)
          + ip_address         = (known after apply)
          + ipv4               = true
          + ipv6               = false
          + ipv6_address       = (known after apply)
          + mac_address        = (known after apply)
          + nat                = true
          + nat_ip_address     = (known after apply)
          + nat_ip_version     = (known after apply)
          + security_group_ids = (known after apply)
          + subnet_id          = (known after apply)
        }

      + placement_policy {
          + placement_group_id = (known after apply)
        }

      + resources {
          + core_fraction = 100
          + cores         = 2
          + memory        = 2
        }

      + scheduling_policy {
          + preemptible = (known after apply)
        }
    }

  # module.instances_fe.yandex_compute_instance.instance_fe["vm-fe-2"] will be created
  + resource "yandex_compute_instance" "instance_fe" {
      + created_at                = (known after apply)
      + description               = "Demo for_each"
      + folder_id                 = (known after apply)
      + fqdn                      = (known after apply)
      + hostname                  = "vm-fe-2"
      + id                        = (known after apply)
      + metadata                  = {
          + "user-data" = <<-EOT
                sysadmin:ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDFaeqemGCwv4VgiWEr4Ljlc+s9BozJSUIlHkmnYFBRo5uNk8qhibML9/q5LYqAhXUmZXw+PjOWOLP9GGyTi6a93qc5/GTuYipgqPAqf/Pc/kw87jm7ePJg4KIrWZ+hbSOguqkYEI9lwEbmaQKItvhp1ormF9df7FIfYGUc7cNJXF2crcrXLdrqFC/AWAfVNzzJUG4AanqGe64aCRnwC4evFOQTckDG2BmSJjCsPCdeL37XNiPT6Q5pTwiF3ani0vz+iM6As880xHOyqiCFDXE3U8PBfqTACzjPdmuYm3jhzLvZiJ0IQQnCT+OY+IW4nVWzucmKmx8KuEv7f3Zgwln1lvE5CwfoovUGf0B2qpOWMd8SZfSxDqcQRUKxNglnxKO2nRbuVTAQGG4HqpFWbqjBfFfeJk8c5vHAmgnTrFU4mF4eUu9bzx/frpZ2S62VxjJH71lBzPP6wGkRZkHm3D/3lQXoD+t5rMn3R6CT7Z72wo98MxVoLVdHWZmxHB60h4c= root@deb11-test50
            EOT
        }
      + name                      = "vm-fe-2"
      + network_acceleration_type = "standard"
      + platform_id               = "standard-v2"
      + service_account_id        = (known after apply)
      + status                    = (known after apply)
      + zone                      = "ru-central1-a"

      + boot_disk {
          + auto_delete = true
          + device_name = (known after apply)
          + disk_id     = (known after apply)
          + mode        = (known after apply)

          + initialize_params {
              + description = (known after apply)
              + image_id    = (known after apply)
              + name        = (known after apply)
              + size        = 20
              + snapshot_id = (known after apply)
              + type        = "network-ssd"
            }
        }

      + network_interface {
          + index              = (known after apply)
          + ip_address         = (known after apply)
          + ipv4               = true
          + ipv6               = false
          + ipv6_address       = (known after apply)
          + mac_address        = (known after apply)
          + nat                = true
          + nat_ip_address     = (known after apply)
          + nat_ip_version     = (known after apply)
          + security_group_ids = (known after apply)
          + subnet_id          = (known after apply)
        }

      + placement_policy {
          + placement_group_id = (known after apply)
        }

      + resources {
          + core_fraction = 100
          + cores         = 2
          + memory        = 2
        }

      + scheduling_policy {
          + preemptible = (known after apply)
        }
    }

  # module.vpc.yandex_resourcemanager_folder.folder[0] will be created
  + resource "yandex_resourcemanager_folder" "folder" {
      + cloud_id    = (known after apply)
      + created_at  = (known after apply)
      + description = "terraform managed"
      + id          = (known after apply)
      + name        = "prod"
    }

  # module.vpc.yandex_vpc_network.this will be created
  + resource "yandex_vpc_network" "this" {
      + created_at                = (known after apply)
      + default_security_group_id = (known after apply)
      + description               = "managed by terraform prod network"
      + folder_id                 = (known after apply)
      + id                        = (known after apply)
      + name                      = "prod"
      + subnet_ids                = (known after apply)
    }

  # module.vpc.yandex_vpc_subnet.this["ru-central1-a"] will be created
  + resource "yandex_vpc_subnet" "this" {
      + created_at     = (known after apply)
      + description    = "managed by terraform prod subnet for zone ru-central1-a"
      + folder_id      = (known after apply)
      + id             = (known after apply)
      + name           = "prod-ru-central1-a"
      + network_id     = (known after apply)
      + v4_cidr_blocks = [
          + "10.128.0.0/24",
        ]
      + v6_cidr_blocks = (known after apply)
      + zone           = "ru-central1-a"
    }

  # module.vpc.yandex_vpc_subnet.this["ru-central1-b"] will be created
  + resource "yandex_vpc_subnet" "this" {
      + created_at     = (known after apply)
      + description    = "managed by terraform prod subnet for zone ru-central1-b"
      + folder_id      = (known after apply)
      + id             = (known after apply)
      + name           = "prod-ru-central1-b"
      + network_id     = (known after apply)
      + v4_cidr_blocks = [
          + "10.129.0.0/24",
        ]
      + v6_cidr_blocks = (known after apply)
      + zone           = "ru-central1-b"
    }

  # module.vpc.yandex_vpc_subnet.this["ru-central1-c"] will be created
  + resource "yandex_vpc_subnet" "this" {
      + created_at     = (known after apply)
      + description    = "managed by terraform prod subnet for zone ru-central1-c"
      + folder_id      = (known after apply)
      + id             = (known after apply)
      + name           = "prod-ru-central1-c"
      + network_id     = (known after apply)
      + v4_cidr_blocks = [
          + "10.130.0.0/24",
        ]
      + v6_cidr_blocks = (known after apply)
      + zone           = "ru-central1-c"
    }

Plan: 9 to add, 0 to change, 0 to destroy.

────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

Note: You didn't use the -out option to save this plan, so Terraform can't guarantee to take exactly these actions if you run "terraform apply" now.
root@deb11-test50:~/olekirs/07-terraform#

```

</details>

---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
````
