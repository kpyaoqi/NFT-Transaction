/*
该文件为配置文件，路由前缀front_xxxx,
xxxx必须要和请求地址前缀/api/xxxx的xxxx保持一致，方便项目开发。
请求地址前缀在部署打包的时候，要去掉前面的api
*/

export const preRouter='/front_nft_pc';//路由前缀
export const IframeUrl='';//外链前缀
export const AJAX_PREFIXCONFIG='/api/nft';//请求地址前缀（部署打包时候请去掉前面api）
export const LOGIN_SUCCESS='/nft_home';//默认登录成功后跳转页面
export const STATIC_IMG='/api/static';//默认登录成功后跳转页面