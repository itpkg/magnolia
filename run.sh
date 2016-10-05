#!/bin/sh

JAVA_HOME=/opt/jdk8
PATH=$JAVA_HOME:$PATH
export JAVA_HOME PATH

java -cp 'magnolia-app-2016.10.05.jar:config' org.springframework.boot.loader.JarLauncher
