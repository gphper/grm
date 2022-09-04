<template>
    <el-scrollbar>
        <el-menu
                class="el-menu-vertical-demo"
                ref="connMenu"
                @open="handleOpen"
                @close="handleClose"
            >
                <el-sub-menu :index="key+'_'+item.key" v-for="item,key in connetions" :key="key">
                
                <template #title>
                    <el-icon><i class="iconfont icon-server"></i></el-icon>
                    <span style="width:60px;">{{item.name}}</span>
                    <el-row :id="'menu#'+key+'_'+item.key">
                        <el-popover trigger="hover" content="关闭连接">
                            <template #reference>
                                <el-button @click.stop="closeDb(key+'_'+item.key)" type="danger" title="关闭连接" circle><i class="iconfont icon-close"></i></el-button>
                            </template>
                        </el-popover>

                        <el-popover trigger="hover" content="打开命令行">
                            <template #reference>
                                <el-button @click.stop="terminalDb('cmd',item.key,0,item.name)" type="success" circle><i class="iconfont icon-terminal"></i></el-button>
                            </template>
                        </el-popover>

                        <el-popover trigger="hover" content="服务信息">
                            <template #reference>
                                <el-button @click.stop="infoServe(item.key,item.name)" type="primary" circle><i class="iconfont icon-info"></i></el-button>
                            </template>
                        </el-popover>

                        <el-popover trigger="hover" content="删除服务">
                            <template #reference>
                                <el-button @click.stop="delServe(item.key)" type="danger" circle><i class="iconfont icon-shanchu"></i></el-button>
                            </template>
                        </el-popover>
                    </el-row>
                </template>
                <div :id="'sub#'+key+'_'+item.key"></div>
                </el-sub-menu>
        </el-menu>
    </el-scrollbar>
</template>

<script>
import { computed, createApp,h, onMounted, ref } from '@vue/runtime-dom'
import { ElButton,ElMenu, ElRow, ElIcon,ElPopover, ElSubMenu,ElDialog,ElInput,ElForm,ElFormItem,ElCol, ElMessageBox } from 'element-plus'
import {NewShell} from '@/utils/terminal.js'
import SubMenu from "@/components/index/SubMenu.vue"
import {getConnList,delConn} from "@/api/base.js"
import {openDb} from "@/api/index.js"
import { useStore } from 'vuex'
import { useRouter } from 'vue-router'


export default{
    name:"MenuSection",
    setup(){

        let connMenu = ref(null);
        
        const store = useStore()
        const router = useRouter()

        const connetions = computed(() => store.state.connList);

        const handleOpen = function(index){
            if(!document.getElementById("menu#"+index).nextElementSibling.getAttribute("style")){
                openDb({"index":index}).then((res) => {
                    genSubMenu("sub#"+index,res.data);
                    document.getElementById("menu#"+index).nextElementSibling.setAttribute("style","display:flex");
                });
            }
        };

        const handleClose = function(){
            
        };
        
        const closeDb = function(index){
            connMenu.value.close(index+'');
            // document.getElementById("menu#"+index).nextElementSibling.setAttribute("style","display:none")
            document.getElementById("menu#"+index).nextElementSibling.removeAttribute("style");
        };
        
        const terminalDb = function(name,sk,db,sname){
            store.commit("setTagsItem", {
                title: sname+':'+db,
                name: name,
                content: "",
                id:name
            });
            store.commit("setCurrentTag", name);

            setTimeout(()=>{
                NewShell(name,sk,db,sname)
            },500);
            
        };

        const infoServe = function(sk,name){
            store.commit("setTagsItem", {
                title: name+"-监控",
                name: "grmmonitor",
                content: "",
                view: true
            });
            store.commit("setCurrentTag", "grmmonitor");
            
            router.push({name:'monitor', params:{ sk }})
        };
    
        const delServe = function(connec_id){
            ElMessageBox.confirm(
                '确定删除该服务信息吗？',
                '操作提示',
                {
                confirmButtonText: '是的',
                cancelButtonText: '再想想',
                type: 'warning',
                }
            ).then(() => {
                delConn({
                    "service_name":connec_id
                }).then((res)=>{
                    console.log(res)
                    if(res.status == 200){
                        store.commit("delConn",res.data.data)
                    }
                })
            })
        }

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
            vdomx.use(store).use(ElMenu).use(ElSubMenu).use(ElButton).use(ElIcon).use(ElRow).use(ElPopover).use(ElDialog).use(ElInput).use(ElForm).use(ElFormItem).use(ElCol).mount(parentx)
        };

        onMounted(()=>{
            getConnList().then((res) => {
                store.state.connList = res.data
            });
        })

        return{
            handleOpen,
            handleClose,
            closeDb,
            terminalDb,
            infoServe,
            delServe,
            genSubMenu,
            connMenu,
            connetions
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

</style>