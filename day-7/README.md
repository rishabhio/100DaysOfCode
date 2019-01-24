Did redis setup on ubuntu 16.04 

steps to install redis on ubuntu 16.04 

1. sudo apt-get update
2. sudo apt-get install build-essential tcl 

3. cd /tmp

4. curl -O http://download.redis.io/redis-stable.tar.gz

5. tar xzvf redis-stable.tar.gz

6. cd redis-stable

7. make 

8. make test

9. sudo make install 

Configure Redis 

10. sudo mkdir /etc/redis

11. sudo cp /tmp/redis-stable/redis.conf /etc/redis

12. sudo nano /etc/redis/redis.conf

13. configuration options 

supervised systemd
dir /var/lib/redis

sudo pico /etc/systemd/system/redis.service

[Unit]
Description=Redis In-Memory Data Store
After=network.target

[Service]
User=redis
Group=redis
ExecStart=/usr/local/bin/redis-server /etc/redis/redis.conf
ExecStop=/usr/local/bin/redis-cli shutdown
Restart=always

[Install]
WantedBy=multi-user.target




sudo adduser --system --group --no-create-home redis 

create /var/lib/redis directory 
sudo mkdir /var/lib/redis

sudo chown redis:redis /var/lib/redis

sudo chmod 770 /var/lib/redis

sudo systemctl start redis
sudo systemctl status redis



redis-cli 







