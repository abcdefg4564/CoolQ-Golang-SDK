// Code generated by cqgen, DO NOT EDIT.
// +build !websocket

package cqp

/*
#include <windows.h>
#include <stdint.h>
#include <stdlib.h>

// #define __stdcall __attribute__((__stdcall__))
#define cq_bool_t int32_t

#define CQAPI(RetType, Name, ...)                                         \
      typedef RetType(__stdcall *Name##_Type)(int32_t ac, ##__VA_ARGS__); \
      Name##_Type Name##_Ptr;                                             \
      RetType Name(__VA_ARGS__);

#define LoadAPI(Name) Name##_Ptr = (Name##_Type)GetProcAddress(hmod, #Name)

int32_t ac; //AccessCode
CQAPI(int32_t, CQ_sendPrivateMsg, int64_t, char *)
CQAPI(int32_t, CQ_sendGroupMsg, int64_t, char *)
CQAPI(int32_t, CQ_sendDiscussMsg, int64_t, char *)
CQAPI(int32_t, CQ_deleteMsg, int64_t)
CQAPI(int32_t, CQ_sendLikeV2, int64_t, int32_t)
CQAPI(int32_t, CQ_setGroupKick, int64_t, int64_t, cq_bool_t)
CQAPI(int32_t, CQ_setGroupBan, int64_t, int64_t, int64_t)
CQAPI(int32_t, CQ_setGroupAnonymousBan, int64_t, char *, int64_t)
CQAPI(int32_t, CQ_setGroupWholeBan, int64_t, cq_bool_t)
CQAPI(int32_t, CQ_setGroupAdmin, int64_t, int64_t, cq_bool_t)
CQAPI(int32_t, CQ_setGroupAnonymous, int64_t, cq_bool_t)
CQAPI(int32_t, CQ_setGroupCard, int64_t, int64_t, char *)
CQAPI(int32_t, CQ_setGroupLeave, int64_t, cq_bool_t)
CQAPI(int32_t, CQ_setGroupSpecialTitle, int64_t, int64_t, char *, int64_t)
CQAPI(int32_t, CQ_setDiscussLeave, int64_t)
CQAPI(int32_t, CQ_setFriendAddRequest, char *, int32_t, char *)
CQAPI(int32_t, CQ_setGroupAddRequestV2, char *, int32_t, int32_t, char *)
CQAPI(int64_t, CQ_getLoginQQ)
CQAPI(char *, CQ_getLoginNick)
CQAPI(char *, CQ_getStrangerInfo, int64_t, cq_bool_t)
CQAPI(char *, CQ_getFriendList, cq_bool_t)
CQAPI(char *, CQ_getGroupList)
CQAPI(char *, CQ_getGroupInfo, int64_t, cq_bool_t)
CQAPI(char *, CQ_getGroupMemberList, int64_t)
CQAPI(char *, CQ_getGroupMemberInfoV2, int64_t, int64_t, cq_bool_t)
CQAPI(char *, CQ_getCookiesV2, char *)
CQAPI(int32_t, CQ_getCsrfToken)
CQAPI(char *, CQ_getAppDirectory)
CQAPI(char *, CQ_getRecordV2, char *, char *)
CQAPI(char *, CQ_getImage, char *)
CQAPI(cq_bool_t, CQ_canSendImage)
CQAPI(cq_bool_t, CQ_canSendRecord)
CQAPI(int32_t, CQ_addLog, int32_t, char *, char *)
extern int32_t __stdcall Initialize(int32_t access_code)
{
    ac = access_code;
    HMODULE hmod = LoadLibrary("CQP.dll");
    LoadAPI(CQ_sendPrivateMsg);
    LoadAPI(CQ_sendGroupMsg);
    LoadAPI(CQ_sendDiscussMsg);
    LoadAPI(CQ_deleteMsg);
    LoadAPI(CQ_sendLikeV2);
    LoadAPI(CQ_setGroupKick);
    LoadAPI(CQ_setGroupBan);
    LoadAPI(CQ_setGroupAnonymousBan);
    LoadAPI(CQ_setGroupWholeBan);
    LoadAPI(CQ_setGroupAdmin);
    LoadAPI(CQ_setGroupAnonymous);
    LoadAPI(CQ_setGroupCard);
    LoadAPI(CQ_setGroupLeave);
    LoadAPI(CQ_setGroupSpecialTitle);
    LoadAPI(CQ_setDiscussLeave);
    LoadAPI(CQ_setFriendAddRequest);
    LoadAPI(CQ_setGroupAddRequestV2);
    LoadAPI(CQ_getLoginQQ);
    LoadAPI(CQ_getLoginNick);
    LoadAPI(CQ_getStrangerInfo);
    LoadAPI(CQ_getFriendList);
    LoadAPI(CQ_getGroupList);
    LoadAPI(CQ_getGroupInfo);
    LoadAPI(CQ_getGroupMemberList);
    LoadAPI(CQ_getGroupMemberInfoV2);
    LoadAPI(CQ_getCookiesV2);
    LoadAPI(CQ_getCsrfToken);
    LoadAPI(CQ_getAppDirectory);
    LoadAPI(CQ_getRecordV2);
    LoadAPI(CQ_getImage);
    LoadAPI(CQ_canSendImage);
    LoadAPI(CQ_canSendRecord);
    LoadAPI(CQ_addLog);
    return 0;
}
int32_t CQ_sendPrivateMsg(int64_t var0, char * var1)
{
    int32_t ret = CQ_sendPrivateMsg_Ptr(ac, var0, var1);
    free(var1);
    return ret;
}
int32_t CQ_sendGroupMsg(int64_t var0, char * var1)
{
    int32_t ret = CQ_sendGroupMsg_Ptr(ac, var0, var1);
    free(var1);
    return ret;
}
int32_t CQ_sendDiscussMsg(int64_t var0, char * var1)
{
    int32_t ret = CQ_sendDiscussMsg_Ptr(ac, var0, var1);
    free(var1);
    return ret;
}
int32_t CQ_deleteMsg(int64_t var0)
{
    int32_t ret = CQ_deleteMsg_Ptr(ac, var0);
    return ret;
}
int32_t CQ_sendLikeV2(int64_t var0, int32_t var1)
{
    int32_t ret = CQ_sendLikeV2_Ptr(ac, var0, var1);
    return ret;
}
int32_t CQ_setGroupKick(int64_t var0, int64_t var1, cq_bool_t var2)
{
    int32_t ret = CQ_setGroupKick_Ptr(ac, var0, var1, var2);
    return ret;
}
int32_t CQ_setGroupBan(int64_t var0, int64_t var1, int64_t var2)
{
    int32_t ret = CQ_setGroupBan_Ptr(ac, var0, var1, var2);
    return ret;
}
int32_t CQ_setGroupAnonymousBan(int64_t var0, char * var1, int64_t var2)
{
    int32_t ret = CQ_setGroupAnonymousBan_Ptr(ac, var0, var1, var2);
    free(var1);
    return ret;
}
int32_t CQ_setGroupWholeBan(int64_t var0, cq_bool_t var1)
{
    int32_t ret = CQ_setGroupWholeBan_Ptr(ac, var0, var1);
    return ret;
}
int32_t CQ_setGroupAdmin(int64_t var0, int64_t var1, cq_bool_t var2)
{
    int32_t ret = CQ_setGroupAdmin_Ptr(ac, var0, var1, var2);
    return ret;
}
int32_t CQ_setGroupAnonymous(int64_t var0, cq_bool_t var1)
{
    int32_t ret = CQ_setGroupAnonymous_Ptr(ac, var0, var1);
    return ret;
}
int32_t CQ_setGroupCard(int64_t var0, int64_t var1, char * var2)
{
    int32_t ret = CQ_setGroupCard_Ptr(ac, var0, var1, var2);
    free(var2);
    return ret;
}
int32_t CQ_setGroupLeave(int64_t var0, cq_bool_t var1)
{
    int32_t ret = CQ_setGroupLeave_Ptr(ac, var0, var1);
    return ret;
}
int32_t CQ_setGroupSpecialTitle(int64_t var0, int64_t var1, char * var2, int64_t var3)
{
    int32_t ret = CQ_setGroupSpecialTitle_Ptr(ac, var0, var1, var2, var3);
    free(var2);
    return ret;
}
int32_t CQ_setDiscussLeave(int64_t var0)
{
    int32_t ret = CQ_setDiscussLeave_Ptr(ac, var0);
    return ret;
}
int32_t CQ_setFriendAddRequest(char * var0, int32_t var1, char * var2)
{
    int32_t ret = CQ_setFriendAddRequest_Ptr(ac, var0, var1, var2);
    free(var0);
    free(var2);
    return ret;
}
int32_t CQ_setGroupAddRequestV2(char * var0, int32_t var1, int32_t var2, char * var3)
{
    int32_t ret = CQ_setGroupAddRequestV2_Ptr(ac, var0, var1, var2, var3);
    free(var0);
    free(var3);
    return ret;
}
int64_t CQ_getLoginQQ()
{
    int64_t ret = CQ_getLoginQQ_Ptr(ac);
    return ret;
}
char * CQ_getLoginNick()
{
    char * ret = CQ_getLoginNick_Ptr(ac);
    return ret;
}
char * CQ_getStrangerInfo(int64_t var0, cq_bool_t var1)
{
    char * ret = CQ_getStrangerInfo_Ptr(ac, var0, var1);
    return ret;
}
char * CQ_getFriendList(cq_bool_t var0)
{
    char * ret = CQ_getFriendList_Ptr(ac, var0);
    return ret;
}
char * CQ_getGroupList()
{
    char * ret = CQ_getGroupList_Ptr(ac);
    return ret;
}
char * CQ_getGroupInfo(int64_t var0, cq_bool_t var1)
{
    char * ret = CQ_getGroupInfo_Ptr(ac, var0, var1);
    return ret;
}
char * CQ_getGroupMemberList(int64_t var0)
{
    char * ret = CQ_getGroupMemberList_Ptr(ac, var0);
    return ret;
}
char * CQ_getGroupMemberInfoV2(int64_t var0, int64_t var1, cq_bool_t var2)
{
    char * ret = CQ_getGroupMemberInfoV2_Ptr(ac, var0, var1, var2);
    return ret;
}
char * CQ_getCookiesV2(char * var0)
{
    char * ret = CQ_getCookiesV2_Ptr(ac, var0);
    free(var0);
    return ret;
}
int32_t CQ_getCsrfToken()
{
    int32_t ret = CQ_getCsrfToken_Ptr(ac);
    return ret;
}
char * CQ_getAppDirectory()
{
    char * ret = CQ_getAppDirectory_Ptr(ac);
    return ret;
}
char * CQ_getRecordV2(char * var0, char * var1)
{
    char * ret = CQ_getRecordV2_Ptr(ac, var0, var1);
    free(var0);
    free(var1);
    return ret;
}
char * CQ_getImage(char * var0)
{
    char * ret = CQ_getImage_Ptr(ac, var0);
    free(var0);
    return ret;
}
cq_bool_t CQ_canSendImage()
{
    cq_bool_t ret = CQ_canSendImage_Ptr(ac);
    return ret;
}
cq_bool_t CQ_canSendRecord()
{
    cq_bool_t ret = CQ_canSendRecord_Ptr(ac);
    return ret;
}
int32_t CQ_addLog(int32_t var0, char * var1, char * var2)
{
    int32_t ret = CQ_addLog_Ptr(ac, var0, var1, var2);
    free(var1);
    free(var2);
    return ret;
}
*/
import "C"
import sc "golang.org/x/text/encoding/simplifiedchinese"

func cString(str string) *C.char {
	gbstr, _ := sc.GB18030.NewEncoder().String(str)
	return C.CString(gbstr)
}

func goString(str *C.char) string {
	utf8str, _ := sc.GB18030.NewDecoder().String(C.GoString(str))
	return utf8str
}

func cBool(b bool) C.int32_t {
	if b {
		return 1
	}
	return 0
}
func sendPrivateMsg(var0 int64, var1 string) int32 {
	return int32(C.CQ_sendPrivateMsg(
		C.int64_t(var0), cString(var1),
	))
}
func sendGroupMsg(var0 int64, var1 string) int32 {
	return int32(C.CQ_sendGroupMsg(
		C.int64_t(var0), cString(var1),
	))
}
func sendDiscussMsg(var0 int64, var1 string) int32 {
	return int32(C.CQ_sendDiscussMsg(
		C.int64_t(var0), cString(var1),
	))
}
func deleteMsg(var0 int64) int32 {
	return int32(C.CQ_deleteMsg(
		C.int64_t(var0),
	))
}
func sendLikeV2(var0 int64, var1 int32) int32 {
	return int32(C.CQ_sendLikeV2(
		C.int64_t(var0), C.int32_t(var1),
	))
}
func setGroupKick(var0 int64, var1 int64, var2 bool) int32 {
	return int32(C.CQ_setGroupKick(
		C.int64_t(var0), C.int64_t(var1), cBool(var2),
	))
}
func setGroupBan(var0 int64, var1 int64, var2 int64) int32 {
	return int32(C.CQ_setGroupBan(
		C.int64_t(var0), C.int64_t(var1), C.int64_t(var2),
	))
}
func setGroupAnonymousBan(var0 int64, var1 string, var2 int64) int32 {
	return int32(C.CQ_setGroupAnonymousBan(
		C.int64_t(var0), cString(var1), C.int64_t(var2),
	))
}
func setGroupWholeBan(var0 int64, var1 bool) int32 {
	return int32(C.CQ_setGroupWholeBan(
		C.int64_t(var0), cBool(var1),
	))
}
func setGroupAdmin(var0 int64, var1 int64, var2 bool) int32 {
	return int32(C.CQ_setGroupAdmin(
		C.int64_t(var0), C.int64_t(var1), cBool(var2),
	))
}
func setGroupAnonymous(var0 int64, var1 bool) int32 {
	return int32(C.CQ_setGroupAnonymous(
		C.int64_t(var0), cBool(var1),
	))
}
func setGroupCard(var0 int64, var1 int64, var2 string) int32 {
	return int32(C.CQ_setGroupCard(
		C.int64_t(var0), C.int64_t(var1), cString(var2),
	))
}
func setGroupLeave(var0 int64, var1 bool) int32 {
	return int32(C.CQ_setGroupLeave(
		C.int64_t(var0), cBool(var1),
	))
}
func setGroupSpecialTitle(var0 int64, var1 int64, var2 string, var3 int64) int32 {
	return int32(C.CQ_setGroupSpecialTitle(
		C.int64_t(var0), C.int64_t(var1), cString(var2), C.int64_t(var3),
	))
}
func setDiscussLeave(var0 int64) int32 {
	return int32(C.CQ_setDiscussLeave(
		C.int64_t(var0),
	))
}
func setFriendAddRequest(var0 string, var1 int32, var2 string) int32 {
	return int32(C.CQ_setFriendAddRequest(
		cString(var0), C.int32_t(var1), cString(var2),
	))
}
func setGroupAddRequestV2(var0 string, var1 int32, var2 int32, var3 string) int32 {
	return int32(C.CQ_setGroupAddRequestV2(
		cString(var0), C.int32_t(var1), C.int32_t(var2), cString(var3),
	))
}
func getLoginQQ() int64 {
	return int64(C.CQ_getLoginQQ())
}
func getLoginNick() string {
	return goString(C.CQ_getLoginNick())
}
func getRawStrangerInfo(var0 int64, var1 bool) string {
	return goString(C.CQ_getStrangerInfo(
		C.int64_t(var0), cBool(var1),
	))
}
func getRawFriendList(var0 bool) string {
	return goString(C.CQ_getFriendList(
		cBool(var0),
	))
}
func getRawGroupList() string {
	return goString(C.CQ_getGroupList())
}
func getRawGroupInfo(var0 int64, var1 bool) string {
	return goString(C.CQ_getGroupInfo(
		C.int64_t(var0), cBool(var1),
	))
}
func getRawGroupMemberList(var0 int64) string {
	return goString(C.CQ_getGroupMemberList(
		C.int64_t(var0),
	))
}
func getRawGroupMemberInfoV2(var0 int64, var1 int64, var2 bool) string {
	return goString(C.CQ_getGroupMemberInfoV2(
		C.int64_t(var0), C.int64_t(var1), cBool(var2),
	))
}
func getCookiesV2(var0 string) string {
	return goString(C.CQ_getCookiesV2(
		cString(var0),
	))
}
func getCsrfToken() int32 {
	return int32(C.CQ_getCsrfToken())
}
func getAppDirectory() string {
	return goString(C.CQ_getAppDirectory())
}
func getRecordV2(var0 string, var1 string) string {
	return goString(C.CQ_getRecordV2(
		cString(var0), cString(var1),
	))
}
func getImage(var0 string) string {
	return goString(C.CQ_getImage(
		cString(var0),
	))
}
func canSendImage() bool {
	return 0 != (C.CQ_canSendImage())
}
func canSendRecord() bool {
	return 0 != (C.CQ_canSendRecord())
}
func addLog(var0 int32, var1 string, var2 string) int32 {
	return int32(C.CQ_addLog(
		C.int32_t(var0), cString(var1), cString(var2),
	))
}