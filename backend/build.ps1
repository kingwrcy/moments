param (
    [string]$envType = "prod",
    [string]$arch = "amd64"
)

if ($envType -ne "dev" -and $envType -ne "prod") {
    Write-Host "Usage: .\build.ps1 [dev|prod] [amd64|386]"
    exit 1
}

if ($arch -ne "amd64" -and $arch -ne "386") {
    Write-Host "Usage: .\build.ps1 [dev|prod] [amd64|386]"
    exit 1
}

$tags = ""
if ($envType -eq "prod") {
    $tags = "prod"
}

$env:GOOS = "windows"
$env:GOARCH = $arch

Write-Host "Building the Go application for $env:GOOS/$env:GOARCH..."
go build -tags $tags -ldflags="-s -w" -o moments.exe

if (Get-Command upx -ErrorAction SilentlyContinue) {
    Write-Host "Compressing the binary with upx..."
    upx --best --lzma moments.exe
} else {
    Write-Host "upx is not installed. Skipping compression."
}

Write-Host "Build complete."
