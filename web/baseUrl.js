let baseUrl= "";   //这里是一个默认的url，可以没有
let wsUrl = "";
switch (process.env.NODE_ENV) {
    case 'prv':
        baseUrl = "/api"  //开发环境url
        wsUrl = "localhost:8088"
        break
    case 'pro':   // 注意这里的名字要和步骤二中设置的环境名字对应起来
        baseUrl = "/"  //测试环境中的url
        wsUrl = window.location.host
        break
}
 
export {baseUrl,wsUrl};