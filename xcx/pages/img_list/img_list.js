import Dialog from '../../miniprogram_npm/@vant/weapp/dialog/dialog';
const app = getApp()
const api = require("../../api/api.js");
Page({
  data: {
    img_list: [],
    page: 1,
    previewList: [],
    show_empty: "none",
    view_start_time: 0,
    view_end_time: 0,
    is_ope:false,
    ope_obj_id:0,
    ope_select:0,

  },
  onLoad() {
    this.getImgList()
  },
  getImgList: function () {
    api.userImageList({
      uid: app.globalData.uid,
      page: this.data.page
    }).then(res => {
      if (res.code == 1) {
        if (res.data.img_list) {
          this.setData({
            img_list: res.data.img_list
          })
        } else {
          this.data.show_empty = "block"
        }
      } else {
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
    wx.previewImage({
      current: currentUrl, // 当前显示图片的http链接
      urls: [currentUrl]// 需要预览的图片http链接列表
    })
  },

  mytouchstart: function (e) { 
    console.log("记录触屏开始时间")
    this.setData({ view_start_time: e.timeStamp })
  },
  mytouchend: function (e) {  
    console.log("记录触屏结束时间")
    this.setData({view_end_time: e.timeStamp })
  },
  deleteitem: function (e) {
    console.log("长按事件内容",e)
    this.setData({
      is_ope:true,
      ope_obj_id:e.currentTarget.dataset.item.id
    })
  },
  onLinkDatail: function (e) {
    console.log("单击事件内容")
    if (this.data.view_end_time - this.data.view_start_time < 350) {
      this.preview(e)
    }
  },
  ope_select:function(e){
    this.setData({
      ope_select:e.currentTarget.dataset.name
    })
  },
  confirm_ope:function(e){
      if (this.data.ope_select == 1){
        this.loveImage(this.data.ope_obj_id)
      }
      if (this.data.ope_select == 2){
        this.deleteImage(this.data.ope_obj_id)
      }
  },
  deleteImage: function (imgId) {
    api.delImg({
      uid:app.globalData.uid,
      id:imgId
    }).then(res=>{
      this.setData({
        is_ope:false,
        ope_select:0,
        ope_obj_id:0
      })
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
  },
  loveImage: function (imgId) {
    api.loveImg({
      uid:app.globalData.uid,
      id:imgId
    }).then(res=>{
      this.setData({
        is_ope:false,
        ope_select:0,
        ope_obj_id:0
      })
      if (res.code == 1){
          Dialog.alert({
            message: '收藏成功~~',
          }).then(() => {
            // on close
          });
        }else{
          Dialog.alert({
            message: '收藏失败了，稍后再试~',
          }).then(() => {
            // on close
          });
        }
    })
  },
  onShow: function () {
    this.onLoad();
  },
})
