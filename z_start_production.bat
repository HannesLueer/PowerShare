cd frontend && call npm install && call npm run build & cd ..

go build -o bin/powershare.exe cmd/main.go

cd bin && start powershare.exe & cd ..

SET edgepath=%appdata%\..\Local\Microsoft\Edge SXS\Application
C:
cd "%edgepath%"
start msedge.exe --ignore-certificate-errors --unsafely-treat-insecure-origin-as-secure=https://localhost:5000/ --allow-insecure-localhost https://localhost:5000/

pause
