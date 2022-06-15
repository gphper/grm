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
                <el-input v-model="ruleForm.key"/>
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

            <el-form-item label="键值" prop="value">
                <el-input
                    v-model="ruleForm.value"
                    :rows="6"
                    type="textarea"
                    placeholder="输入键值"
                />
            </el-form-item>

            <el-form-item>
                <el-button type="primary" @click="onSubmit ">确定</el-button>
            </el-form-item>
        </el-form>

    </el-dialog>
</template>

<script>
import { reactive, ref } from '@vue/reactivity';
import { computed, watch } from '@vue/runtime-core';
import { useStore } from 'vuex';
export default {
    name:"DataForm",
    setup(props) {

        const store = useStore()
        const dialogFormVisible = computed(() => store.state.showDataForm)

        let ruleFormRef = ref(null);
        let ruleForm = reactive({
            key: '',
            type: '',
            hkey: '',
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
                value: 'string',
                label: 'String',
            },
            {
                value: 'list',
                label: 'List',
            },
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
        ])

        const onSubmit = ()=>{
            console.log(ruleForm);
        }

        const onReset = ()=>{
            store.commit("switchDataForm")
            ruleFormRef.value.resetFields();
        }

        watch(
            () => props.visible,
            (value) => {
                dialogFormVisible.value = value
            }
        )

        return{
            options,
            dialogFormVisible,
            rules,
            ruleFormRef,
            ruleForm,
            onSubmit,
            onReset
        }
    },
}
</script>