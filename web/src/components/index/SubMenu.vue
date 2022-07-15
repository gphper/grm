<template>
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
                <span style="width:50px;">{{item.db}} (<span :id="'dbnum#'+key+'-'+item.servicekey">{{item.keys}}</span>)</span>
                <el-row :id="'menu#'+key+'-'+item.servicekey">
                    <el-popover trigger="hover" content="重载">
                        <template #reference>
                            <el-button @click.stop="closeDb('1')" type="success" title="重载" circle><i class="iconfont icon-zhongqi"></i></el-button>
                        </template>
                    </el-popover>

                    <el-popover trigger="hover" content="打开命令行">
                        <template #reference>
                            <el-button @click.stop="terminalDb('cmd')" type="success" circle><i class="iconfont icon-terminal"></i></el-button>
                        </template>
                    </el-popover>

                    <el-popover trigger="hover" content="新增键值">
                        <template #reference>
                            <el-button @click.stop="append(item.servicekey,key)" type="primary" circle><i class="iconfont icon-add1"></i></el-button>
                        </template>
                    </el-popover> 
                </el-row>
            </template>
            <div :id="'tree#'+key+'-'+item.servicekey"></div>
    </el-sub-menu>
</el-menu>

</template>

<script>

import store from '@/store/index.js'
import {dbKeysList} from "@/utils/tree.js"

export default {
    name:"SubMenu",
    props:{
        data:{
            type:Object
        }
    },
    setup(props) {

        const handleOpen = function(index){
            dbKeysList(index)
        };

        const handleClose = function(){
        };

        const append = (sk,db) => {
            // const newChild = { id: 10, label: 'testtest', children: [] }
            // if (!data.children) {
            //     data.children = []
            // }
            // data.children.push(newChild)
            store.commit("switchDataForm",{sk:sk,db:db});
        }

        return {
            props,
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