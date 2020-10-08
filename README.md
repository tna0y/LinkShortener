# LinkShortener
Basic link shortener with Telegram bot interface

# Design draft

## Models

### Link
- ID (link alias)
- Owner ID
- Destination
- Created
- Expiration

### User
Could be more useful later if web UI is implemented
- ID
- Telegram user ID

## Services

### Links
CRUD-like service around the Link model
Endpoints:
- Create Link
- Delete Link
- Get Link by it's ID
- List Links for a given user

### Users
CRUD-like service around the User model
- Create User
- Get User by telegram ID

### Redirector
Http service that works as a paroxy, redirecting requests to designated link locations.
Service has a single endpoint used for redirection.

### Bot
Control plane of the link shortener.
Authentication is based upon telegram API, users are only allowed to manage links they have created.
Commands:
- Create link
- Delete link
- List links


## DB Rationale

Reads are expected to produce much more load than writes. 
Link shortener should work swiftly across the globe thus read-only replicas in different regions should be an option. 

Models are well structured thus we don't have a need for NoSQL.

PostgreSQL is a reasonable choice since we can handle read load via read-only replicas and in case write load gets too high clientside sharding should also do the trick.

Cassandra could also be an interesting choice in terms of clientside load balancing and cross-dc replication. Automatic link expiration could also be achieved with this DB.

For the sake of implementation simplicity we will go with PostgreSQL.

