<template>
    <el-dialog v-model="dialogFormVisible" @close="onReset" title="新建Redis连接">
        <el-form
            :rules="rules"
            ref="ruleFormRef"
            :model="ruleForm"
            label-width="120px"
            status-icon
        >
            <el-form-item label="连接名称" prop="connec_name">
                <el-input v-model="ruleForm.connec_name"/>
            </el-form-item>

            <el-form-item label="地址">
                <el-row :gutter="20">
                <el-col :span="12">
                    <el-input
                    v-model="ruleForm.ip"
                    label="Ip"
                    placeholder="First Name"
                    />
                </el-col>
                <el-col :span="12">
                    <el-input
                    v-model="ruleForm.port"
                    label="Port"
                    placeholder="Last Name"
                    />
                </el-col>
                </el-row>
            </el-form-item>

            <el-form-item label="密码" prop="password">
                <el-input type="password" show-password v-model="ruleForm.password"/>
            </el-form-item>

            <el-form-item label="启用SSH连接">
                <el-switch v-model="ruleForm.ssh" />
            </el-form-item>

            <el-form-item v-if="ruleForm.ssh" label="SSH地址">
                <el-row :gutter="20">
                <el-col :span="12">
                    <el-input
                    v-model="ruleForm.ssh_ip"
                    label="SshIp"
                    placeholder="First Name"
                    />
                </el-col>
                <el-col :span="12">
                    <el-input
                    v-model="ruleForm.ssh_port"
                    label="SshPort"
                    placeholder="Last Name"
                    />
                </el-col>
                </el-row>
            </el-form-item>
            <el-form-item v-if="ruleForm.ssh" label="SSH用户" prop="ssh_username">
                <el-input v-model="ruleForm.ssh_username"/>
            </el-form-item>
            <el-form-item v-if="ruleForm.ssh" label="SSH密码" prop="ssh_password">
                <el-input type="password" show-password v-model="ruleForm.ssh_password"/>
            </el-form-item>

            <el-form-item>
                <el-button type="success" @click="testConn">测试连接</el-button>
                <el-button type="primary" @click="onSubmit ">确定</el-button>
            </el-form-item>
        </el-form>

    </el-dialog>
</template>

<script>
import { reactive, ref } from "@vue/reactivity"
import { watch } from '@vue/runtime-core';

export default {
    name:"ConnForm",
    props:{
        visible:{
            type:Boolean
        }
    },
    setup(props,context) {
        let dialogFormVisible = ref(false)
        let ruleFormRef = ref(null);
        let ruleForm = reactive({
            connec_name: '',
            ip:'127.0.0.1',
            port:2379,
            password:'',
            ssh: false,
            ssh_ip:'127.0.0.1',
            ssh_port:22,
            ssh_username:'',
            ssh_password:''
        })
        let rules = {
            connec_name:[
                {required:true,message: '请输入连接名称', trigger: 'blur' }
            ],
            ip:[
                {required:true,message: '请输入IP地址', trigger: 'blur' }
            ],
            port:[
                {required:true,message: '请输入端口号', trigger: 'blur' }
            ],
        }

        const testConn = ()=>{
            console.log(ruleForm);
        }

        const onSubmit = ()=>{
            console.log(ruleForm);
        }

        const onReset = ()=>{
            context.emit("close");
            ruleFormRef.value.resetFields();
        }


        watch(
            () => props.visible,
            (value) => {
                dialogFormVisible.value = value
            }
        )

        return{
            dialogFormVisible,
            rules,
            ruleFormRef,
            ruleForm,
            testConn,
            onSubmit,
            onReset
        }
    },
}
</script>