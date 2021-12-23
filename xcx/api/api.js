
import {request,getUrl}  from '../utils/util';

const http = getUrl()
// 登录
const getUserLogin = data => {
    // data.str=""
    return request({data:data,url:http+'/login'})
};
// 授权
const getUserSaveInfo = data => {
    return request({data:data,url:http+'/user/save/info'})
};

const getUploadToen = data => {
    return request({data:data,url:http+'/upload/token'})
};
const userImageAdd = data => {
    return request({data:data,url:http+'/user/image/add'})
};
const userImageList = data => {
    return request({data:data,url:http+'/user/image/list'})
};
const delImg = data => {
    return request({data:data,url:http+'/user/image/del'})
};
const loveImg = data => {
    return request({data:data,url:http+'/user/image/love'})
};

module.exports = {
    getUserSaveInfo,
    getUserLogin,
    getUploadToen,
    userImageAdd,
    userImageList,
    delImg,
    loveImg
}