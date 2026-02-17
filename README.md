# Apache Answer - Enhanced Edition

## Overview

This is a customized version of Apache Answer featuring a hierarchical tags system with three levels (Offerings → Specializations → Topics) that organizes questions in a sortable, structured way.

## Jan 14 2026 - MacOS Development Setup

Currently working on MacOS with the following

`brew install wget gnupg node fuse`

`brew install node@20`

`go install go.uber.org/mock/mockgen@latest`

`go install github.com/google/wire/cmd/wire@0.5.0`

`make generate`

`make ui`

`make build`


Note - if you get this error

```
config file path:  answer-data/conf/config.yaml
Answer is starting..........................
2026-01-14 15:14:55.856	INFO	data/data.go:118	try to load cache file from ./answer-data/cache/cache.db
2026-01-14 15:14:55.858	INFO	cron/cron.go:63	cron job manager start
2026-01-14 15:14:55.858	INFO	cron/cron.go:99	clean up uploads cron enabled
answer Version: 1.6.0  Revision: ef0b3533
2026-01-14 15:14:57.009	ERROR	router/ui.go:138	open build/index.html: file does not exist
2026-01-14 15:14:58.379	ERROR	router/ui.go:138	open build/index.html: file does not exist
2026-01-14 15:15:02.627	ERROR	router/ui.go:138	open build/index.html: file does not exist
```

Then 

`cd ui`

`npm i`

`npm run build`

then run 

`ls build`

if there is an `index.html` file in that folder the error is resolved

note for frontend build it is also possible to run

`cd ui`

`pnpm start`

and then access the UI on a development server at `localhost:3000`

then to build the backend with the ui

`cd ..`

`make build`

`./answer run -C ./answer-data`

now the build should be complete

note that for some reason running `pnpm build` seems to have failed to generate the index.html file in this directory

if it says
`panic: stat answer-data/conf/config.yaml: no such file or directory`

then run 
`./answer init -C ./answer-data`

and access localhost:80 to install

set the database file to `answer.db`

set the contact email to `admin@admin.com`

set the admin account to `admin` `adminadmin` and email `admin@admin.com`

then when signing in use account

email
`admin@admin.com`
password
`adminadmin`

## Feb 17 2026 - Demo Users

create second user with credentials

user@user.com
useruser


## Feb 17 2026 - Setup SMTP Server

Run `docker run --rm -p 1025:1025 -p 8025:8025 axllent/mailpit:latest`

Then login as an admin and go to the SMTP settings at http://localhost/admin/smtp

Then set 

From email no-reply@answer-local.test
From name Apache Answer (Local)
SMTP host localhost
Encryption none
SMTP port 1025

To test the SMTP server is working, attempt to register as a new user and check mailpit at http://localhost:8025/ to check if the verification email appears