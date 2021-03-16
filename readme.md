## UI test chromedp

Эксперимент по написанию Web-UI тестов на Go с использованием библиотеки  
https://github.com/chromedp/chromedp

#### Steps:
1) пример из https://github.com/chromedp/examples
2) go get -u github.com/chromedp/chromedp
3) Goland - run - запускается сразу 
4) go build -o uitest.exe
5) docker build -t uitest .
6) docker run -it --entrypoint /bin/sh uitest
7) docker run -it uitest sh

#### Результаты:
- alpine + chromium + uitest - 380MB
- alpine + uitest  - 18.7MB
- scratch + uitest - 13.1MB
- uitest - 12.5MB

#### Итого:
- Рабочий вариант весит 380Мб
- Генератора, как в Playwright нет
- Реализовать паттерн Page Object возможно аналогично https://github.com/sayems/golang.webdriver  
- Для написания небольших задачек подходит, но при условии запуска на локальном ПК, где уже есть Chrome
- Для крупного проекта не думаю, что стоит