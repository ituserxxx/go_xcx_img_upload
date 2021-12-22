import Dialog from '../../miniprogram_npm/@vant/weapp/dialog/dialog';
const app = getApp()
const api =  require("../../api/api.js") ;
Page({
  data: {
    img_list:[],
    page:1,
    previewList:[],
    show_empty:"none",
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
          if (res.data.img_list){
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
    var imgId = e.currentTarget.dataset.item.id
    // wx.showModal({
    //   title: '提示',
    //   content: '确定要删除此图片吗？',
    //   success: function (res) {
    //     if (res.confirm) {
    //       api.delImg({
    //         uid:app.globalData.uid,
    //         id:imgId
    //       }).then(res=>{
    //         if (res.code == 1){
    //             Dialog.alert({
    //               message: '删除成功~~',
    //             }).then(() => {
    //               // on close
    //             });
    //           }else{
    //             Dialog.alert({
    //               message: '删除失败了，稍后再试~',
    //             }).then(() => {
    //               // on close
    //             });
    //           }
    //       })
    //     } else if (res.cancel) {
    //        return false;       
    //       }
    //   }
    // })
  },
  mytouchstart: function (e) {  //记录触屏开始时间
    // this.setData({start:e.timeStamp })
    console.log("记录触屏开始时间")

  },
  mytouchend: function (e) {  //记录触屏结束时间
    // this.setData({end: e.timeStamp })
    console.log("记录触屏结束时间")
  }, 
 deleteitem:function(e) {  
   console.log("长按事件内容")  
  },
 onLinkDatail:function(e) {
  console.log("单击事件内容")
  // this.preview(e)
  //  if (_that.data.end - _that.data.start < 350){
     
  //  }
  }

})
