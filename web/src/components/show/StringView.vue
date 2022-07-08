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
                        <el-button type="danger" text bg  @click="delKey(props.xkey)">删除</el-button>
                        <el-button type="warning" text bg @click="reload">重载数据</el-button>
                    </el-col>
                </el-row>
            </template>
            
            <OneDetail :data="data"></OneDetail>
            
        </el-card>
    </div>
</template>

<script>
import OneDetail from "@/components/one/OneDetail.vue"
import { onMounted, ref } from '@vue/runtime-core';
import { showString } from "@/api/string.js"
import { delKeys } from "@/api/index.js"
import store from '@/store/index.js'
import { ElMessage } from 'element-plus';

export default {
    name: "StringView",
    props: {
        xkey: {
            type: String
        },
        sk:{
            type:String
        },
        db:{
            type:Number
        }
    },
    components:{
        OneDetail
    },
    setup(props,{emit}) {
    
        let data = ref('');

        const reload = ()=>{
            showString({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
            }).then((res)=>{
                data.value = res.data.data
            })
        }

        const delKey = (key)=>{
            delKeys({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
            }).then((res)=>{
                if(res.data.status == 1){
                    //删除成功
                    ElMessage({
                        message: "删除成功",
                        type: 'success',
                    })
                }
            })

            store.commit("delTagsItem",key)
            emit("del",key)
        }

        onMounted(reload)

        return {
            props,
            data,
            reload,
            delKey
        };
    },
}
</script>