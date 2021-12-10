Object Pool Pattern says that
### to reuse the object that are expensive to create".

Basically, an Object pool is a container which contains a specified amount of objects. When an object is taken from the pool, it is not available in the pool until it is put back. Objects in the pool have a lifecycle: creation, validation and destroy.

As and when needed, a client can request an object from the pool, use it, and return it to the pool. The object in the pool is never destroyed.
