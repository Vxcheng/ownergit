/*    */ package com.chinamobile.opertions.project.pushdata;
/*    */ 
/*    */ import java.io.IOException;
/*    */ import java.io.UnsupportedEncodingException;
/*    */ import org.apache.http.HttpEntity;
/*    */ import org.apache.http.client.ClientProtocolException;
/*    */ import org.apache.http.client.config.RequestConfig;
/*    */ import org.apache.http.client.methods.CloseableHttpResponse;
/*    */ import org.apache.http.client.methods.HttpPost;
/*    */ import org.apache.http.client.methods.HttpUriRequest;
/*    */ import org.apache.http.entity.StringEntity;
/*    */ import org.apache.http.impl.client.CloseableHttpClient;
/*    */ import org.apache.http.impl.client.HttpClients;
/*    */ import org.apache.http.util.EntityUtils;
/*    */ 
/*    */ 
/*    */ 
/*    */ 
/*    */ 
/*    */ 
/*    */ 
/*    */ 
/*    */ public class HttpUtils
/*    */ {
/*    */   private static final int CONNECT_TIME_OUT = 5000;
/*    */   private static final int REQUEST_TIME_OUT = 15000;
/*    */   private static final int SOCKET_TIME_OUT = 60000;
/*    */   
/*    */   public static String postJson(String url, String jsonData) {
/*    */     try {
/* 31 */       CloseableHttpClient client = HttpClients.createDefault();
/* 32 */       HttpPost httpPost = new HttpPost(url);
/* 33 */       httpPost.setHeader("Content-Type", "application/json;charset=UTF-8");
/*    */ 
/*    */ 
/*    */       
/* 37 */       RequestConfig conf = RequestConfig.custom().setConnectTimeout(5000).setSocketTimeout(60000).setConnectionRequestTimeout(15000).build();
/* 38 */       httpPost.setConfig(conf);
/*    */       
/* 40 */       StringEntity se = new StringEntity(jsonData, "UTF-8");
/* 41 */       httpPost.setEntity((HttpEntity)se);
/*    */       
/* 43 */       CloseableHttpResponse response = client.execute((HttpUriRequest)httpPost);
/* 44 */       HttpEntity entity = response.getEntity();
/* 45 */       String result = EntityUtils.toString(entity, "UTF-8");
/* 46 */       return result;
/* 47 */     } catch (UnsupportedEncodingException e) {
/* 48 */       e.printStackTrace();
/* 49 */     } catch (ClientProtocolException e) {
/* 50 */       e.printStackTrace();
/* 51 */     } catch (IOException e) {
/* 52 */       e.printStackTrace();
/*    */     } 
/* 54 */     return null;
/*    */   }
/*    */ }


/* Location:              C:\Users\wuhongwei\Desktop\chinamobile-pushdata-1.0.0.jar!\com\chinamobile\opertions\project\pushdata\HttpUtils.class
 * Java compiler version: 5 (49.0)
 * JD-Core Version:       1.1.3
 */