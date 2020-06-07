/*    */ package com.chinamobile.opertions.project.pushdata;
/*    */ 
/*    */ import com.alibaba.fastjson.JSONObject;
/*    */ 
/*    */ 
/*    */ 
/*    */ 
/*    */ 
/*    */ 
/*    */ public class PushUtils
/*    */ {
/* 12 */   private static String accessToken = null;
/*    */   
/*    */   public static String push(String loginUrl, String userName, String password, String pushUrl, String requestJsonStr) {
/* 15 */     String result = null;
/*    */     
/*    */     try {
/* 18 */       if (accessToken == null || accessToken.length() == 0) {
/*    */         
/* 20 */         result = _exeLogin(userName, password, requestJsonStr, loginUrl, pushUrl);
/*    */       } else {
/* 22 */         result = _push(pushUrl, requestJsonStr);
/* 23 */         JSONObject pushResultJson = JSONObject.parseObject(result);
/* 24 */         Integer pushCode = pushResultJson.getInteger("code");
/*    */         
/* 26 */         if (pushCode.intValue() == -1)
/*    */         {
/* 28 */           result = _exeLogin(userName, password, requestJsonStr, loginUrl, pushUrl);
/*    */         }
/*    */       } 
/* 31 */     } catch (Exception e) {
/* 32 */       e.printStackTrace();
/*    */     } 
/* 34 */     return result;
/*    */   }
/*    */   
/*    */   private static String _exeLogin(String userName, String password, String requestJsonStr, String loginUrl, String pushUrl) {
/* 38 */     String result = null;
/*    */     
/* 40 */     result = _login(loginUrl, userName, password);
/* 41 */     JSONObject loginResultJson = JSONObject.parseObject(result);
/* 42 */     Integer loginCode = loginResultJson.getInteger("code");
/*    */     
/* 44 */     if (loginCode.intValue() == 0) {
/* 45 */       accessToken = loginResultJson.getString("data");
/*    */       
/* 47 */       result = _push(pushUrl, requestJsonStr);
/*    */     } 
/* 49 */     return result;
/*    */   }
/*    */   
/*    */   private static String _push(String pushUrl, String requestJsonStr) {
/* 53 */     JSONObject pushJson = new JSONObject();
/* 54 */     pushJson.put("accessToken", accessToken);
/* 55 */     pushJson.put("requestJsonStr", requestJsonStr);
/* 56 */     return HttpUtils.postJson(pushUrl, pushJson.toJSONString());
/*    */   }
/*    */   
/*    */   private static String _login(String loginUrl, String userName, String password) {
/* 60 */     JSONObject loginJson = new JSONObject();
/* 61 */     loginJson.put("userName", userName);
/* 62 */     loginJson.put("password", password);
/* 63 */     return HttpUtils.postJson(loginUrl, loginJson.toJSONString());
/*    */   }
/*    */ }


/* Location:              C:\Users\wuhongwei\Desktop\chinamobile-pushdata-1.0.0.jar!\com\chinamobile\opertions\project\pushdata\PushUtils.class
 * Java compiler version: 5 (49.0)
 * JD-Core Version:       1.1.3
 */