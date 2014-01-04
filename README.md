histri
======
**histri** is a Go ([golang](http://golang.org)) library for capturing events and user activity, and efficiently storing them in [PostgreSql](http://www.postgresql.org/). 

PostgreSql is used as the data store because it is mature, has a large existing ecosystem of management and reporting tools and profressionals that know how to use it, and it has excellent support for storing and querying JSON data (which is handy for extended event metadata).

Golang is used because it is an excellent choice for writing fast, robust network servers and it is easy to write clean, testable code with.

The goal is for histri to eventually consist of:
- The above mentioned library
- A server daemon with a built-in HTTP service that event data can be POSTed to (the Capture service).
- A command-line utility for managing and querying the event data (the Management CLI).
- (Longer-term) Another HTTP service for generating reports, charts and data feeds (the Analytics service).

The Capture service is the first focus of attention for the project.

The goals for the Capture service are:
- It should be able to efficiently capture many thousands of user activity events per second with a single daemon.
- It should be lightweight enough that it is an option to run it on the same server as the PostgreSql database process.
- It should optionally accept a timestamp (in UTC) for incoming events so that it can be used to backfill previously recorded events or accept data from a log replay.
- It should accept extended event metadata without the need to preplan the schema for that extra data.

Please submit bugs to the [Issues](https://github.com/nathanwdavis/histri/issues) page for the project. I welcome contributed bug fixes. If you find a bug and want to contribute a fix, please first open an Issue, then proceed with a pull request.
