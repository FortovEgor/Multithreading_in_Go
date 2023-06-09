## Многопоточность Go

Данный репозиторий содержит примеры с пояснением ключевых элементов многопоточности в Go. В Go нет как таковых потоков и процессов, есть горутины - легковесные потоки/процессы, которыми управляет планировщик горутин. Горутины сами распределяются на потоки и процессы ОС, разработчику не нужно об этом думать. Такая концепция помогает сэкономить ресурсы компьютера и освободить код от ошибок программиста при создании потоков/процессов.

### Основные элементы синхронизации между горутинами - это каналы. 
### Канал - это потокобезопасный способ передачи данных между горутинами. Каналы бывают двух типов: 
- Небуферизированные (используются в 99% случаев)
- Буферизированные (1% случаев) (<=> небуферизированный канал + буфер)

### Другие элементы многопоточности в Go:
- Атомарные операции (пакет sync/atomic, работает быстрее мьютексов)
- Мьютексы (пакет sync, нужен для локальной блокировки участка кода, который может обрабатывать лишь одна горутина/процесс/поток)
- WaitGroupes (похожи на семафоры в POSIX и SYSTEM V, позволяют контролировать число горутин, которое точно должно выполнится; хороший аналог fmt.Scan() в конце файла, который используется для того, чтобы дать всем запущенным горутинам завершить выполнение)

### Важное практическое замечание!
Блокировка (deadlock) может возникать в любой горутине, однако компилятор считает за блокировку только блокировки, которые происходят в ГЛАВНОЙ горутине.

### Опции при запуске программы
Для отслеживания состояния гонки в приведенных примерах (при наличии взаимодействия горутин с общим ресурсом) программы следует запускать с флагом -race:
```sh
go run -race prog_name.go
```


##### "А зачем?"
Многопоточность - важная тема в Go, ведь данный язык создан по большей части для создания веб-сервисов (в частности, микросервисов), а в них многопоточность начинается с самого начала (например, обработка входящих соединений). К сожалению, в сети нет годных примеров, которые помогли бы программисту разобраться с этой темой. Теперь этот пробел заполнен :)
