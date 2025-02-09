# Container monitoring

Приложение для мониторинга контейнеров, которое получает IP адреса контейнеров Docker, пингует их с определенными
интервалами и помещает данные в базу данных.

Приложение состоит из 4 сервисов:
- Backend-сервис - RESTful API для запроса данных из БД и добавления туда новых данных;
- Frontend-сервис - отображение данных, взятых через API Backend, в виде таблицы: IP адрес, время пинга, дата 
последней успешной попытки;
- Pinger сервис - пинг контейнеров по IP адресам и отправка данных через API Backend;
- База данных - PostgreSQL.

Для каждого сервиса созданы Dockerfile и общий файл compose, которые собирают образы и запускают сервисы. После этого
можно зайти через http на порт 3000 и увидеть данные о результатах опроса на динамически формируемой веб-странице.

### Запуск

Для запуска приложения необходимо:
1. Склонировать репозиторий
```bash
git clone https://github.com/DeaPacis/container_monitoring.git
```
2. Перейти в директорию
```bash
cd container_monitoring
```
3. Создать .env файл на основе .env.example файла
```bash
cp .env.example .env
```
4. Запустить compose файл
```bash
docker compose up -d
```
5. Открыть веб-страницу http://localhost:3000/