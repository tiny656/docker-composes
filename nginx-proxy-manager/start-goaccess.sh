#!/bin/sh

# 构建日志文件参数列表
LOG_FILES=""
for log in /opt/logs/proxy-host-*_access.log; do
    LOG_FILES="$LOG_FILES --log-file=$log"
done

# 启动 GoAccess websocket with 443 port
exec goaccess \
    $LOG_FILES \
    --log-format='[%d:%t %^] %C %^ %s %^ %m %^ %v "%U" [%^ %h] [%^ %b] %^"%u" "%R"' \
    --date-format='%d/%b/%Y' \
    --time-format='%T' \
    --no-strict-status \
    --real-time-html \
    --output=/report/index.html \
    --ws-url=wss://goaccess-ws.<your domain> \
    --port=443 \
    --ssl-cert=/certs/fullchain.pem \
    --ssl-key=/certs/privkey.pem \
    --no-global-config