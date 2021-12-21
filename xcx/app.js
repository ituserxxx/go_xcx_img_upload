const api =  require('./api/api') ;
// app.js
App({
  globalData: {
    uid: 0
  },
  onLaunch() {
    // 登录
    wx.login({
      success: res => {
        api.getUserLogin({
          code:res.code
        }).then(res=>{
          this.globalData.uid = res.data
        }).catch(e=>{
          console.log(res)
        })
      }
    })
  },
})
