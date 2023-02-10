# Favour's Busha Assessment

This is my attempt at the assessment for the backend developer position at Busha. This application uses Redis, Postgres, and was hosted in Heroku. It also utilised gin for routing.

## Getting Started
The base url for this application is https://neo-swapi.herokuapp.com/. The documentation for this service can be found at: https://documenter.getpostman.com/view/17909079/2s935soh6Z

## Dockerization
I was able to dockerize the environment: docker.io/library/busha-assessment:latest. The Dockerfile should also be available in the project

To clarify:

### Tasks
- List an array of movies containing the name, opening crawl and comment count -- Done
- Add a new comment for a movie -- Done
- List the comments for a movie -- Done
- Get list of characters for a movie -- Done

### General requirements

- The application should have basic documentation that lists available endpoints and methods along with their request and response signatures -- Done:  https://documenter.getpostman.com/view/17909079/2s935soh6Z
- Keep your application source code on a public Github repository -- Done
- Deploy the API endpoints and provide a live demo URL of the API documentation. Heroku is a good option. -- Done:  https://neo-swapi.herokuapp.com/
- Bonus, but not mandatory, if you can dockerize the development environment. -- Done, Dockerfile in project

### Data requirements

- The movie data should be fetched online from `**[https://swapi.dev](https://swapi.dev)**` -- Done
- Movie names in the movie list endpoint should be sorted by release date from earliest to newest and each movie should be listed along with opening crawls and count of comments. -- Done
- Data fetched from `**[https://swapi.dev](https://swapi.dev)`** should be cached with Redis and then accessed from the cache for subsequent requests. -- Done
- Comments should be stored in a Postgres database. -- Done
- Error responses should be returned in case of errors. -- Done

### Character list requirements

- Endpoint should accept sort parameters to sort by one of name, gender or height in ascending or descending order. -- Done
- Endpoint should also accept a filter parameter to filter by gender. -- Done
- The response should also return metadata that contains the total number of characters that match the criteria along with the total height of the characters that match the criteria -- Done
- The total height should be provided both in cm and in feet/inches. For instance, 170cm makes 5ft and 6.93 inches. -- Done

### Comment requirements

- Comment list should be retrieved in reverse chronological order -- Done
- Comments should be retrieved along with the public IP address of the commenter and UTC date&time they were stored -- Done
- Comment length should be limited to 500 characters -- Done



