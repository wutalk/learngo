#!/bin/bash

echo "enable system proxy"
sudo cp -pf /etc/environment.proxy /etc/environment

echo "enable docker proxy"
cp -pf /etc/systemd/system/docker.service.d/http-proxy.conf.proxy /etc/systemd/system/docker.service.d/http-proxy.conf
systemctl daemon-reload
systemctl restart docker

echo "swith git config to gerrite"
cp -pf /home/owuv/.gitconfig-neo /home/owuv/.gitconfig

echo "logout to take effect"

