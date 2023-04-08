In this example, Repository has a dependency on DB, which is used to retrieve data. The constructor for Repository takes a pointer to a Database instance, which is used to initialize the db field of the struct.  
  
By using a mock implementation of DB, we can isolate the behavior of Repository and test it without relying on the actual implementation of DB.  
  
we create a mock implementation of DB called MockDB. This mock implementation has a GetFunc field, which is used to simulate the behavior of Get in our tests.