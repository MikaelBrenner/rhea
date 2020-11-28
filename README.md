# rhea

**rhea** aims to be a blazing fast reverse-proxy with a focus on caching to take some of the heavy lifting off your server.


# Main goal
To have a small set of features that work amazingly well.

 There are plenty of reverse proxies available, and this project does not aim at replacing their full functionality.

The focus here is to be a very thin layer of functionality that increases the performance of your application(if your responses are cacheable in nature) without the user having to go through pages and pages of documentation. 

# Goals
* To have simple, human readable configuration files
* To require no change in the server to achieve maximal performance
* To allow for routes to be fully configurable, including number of cached requests, timeout of refreshing the cache and subroute configuration, both for inclusion or removal



