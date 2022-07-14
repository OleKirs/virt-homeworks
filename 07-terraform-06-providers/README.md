# Домашнее задание к занятию "7.6. Написание собственных провайдеров для Terraform."

Бывает, что 
* общедоступная документация по терраформ ресурсам не всегда достоверна,
* в документации не хватает каких-нибудь правил валидации или неточно описаны параметры,
* понадобиться использовать провайдер без официальной документации,
* может возникнуть необходимость написать свой провайдер для системы используемой в ваших проектах.   

## Задача 1. 
Давайте потренируемся читать исходный код AWS провайдера, который можно склонировать от сюда: 
[https://github.com/hashicorp/terraform-provider-aws.git](https://github.com/hashicorp/terraform-provider-aws.git).
Просто найдите нужные ресурсы в исходном коде и ответы на вопросы станут понятны.  


1. Найдите, где перечислены все доступные `resource` и `data_source`, приложите ссылку на эти строки в коде на 
гитхабе. 

### Ответ:  
1.1. Доступные `resource` перечислены в `ResourcesMap` [в текущей версии - начиная со строки 906](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/provider/provider.go#L906)  

> ResourcesMap: map[string]*schema.Resource{  

1.2. Доступные `data_source` перечислены в `DataSourcesMap` [в данной версии - начиная со строки 412](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/provider/provider.go#L412)  

> DataSourcesMap: map[string]*schema.Resource{

2. Для создания очереди сообщений SQS используется ресурс `aws_sqs_queue` у которого есть параметр `name`. 
* С каким другим параметром конфликтует `name`? Приложите строчку кода, в которой это указано.

>  Строка 87 в файле [queue.go](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/service/sqs/queue.go#L87)  
>  (строка `ConflictsWith: []string{"name_prefix"}` в блоке `name`)
> <pre><code>...
>"name": {  
> Type:          schema.TypeString,
>   Optional:      true,  
>   Computed:      true,  
>   ForceNew:      true,  
>   ConflictsWith: []string{"name_prefix"},
> ... 
> </code></pre>


* Какая максимальная длина имени? 
 
> Судя по отдельным частям кода, валидная длина имени с суффиксом `.fifo`
> не должна превышать 80 символов и подчиняться регулярному выражению:
> `^[0-9A-Za-z-_]+(\.fifo)?$`    
> Но в текущей версии провайдера проверки максимальной длины `name`
> в файле ресурса не нашёл иначе, чем в определении функции
    > `resourceQueueCustomizeDiff`
> [строка 427](https://github.com/hashicorp/terraform-provider-aws/blob/b9fd7aba413d3967d89f8d873432f910e5905bea/internal/service/sqs/queue.go#L427),
> 
> Однако, в файле [main/internal/service/schemas/schema.go](https://github.com/hashicorp/terraform-provider-aws/blob/main/internal/service/schemas/schema.go)
> есть схема с описанием `name` ( ( [строка 58](https://github.com/hashicorp/terraform-provider-aws/blob/b9fd7aba413d3967d89f8d873432f910e5905bea/internal/service/schemas/schema.go#L58) ). )
> в которой указано определение `validation.StringLenBetween(1, 385),`:
> ```commandline
>"name": {
>		Type:     schema.TypeString,
>		Required: true,
>		ForceNew: true,
>		ValidateFunc: validation.All(
>		    validation.StringLenBetween(1, 385),
>		    validation.StringMatch(regexp.MustCompile(`^[\.\-_A-Za-z@]+`), ""),
>		),
>``` 
> Не совсем понял механизмы импорта схемы, но полагаю,
> что максимальная длина имени сейчас технически составляет 385 символов или не определена (не валидируется).
>  
  

* Какому регулярному выражению должно подчиняться имя? 
 
> Аналогично, [/internal/service/schemas/schema.go, строка 59](https://github.com/hashicorp/terraform-provider-aws/blob/b9fd7aba413d3967d89f8d873432f910e5905bea/internal/service/schemas/schema.go#L58)
> ```
>  validation.StringMatch(regexp.MustCompile(`^[\.\-_A-Za-z@]+`), ""),
> ```
   

В более ранних версиях провайдера была отдельная функция `validateSQSQueueName` (есть в коммите c54155effe, в файле `./aws/validators.go`):
```commandline
func validateSQSQueueName(v interface{}, k string) (ws []string, errors []error) {
	value := v.(string)
	if len(value) > 80 {
		errors = append(errors, fmt.Errorf("%q cannot be longer than 80 characters", k))
	}

	if !regexp.MustCompile(`^[0-9A-Za-z-_]+(\.fifo)?$`).MatchString(value) {
		errors = append(errors, fmt.Errorf("only alphanumeric characters and hyphens allowed in %q", k))
	}
	return
}
```

Но сейчас её не нашёл, т.к. была удалена в коммите f47d5cb92b

```commandline
$ git log -S validateSQSQueueName --oneline
f47d5cb92b r/aws_sqs_queue: append .fifo suffix for FIFO queue if name unspecified (#17164)
c54155effe allow SQS name validation during terraform validate
5e1cdcca42 provider/aws: Added SQS FIFO queues (#10614)
```

## Задача 2. (Не обязательно) 
В рамках вебинара и презентации мы разобрали как создать свой собственный провайдер на примере кофемашины. 
Также вот официальная документация о создании провайдера: 
[https://learn.hashicorp.com/collections/terraform/providers](https://learn.hashicorp.com/collections/terraform/providers).


1. Проделайте все шаги создания провайдера.
2. В виде результата приложение ссылку на исходный код.
3. Попробуйте скомпилировать провайдер, если получится то приложите снимок экрана с командой и результатом компиляции.   

---

### Как cдавать задание

Выполненное домашнее задание пришлите ссылкой на .md-файл в вашем репозитории.

---
