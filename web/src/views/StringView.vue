<template>
    <div>
        <el-card class="box-card">
            <template #header>
                <el-row :gutter="20">
                    <el-col :span="12">
                        <el-input  disabled :value="props.xkey">
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
            <div class="toolbox">键值：xxxKB</div>
            <el-input v-model="data" :autosize="{ minRows:10, maxRows:12 }" type="textarea" placeholder="Please input" disabled />
        </el-card>
    </div>
</template>

<style scoped>
.toolbox{
    padding: 10px;
    width: 100%;
}
</style>

<script>
import { ref } from '@vue/reactivity'

export default {
    name:"StringView",
    props:{
        xkey:{
            type:String
        },
    },
    setup(props) {

        //TODO 获取data值
        let data = ref('')
        data = "world"

        return {
            props,
            data
        }
    },
}
</script>