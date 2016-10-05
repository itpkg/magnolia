# Magnolia
A collections of web applications.


## Prepare database

```
psql -U postgres
CREATE DATABASE db-name WITH ENCODING = 'UTF8';
CREATE USER user-name WITH PASSWORD 'change-me';
GRANT ALL PRIVILEGES ON DATABASE db-name TO user-name;
```

## Deployment

```
make clean
make
cd dist
ls config # edit configuration file
./run.sh

firefox http://your-hostname:8080
```
