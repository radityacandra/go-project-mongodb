db.createCollection("users");
db.createUser(
  {
    user: "applicationUser",
    pwd: "password",
    roles: [
       { role: "readWrite", db: "database" }
    ]
  }
)
