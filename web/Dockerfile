FROM centos:7
COPY ./moodle.conf /etc/nginx/conf.d/moodle.conf

RUN mkdir -p /var/www/html/moodle/www && mkdir -p /var/www/html/moodle/moodledata

#install nginx
RUN yum update -y && \
yum install -y epel-release && \
yum install -y nginx && \
yum install -y http://rpms.famillecollet.com/enterprise/remi-release-7.rpm && \
yum install -y php-common \
php72-php-fpm \
php72-php-mysql \
php72-php-pecl-memcache \
php72-php-pecl-memcached \
php72-php-gd \
php72-php-mbstring \
php72-php-mcrypt \
php72-php-xml \
php72-php-pecl-apc \
php72-php-pear \
php72-php-pdo \
php72-php-intl \
php72-php-zip \
php72-php-pgsql \
php72-php-iconv \
php72-php-cli \
php72-php-curl \
php72-php-openssl \
php72-php-tokenizer \
php72-php-xmlrpc \
php72-php-soap \
php72-php-ctype \
php72-php-simplexml \
php72-php-spl \
php72-php-pcre \
php72-php-dom \
php72-php-json \
php72-php-pecl-redis \
php72-php-redis \
php72-php-opcache \
php-redis \
unzip && \
sed -i 's/apache/nginx/g' /etc/opt/remi/php72/php-fpm.d/www.conf && \
sed -i 's/default_server/ /g' /etc/nginx/nginx.conf

COPY ./moodle /var/www/html/moodle/www
RUN chown -R nginx:nginx /var/www/html/moodle

#ENTRYPOINT "/opt/remi/php72/root/usr/sbin/php-fpm && /usr/sbin/nginx"
COPY ./startup.sh /usr/local/
RUN chmod 777 /usr/local/startup.sh
ENTRYPOINT ["/usr/local/startup.sh"]


EXPOSE 80
