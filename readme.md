## Project Structure
![image](https://user-images.githubusercontent.com/73453940/213455823-82ee07b9-516f-49bd-ae66-92ee65bd3e1a.png)
- This project follow Clean Architecture. There are 3 main layer: Entities Layer, Usecase Layer and Presenter Layer
- All layer will communicate to each other using interfaces -> Layer inside will not be coupled with ouside layer, easy to maintain and scale
- Outside layers can communicate to inside layers and same level layers by implementing interface -> Resuable
## Project Structure
```
├─internal:
│   └─offer: Domain's name
│       └─io: Iomodel
├───external:
│   ├───common: Common Structs, Varibles stuffs
├─cmd: Main package, where to run project
├─delivery: Presenter Layer, communicate with Client-server or Front-end
│   └─offer: Domain's name
│       └─io: Input-Output Model
├───external: Shared package, define logic which use by other's package
│   ├───common: Structs, varibles stuffs
│   └───request: RestAPI stuffs
├───internal: Bussiness package, where to define bussiness logic and communicate to database
│   └───offers: Big domain's name
│       └───merchant: Domain's name
│           ├───storage: Entities layer, communicate with database
│           └───usecase: Bussiness layer, use storage layer
└───providers: Interact with providers
    ├───ascenda: Provider's name
    └───base_api: Provider APi convert
 ```
## How to run
- Install Golang: [Golang Dev](https://go.dev/doc/install)
- Run these command
```bash
cd cmd
```

```bash
go run .
```

- API GET 
```
http://localhost:8080/api/v1/offers/near_by?checkin=2019-12-25&lat=324324&lon=412&rad=20
```


  
