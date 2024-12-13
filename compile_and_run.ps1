
echo "Building..."
go build cmd\main\main.go

echo "Running..."
echo "Ctrl+C to stop"
.\main.exe
echo "Stopped"

rm "main.exe" -force
echo "main.exe removed"
