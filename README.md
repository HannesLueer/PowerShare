# PowerShare

## About the App

Neben kommerziellen Ladestationen finden private Lademöglichkeiten (Wallboxen) zunehmend Verbreitung. Kommerzielle Ladestationen können über diverse Apps gefunden werden; dies gilt jedoch nicht für private Wallboxen. Zur Erhöhung der Ladeinfrastrukturdichte sollen auch private Wallboxen für E-Mobilisten zugänglich gemacht werden. Daher soll ein Programm entwickelt werden, welches über mobile Endgeräte ermöglicht, private Ladestationen zu finden und ein Nutzungsentgelt für die Energieentnahme zwischen den Nutzern auszutauschen.

In addition to commercial charging stations, private charging options are becoming increasingly popular. Commercial charging stations can be found via various apps; however, this does not apply to private charging. To increase the charging infrastructure density, private chargers should also be made accessible to EV drivers. Therefore, an app is to be developed that enables users to find private charging stations via mobile devices and to exchange a usage fee for the energy withdrawal between the users.


## Run the Frontend

First you have to install all required packages via `npm install`.

To start the app during the **development** process with hot-reload `npm run dev` is used to deliver the app on localhost and `npm run dev-host` to deliver it network-wide.

To use the app in a **production** environment, first run `npm run build` to build the app and then run `npm run preview` to start the app. Run `npm run preview-host` to deliver it network-wide.

To specify the port of the app, you can add `-- --port <your-port>` to the command when you start the app .


## Run the Backend

PostgreSQL is required to use the PowerShare backend. You can [download it here](https://www.enterprisedb.com/downloads/postgres-postgresql-downloads). 
Create an empty database and enter the appropriate data into the db.env file. 

Before using the server, the certificate and key must be stored in the path specified in server.env, since the server only allows HTTPS.

The backend code can be compiled using the `go build cmd/main.go` command or compiled and executed directly using `go run cmd/main.go`.

### Configuration
Configurations like the port number or the paths of the SSL certificates can be specified in the config/server.env file. 
For the information needed for the database connection, use the config/db.env file.
If there are no custom files yet, the server will create default files at the first startup.


## Easy start

Batch files have been created for Windows users to make it easier to start the software.

### z_start_dev-host.bat
This file can be used to start the frontend network-wide in hot reload mode and open it in MS Edge.

### z_start_production.bat
This is used to build the productive version of the frontend, compile the Go Server and start it. 
The page is then opened in MS Edge Canary.