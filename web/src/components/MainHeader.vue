<template>
    <div>
        <el-row>
            <el-col :span="4"><span class="logo">Go Redis Manager</span></el-col>
            <el-col :span="15"></el-col>
            <el-col :span="3"><el-button type="primary" plain @click="dialogFormVisible = true"><i class="iconfont icon-add" style="font-size:20px;">&nbsp;</i>创建Redis连接</el-button></el-col>
            <el-col :span="2"><el-button type="danger" plain @click="logout"><i class="iconfont icon-close" style="font-size:20px;">&nbsp;</i>退出登录</el-button></el-col>
        </el-row>
        <ConnForm :visible="dialogFormVisible" @close="close"></ConnForm>
    </div>
</template>

<script>
import { ref } from "@vue/reactivity"
import ConnForm from "@/components/index/ConnForm.vue"
import { useRouter } from "vue-router";

export default({
    name:"MainHeader",
    components:{
        ConnForm
    },
    setup() {
        let dialogFormVisible = ref(false);

        const router = useRouter()

        const close = ()=>{
            dialogFormVisible.value = false;
        }

        const logout = ()=>{
            sessionStorage.clear('auth')
            router.push('/login')
        }

        return{
            close,
            logout,
            dialogFormVisible,
        }
    },
})
</script>

<style scoped>
.logo{
    padding: 7px;
    background: #409EFF;
    color: #fff;
    font-weight: bold;
    font-size: 20px;
    border-radius: 6px;
    -moz-user-select: none;
    -webkit-user-select: none;
    -ms-user-select: none;
    -khtml-user-select: none;
    user-select: none;
}
</style>