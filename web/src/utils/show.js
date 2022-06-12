import StringView from '@/components/show/StringView.vue';
import ListView from '@/components/show/ListView.vue'
import { ElButton, ElCard, ElCol, ElDivider, ElInput, ElRow, ElSelect, ElTable } from 'element-plus';
import { createApp, h } from 'vue';

const ShowString = (key,id)=>{
    const vdom = createApp({    
        setup() {
        },
        render() {
        
            return h(
                StringView,
                {
                    xkey:key,
                },
                null
            )
        }
        
        });

        setTimeout(()=>{
            const parent = document.getElementById(id)
            vdom.use(ElButton).use(ElInput).use(ElRow).use(ElCol).use(ElCard).use(ElSelect).mount(parent)
        },500);
}


const ShowList = (key,id)=>{
    const vdom = createApp({    
        setup() {
        },
        render() {
        
            return h(
                ListView,
                {
                    xkey:key,
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