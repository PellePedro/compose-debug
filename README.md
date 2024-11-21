# compose-debug


## Dashboard Backend
Generate static files of the react dashboard application
```
npm run build
```

## Dashboard Backend
In order to debug and set breakpoint in the typescript code update tsconfig.json according to:

```
{
  "compilerOptions": {
    "target": "es2016",
    /* Modules */
    "module": "commonjs",
    "outDir": "./dist",
    "esModuleInterop": true,
    "forceConsistentCasingInFileNames": true ,
    /* Type Checking */
    "strict": true ,
    "skipLibCheck": true,
    "sourceMap": true, // Generates source maps
    "rootDir": "./src" 
  }
}
```

### Install dev dependency
```
npm install --save-dev typescript ts-node
```
