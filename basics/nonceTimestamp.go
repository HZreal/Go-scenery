package main

// 防重放逻辑

// String token = request.getHeader("token");
// String timestamp = request.getHeader("timestamp");
// String nonceStr = request.getHeader("nonceStr");
//
// String url = request.getHeader("url");
//
// String signature = request.getHeader("signature");
//
//
// if(StringUtil.isBlank(token) || StringUtil.isBlank(timestamp) || StringUtil.isBlank(nonceStr) || StringUtil.isBlank(url)
// || StringUtil.isBlank(signature))
// {
//    return;
// }
//
// UserTokenInfo userTokenInfo = TokenUtil.getUserTokenInfo(token);
//
// if(userTokenInfo == null){
//    return;
// }
//
// if(!request.getRequestURI().equal(url)){
// return;
// }
//
// if(DateUtil.getSecond()-DateUtil.toSecond(timestamp) > 60){
//    return;
// }
//
// if(RedisUtils.haveNonceStr(userTokenInfo,nonceStr)){
//    return;
// }
//
// String stringB = SignUtil.signature(token, timestamp, nonceStr, url, request);
// if(!signature.equals(stringB)){
//    return;
// }
// RedisUtils.saveNonceStr(userTokenInfo,nonceStr,60);
