import StringView from '@/components/show/StringView.vue'
import { ElButton, ElCard, ElCol, ElInput, ElRow, ElSelect } from 'element-plus';
import { createApp, h } from 'vue';
// import OneText from '@/components/OneText.vue'

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

export {ShowString}