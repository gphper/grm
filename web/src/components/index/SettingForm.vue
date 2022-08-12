<template>
    <el-dialog v-model="dialogFormVisible" @close="onReset" title="设置">
        
        <el-row>
            <el-col :span="20">
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
            </el-col>
        </el-row>
        
    </el-dialog>
</template>

<script>
import { reactive, ref } from '@vue/reactivity';
import { computed, watch } from '@vue/runtime-core';
import { useStore } from 'vuex';
import { setting } from "@/api/index.js"
import { ElMessageBox } from 'element-plus';

export default {
    name:"SettingForm",
    setup() {
        const store = useStore()
        const dialogFormVisible = computed(() => store.state.showSettingForm)

        watch(
            ()=> [store.state.tree,store.state.separator],
            (value) => {
                settingForm.tree = value[0]
                settingForm.separator = value[1]
            }
        )
        if (dialogFormVisible == true){
            console.log("dssadasdas")
        }

        let settingFormRef = ref(null);
        let settingForm = reactive({
            tree:true,
            separator:":"
        })

        const onSubmit = ()=>{
            setting(settingForm).then((res)=>{
                if (res.status == 200){
                    ElMessageBox.confirm(
                        '设置成功需刷新后生效,是否现在刷新？',
                        '操作成功',
                        {
                        confirmButtonText: '现在刷新',
                        cancelButtonText: '稍后手动操作',
                        type: 'success',
                        }
                    ).then(() => {
                        window.location.reload()
                    }).catch(() => {
                        store.commit("switchSettingForm",{
                            visiable:false,
                        });
                    })
                }
            })
        }

        const onReset = ()=>{
            store.commit("switchSettingForm",{
                visiable:false,
            })
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