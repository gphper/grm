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

                        <el-tooltip
                            class="box-item"
                            effect="dark"
                            :content="ttl"
                            placement="top-start"
                        >
                            <el-button text bg @click="ttlVisible = true">TTL:<span style="width: 22px;overflow: hidden;">{{ttl}}</span></el-button>
                        </el-tooltip>
                        <el-button type="danger" text bg  @click="delKey(props.sk+props.db+props.xkey)">删除</el-button>
                        <el-button type="warning" text bg @click="reload">重载数据</el-button>
                    </el-col>
                </el-row>
            </template>
            <OneDetail :data="data"></OneDetail>
            <TtlForm @close="close" @reload="reload" :visiable="ttlVisible" :sk="props.sk" :db="props.db" :keyx="props.xkey"></TtlForm>
        </el-card>
    </div>
</template>

<script>
import OneDetail from "@/components/one/OneDetail.vue"
import TtlForm from "@/components/one/TtlForm.vue";
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
        OneDetail,
        TtlForm
    },
    setup(props,{emit}) {
    
        let data = ref('');
        let ttl = ref(0);
        let ttlVisible = ref(false)

        const reload = ()=>{
            showString({
                id:props.xkey,
                sk:props.sk,
                db:props.db,
            }).then((res)=>{
                data.value = res.data.data
                ttl.value = res.data.ttl
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

        const close = ()=>{
            ttlVisible.value = false
        }

        onMounted(reload)

        return {
            props,
            data,
            ttl,
            ttlVisible,
            reload,
            delKey,
            close
        };
    },
}
</script>