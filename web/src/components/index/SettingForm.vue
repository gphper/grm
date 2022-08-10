<template>
    <el-dialog v-model="dialogFormVisible" @close="onReset" title="设置">
        <el-form
            ref="settingFormRef"
            :model="settingForm"
            label-width="120px"
            status-icon
        >

            <el-form-item label="树状结构" prop="tree">
                <el-switch v-model="settingForm.tree" />
            </el-form-item>

            <el-form-item label="分隔符" prop="separator">
                <el-input v-model="settingForm.separator" placeholder=":"/>
            </el-form-item>

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

export default {
    name:"SettingForm",
    setup() {
        const store = useStore()
        const dialogFormVisible = computed(() => store.state.showSettingForm)

        let settingFormRef = ref(null);
        let settingForm = reactive({
            tree:true,
            separator:":"
        })

        const onSubmit = ()=>{
            console.log(settingForm)
        }

        const onReset = ()=>{
            store.commit("switchSettingForm")
        }

        return{
            settingForm,
            settingFormRef,
            dialogFormVisible,
            onSubmit,
            onReset,
        }
    },
}
</script>