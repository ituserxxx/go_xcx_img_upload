const formatTime = date => {
  const year = date.getFullYear()
  const month = date.getMonth() + 1
  const day = date.getDate()
  const hour = date.getHours()
  const minute = date.getMinutes()
  const second = date.getSeconds()

  return `${[year, month, day].map(formatNumber).join('/')} ${[hour, minute, second].map(formatNumber).join(':')}`
}

const formatNumber = n => {
  n = n.toString()
  return n[1] ? n : `0${n}`
}

module.exports = {
  formatTime
}

const getUrl  = ()=>{
  return 'https://jstgs-image.vtian.top/api/v1'
}
const request = data => {
  return new Promise(function(resolve, reject){
    wx.request({
      url: data.url ,// app.url+'/pet/service_user/list',
      data: data.data||{},
      header: {
        'content-type': 'application/json' 
      },
      method: 'post',
      success:function(ref){
          if(ref.data.code==1){
            resolve(ref.data);
          }else{
            reject(ref.data);
            wx.showToast({
              title: ref.data.msg,
              icon: 'none', 
              duration: 2000     
            }) 
          }
      },
      fail:function(){
        reject({msg:'网络错误'});
        wx.showToast({
          title: '网络错误',
          icon: 'none',    //如果要纯文本，不要icon，将值设为'none'
          duration: 2000     
        }) 
      },
    });
})
}
module.exports = {
  formatTime,
  request,
  getUrl
}
