import { createApp,h } from '@vue/runtime-dom'
import MenuTree from '@/components/MenuTree.vue'
import { ElTree,ElButton,ElDialog,ElInput,ElSelect,ElForm,ElOption, ElTreeV2 } from 'element-plus'
import store from '@/store/index.js'


ElTreeV2.setup1 = ElTreeV2.setup;
ElTreeV2.setup = (props, ctx) => {
    const obj = ElTreeV2.setup1(props, ctx);
    obj['itemSize'] = 38
    return obj;
};

const genNode = function(indexs,matchs,cursors){
    const vdom = createApp({
    setup() {
        let index = indexs;
        let match = matchs;
        let cursor = cursors;
        return {
            index,
            match,
            cursor,
        }
    },
    render() {
        return h(
            MenuTree,
            {
                index:this.index,
                match:this.match,
                cursor:this.cursor
            },
            null
        )
    }
    
    });
    
    if(store.state.globalTree["tree#"+indexs]){
        store.state.globalTree["tree#"+indexs].unmount();
    }
    store.commit("setGlobalTree",{key:"tree#"+indexs,vm:vdom})
    const parent = document.getElementById("tree#"+indexs)
    vdom.use(store).use(ElTree).use(ElButton).use(ElDialog).use(ElInput).use(ElOption).use(ElForm).use(ElSelect).use(ElTreeV2).mount(parent)
};


export {genNode}