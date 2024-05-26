<template>
  <div>
    <div class="background">
      <img :src="imgSrc" width="100%" height="100%" alt=""/>
    </div>
    <div class="login_container"   >
      <!-- 登录块 -->
      <div class="login_box">
        <!-- 表单区域 -->
        <h2>欢迎光临dyy米其林餐厅</h2>

        <el-form ref="loginFormRef" :model="loginForm"   :rules="loginRules" :v-model="login" class="login_form" label-width="0">
          <!-- 用户id -->

          <el-form-item prop="id">
            <el-input v-model="loginForm.id" prefix-icon="iconfont icon-yonghu" class="in"></el-input>
          </el-form-item>
          <!-- 密码 -->
          <el-form-item  prop="password">
            <el-input v-model="loginForm.password" prefix-icon="iconfont icon-mima"  show-password class="in"></el-input>
          </el-form-item>

          <!-- 按钮 -->
          <el-form-item  class="btns">
            <el-button type="primary" @click="login" >提交</el-button>
            <el-button type="info" @click="resetLoginForm()">重置</el-button>
          </el-form-item>
        </el-form>
      </div>
    </div>
  </div>
</template>

<script>
export default{
  data(){
    return {
      imgSrc:require('./assets/backgg.jpg'),
      //表单数据
      loginForm:{
        id:"",
        password:""
      },
      //验证对象
      loginRules:{
        //校验id
        id:[
          { required: true, message: '用户id为必填项', trigger: 'blur' },
          { min: 0, max: 10, message: '长度在 0 到 10 个字符', trigger: 'blur' }
        ],
        //校验密码
        password:[
          { required: true, message: '用户密码为必填项', trigger: 'blur' },
          { min: 6, max: 10, message: '长度在 6 到 10 个字符', trigger: 'blur' }
        ],
      },
    };
  },
  methods: {
    resetLoginForm(){//重置用户名和密码方法
      this.$refs.loginFormRef.resetFields();
    },
    login(){
      this.$refs.loginFormRef.validate(async valid =>{
        //1.验证失败则结束
        if(!valid) return;
        const {data:res} = await this.$http.post("login",this.loginForm);
        //2.是否登入后台
        if(res.flag=="ok"){
          //3.登陆成功，跳转到首页
          this.$message.success("操作成功");

          //4.存储user信息
          window.sessionStorage.setItem("user",res.user);
          console.log(res.user);
          window.sessionStorage.setItem("userid",res.user.id);
          //5.进入首页
          this.$router.push({path:"/main"});
        }else{
          this.$message.error("操作失败");
        }
      })
    },
  },

}
</script>

<style  scoped>
.login_container{
  background-color:palegreen;
}
.login_box{
  font-size: 28px;
  width: 450px;
  height: 300px;
  /* background-color: #ffffff; */
  background: rgba(255,255,255,0.3);
  border: 5px solid black;
  opacity: 0.8;
  filter: alpha(opacity=90);
  border-radius: 20px;
  position:fixed;
  left: 50%;
  top:50%;
  transform: translate(-50%,-50%);

}
.avatar_box{
  width: 130px;
  height: 130px;
  border:-30px solid #eee;
  border-radius: 50%;
  padding: 0px;
  box-shadow: 0 0 2px #ddd;
  left:50%;

}
.btns{
  display: flex;
  justify-content: space-around;
}
.login_form{
  position: absolute;
  bottom: 0%;
  width: 100%;
  border: 30px;
}
.background{
  width:100%;
  height:100%;  /**宽高100%是为了图片铺满屏幕 */
  z-index:-1;
  position: fixed;
  top:0;
  left:0;
}
.in{
  margin-left: 10%;
  margin-right: 10%;
  width: 350px;
}
h1,h2,h3{



  color:rgb(7, 7, 7);

  text-align:center;

}
</style>
