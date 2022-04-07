# GOHARE - Flight data tracking

Go'Hare is a collection of flight tracking tools.
The name is derived from the main implementation language Go and the Chicago O’Hare International Airport.

As a long term goal, this should become a comprehensive tool supporting a large variety of data sources, like

* ADSB
* ASTERIX
* OPENSKY
* FLIGHTRADAR 24

and possibilities to view the current state, like a map or tables.

The idea is, to be able to run different connectors as separate services, being able to write tracking data to a separate storage instance. 
Different Tools or Viewers (e.g. some webapp, or Xplane, or MSFS 2020) could connect to display current data.

At this stage, this project is purely a hobby, one of the famous side projects of a professional Software Engineer :).
It is roughly based on my professional experience in this field and was supposed to be used for learning various concepts of Go.
This whole project is more like a PoC, e.g. showing only limited data.

### How it works (more or less)
* opensky connector, reads data from opensky and writes them into a trackstore
* separate trackstore and gRPC server
* separate webapp, fetches data cia gRPC client from the trackstore gRPC server

## Planned changes
* separate trackstore from trackers
* webApp single track view should get streaming data updates
* maybe a better decoupling, to not pass gRPC client to the web but abstract it away
* gRPC streaming for getting updates of a flight (new positions)
* trackers (e.g. openskyReader would be a tracker) can write new updates to the store
* maybe use REDIS or something different as trackstore solution 
* generally make it more GO-like and less Java-like, e.g. i am not sure if its a good idea to have a separate track package, if its actually part of the trackstore. To have it modular, it should be somehow exported, so that eventually all separate readers can write their own tracks
    * But the problem would be how to deal with duplicates, or several readers having the same flights. Those will need to be merged as part of the store. Maybe this is out of scope of this PoC tool :) 
    