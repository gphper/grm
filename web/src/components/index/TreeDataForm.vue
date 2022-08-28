<template>
    <el-dialog v-model="dialogFormVisible" @close="onReset" title="添加新键">
        <el-form
            :rules="rules"
            ref="ruleFormRef"
            :model="ruleForm"
            label-width="120px"
            status-icon
            destroy-on-close
        >
            <el-form-item label="键名" prop="key">
                <el-input v-model="ruleForm.key" placeholder="键名">
                    <template #prepend>{{pre}}:</template>
                </el-input>
            </el-form-item>

            <el-form-item label="类型" prop="type">
                <el-select v-model="ruleForm.type" placeholder="类型">
                    <el-option
                    v-for="item in options"
                    :key="item.value"
                    :label="item.label"
                    :value="item.value"
                    />
                </el-select>
            </el-form-item>

            <el-form-item v-if="ruleForm.type=='zset'" label="分数" prop="score">
                <el-input type="number" v-model="ruleForm.score" placeholder="分值" />
            </el-form-item>

            <el-form-item v-if="ruleForm.type=='hash'" label="键名" prop="hkey">
                <el-input v-model="ruleForm.hkey" placeholder="键值" />
            </el-form-item>

            <el-form-item v-if="ruleForm.type=='stream'" label="ID" prop="id">
                <el-input v-model="ruleForm.id" placeholder="键值" />
            </el-form-item>

            <el-form-item label="键值" prop="value">
                <el-input
                    v-model="ruleForm.value"
                    :rows="6"
                    type="textarea"
                    placeholder="输入键值"
                />
            </el-form-item>

            <el-input v-model="ruleForm.sk" type="hidden"/>
            <el-input v-model="ruleForm.db" type="hidden"/>

            <el-form-item>
                <el-button type="primary" @click="onSubmit ">确定</el-button>
            </el-form-item>
        </el-form>

    </el-dialog>
</template>

<script>
import { reactive, ref } from '@vue/reactivity';
import { addHash } from '@/api/hash.js';
import { addList } from '@/api/list.js';
import { addSet } from '@/api/set.js';
import { addZset } from '@/api/zset.js';
import { addStream } from '@/api/stream.js';
import { addString } from '@/api/string.js';
import { watch } from '@vue/runtime-core';
import { ElMessage } from 'element-plus';

export default {
    name:"TreeDataForm",
    props:{
        sk:{
            type:String
        },
        db:{
            type:Number
        },
        prefix:{
            type:String
        },
        show:{
            type:Boolean
        },
        data:{
            type:Object
        }
    },
    setup(props,{emit}) {

        let dialogFormVisible = ref(props.show)
        let pre = ref(props.prefix)
        let ruleFormRef = ref(null);
        let ruleForm = reactive({
            key: '',
            type: '',
            hkey: '',
            id:'*',
            score: 0,
            value: ''
        })
        let rules = {
            key:[
                {required:true,message: '请输入键值', trigger: 'blur' }
            ],
            type:[
                {required:true,message: '请选择类型', trigger: 'blur' }
            ]
        }

       let options = ref([
            {
                value: 'set',
                label: 'Set',
            },
            {
                value: 'zset',
                label: 'Zset',
            },
            {
                value: 'hash',
                label: 'Hash',
            },
            {
                value: 'list',
                label: 'List',
            },
            {
                value: 'string',
                label: 'String',
            },
            {
                value: 'stream',
                label: 'Stream',
            },
        ])


        watch(
            ()=> [props.show,props.prefix],
            (value) => {
                dialogFormVisible.value = value[0]
                pre.value = value[1]
            }
        )


        const onSubmit = ()=>{

            let all_key = pre.value+':'+ruleForm.key
            
            switch(ruleForm.type){
                case "string":
                    addString({
                        id:all_key,
                        sk:props.sk,
                        db:props.db,
                        item:ruleForm.value,
                    }).then(()=>{
                        ElMessage({
                            message: "添加成功",
                            type: 'success',
                        });
                        emit("append",props.data,ruleForm.key)
                        onReset()
                    });
                    break;
                case "list":
                    addList({
                        id:all_key,
                        sk:props.sk,
                        db:props.db,
                        item:ruleForm.value,
                    }).then(()=>{
                        ElMessage({
                            message: "添加成功",
                            type: 'success',
                        });
                        emit("append",props.data,ruleForm.key)
                        onReset()
                    });
                    break;
                case "set":
                    addSet({
                        id:all_key,
                        sk:props.sk,
                        db:props.db,
                        item:ruleForm.value,
                    }).then(()=>{
                        ElMessage({
                            message: "添加成功",
                            type: 'success',
                        });
                        emit("append",props.data,ruleForm.key)
                        onReset()
                    });
                    break;   
                case "zset":
                    addZset({
                        id:all_key,
                        sk:props.sk,
                        db:props.db,
                        score:ruleForm.score,
                        item:ruleForm.value,
                    }).then(()=>{
                        ElMessage({
                            message: "添加成功",
                            type: 'success',
                        });
                        emit("append",props.data,ruleForm.key)
                        onReset()
                    });
                    break;    
                case "hash":
                    addHash({
                        id:all_key,
                        sk:props.sk,
                        db:props.db,
                        itemk:ruleForm.hkey,
                        itemv:ruleForm.value
                    }).then(()=>{
                        ElMessage({
                            message: "添加成功",
                            type: 'success',
                        });
                        emit("append",props.data,ruleForm.key)
                        onReset()
                    });
                    break;    
                case "stream":
                    addStream({
                        id:all_key,
                        sk:props.sk,
                        db:props.db,
                        idx:ruleForm.id,
                        item:ruleForm.value,
                    }).then(()=>{
                        ElMessage({
                            message: "添加成功",
                            type: 'success',
                        });
                        emit("append",props.data,ruleForm.key)
                        onReset()
                    });
                    break;
            }
            
        }

        const onReset = ()=>{
            emit('reset')
            ruleFormRef.value.resetFields();
        }

        return{
            pre,
            rules,
            options,
            ruleForm,
            ruleFormRef,
            dialogFormVisible,
            onSubmit,
            onReset,
        }
    },
}
</script>