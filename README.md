# PowerShare

## About the App

Neben kommerziellen Ladestationen finden private Lademöglichkeiten (Wallboxen) zunehmend Verbreitung. Kommerzielle Ladestationen können über diverse Apps gefunden werden; dies gilt jedoch nicht für private Wallboxen. Zur Erhöhung der Ladeinfrastrukturdichte sollen auch private Wallboxen für E-Mobilisten zugänglich gemacht werden. Daher soll ein Programm entwickelt werden, welches über mobile Endgeräte ermöglicht, private Ladestationen zu finden und ein Nutzungsentgelt für die Energieentnahme zwischen den Nutzern auszutauschen.

In addition to commercial charging stations, private charging options are becoming increasingly popular. Commercial charging stations can be found via various apps; however, this does not apply to private charging. To increase the charging infrastructure density, private chargers should also be made accessible to EV drivers. Therefore, an app is to be developed that enables users to find private charging stations via mobile devices and to exchange a usage fee for the energy withdrawal between the users.


## Run the App

To start the app during the **development** process with hot-reload `npm run dev` is used to deliver the app on localhost and `npm run dev-host` to deliver it network-wide.

To use the app in a **production** environment, first run `npm run build` to build the app and then run `npm run preview` to start the app. Run `npm run preview-host` to deliver it network-wide.

To specify the port of the app, you can add `-- --port <your-port>` to the command when you start the app .
