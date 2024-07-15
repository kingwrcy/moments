package handler

import "github.com/samber/do/v2"

type CommentHandler struct {
	base BaseHandler
}

func NewCommentHandler(injector do.Injector) *CommentHandler {
	return &CommentHandler{do.MustInvoke[BaseHandler](injector)}
}
