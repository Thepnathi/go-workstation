# Handler Interface

* Another way to register code to handle an HTTP request is to use the Handler interface
* Any type that implements the interface http.Handler can be be passed to the http.Handle function, along with 
pattern string