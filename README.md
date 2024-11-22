# compose-debug

## Dashboard
### Backend

Install dependencies:
```bash
npm install
```

Update tsconfig.json according to
```
{
  "compilerOptions": {
    "target": "es2016",
    "module": "commonjs",
    "outDir": "./dist",
    "esModuleInterop": true,
    "forceConsistentCasingInFileNames": true,
    "strict": true ,
    "skipLibCheck": true,
    "sourceMap": true,
    "rootDir": "./src" 
  }
}
```
### Create VScode config
```
mkdir -p .vsconfig
```

```
{
    "version": "0.2.0",
    "configurations": [
      {
        "type": "node",
        "request": "launch",
        "name": "Debug Dashboard Backend",
        "runtimeArgs": ["-r", "ts-node/register"],
        "args": ["${workspaceFolder}/src/server.ts"],
        "skipFiles": ["<node_internals>/**"],
        "sourceMaps": true
      }
    ]
}
```

### Frontend

Install dependencies:
```bash
npm install
```

Build release bundle
```bash
npm run build
```

## Docker-Debug
- Update environment variables in moby/.env to match the skyramp and dashboard location
- Run setup-debug-worker.sh and "Build Devcontainer"
- Run setup-debug-worker.sh and "Docker Compose Up"



