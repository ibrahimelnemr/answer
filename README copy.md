# Apache Answer - Enhanced Edition

## What This Is

This is a customized version of Apache Answer featuring a hierarchical tags system with three levels (Offerings → Specializations → Topics) that organizes questions in a sortable, structured way. The system includes role-based access (admins create tags, users create questions), an accordion browse view with filtering, and comes with sample data for demo purposes. Additionally, it includes streamlined Docker deployment options for both interactive development and automated production environments.

## Windows Users

**Having trouble with `make` on Windows?** See [WINDOWS_BUILD.md](WINDOWS_BUILD.md) for Windows-specific build instructions using batch scripts.

## Getting Started

### Interactive Build (Dockerfile.local)
For development where you manually build and run:
```bash
./docker-local-start.sh              # Build image and enter container
make build                           # Inside: build Linux binary
./answer run -C ./answer-data        # Inside: start server
```
Access at http://localhost:9080

### Automated Build (Dockerfile)
For production where everything runs automatically:
```bash
./docker-start.sh                    # Build image, start server automatically
```
Access at http://localhost:9080

### Local Development (No Docker)

**Linux/Mac:**
```bash
make build                           # Build backend
./answer run -C ./answer-data        # Run backend

cd ui && pnpm install && pnpm start  # In another terminal: run frontend
```

**Windows (cmd.exe):**
```cmd
build-all.bat                        # Build both UI and backend
answer.exe run -C .\answer-data      # Run backend

REM Or build separately:
build-ui.bat                         # Build UI only
build.bat                            # Build backend only
```

## Docker Hub Publishing

Both scripts support pushing to Docker Hub. Ensure you're logged in first:
```bash
docker login
```

Then run with your Docker Hub username:
```bash
./docker-local-start.sh push <username>    # Push apache-local image
./docker-start.sh push <username>          # Push apache-automated image
```

## Update - Jan 14 2025

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



