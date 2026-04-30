# expense-tracker-cli

projectId : https://roadmap.sh/projects/expense-tracker
```
expense-tracker-cli
├─ README.md
├─ cmd
│  └─ expense-tracker
│     └─ main.go
├─ go.mod
├─ go.sum
├─ interfaces
│  └─ cli
│     ├─ handler.go
│     ├─ output.go
│     └─ parser.go
└─ internal
   ├─ application
   │  └─ expense-tracker
   │     ├─ dto
   │     │  ├─ input.go
   │     │  └─ output.go
   │     └─ service.go
   ├─ config
   ├─ domain
   │  └─ expense-tracker
   │     ├─ entity.go
   │     ├─ errrors.go
   │     └─ repository.go
   └─ infrastructure
      ├─ expense-tracker
      │  └─ service_implementation.go
      └─ persistence
         └─ json
            ├─ expense_repository.go
            └─ storage.json

```