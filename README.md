## What is GG-CMS?
GG-CMS (GoGo-Content Management System) is a content management system in short, allows you to create your own blog and manage your own content.

Right now is early stage, but the plan is to add more features in futures versions

Here is a list of the features in the current version **V0.1**:
* Manage Posts (Create, Delete, Update and Get)
* Manage Users with admin role
* JWT (Json Web Token) Implementation for security

**Note:**
**This is only the Back-End of the complete project, please click in the link below to go to the Front-End repository and demo page link:**

* [gg-cms-front-end](https://github.com/elisardofelix/gg-cms-front-end)
* [Demo Page](https://gg-cms-demo.devland.bid/)


## Implementation GG-CMS
For implementing this back-end you need to have the following considerations:
* All the configuration is via Environment variables which mean you need to setup first:
    * `PORT` : Is optional because the mayor part of the server implementation use that variable to set the port automatically
    * `GGCMSDBString` : Is to set the server connection string to MongoDB, here are sine examples:
        * Set with only server ip or name `mongodb://<IP OR SERVERNAME>` or with credentials to access the MongoDB server `mongodb://[username]:[password]@<IP OR SERVERNAME>`
    * `GGCMSDBProd` : Here is the database that is use in the enviroment
    * `GGCMSDBTest` : Here is the database that is use for the integrations testing **Note: the database is created at the start and destroyed at the end of all tests**
    * `ESPECIAL_USER` : Is a user that going to work only if you not have any user in the database. Example of how to setup:
        * `[username]:[password]`
    * `JWT_SECRET` : Was created to change the default secret key used to encode the JWT.

### How to set Environment Variables in Windows
In Windows, you only need to set each variable enter the following command in the command prompt before to run the application `set [VariableName]=[Value]` if you want to know more about it below is a link with an article explaining it:

[Article Link](https://ss64.com/nt/set.html)

### How to set Environment Variables in Windows
In Linux, you only need to set each variable enter the following command in the terminal before to run the application `export [VariableName]=[Value]` if you want to know more about it below is a link with an article explaining it:

[Article Link](https://www.cyberciti.biz/faq/set-environment-variable-linux/)

### How to run tests and app
Jus apply the following commands line by line:

```
[Set all the environment variables]
#This is a comment ;)
#Then install all the package required with the following command
go get -d -v
#Then run all the project tests
go test ./...
#Then run the application
go run .
```

## Contributions
If you have any improve proposal to the code, please no doubt of sending me your pull request and I check to apply in the repository as soon as possible.

Please feel free to contact  me 😊: [elisardofelix@gmail.com](mailto:elisardofelix@gmail.com)

Peace ❤

