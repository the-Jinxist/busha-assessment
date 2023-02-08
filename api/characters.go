package api

import "github.com/gin-gonic/gin"

func (s *Server) getCharacters(ctx *gin.Context) {

	//check if redis already have something stored for this endpoint and param combination,
	//if, it does return the redis version

	//if not, make the call to the swapi service

	//after making the call to the swapi service
	//get the list

	//parse the sort/filter elements if any

	//Switch between the logic to filter

	//filter the list

	//save to redis using a combination of the endpoint and the params

	//send to user
}
