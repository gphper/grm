<template>
    <el-dialog v-model="dialogFormVisible" @close="onReset" title="添加新键">
        <el-form
            :rules="rules"
            ref="ruleFormRef"
            :model="ruleForm"
            label-width="120px"
            status-icon
        >
            <el-form-item label="键名" prop="key">
                <el-input v-if="root == true" v-model="ruleForm.key" placeholder="键名"/>

                <el-input v-if="root == false" v-model="ruleForm.key" placeholder="键名">
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
import { computed } from '@vue/runtime-core';
import { useStore } from 'vuex';
import {dbKeysList} from "@/utils/tree.js"
import { addHash } from '@/api/hash.js';
import { addList } from '@/api/list.js';
import { addSet } from '@/api/set.js';
import { addZset } from '@/api/zset.js';
import { addStream } from '@/api/stream.js';
import { addString } from '@/api/string.js';

export default {
    name:"DataForm",
    setup() {
        const store = useStore()
        const dialogFormVisible = computed(() => store.state.showDataForm)
        const sk = computed(() => store.state.sk)
        const db = computed(() => store.state.db)
        const root = computed(()=>store.state.root)
        const pre = computed(()=>store.state.pre)

        let ruleFormRef = ref(null);
        let ruleForm = reactive({
            key: '',
            type: '',
            hkey: '',
            id:'*',
            score: 0,
            value: '',
            sk:sk,
            db:db
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

    
        const onSubmit = ()=>{

            let index = ruleForm.db +'-'+ruleForm.sk
            if(root.value == false){
                ruleForm.key = pre.value+':'+ruleForm.key
            }
            
            switch(ruleForm.type){
                case "string":
                    addString({
                        id:ruleForm.key,
                        sk:ruleForm.sk,
                        db:ruleForm.db,
                        item:ruleForm.value,
                    }).then(()=>{
                        dbKeysList(index)
                        onReset()
                    });
                    break;
                case "list":
                    addList({
                        id:ruleForm.key,
                        sk:ruleForm.sk,
                        db:ruleForm.db,
                        item:ruleForm.value,
                    }).then(()=>{
                        dbKeysList(index)
                        onReset()
                    });
                    break;
                case "set":
                    addSet({
                        id:ruleForm.key,
                        sk:ruleForm.sk,
                        db:ruleForm.db,
                        item:ruleForm.value,
                    }).then(()=>{
                        dbKeysList(index)
                        onReset()
                    });
                    break;   
                case "zset":
                    addZset({
                        id:ruleForm.key,
                        sk:ruleForm.sk,
                        db:ruleForm.db,
                        score:ruleForm.score,
                        item:ruleForm.value,
                    }).then(()=>{
                        dbKeysList(index)
                        onReset()
                    });
                    break;    
                case "hash":
                    addHash({
                        id:ruleForm.key,
                        sk:ruleForm.sk,
                        db:ruleForm.db,
                        itemk:ruleForm.hkey,
                        itemv:ruleForm.value
                    }).then(()=>{
                        dbKeysList(index)
                        onReset()
                    });
                    break;    
                case "stream":
                    addStream({
                        id:ruleForm.key,
                        sk:ruleForm.sk,
                        db:ruleForm.db,
                        idx:ruleForm.id,
                        item:ruleForm.value,
                    }).then(()=>{
                        dbKeysList(index)
                        onReset()
                    });
                    break;     
            }
            
        }

        const onReset = ()=>{
            store.commit("switchDataForm",{sk:"",db:0,show:false})
            ruleFormRef.value.resetFields();
        }

        return{
            pre,
            root,
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