<template>
    <el-menu
        default-active="2"
        class="el-menu-vertical-demo"
        ref="connMenu"
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
                        <el-button @click.stop="closeDb('1')" type="danger" title="关闭连接" circle><i class="iconfont icon-close"></i></el-button>
                    </template>
                </el-popover>

                <el-popover trigger="hover" content="打开命令行">
                    <template #reference>
                        <el-button @click.stop="terminalDb('cmd')" type="success" circle><i class="iconfont icon-terminal"></i></el-button>
                    </template>
                </el-popover>

                <el-popover trigger="hover" content="服务信息">
                    <template #reference>
                        <el-button @click.stop="infoServe('华为云地址连接')" type="primary" circle><i class="iconfont icon-info"></i></el-button>
                    </template>
                </el-popover>
            </el-row>
            <el-row>
                <i v-if="point == 'top'" class="iconfont icon-xiangshangjiantou"></i>
                <i v-if="point == 'bottom'" class="iconfont icon-xiangxiajiantou"></i>
            </el-row>
          </template>
          <div id="sub#1"></div>
        </el-sub-menu>
      </el-menu>
</template>

<script>
import { createApp,h, ref } from '@vue/runtime-dom'
import { ElButton,ElMenu, ElRow, ElIcon,ElPopover, ElSubMenu } from 'element-plus'

import store from '@/store/index.js'
import {term,fitAddon} from '@/utils/terminal.js'
import SubMenu from "@/components/index/SubMenu.vue"
// import { ElMenu, ElSubMenu } from 'element-plus/lib/components'

export default{
    name:"MenuSection",
    setup(props,context){

        let connMenu = ref(null);
        let point = ref('')
        
        const handleOpen = function(index){
            genSubMenu("sub#"+index,subData);
            //主菜单操作
            point.value = "bottom";
        };

        const handleClose = function(){
            point.value = "top";
        };
        
        const closeDb = function(index){
            connMenu.value.close(index);
            point.value = "close";
        };
        
        const terminalDb = function(name){
            store.commit("setTagsItem", {
                title: name,
                name: name,
                content: "",
                id:name
            });
            store.commit("setCurrentTag", name);

            setTimeout(()=>{
                term.loadAddon(fitAddon)
                term.open(document.getElementById(name));
                fitAddon.fit();
                term.focus();
            },500);
            
        };

        const infoServe = function(connec_id){
            context.emit("info",connec_id);
        };
        
        let subData = ref([
            {
                db:0,
                keys:10
            },
            {
                db:1,
                keys:0
            },
        ]);

        const genSubMenu = function(id,dataS){
            const vdomx = createApp({
            setup() {
                const data = dataS
                return { data }
            },
            render() {
            
                return h(
                    SubMenu,
                    {
                        data:this.data,
                    },
                    {}
                )
            }
            
            });
            
            const parentx = document.getElementById(id)
            vdomx.use(store).use(ElMenu).use(ElSubMenu).use(ElButton).use(ElIcon).use(ElRow).use(ElPopover).mount(parentx)
        };

        return{
            handleOpen,
            handleClose,
            closeDb,
            terminalDb,
            infoServe,
            genSubMenu,
            subData,
            connMenu,
            point
        }
    }

    
}
</script>

<style type="text/css" scoped>
.el-row{
    padding-left: 15%;
}
:deep(.el-sub-menu .el-menu){
    padding-left: 30px;
}

.el-tree .el-tree-node__expand-icon.expanded
{
    -webkit-transform: rotate(0deg);
    transform: rotate(0deg);
}

:deep(.el-tree-node__content){
    height: 38px;
}

:deep(.el-sub-menu__icon-arrow){
    display: none;
}

:deep(.db .el-sub-menu__title){
    /* padding-left: 10px; */
    padding-left: 15px !important;
}
</style>