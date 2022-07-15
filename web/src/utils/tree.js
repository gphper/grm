import { createApp,h, ref } from '@vue/runtime-dom'
import MenuTreeSolt from '@/components/MenuTreeSolt.vue'
import { ElTree,ElButton } from 'element-plus'
import {getKeys} from '@/api/index.js'
import store from '@/store/index.js'


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
                default: ({node,data})=>h(MenuTreeSolt,
                {
                    node,
                    data
                })
            }
        )
    }
    
    });
    
    const parent = document.getElementById(id)
    vdom.use(store).use(ElTree).use(ElButton).mount(parent)
};

const dbKeysList = (index)=>{
    getKeys({"index":index}).then((res) => {
        document.getElementById("menu#"+index).nextElementSibling.setAttribute("style","display:flex")
        genNode("tree#"+index,ref(res.data.data));
        document.getElementById("dbnum#"+index).innerHTML = res.data.count;
    });
}

export {dbKeysList}