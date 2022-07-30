<template>
    <el-dialog v-model="ttlFormVisible" @close="onReset" title="修改TTL">
        <el-form
            :rules="rules"
            ref="ttlFormRef"
            :model="ttlForm"
            label-width="120px"
            status-icon
        >
            <el-form-item label="TTL" prop="ttl">
                <el-input type="number" v-model="ttlForm.ttl" placeholder="TTL" />
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="onSubmit ">确定</el-button>
            </el-form-item>
        </el-form>

    </el-dialog>
</template>

<script>
import { reactive, ref } from '@vue/reactivity';
import { watch } from '@vue/runtime-core';
import { ttlKey } from "@/api/index.js"
import { ElMessage } from 'element-plus';
export default {
    name:"TtlForm",
    props:{
        visiable:{
            type:Boolean
        },
        keyx:{
            type:String
        },
        sk:{
            type:String
        },
        db:{
            type:Number
        }
    },
    emits:{
        'reload':null
    },
    setup(props,{emit}) {

        let ttlFormRef = ref(null);
        let ttlFormVisible = ref(false);

        let ttlForm = reactive({
            ttl:0,
            id:props.keyx,
            sk:props.sk,
            db:props.db,
        })
        let rules = {
            ttl:[
                {required:true,message: '请输入TTL', trigger: 'blur' }
            ],
        }

        const onReset = ()=>{
            ttlFormRef.value.resetFields();
            emit("close")
        }

        const onSubmit = ()=>{
            ttlKey(ttlForm).then((res)=>{
                if(res.data.result){
                    ElMessage({
                        message: "设置成功",
                        type: 'success',
                    })
                    emit("reload")
                    onReset()
                }
            })
        }

        watch(
            () => props.visiable,
            (value) => {
                ttlFormVisible.value = value
            }
        )


        return{
            rules,
            ttlForm,
            ttlFormRef,
            ttlFormVisible,
            onSubmit,
            onReset
        }
    },
}
</script>