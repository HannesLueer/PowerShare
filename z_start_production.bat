cd frontend && call npm install && call npm run build & cd ..

go build -o bin/powershare.exe cmd/main.go

cd bin && start powershare.exe & cd ..

start msedge https://localhost:5000/

pause
