# GORM JWT API

This is a skeleton application to build a JWT auth system using GORM, and GO.

## .env file info
When deploying, please rename .sampleenv to .env


PORT= *Any port number you would like, 8000 or higher*

DB_ENGINE= *Options SQLITE POSTGRESQL MYSQL*

POSTGRESDSN=*"Please check DSN Format from https://gorm.io/docs/connecting_to_the_database.html#PostgreSQL"*

MYSQLDSN=*"Please check DSN Format from https://gorm.io/docs/connecting_to_the_database.html#MySQL"*

SQLITEDBNAME="./db-data/*filename.db*"

AUTHTYPE="*INTERNAL or LDAP*"

LDAPSERVER="*ldap.example.com*"

LDAPPORT="*389*"

LDAPBASEDN="*OU=Users,DC=example,DC=com*"

LDAPSECURITY="*auth.SecurityStartTLS*"