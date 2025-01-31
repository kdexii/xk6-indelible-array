# xk6-indelible-array

Это небольшое расширение для k6 создано для более быстрой и удобной работы с массиов данных.

Самое главное отличие от SharedArray, это то, что в него можно добавлять данные во время теста и получать весь список данных в других сценариях.

На данный момент имеется всего 3 метода:
```javascript
isMyArray.initIndelibleArray() // Инициализировать массив данных
isMyArray.pushToIndelibleArray("123") // Добавить в существующий массив любой тип данных
isMyArray.getIndelibleArray() // Вывести весь массив в формате ["123", 123, 12.3]
```
## Требования

Для запуска скрипта потребуется: 📋

- [k6](https://k6.io/docs/) версии 0.45.0 или выше. ✅
- [golang](https://go.dev/dl/) версии 1.23 или выше. ✅

## Установка

1. Сборка k6 с необходимыми расширениями: 🏗️⚙️🚀

```bash
go install go.k6.io/xk6/cmd/xk6@latest
```

```bash
xk6 build --with github.com/kdexii/xk6-indelible-array@latest
```

## Пример скрипта k6

#### script.js

```javascript
import isMyArray from 'k6/x/xk6-indelible-array';

isMyArray.initIndelibleArray()
let isMyArrayDurationMinutes = 1

export let options = {
    scenarios: {
        add: {
            executor: 'constant-arrival-rate',
            rate: 2,
            timeUnit: '2s',
            duration: `${isMyArrayDurationMinutes}m`,
            preAllocatedVUs: 10,
            exec: 'addId',
        },
        get: {
            executor: 'constant-arrival-rate',
            rate: 1,
            timeUnit: '5s',
            duration: `${isMyArrayDurationMinutes}m`,
            preAllocatedVUs: 10,
            exec: 'getId',
            startTime: '2s',
        },
    },
};


export function addId() {
    isMyArray.pushToIndelibleArray(`ids_${__VU}`)
}

export function getId() {
    for (const currentId of isMyArray.getIndelibleArray()) {
        console.log(currentId)
    }
}
```

## Пример вывода в консоль
```bash
INFO[0008] ids_15                                        source=console
INFO[0008] ids_16                                        source=console
INFO[0008] ids_12                                        source=console
```

## Запуск теста
Запустите скрипт с помощью k6: 🚀📊

```bash
./k6 run script.js
```