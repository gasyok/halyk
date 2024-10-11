#!/bin/bash

# Получение IP в зависимости от ОС
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    IP=$(hostname -I | awk '{print $1}')
elif [[ "$OSTYPE" == "darwin"* ]]; then
    IP=$(ifconfig | grep "inet " | grep -v 127.0.0.1 | awk '{print $2}' | head -n 1)
else
    echo "Неизвестная ОС. Скрипт поддерживает только Linux и MacOS."
    exit 1
fi

# Проверяем, есть ли .env файл
if [ ! -f .env ]; then
    touch .env
fi

# Функция для обновления или добавления строки в .env
update_or_add_env_var() {
    local key="$1"
    local value="$2"

    if grep -q "^$key=" .env; then
        sed -i '' "s/^$key=.*/$key=$value/" .env
    else
        echo "$key=$value" >>.env
    fi
}

# Заполняем нужные переменные
update_or_add_env_var "EXTERNAL_HOST" "$IP"
update_or_add_env_var "EXTERNAL_REST_PORT" "7445"
update_or_add_env_var "EXTERNAL_TIMEOUT" "5s"
update_or_add_env_var "EXTERNAL_HOST_COPY" "$IP"
update_or_add_env_var "EXTERNAL_REST_PORT_COPY" "7447"
update_or_add_env_var "DIVISION" "80.0"
update_or_add_env_var "SERVICE_REST_PORT" "8445"

# Вывод информации
echo "Заполнено .env:"
echo "EXTERNAL_HOST=$IP"
echo "EXTERNAL_REST_PORT=7445"
echo "EXTERNAL_TIMEOUT=5s"
echo "EXTERNAL_HOST_COPY=$IP"
echo "EXTERNAL_REST_PORT_COPY=7447"
echo "DIVISION=80.0"
echo "SERVICE_REST_PORT=8445"
