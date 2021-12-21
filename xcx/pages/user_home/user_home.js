const app = getApp()
const api =  require("../../api/api.js") ;
Page({
  data: {
    img_list:[],
    page:1,
    previewList:[],
  },
  onLoad() {
      this.getImgList()
  },
  getImgList:function(){
    api.userImageList({
      uid:app.globalData.uid,
      page:this.data.page
    }).then(res=>{
      if (res.code == 1){
        this.setData({
          img_list:res.data.img_list
        })
      }
    })
  },
  //预览图片，放大预览
  preview(event) {
    console.log(event)
    let currentUrl = event.target.dataset.item.img_url

    wx.previewImage({
      // current: currentUrl, // 当前显示图片的http链接
      url:currentUrl,
      // urls: this.data.img_list // 需要预览的图片http链接列表
    })
  }
})
