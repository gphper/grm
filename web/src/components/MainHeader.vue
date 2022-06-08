<template>
    <div>
        <el-row>
            <el-col :span="4"><span class="logo">Go Redis Manager</span></el-col>
            <el-col :span="17"></el-col>
            <el-col :span="3"><el-button type="primary" plain @click="dialogFormVisible = true"><i class="iconfont icon-add" style="font-size:20px;">&nbsp;</i>创建Redis连接</el-button></el-col>
        </el-row>
        <el-dialog v-model="dialogFormVisible" title="新建Redis连接">
            redis连接表单
        </el-dialog>
    </div>
</template>

<script>
import { ref } from "@vue/reactivity"

export default({
    name:"MainHeader",
    setup() {
        let dialogFormVisible = ref(false);
        
        return{
           dialogFormVisible 
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