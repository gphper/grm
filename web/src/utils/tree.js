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

const genNode = function(index){
    const vdom = createApp({
    setup() {
        let inde = index;
        return {
            inde
        }
    },
    render() {
        return h(
            MenuTree,
            {
                index:this.inde
            },
            null
        )
    }
    
    });
    
    if(store.state.globalTree["tree#"+index]){
        store.state.globalTree["tree#"+index].unmount();
    }
    store.commit("setGlobalTree",{key:"tree#"+index,vm:vdom})
    const parent = document.getElementById("tree#"+index)
    vdom.use(store).use(ElTree).use(ElButton).use(ElDialog).use(ElInput).use(ElOption).use(ElForm).use(ElSelect).use(ElTreeV2).mount(parent)
};


const dbKeysList = (index)=>{

}

export {dbKeysList,genNode}