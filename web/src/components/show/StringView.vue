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
            
            <OneDetail :data="data"></OneDetail>
            
        </el-card>
    </div>
</template>

<script>
import OneDetail from "@/components/one/OneDetail.vue"

export default {
    name: "StringView",
    props: {
        xkey: {
            type: String
        },
    },
    components:{
        OneDetail
    },
    setup(props) {
        //TODO 获取data值
        let data = `{"username":"gphper","nickname":"GPHER"}`;
        
        return {
            props,
            data,
        };
    },
}
</script>