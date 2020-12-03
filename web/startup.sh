#!/bin/bash
echo "Starting Services" && \
/opt/remi/php72/root/usr/sbin/php-fpm >> startup.log && \
echo "Started php-fpm" && \
/usr/sbin/nginx >> startup.log && \
echo "Started Nginx" && \

tailf /var/log/nginx/error.log
