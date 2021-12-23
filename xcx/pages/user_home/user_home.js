import Dialog from '../../miniprogram_npm/@vant/weapp/dialog/dialog';
const app = getApp()
const api =  require("../../api/api.js") ;
Page({
  data: {
    img_list:[],
    page:1,
    show_empty:"none",
    default_tab_name : "love"
  },
  onLoad() {
      this.getImgList(3)
  },
  getImgList:function(status_id){
    ////1-正常，2-删除,3-喜欢的(api接口参数)
    api.userImageList({
      uid:app.globalData.uid,
      page:this.data.page,
      status_id:status_id,
    }).then(res=>{
      if (res.code == 1){
          if (res.data.img_list != null){
            this.setData({
              img_list:res.data.img_list
            })
          }else{
            this.data.show_empty = "block"
          }
        }else{
          Dialog.alert({
            message: res.msg,
          }).then(() => {
            // on close
          });
        }
    })
  },
  //预览图片，放大预览
  preview(event) {
    let currentUrl = event.currentTarget.dataset.item.img_url
    console.log(event.currentTarget.dataset.item.img_url)
    wx.previewImage({
      current: currentUrl, // 当前显示图片的http链接
      urls: [currentUrl]// 需要预览的图片http链接列表
    })
  },
  deleteImage: function (e) {
    if (this.data.default_tab_name == "rubbish"){
      return
    }
    var imgId = e.currentTarget.dataset.item.id
    wx.showModal({
      title: '提示',
      content: '确定要删除此图片吗？',
      success: function (res) {
        if (res.confirm) {
          api.delImg({
            uid:app.globalData.uid,
            id:imgId
          }).then(res=>{
            if (res.code == 1){
                Dialog.alert({
                  message: '删除成功~~',
                }).then(() => {
                  // on close
                });
              }else{
                Dialog.alert({
                  message: '删除失败了，稍后再试~',
                }).then(() => {
                  // on close
                });
              }
          })
        } else if (res.cancel) {
           return false;       
          }
      }
    })
  },
  change_tab:function(e){
      this.setData({
        img_list:[]
      })
      var name = e.detail.name
      this.setData({
        default_tab_name:name
      })
      if (name == "rubbish"){
        this.getImgList(2)
      }
      if (name == "love") {
        this.getImgList(3)
    }
  },
  onShow: function () {
    this.onLoad();
  },
})
