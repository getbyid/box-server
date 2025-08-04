# HTTP-сервер для сжатых веб-сайтов

Архивация веб-сайтов позволяет сохранить информацию от исчезновения или блокировки. Вы можете создавать локальные копии сайтов разными способами, включая утилиты [wget](https://www.gnu.org/software/wget/) и [httrack](https://www.httrack.com/). При этом появляется несколько сложностей при дальнейшем хранении и открытии копий сайтов:

+ Веб-сайт может состоять из множества папок и небольших файлов. Удобнее упаковать всё это в единый архив.
+ Часть ссылок внутри сайта может быть указана от корня сайта. Нормальный просмотр оффлайн возможен через локальный веб-сервер для статических файлов, что-то типа `python3 -m http.server 8080` или `php -S localhost:8080`.

Как вариант решения — эта небольшая утилита, которая:

+ Запускает веб-сервер на указанном порту (по умолчанию `8080`)
+ При запросе корня каталога выдаёт индексный файл из него (по умолчанию `index.html`)
+ Если в корне архива обнаружится единственный каталог, то перемещает корень веб-сайта внутрь этого каталога

## Пример использования

Итак, делаем копию сайта:

~~~bash
wget -mpEk https://gobyexample.com/
~~~

Упаковываем в zip-архив:

~~~bash
zip -r gobyexample.zip gobyexample.com/
~~~

Запускаем оффлайн на любой системе, хоть даже на Android (через Termux):

~~~bash
$ box-server gobyexample.zip 
2025/08/04 12:34:00 Server is starting on http://localhost:8080
2025/08/04 12:34:56 GET /
2025/08/04 12:34:56 --> text/html 7257 bytes
2025/08/04 12:34:56 GET /site.css
2025/08/04 12:34:56 --> text/css 5589 bytes
2025/08/04 12:34:56 GET /favicon.ico
2025/08/04 12:34:56 not found: /favicon.ico
~~~

## Выбор названия

Первый вариант названия утилиты появился из подобия: раз веб-сайты раздаются через **веб-сервер**, значит zip-файлы с сайтами внутри можно считать zip-сайтами и раздавать через **zip-server**. Этот вариант сомнителен из-за опасной близости с утилитами zip, да и для автодополнения в bash необходимо 4 символа:

~~~bash
$ zip<TAB><TAB>
zip         zipcloak    zipdetails  zipgrep     zipinfo     zipnote     zipsplit  
~~~

Бережно хранимый архив с сайтом внутри можно считать цифровой шкатулкой (box) и для открытия использовать **box-server**. Для автодополнения в bash достаточно 3 символа:

~~~bash
$ box-server 
Usage: box-server website.zip
Options:
  -index string
        file for root path (default "index.html")
  -port int
        listen port (default 8080)
~~~

## Подобные проекты

Я не нашёл готовых решений и поэтому сочинил этот небольшой проект. Всё же есть несколько интересных работ по данной теме:

### [zipweb](https://github.com/jgreco/zipweb)

Ровно та же самая задача! Но я ничего не смыслю в Scheme, да и репозиторий заброшен:

> Possible usecase: You've archived a website using wget and zipped it so it wouldn't clutter your harddrive, but you still want to browse your archive.

### [redbean](https://redbean.dev/)

Крайне интересная идея! Достаточно взять исполняемый COM-файл веб-сервера, который также является zip-архивом и добавить к нему командой `zip` каталог с файлами сайта.

> Basic idea is if you want to build a web app that runs anywhere, then you download the redbean.com file, put your .html and .lua files inside it using the zip command, and you've got a hermetic app you deploy and share.
