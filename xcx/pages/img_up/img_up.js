import Dialog from '../../miniprogram_npm/@vant/weapp/dialog/dialog';
const qiniu = require("../../utils/qiniuUploader")
const api =  require("../../api/api.js") ;
const app = getApp()
Page({

  /**
   * 页面的初始数据
   */
  data: {
    fileList: [
      //实例
      // {
      //   url: 'https://img.yzcdn.cn/vant/leaf.jpg',
      //   status: 'uploading',
      //   message: '上传中',
      // }
    ],
    saveList:[],
    qiniuOption:{}
  },

  getQinniuOption:function(){
    wx.showLoading({
      title: '初始化ing',
    })
    api.getUploadToen().then(res=>{
      wx.hideLoading()
      var options = {
        region:"ECN",
        uptoken:res.data.token,
        uptokenURL:res.data.region,
        domain:res.data.domain,
      }
      this.setData({
        qiniuOption:options
      })
    }).catch(e=>{
      Dialog.alert({
        message: "初始化失败了，请联系管理员~",
      }).then(() => {
        // on close
      });
    })
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function () {
      this.getQinniuOption()
  },

  deleteImg:function(res){
    var index = res.detail.index
    var oldData = this.data.fileList
    oldData.splice(index,1)
    this.setData({
      fileList:oldData
    })
  },
  selectImg:function(res){
    var imgList = []
    console.log(res.detail)
    for (let key in res.detail.file) {
      imgList.push({url:res.detail.file[key].url});
    }
    this.setData({
      fileList :imgList
    })
  },
  doSave:function(){
    var data = {
      img_list:this.data.saveList,
      uid:app.globalData.uid,
    }
    api.userImageAdd(data).then(res=>{
      if (res.code == 1){
        Dialog.alert({
          message: "保存成功啦~",
        }).then(() => {
          // on close
        });
        this.setData({
          fileList:[],
          saveList:[]
        })
      }else{
        wx.showToast({
          title: '保存失败',
          icon: 'error',
          duration: 2000
        })
      }
    }).catch(e=>{
        console.log("api.userImageAdd",e)
    })
  },
  
  commit() {
    if(this.data.fileList.length<1){
      return;
    }
    //初始化七牛云配置
    qiniu.init(this.data.qiniuOption)
    let that = this;
    let imgList = that.data.fileList;
    let count = 0; //count记录当前已经上传到第几张图片
    wx.showToast({
      title: '上传中ing',
      icon: 'loading',
    })
    // 上传多张
    for (let i = 0; i < imgList.length; i++) {
      qiniu.upload(
        imgList[i].url,
        res => {
          count++;
          that.data.saveList.push(res.imageURL);
          if (count == imgList.length) {
            wx.hideToast()
            //全部上传完成  调用保存接口
            that.doSave();
          }
        },
        error => {
          console.error("error: " + JSON.stringify(error));
        },
        null, // 可以使用上述参数，或者使用 null 作为参数占位符
        progress => {
          console.log("上传进度", progress);
        },
        cancelTask => {
          // that.setData({ cancelTask });
          // this.cancelTask = cancelTask;
        }
      );
    }
  },
  /**
   * 生命周期函数--监听页面初次渲染完成
   */
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {

  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {

  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {

  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {

  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {

  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {

  }
})