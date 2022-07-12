let baseUrl= "";   //这里是一个默认的url，可以没有
switch (process.env.NODE_ENV) {
    case 'prv':
        baseUrl = "/api"  //开发环境url
        break
    case 'pro':   // 注意这里的名字要和步骤二中设置的环境名字对应起来
        baseUrl = "/"  //测试环境中的url
        break
}
 
export default baseUrl;