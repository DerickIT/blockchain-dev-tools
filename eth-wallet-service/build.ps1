# build.ps1

param(
    [string]$target = "help"
)

$projectName = "eth-wallet-service"
$mainFile = "cmd\main.go"
$binary = "$projectName.exe"

function Init {
    # 初始化 go.mod
    if (-not (Test-Path go.mod)) {
        go mod init github.com/derickit/eth-wallet-service
    }
    go mod tidy
    # 生成 Swagger 文档
    Swag
}

function Build {
    Init
    go build -o $binary $mainFile
}

function Run {
    Build
    if (Test-Path $binary) {
        & ".\$binary"
    } else {
        Write-Host "Build failed. Unable to run."
    }
}

function Clean {
    Remove-Item -Force $binary -ErrorAction SilentlyContinue
    Remove-Item -Recurse -Force .\docs -ErrorAction SilentlyContinue
    Remove-Item -Force go.mod -ErrorAction SilentlyContinue
    Remove-Item -Force go.sum -ErrorAction SilentlyContinue
}

function Test {
    Init
    go test .\...
}

function Swag {
    # 确保 swag 命令存在
    if (Get-Command swag -ErrorAction SilentlyContinue) {
        swag init -g $mainFile
    } else {
        Write-Host "Swag is not installed. Please install it with: go install github.com/swaggo/swag/cmd/swag@latest"
        exit 1
    }
}

function Deps {
    Init
    go mod download
}

function UpdateDeps {
    Init
    go get -u .\...
}

function Help {
    Write-Host "Available targets:"
    Write-Host "  init        - Initialize the project (generate docs and setup go.mod)"
    Write-Host "  build       - Compile the project"
    Write-Host "  run         - Compile and run the project"
    Write-Host "  clean       - Remove compiled files and go.mod"
    Write-Host "  test        - Run tests"
    Write-Host "  swag        - Generate Swagger documentation"
    Write-Host "  deps        - Download dependencies"
    Write-Host "  update-deps - Update dependencies"
}

switch ($target) {
    "init" { Init }
    "build" { Build }
    "run" { Run }
    "clean" { Clean }
    "test" { Test }
    "swag" { Swag }
    "deps" { Deps }
    "update-deps" { UpdateDeps }
    default { Help }
}
