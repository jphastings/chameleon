# Chameleon
A simple web container for testing routing.

If you're messing about with docker containers and routing it can be handy to spin up a distictive container with a webserver in it.

1. Start a chameleon server

    ```bash
    $ docker run -e "SERVICE_NAME=hello" jphastings/chameleon
    2015/04/05 22:38:49 Starting webserver: hello
    ```
    
2. View the webserver and see the source:

    ```bash
    $ curl 127.0.0.1:49123/world
    hello: /world
    ```
