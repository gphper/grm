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
                <span style="width:50px;">{{item.db}} ({{item.keys}})</span>
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
                            <el-button @click.stop="infoServe('华为云地址连接')" type="primary" circle><i class="iconfont icon-add1"></i></el-button>
                        </template>
                    </el-popover>
                </el-row>
            </template>
            <div :id="'tree#'+key+'-'+item.servicekey"></div>
    </el-sub-menu>
    <!-- <MenuNode></MenuNode> -->
</el-menu>

</template>

<script>
import { createApp,h, ref } from '@vue/runtime-dom'
import MenuTreeSolt from '@/components/MenuTreeSolt.vue'
import { ElTree,ElButton } from 'element-plus'
import {getKeys} from '@/api/index.js'
import store from '@/store/index.js'

export default {
    name:"SubMenu",
    props:{
        data:{
            type:Object
        }
    },
    setup(props) {

        const handleOpen = function(index){
            
            getKeys({"index":index}).then((res) => {
                document.getElementById("menu#"+index).nextElementSibling.setAttribute("style","display:flex")
                genNode("tree#"+index,ref(res.data));
            });
        };

        const handleClose = function(){
        };

        const genNode = function(id,dataS){
            const vdom = createApp({    
            setup() {
                const data = dataS
                return { data }
            },
            render() {
            
                return h(
                    ElTree,
                    {
                        data:this.data,
                    },
                    {
                        default: ({node,data})=>h(MenuTreeSolt,
                        {
                            node,
                            data
                        })
                    }
                )
            }
            
            });
            
            const parent = document.getElementById(id)
            vdom.use(store).use(ElTree).use(ElButton).mount(parent)
        };

        return {
            props,
            handleOpen,
            handleClose,
            genNode
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