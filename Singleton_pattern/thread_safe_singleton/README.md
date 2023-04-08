To create a thread-safe variant of the Singleton pattern in Golang, we can use the sync.Once package. We modify the instance variable to be of type sync.Once, and then use its Do() method to create a new instance of the Singleton struct only once.