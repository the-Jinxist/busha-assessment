package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	database "github.com/the-Jinxist/busha-assessment/database/sqlc"
)

type ListCommentsRequest struct {
	Limit   int64 `form:"limit" binding:"number"`
	Offset  int64 `form:"offset" binding:"number"`
	MovieID int64 `form:"movie_id" binding:"required,max=6"`
}

func (s *Server) getComments(ctx *gin.Context) {
	var request ListCommentsRequest
	err := ctx.ShouldBindQuery(&request)

	if err != nil {
		log.Printf("error while binding query: %s", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	var finalLimit = request.Limit
	if request.Limit == 0 {
		finalLimit = 10
	}

	arg := database.ListCommentsParams{
		MovieID: strconv.Itoa(int(request.MovieID)),
		Limit:   int32(finalLimit),
		Offset:  int32(request.Offset),
	}

	comments, err := s.store.ListComments(ctx, arg)
	if err != nil {
		log.Printf("error while listing comments: %s", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   comments,
	})
}

type AddCommentRequest struct {
	// MovieID          string `json:"movie_id"`
	// Comment          string `json:"comment"`
	MovieID int64  `json:"movie_id" binding:"required,max=6"`
	Comment string `json:"comment" binding:"required,max=500"`
}

func (s *Server) postComment(ctx *gin.Context) {

	//Enhancement: Check if a movie that corresponds to your id is available

	var request AddCommentRequest
	err := ctx.ShouldBindJSON(&request)

	if err != nil {
		log.Printf("error while binding json: %s", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	arg := database.CreateCommentParams{
		MovieID:          strconv.Itoa(int(request.MovieID)),
		Comment:          request.Comment,
		CommentIpAddress: ctx.ClientIP(),
	}

	comments, err := s.store.CreateComment(ctx, arg)
	if err != nil {
		log.Printf("error while getting comments: %s", err)
		ctx.JSON(http.StatusBadRequest, errorResponse(err, http.StatusBadRequest))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   comments,
	})
}
