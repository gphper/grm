<template>
    <div>

        <el-menu 
            default-active="2"
            class="el-menu-vertical-demo"
            ref="connMenu"
            @open="handleOpen"
            @close="handleClose"
        >
                <el-sub-menu v-for="item,key in props.data" :key="key" class="db" :index="key+'-'+item.servicekey">
                    <template #title>
                        <el-icon><i class="iconfont icon-database"></i></el-icon>
                        <span style="width:50px;">{{item.db}} (<span :id="'dbnum#'+key+'-'+item.servicekey">0</span>/{{item.keys}})</span>
                        <el-row :id="'menu#'+key+'-'+item.servicekey">
                            <el-popover trigger="hover" content="重载">
                                <template #reference>
                                    <el-button @click.stop="reload(item.servicekey,key)" type="success" title="重载" circle><i class="iconfont icon-zhongqi"></i></el-button>
                                </template>
                            </el-popover>

                            <el-popover trigger="hover" content="搜索">
                                <template #reference>
                                    <el-button @click.stop="search(item.servicekey,key)" type="success" circle><i class="iconfont icon-chaxun"></i></el-button>
                                </template>
                            </el-popover>

                            <el-popover trigger="hover" content="新增键值">
                                <template #reference>
                                    <el-button @click.stop="append(item.servicekey,key,true)" type="primary" circle><i class="iconfont icon-add1"></i></el-button>
                                </template>
                            </el-popover>

                        </el-row>
                    </template>
                    
                    <div :id="'page#'+key+'-'+item.servicekey"></div>
                    <div :id="'tree#'+key+'-'+item.servicekey"></div>
            </el-sub-menu>
        </el-menu>


        <el-dialog v-model="searchVisiable" title="搜索">
            <el-form
                :rules="rules"
                ref="ruleFormRef"
                :model="searchForm"
                label-width="120px"
                status-icon
            >
                <el-form-item label="匹配模式" prop="math">
                    <el-input
                        v-model="searchForm.match"
                        :rows="6"
                        type="input"
                        placeholder="匹配模式"
                    />
                </el-form-item>

                <el-input v-model="searchForm.sk" type="hidden"/>
                <el-input v-model="searchForm.db" type="hidden"/>

                <el-form-item>
                    <el-button type="primary" @click="onSubmit ">确定</el-button>
                </el-form-item>
            </el-form>

        </el-dialog>

    </div>
</template>

<script>

import store from '@/store/index.js'
import {dbKeysList} from "@/utils/tree.js"
import { reactive, ref } from '@vue/reactivity'

export default {
    name:"SubMenu",
    props:{
        data:{
            type:Object
        }
    },
    setup(props) {

        let searchVisiable = ref(false)

        let searchForm = reactive({
            sk:'',
            db:'',
            match: '*'
        });

        let rules = {
            match:[
                {required:true,message: '请输入匹配值', trigger: 'blur' }
            ]
        }

        const handleOpen = function(index){
            dbKeysList(index,"*",0)
        };

        const handleClose = function(){
        };

        const append = (sk,db,root) => {
            store.commit("switchDataForm",{sk:sk,db:db,show:true,root:root,pre:''});
        }

        const search = (sk,db) => {
            searchVisiable.value = !searchVisiable.value
            searchForm.sk = sk
            searchForm.db = db
        }

        const onSubmit = ()=>{
            dbKeysList(searchForm.db+'-'+searchForm.sk,searchForm.match,0)
            searchVisiable.value = false
        }

        const reload = (index,db)=>{
            dbKeysList(db+'-'+index,"*",0)
        }

        return {
            rules,
            props,
            searchForm,
            searchVisiable,
            reload,
            search,
            onSubmit,
            handleOpen,
            handleClose,
            append
        }
    },
}
</script>

<style scoped>
.el-row{
    padding-left: 15%;
}

.el-sub-menu .el-menu {
    padding-left: 0px !important;
}


.el-sub-menu__icon-arrow {
    display: flex;
}
</style>