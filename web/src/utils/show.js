import StringView from '@/views/StringView.vue'
import { ElButton, ElCard, ElCol, ElInput, ElRow } from 'element-plus';
import { createApp, h } from 'vue';

const ShowString = (key,data,id)=>{
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
            vdom.use(ElButton).use(ElInput).use(ElRow).use(ElCol).use(ElCard).mount(parent)
        },500);
}

export {ShowString}