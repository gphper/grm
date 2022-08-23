<!--
 * @Description: 
 * @Author: gphper
 * @Date: 2022-03-06 13:39:26
-->
<template>

 <el-container  class="login-bg">
    <el-card class="box-card" :style="{marginTop:top}">
      <template #header>
        <div class="card-header">
          <div class="header-title">GRM登录</div>
        </div>
      </template>
      <el-form ref="formRef" :model="form" :rules="rules">
          <el-form-item prop="username">
            <el-input v-model="form.username" placeholder="用户名"></el-input>
          </el-form-item>
          <el-form-item prop="password">
            <el-input type="password" show-password v-model="form.password" placeholder="密码"></el-input>
          </el-form-item>
          
          <div class="pannel-footer">
              <el-form-item>
                <el-button type="primary" @click="onSubmit">登录</el-button>
                <el-button type="primary" @click="onReset()">重置</el-button>
              </el-form-item>
          </div>
      </el-form>

    </el-card>
 </el-container>
    

</template>

<style scoped>
  .box-card{
    width: 350px;
    height: 250px;
    margin:0 auto;
  }

  .header-title{
    text-align: center;
  }

  .pannel-footer{
    padding-left: 90px;
  }

  .login-bg{
    background-color: #606266;
  }
</style>

<script>
import { reactive, ref } from '@vue/reactivity'
import { onMounted } from '@vue/runtime-core'
import { login } from "@/api/user.js"
import { useRouter } from 'vue-router'

export default {
  name:"LoginView",
  setup() {

    const form = reactive({
      username:"",
      password:""
    })

    const router = useRouter()

    const top = ref(0)

    const formRef = ref(null)

    onMounted(()=>{
      document.title = "go redis manager";
      top.value = document.documentElement.clientHeight/2-100 + 'px';
    })

    const rules = {
      username:[
        {required:true,message: '请输入用户名', trigger: 'blur' }
      ],
      password:[
        {required:true,message: '请输入密码', trigger: 'blur' }
      ]
    }

    function onSubmit(){
      login(form).then((res)=>{
        if(res.status){
          sessionStorage.setItem('auth',JSON.stringify(res.data))
          router.push('/')
        }
      })
    }

    function onReset(){
      formRef.value.resetFields()
    }

    return {onSubmit,onReset,rules,form,formRef,top}
  },
}
</script>

