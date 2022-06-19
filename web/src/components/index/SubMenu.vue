<template>
    <el-menu 
        default-active="2"
        class="el-menu-vertical-demo"
        ref="connMenu"
        @open="handleOpen"
        @close="handleClose"
    >
        <el-sub-menu class="db" index="1-1">
            <template #title>
                <el-icon><i class="iconfont icon-database"></i></el-icon>
                <span>DB0 (30)</span>
                <el-row>
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
                <el-row>
                    <i v-if="point == 'top'" class="iconfont icon-xiangshangjiantou"></i>
                    <i v-if="point == 'bottom'" class="iconfont icon-xiangxiajiantou"></i>
                </el-row>
            </template>
            <div id="tree#1-1"></div>
    </el-sub-menu>
</el-menu>

</template>

<script>
import { createApp,h, ref } from '@vue/runtime-dom'
import MenuTreeSolt from '@/components/MenuTreeSolt.vue'
import { ElTree,ElButton } from 'element-plus'
import store from '@/store/index.js'

export default {
    name:"SubMenu",
    props:{
        data:{
            type:Object
        }
    },
    setup(props) {

        let point = ref('')

        let dataS = ref([
            {
                id: 1,
                label: 'Level one 1',
                children: [
                {
                    id: 4,
                    label: 'Level two 1-1',
                    children: [
                    {
                        id: 9,
                        label: 'string_key',
                    },
                    {
                        id: 10,
                        label: 'list_key',
                    },
                    ],
                },
                ],
            },
            {
                id: 2,
                label: 'Level one 2',
                children: [
                {
                    id: 5,
                    label: 'Level two 2-1',
                },
                {
                    id: 6,
                    label: 'Level two 2-2',
                },
                ],
            },
            {
                id: 3,
                label: 'Level one 3',
                children: [
                {
                    id: 7,
                    label: 'Level two 3-1',
                },
                {
                    id: 8,
                    label: 'Level two 3-2',
                },
                ],
            },
        ]);

        const handleOpen = function(index){
            genNode("tree#"+index,dataS);
            point.value = "bottom";
        };

        const handleClose = function(){
            point.value = "top";
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
            point,
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
</style>