// Package dispatcher code generated by oapi sdk gen
/*
 * MIT License
 *
 * Copyright (c) 2022 Lark Technologies Pte. Ltd.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice, shall be included in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package dispatcher

import (
	"context"
	"github.com/larksuite/oapi-sdk-go/v3/service/moments/v1"
)
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2CommentCreatedV1(handler func(ctx context.Context, event *larkmoments.P2CommentCreatedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["moments.comment.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "moments.comment.created_v1")
	}
	dispatcher.eventType2EventHandler["moments.comment.created_v1"] = larkmoments.NewP2CommentCreatedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2CommentDeletedV1(handler func(ctx context.Context, event *larkmoments.P2CommentDeletedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["moments.comment.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "moments.comment.deleted_v1")
	}
	dispatcher.eventType2EventHandler["moments.comment.deleted_v1"] = larkmoments.NewP2CommentDeletedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2DislikeCreatedV1(handler func(ctx context.Context, event *larkmoments.P2DislikeCreatedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["moments.dislike.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "moments.dislike.created_v1")
	}
	dispatcher.eventType2EventHandler["moments.dislike.created_v1"] = larkmoments.NewP2DislikeCreatedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2DislikeDeletedV1(handler func(ctx context.Context, event *larkmoments.P2DislikeDeletedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["moments.dislike.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "moments.dislike.deleted_v1")
	}
	dispatcher.eventType2EventHandler["moments.dislike.deleted_v1"] = larkmoments.NewP2DislikeDeletedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2PostCreatedV1(handler func(ctx context.Context, event *larkmoments.P2PostCreatedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["moments.post.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "moments.post.created_v1")
	}
	dispatcher.eventType2EventHandler["moments.post.created_v1"] = larkmoments.NewP2PostCreatedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2PostDeletedV1(handler func(ctx context.Context, event *larkmoments.P2PostDeletedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["moments.post.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "moments.post.deleted_v1")
	}
	dispatcher.eventType2EventHandler["moments.post.deleted_v1"] = larkmoments.NewP2PostDeletedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2PostStatisticsUpdatedV1(handler func(ctx context.Context, event *larkmoments.P2PostStatisticsUpdatedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["moments.post_statistics.updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "moments.post_statistics.updated_v1")
	}
	dispatcher.eventType2EventHandler["moments.post_statistics.updated_v1"] = larkmoments.NewP2PostStatisticsUpdatedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2ReactionCreatedV1(handler func(ctx context.Context, event *larkmoments.P2ReactionCreatedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["moments.reaction.created_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "moments.reaction.created_v1")
	}
	dispatcher.eventType2EventHandler["moments.reaction.created_v1"] = larkmoments.NewP2ReactionCreatedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2ReactionDeletedV1(handler func(ctx context.Context, event *larkmoments.P2ReactionDeletedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["moments.reaction.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "moments.reaction.deleted_v1")
	}
	dispatcher.eventType2EventHandler["moments.reaction.deleted_v1"] = larkmoments.NewP2ReactionDeletedV1Handler(handler)
	return dispatcher
}