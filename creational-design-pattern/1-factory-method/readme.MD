A Factory Pattern or Factory Method Pattern says that just 
## Define an interface or abstract class for creating an object but let the subclasses decide which class to instantiate
OR
## The client only interacts with a factory struct and tells the kind of instances that needs to be created. The factory class interacts with the corresponding concrete structs and returns the correct instance back.

When ?
1. When a class doesn't know what sub-classes will be required to create

==> ``The main func acts as a client and instead of directly interacting with domestic or commericial, it relies on getPlanFactory to create instances of domestic and commericial``