<template>
    <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        @open="handleOpen"
        @close="handleClose"
      >
        <el-sub-menu index="1">
          <template #title>
            <el-icon><i class="iconfont icon-server"></i></el-icon>
            <span>华为云地址连接</span>
            
            <el-row>
                <el-popover trigger="hover" content="关闭连接">
                    <template #reference>
                        <el-button @click.stop="closeDb" type="danger" title="关闭连接" circle><i class="iconfont icon-close"></i></el-button>
                    </template>
                </el-popover>

                <el-popover trigger="hover" content="打开命令行">
                    <template #reference>
                        <el-button @click.stop="terminalDb('cmd')" type="success" circle><i class="iconfont icon-terminal"></i></el-button>
                    </template>
                </el-popover>

                <el-popover trigger="hover" content="服务信息">
                    <template #reference>
                        <el-button @click.stop="infoServe" type="primary" circle><i class="iconfont icon-info"></i></el-button>
                    </template>
                </el-popover>
            </el-row>
          </template>
          
          <div id="tree#1"></div>
        </el-sub-menu>
      </el-menu>
</template>

<script>
import { createApp,h, ref } from '@vue/runtime-dom'
import { ElTree,ElButton } from 'element-plus'
import TreeSolt from '@/components/TreeSolt.vue'
import store from '@/store/index.js'
import {term,fitAddon} from '@/utils/terminal.js'

export default{
    name:"MenuSection",
    setup(){

        let data = ref([
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
                        label: 'Level three 1-1-1',
                    },
                    {
                        id: 10,
                        label: 'Level three 1-1-2',
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
                genNode("tree#"+index,data);
        };
        const handleClose = function(){
                console.log("close")
        };
        const closeDb = function(){
                console.log("close db")
        };
        const terminalDb = function(name){
            console.log(name)
            store.commit("setTagsItem", {
                title: name,
                name: name,
                content: "",
                xterm:name
            });
            store.commit("setCurrentTag", name);

            setTimeout(()=>{
                term.loadAddon(fitAddon)
                term.open(document.getElementById(name));
                fitAddon.fit();
                term.focus();
            },500);
            
        };

        const infoServe = function(){
                console.log("serve info")
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
                        default: ({node,data})=>h(TreeSolt,
                        {
                            node,
                            data
                        })
                    }
                )
            }
            
            });
            
            const parent = document.getElementById(id)
            vdom.use(store).use(ElButton).mount(parent)
        };
        

        return{
            handleOpen,
            handleClose,
            closeDb,
            terminalDb,
            infoServe,
            data
        }
    }

    
}
</script>

<style type="text/css" scoped>
.el-row{
    padding-left: 15%;
}
::v-deep .el-sub-menu .el-menu{
    padding-left: 30px;
}

.el-tree .el-tree-node__expand-icon.expanded
{
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
}

::v-deep .el-tree-node__content {
    height: 38px;
}
</style>