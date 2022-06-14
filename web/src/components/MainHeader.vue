<template>
    <div>
        <el-row>
            <el-col :span="4"><span class="logo">Go Redis Manager</span></el-col>
            <el-col :span="17"></el-col>
            <el-col :span="3"><el-button type="primary" plain @click="dialogFormVisible = true"><i class="iconfont icon-add" style="font-size:20px;">&nbsp;</i>创建Redis连接</el-button></el-col>
        </el-row>
        <ConnForm :visible="dialogFormVisible" @close="close"></ConnForm>
    </div>
</template>

<script>
import { ref } from "@vue/reactivity"
import ConnForm from "@/components/index/ConnForm.vue"

export default({
    name:"MainHeader",
    components:{
        ConnForm
    },
    setup() {
        let dialogFormVisible = ref(false);

        const close = ()=>{
            dialogFormVisible.value = false;
        }

        return{
            close,
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