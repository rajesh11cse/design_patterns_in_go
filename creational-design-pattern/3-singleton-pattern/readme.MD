Singleton Pattern says that just
### define a class that has only one instance and provides a global point of access to it In other words, a class must ensure that only single instance should be created and single object can be used by all other classes.

There are two forms of singleton design pattern

1. Early Instantiation: creation of instance at load time.
2. Lazy Instantiation: creation of instance when required.

### Examples
1. DB instance – we only want to create only one instance of DB object and that instance will be used throughout the application. 
2. Logger instance – again only one instance of the logger should be created and it should be used throughout the application.
