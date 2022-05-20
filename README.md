# Stumble

## Create a Dating App

Imagine there is a dating app “stumble” where there are many
users. There is also a feature for a user to like as many users as he/
she wants. You are given the list of these users in the JSON file.
Each user will have id (int mandatory), name (string mandatory),
location (float mandatory), gender (string optional), email (string
optional).

Also you are given a list of likes. (eg. A likes B- C likes A- B likes A- D
likes C,( Create a database model in any RDBM` for the dating app
using Golang’s ORM of your choice.

ORM used - <b>GORM</b>


Tasks - 

- After creating the database model- import the user’s JSON file
and “likes” data in the DB.
- Create an endpoint to retrieve all the matches.(Match is when X
likes Y & Y likes X)
- Given user X and distance k- create an endpoint to retrieve all
the users within distance k from user X. (Assume you are given the
distance of each user in the 1D coordinate system)
- Given input query q (string)- create an endpoint to retrieve all the
users which have q in their name


## [API Documentation](https://documenter.getpostman.com/view/20830684/Uyxogiaa)


## Deployment Done On [Heroku](https://stumble-date.herokuapp.com/)
