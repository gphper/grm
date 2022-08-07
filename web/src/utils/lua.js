import { useStore } from 'vuex';
import LuaView from '@/components/show/LuaView.vue'
import { createApp, h } from 'vue';
import { ElButton, ElCol, ElInput, ElRow } from 'element-plus';


const showLua = (id,sk,db)=>{
    const store = useStore()
    const vdom = createApp({
        setup() {
        },
        render() {
            return h(
                LuaView,
                {
                   sk:sk,
                   db:db,
                },
                null
            )
        }
    });

    setTimeout(()=>{
        const parent = document.getElementById(id)
        vdom.use(store).use(ElRow).use(ElCol).use(ElButton).use(ElInput).mount(parent)
    },500);
}

export {showLua}