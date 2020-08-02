### On MacOS Catalina don't download the clone zip file from Chrome. Binary will not execute, have to create it using make.

### 4 REST API created for Load,Add,List and emaillisting as a prototype. 


### Build

#### Local Build

```bash
make acme
```

#### Linux Build

```bash
make linux
```

### Run

#### Local

After making the local build, execute the following

```bash
./acme_osx --config.file=$(pwd)/acme.yaml
```

#### Linux

After making the linux build, execute the following

```bash
./acme_linux --config.file=$(pwd)/acme.yaml
```

The server would be listening at port `:10000`

### Commands

#### List Products

```bash
curl -X GET 'http://localhost:10000/product'
```

#### Load Product

```bash
curl -X POST --header "Content-Type: application/json" \
  --data '{"customer_id":"cust13","product_name":"hosting","domain":"abcd.com","start_date":"2021-1-1","duration_months":12}' \
  http://localhost:10000/product/
```
#### Add Product

```bash
curl -X PUT --header "Content-Type: application/json" \
  --data '{"customer_id":"cust13","product_name":"hosting","domain":"abcd.com","duration_months":12}' \
  http://localhost:10000/product/add
```

#### List email dates

```bash
curl -X GET 'http://localhost:10000/product/email'
```
