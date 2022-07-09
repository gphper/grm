import StringView from '@/components/show/StringView.vue';
import ListView from '@/components/show/ListView.vue'
import { ElButton, ElCard, ElCol, ElDialog, ElDivider, ElForm, ElInput, ElMessage, ElRow, ElSelect, ElTable, ElTooltip } from 'element-plus';
import { createApp, h } from 'vue';
import { useStore } from 'vuex';

const ShowString = (key,id,sk,db,func)=>{
        const store = useStore()
        const vdom = createApp({
            setup() {
            },
            render() {
            
                return h(
                    StringView,
                    {
                        xkey:key,
                        sk:sk,
                        db:db,
                        onDel:func
                    },
                    null
                )
            }
        });

        setTimeout(()=>{
            const parent = document.getElementById(id)
            vdom.use(ElButton).use(ElInput).use(ElRow).use(ElCol).use(ElCard).use(ElSelect).use(ElTooltip).use(ElDialog).use(ElForm).use(ElMessage).use(store).mount(parent)
        },500);
}


const ShowList = (key,id,sk,db)=>{
    const vdom = createApp({    
        setup() {
        },
        render() {
        
            return h(
                ListView,
                {
                    xkey:key,
                    sk:sk,
                    db:db
                },
                null
            )
        }
        
        });

        setTimeout(()=>{
            const parent = document.getElementById(id)
            vdom.use(ElButton).use(ElInput).use(ElRow).use(ElCol).use(ElCard).use(ElSelect).use(ElTable).use(ElDivider).mount(parent)
        },500);
}

export {ShowString,ShowList}