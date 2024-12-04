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
	"github.com/larksuite/oapi-sdk-go/v3/service/drive/v1"
)
// 多维表格字段变更
//
// - 多维表格字段变更
//
// - 事件描述文档链接:https://open.feishu.cn/document/uAjLw4CM/ukTMukTMukTM/reference/drive-v1/file/events/bitable_field_changed
func ( dispatcher * EventDispatcher ) OnP2FileBitableFieldChangedV1(handler func(ctx context.Context, event *larkdrive.P2FileBitableFieldChangedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.bitable_field_changed_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.bitable_field_changed_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.bitable_field_changed_v1"] = larkdrive.NewP2FileBitableFieldChangedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2FileBitableRecordChangedV1(handler func(ctx context.Context, event *larkdrive.P2FileBitableRecordChangedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.bitable_record_changed_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.bitable_record_changed_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.bitable_record_changed_v1"] = larkdrive.NewP2FileBitableRecordChangedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2FileCreatedInFolderV1(handler func(ctx context.Context, event *larkdrive.P2FileCreatedInFolderV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.created_in_folder_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.created_in_folder_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.created_in_folder_v1"] = larkdrive.NewP2FileCreatedInFolderV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/file-deleted-completely
func ( dispatcher * EventDispatcher ) OnP2FileDeletedV1(handler func(ctx context.Context, event *larkdrive.P2FileDeletedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.deleted_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.deleted_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.deleted_v1"] = larkdrive.NewP2FileDeletedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/file-edited
func ( dispatcher * EventDispatcher ) OnP2FileEditV1(handler func(ctx context.Context, event *larkdrive.P2FileEditV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.edit_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.edit_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.edit_v1"] = larkdrive.NewP2FileEditV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/file-collaborator-add
func ( dispatcher * EventDispatcher ) OnP2FilePermissionMemberAddedV1(handler func(ctx context.Context, event *larkdrive.P2FilePermissionMemberAddedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.permission_member_added_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.permission_member_added_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.permission_member_added_v1"] = larkdrive.NewP2FilePermissionMemberAddedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:
func ( dispatcher * EventDispatcher ) OnP2FilePermissionMemberAppliedV1(handler func(ctx context.Context, event *larkdrive.P2FilePermissionMemberAppliedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.permission_member_applied_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.permission_member_applied_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.permission_member_applied_v1"] = larkdrive.NewP2FilePermissionMemberAppliedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/file-collaborator-remove
func ( dispatcher * EventDispatcher ) OnP2FilePermissionMemberRemovedV1(handler func(ctx context.Context, event *larkdrive.P2FilePermissionMemberRemovedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.permission_member_removed_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.permission_member_removed_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.permission_member_removed_v1"] = larkdrive.NewP2FilePermissionMemberRemovedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/file-read
func ( dispatcher * EventDispatcher ) OnP2FileReadV1(handler func(ctx context.Context, event *larkdrive.P2FileReadV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.read_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.read_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.read_v1"] = larkdrive.NewP2FileReadV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/file-title-update
func ( dispatcher * EventDispatcher ) OnP2FileTitleUpdatedV1(handler func(ctx context.Context, event *larkdrive.P2FileTitleUpdatedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.title_updated_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.title_updated_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.title_updated_v1"] = larkdrive.NewP2FileTitleUpdatedV1Handler(handler)
	return dispatcher
}
// 
//
// - 
//
// - 事件描述文档链接:https://open.feishu.cn/document/ukTMukTMukTM/uUDN04SN0QjL1QDN/event/delete-file-to-trash-can
func ( dispatcher * EventDispatcher ) OnP2FileTrashedV1(handler func(ctx context.Context, event *larkdrive.P2FileTrashedV1) error) * EventDispatcher{
	_, existed := dispatcher.eventType2EventHandler["drive.file.trashed_v1"]
	if existed {
		panic("event: multiple handler registrations for " + "drive.file.trashed_v1")
	}
	dispatcher.eventType2EventHandler["drive.file.trashed_v1"] = larkdrive.NewP2FileTrashedV1Handler(handler)
	return dispatcher
}