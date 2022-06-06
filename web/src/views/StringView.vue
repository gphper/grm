<template>
    <el-row :gutter="20">
        <el-col :span="12">
            <el-input v-model="input1" disabled placeholder="Please input">
                <template #prepend>STRING:</template>
            </el-input>
        </el-col>
        <el-col :span="8" :offset="2">
            <el-button text bg>重命名</el-button>
            <el-button text bg>TTL</el-button>
            <el-button type="danger" text bg>删除</el-button>
            <el-button type="warning" text bg>重载数据</el-button>
        </el-col>
    </el-row>
</template>

<script>
export default {
    name:"StringView",
    setup() {
        
    },
}
</script>